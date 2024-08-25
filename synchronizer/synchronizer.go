package synchronizer

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/0xPolygonHermez/zkevm-node/etherman"
	"github.com/0xPolygonHermez/zkevm-node/event"
	"github.com/0xPolygonHermez/zkevm-node/log"
	"github.com/0xPolygonHermez/zkevm-node/state"
	stateMetrics "github.com/0xPolygonHermez/zkevm-node/state/metrics"
	"github.com/0xPolygonHermez/zkevm-node/synchronizer/actions"
	"github.com/0xPolygonHermez/zkevm-node/synchronizer/actions/processor_manager"
	syncCommon "github.com/0xPolygonHermez/zkevm-node/synchronizer/common"
	"github.com/0xPolygonHermez/zkevm-node/synchronizer/common/syncinterfaces"
	"github.com/0xPolygonHermez/zkevm-node/synchronizer/l1_parallel_sync"
	"github.com/0xPolygonHermez/zkevm-node/synchronizer/l1event_orders"
	"github.com/0xPolygonHermez/zkevm-node/synchronizer/l2_sync/l2_shared"
	"github.com/0xPolygonHermez/zkevm-node/synchronizer/l2_sync/l2_sync_etrog"
	"github.com/0xPolygonHermez/zkevm-node/synchronizer/metrics"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/jackc/pgx/v4"
)

const (
	// ParallelMode is the value for L1SynchronizationMode to run in parallel mode
	ParallelMode = "parallel"
	// SequentialMode is the value for L1SynchronizationMode to run in sequential mode
	SequentialMode         = "sequential"
	maxBatchNumber         = ^uint64(0)
	timeOfLiveBatchOnCache = 5 * time.Minute
)

// Synchronizer connects L1 and L2
type Synchronizer interface {
	Sync() error
	Stop()
}

// TrustedState is the struct that contains the last trusted state root and the last trusted batches
type TrustedState struct {
	LastTrustedBatches []*state.Batch
	LastStateRoot      *common.Hash
}

// ClientSynchronizer connects L1 and L2
type ClientSynchronizer struct {
	isTrustedSequencer bool
	etherMan           syncinterfaces.EthermanFullInterface
	latestFlushID      uint64
	// If true the lastFlushID is stored in DB and we don't need to check again
	latestFlushIDIsFulfilled bool
	syncBlockProtection      rpc.BlockNumber
	etherManForL1            []syncinterfaces.EthermanFullInterface
	state                    syncinterfaces.StateFullInterface
	pool                     syncinterfaces.PoolInterface
	ethTxManager             syncinterfaces.EthTxManager
	zkEVMClient              syncinterfaces.ZKEVMClientInterface
	eventLog                 syncinterfaces.EventLogInterface
	ctx                      context.Context
	cancelCtx                context.CancelFunc
	genesis                  state.Genesis
	cfg                      Config
	// Id of the 'process' of the executor. Each time that it starts this value changes
	// This value is obtained from the call state.GetStoredFlushID
	// It starts as an empty string and it is filled in the first call
	// later the value is checked to be the same (in function checkFlushID)
	proverID string
	// Previous value returned by state.GetStoredFlushID, is used for decide if write a log or not
	previousExecutorFlushID  uint64
	l1SyncOrchestration      *l1_parallel_sync.L1SyncOrchestration
	l1EventProcessors        *processor_manager.L1EventProcessors
	syncTrustedStateExecutor syncinterfaces.SyncTrustedStateExecutor
	halter                   syncinterfaces.CriticalErrorHandler
}

// NewSynchronizer creates and initializes an instance of Synchronizer
func NewSynchronizer(
	isTrustedSequencer bool,
	ethMan syncinterfaces.EthermanFullInterface,
	etherManForL1 []syncinterfaces.EthermanFullInterface,
	st syncinterfaces.StateFullInterface,
	pool syncinterfaces.PoolInterface,
	ethTxManager syncinterfaces.EthTxManager,
	zkEVMClient syncinterfaces.ZKEVMClientInterface,
	eventLog syncinterfaces.EventLogInterface,
	genesis state.Genesis,
	cfg Config,
	runInDevelopmentMode bool) (Synchronizer, error) {
	ctx, cancel := context.WithCancel(context.Background())
	metrics.Register()
	syncBlockProtection, err := decodeSyncBlockProtection(cfg.SyncBlockProtection)
	if err != nil {
		log.Errorf("error decoding syncBlockProtection. Error: %v", err)
		cancel()
		return nil, err
	}
	log.Info("syncBlockProtection: ", syncBlockProtection)
	res := &ClientSynchronizer{
		isTrustedSequencer:      isTrustedSequencer,
		state:                   st,
		etherMan:                ethMan,
		etherManForL1:           etherManForL1,
		pool:                    pool,
		ctx:                     ctx,
		cancelCtx:               cancel,
		ethTxManager:            ethTxManager,
		zkEVMClient:             zkEVMClient,
		eventLog:                eventLog,
		genesis:                 genesis,
		cfg:                     cfg,
		proverID:                "",
		previousExecutorFlushID: 0,
		l1SyncOrchestration:     nil,
		l1EventProcessors:       nil,
		syncBlockProtection:     syncBlockProtection,
		halter:                  syncCommon.NewCriticalErrorHalt(eventLog, 5*time.Second), //nolint:gomnd
	}

	if !isTrustedSequencer {
		log.Info("Permissionless: creating and Initializing L2 synchronization components")
		L1SyncChecker := l2_sync_etrog.NewCheckSyncStatusToProcessBatch(res.zkEVMClient, res.state)
		sync := &res
		//syncTrustedStateEtrog := l2_sync_etrog.NewSyncTrustedBatchExecutorForEtrog(res.zkEVMClient, res.state, res.state, res,
		//	syncCommon.DefaultTimeProvider{}, L1SyncChecker, cfg.L2Synchronization)
		executorSteps := l2_sync_etrog.NewSyncTrustedBatchExecutorForEtrog(res.state, *sync)
		executor := l2_shared.NewProcessorTrustedBatchSync(executorSteps, syncCommon.DefaultTimeProvider{}, L1SyncChecker, cfg.L2Synchronization)
		if cfg.L2Synchronization.CheckLastL2BlockHashOnCloseBatch {
			log.Infof("Adding check of L2Block hash on close batch when sync from trusted node")
			executor.AddPostChecker(l2_shared.NewPostClosedBatchCheckL2Block(res.state))
		}

		syncTrustedStateEtrog := l2_shared.NewTrustedBatchesRetrieve(executor, zkEVMClient, res.state, *sync, *l2_shared.NewTrustedStateManager(syncCommon.DefaultTimeProvider{}, timeOfLiveBatchOnCache))
		res.syncTrustedStateExecutor = l2_shared.NewSyncTrustedStateExecutorSelector(map[uint64]syncinterfaces.SyncTrustedStateExecutor{
			uint64(state.FORKID_ETROG):      syncTrustedStateEtrog,
			uint64(state.FORKID_ELDERBERRY): syncTrustedStateEtrog,
			uint64(state.FORKID_9):          syncTrustedStateEtrog,
		}, res.state)
	}
	var l1checkerL2Blocks *actions.CheckL2BlockHash
	if cfg.L1SyncCheckL2BlockHash {
		if !isTrustedSequencer {
			log.Infof("Permissionless: L1SyncCheckL2BlockHash is enabled")
			initialL2Block, err := res.state.GetLastL2BlockNumber(res.ctx, nil)
			if errors.Is(err, state.ErrStateNotSynchronized) {
				initialL2Block = 1
				log.Info("State is empty, can't get last L2Block number. Using %d as initial L2Block number", initialL2Block)
			} else if err != nil {
				log.Errorf("error getting last L2Block number from state. Error: %v", err)
				return nil, err
			}
			l1checkerL2Blocks = actions.NewCheckL2BlockHash(res.state, res.zkEVMClient, initialL2Block, cfg.L1SyncCheckL2BlockNumberhModulus)
		} else {
			log.Infof("Trusted Node can't check L2Block hash, ignoring parameter")
		}
	}

	res.l1EventProcessors = defaultsL1EventProcessors(res, l1checkerL2Blocks)
	switch cfg.L1SynchronizationMode {
	case ParallelMode:
		log.Info("L1SynchronizationMode is parallel")
		res.l1SyncOrchestration = newL1SyncParallel(ctx, cfg, etherManForL1, res, runInDevelopmentMode)
	case SequentialMode:
		log.Info("L1SynchronizationMode is sequential")
	default:
		log.Fatalf("L1SynchronizationMode is not valid. Valid values are: %s, %s", ParallelMode, SequentialMode)
	}

	return res, nil
}

func decodeSyncBlockProtection(sBP string) (rpc.BlockNumber, error) {
	switch sBP {
	case "latest":
		return rpc.LatestBlockNumber, nil
	case "finalized":
		return rpc.FinalizedBlockNumber, nil
	case "safe":
		return rpc.SafeBlockNumber, nil
	default:
		return 0, fmt.Errorf("error decoding SyncBlockProtection. Unknown value")
	}
}

var waitDuration = time.Duration(0)

func newL1SyncParallel(ctx context.Context, cfg Config, etherManForL1 []syncinterfaces.EthermanFullInterface, sync *ClientSynchronizer, runExternalControl bool) *l1_parallel_sync.L1SyncOrchestration {
	chIncommingRollupInfo := make(chan l1_parallel_sync.L1SyncMessage, cfg.L1ParallelSynchronization.MaxPendingNoProcessedBlocks)
	cfgConsumer := l1_parallel_sync.ConfigConsumer{
		ApplyAfterNumRollupReceived: cfg.L1ParallelSynchronization.PerformanceWarning.ApplyAfterNumRollupReceived,
		AceptableInacctivityTime:    cfg.L1ParallelSynchronization.PerformanceWarning.AceptableInacctivityTime.Duration,
	}
	L1DataProcessor := l1_parallel_sync.NewL1RollupInfoConsumer(cfgConsumer, sync, chIncommingRollupInfo)

	cfgProducer := l1_parallel_sync.ConfigProducer{
		SyncChunkSize:                              cfg.SyncChunkSize,
		TtlOfLastBlockOnL1:                         cfg.L1ParallelSynchronization.RequestLastBlockPeriod.Duration,
		TimeoutForRequestLastBlockOnL1:             cfg.L1ParallelSynchronization.RequestLastBlockTimeout.Duration,
		NumOfAllowedRetriesForRequestLastBlockOnL1: cfg.L1ParallelSynchronization.RequestLastBlockMaxRetries,
		TimeForShowUpStatisticsLog:                 cfg.L1ParallelSynchronization.StatisticsPeriod.Duration,
		TimeOutMainLoop:                            cfg.L1ParallelSynchronization.TimeOutMainLoop.Duration,
		MinTimeBetweenRetriesForRollupInfo:         cfg.L1ParallelSynchronization.RollupInfoRetriesSpacing.Duration,
	}
	// Convert EthermanInterface to l1_sync_parallel.EthermanInterface
	etherManForL1Converted := make([]l1_parallel_sync.L1ParallelEthermanInterface, len(etherManForL1))
	for i, etherMan := range etherManForL1 {
		etherManForL1Converted[i] = etherMan
	}
	l1DataRetriever := l1_parallel_sync.NewL1DataRetriever(cfgProducer, etherManForL1Converted, chIncommingRollupInfo)
	l1SyncOrchestration := l1_parallel_sync.NewL1SyncOrchestration(ctx, l1DataRetriever, L1DataProcessor)
	if runExternalControl {
		log.Infof("Starting external control")
		externalControl := newExternalCmdControl(l1DataRetriever, l1SyncOrchestration)
		externalControl.start()
	}
	return l1SyncOrchestration
}

// CleanTrustedState Clean cache of TrustedBatches and StateRoot
func (s *ClientSynchronizer) CleanTrustedState() {
	if s.syncTrustedStateExecutor != nil {
		s.syncTrustedStateExecutor.CleanTrustedState()
	}
}

// IsTrustedSequencer returns true is a running in a trusted sequencer
func (s *ClientSynchronizer) IsTrustedSequencer() bool {
	return s.isTrustedSequencer
}

func rollback(ctx context.Context, dbTx pgx.Tx, err error) error {
	rollbackErr := dbTx.Rollback(ctx)
	if rollbackErr != nil {
		log.Errorf("error rolling back state. RollbackErr: %v,because err: %s", rollbackErr, err.Error())
		return rollbackErr
	}
	return err
}

func (s *ClientSynchronizer) isGenesisProcessed(ctx context.Context, dbTx pgx.Tx) (bool, *state.Block, error) {
	lastEthBlockSynced, err := s.state.GetLastBlock(ctx, dbTx)
	if err != nil && errors.Is(err, state.ErrStateNotSynchronized) {
		return false, lastEthBlockSynced, nil
	}

	if lastEthBlockSynced.BlockNumber >= s.genesis.RollupBlockNumber {
		log.Infof("Genesis block processed. Last block synced: %d >= genesis %d", lastEthBlockSynced.BlockNumber, s.genesis.RollupBlockNumber)
		return true, lastEthBlockSynced, nil
	}
	log.Warnf("Genesis block not processed yet. Last block synced: %d < genesis %d", lastEthBlockSynced.BlockNumber, s.genesis.RollupBlockNumber)
	return false, lastEthBlockSynced, nil
}

// getStartingL1Block find if need to update and if yes the starting point:
// bool -> need to process blocks
// uint64 -> first block to synchronize
// error -> error
// 1. First try to get last block on DB, if there are could be fully synced or pending blocks
// 2. If DB is empty the LxLy upgrade block as starting point
func (s *ClientSynchronizer) getStartingL1Block(ctx context.Context, genesisBlockNumber, rollupManagerBlockNumber uint64, dbTx pgx.Tx) (bool, uint64, error) {
	lastBlock, err := s.state.GetLastBlock(ctx, dbTx)
	if err != nil && errors.Is(err, state.ErrStateNotSynchronized) {
		// No block on DB
		upgradeLxLyBlockNumber := rollupManagerBlockNumber
		if upgradeLxLyBlockNumber == 0 {
			upgradeLxLyBlockNumber, err = s.etherMan.GetL1BlockUpgradeLxLy(ctx, genesisBlockNumber)
			if err != nil && errors.Is(err, etherman.ErrNotFound) {
				log.Infof("sync pregenesis: LxLy upgrade not detected before genesis block %d, it'll be sync as usual. Nothing to do yet", genesisBlockNumber)
				return false, 0, nil
			} else if err != nil {
				log.Errorf("sync pregenesis: error getting LxLy upgrade block. Error: %v", err)
				return false, 0, err
			}
		}
		if rollupManagerBlockNumber >= genesisBlockNumber {
			log.Infof("sync pregenesis: rollupManagerBlockNumber>=genesisBlockNumber  (%d>=%d). Nothing in pregenesis", rollupManagerBlockNumber, genesisBlockNumber)
			return false, 0, nil
		}
		log.Infof("sync pregenesis: No block on DB, starting from LxLy upgrade block (rollupManagerBlockNumber) %d", upgradeLxLyBlockNumber)
		return true, upgradeLxLyBlockNumber, nil
	} else if err != nil {
		log.Errorf("Error getting last Block on DB err:%v", err)
		return false, 0, err
	}
	if lastBlock.BlockNumber >= genesisBlockNumber-1 {
		log.Warnf("sync pregenesis: Last block processed is %d, which is greater or equal than the previous genesis block %d", lastBlock, genesisBlockNumber)
		return false, 0, nil
	}
	log.Infof("sync pregenesis: Continue processing pre-genesis blocks, last block processed on DB is %d", lastBlock.BlockNumber+1)
	return true, lastBlock.BlockNumber + 1, nil
}

func (s *ClientSynchronizer) synchronizePreGenesisRollupEvents(syncChunkSize uint64, ctx context.Context) error {
	// Sync events from RollupManager that happen before rollup creation
	startTime := time.Now()
	log.Info("synchronizing events from RollupManager that happen before rollup creation")
	needToUpdate, fromBlock, err := s.getStartingL1Block(ctx, s.genesis.RollupBlockNumber, s.genesis.RollupManagerBlockNumber, nil)
	if err != nil {
		log.Errorf("sync pregenesis: error getting starting L1 block. Error: %v", err)
		return err
	}
	if !needToUpdate {
		log.Infof("sync pregenesis: No need to process blocks before the genesis block %d", s.genesis.RollupBlockNumber)
		return nil
	}
	toBlockFinal := s.genesis.RollupBlockNumber - 1
	log.Infof("sync pregenesis: starting syncing pre genesis LxLy events from block %d to block %d (total %d blocks) chunk size %d",
		fromBlock, toBlockFinal, toBlockFinal-fromBlock+1, syncChunkSize)
	for i := fromBlock; true; i += syncChunkSize {
		toBlock := min(i+syncChunkSize-1, toBlockFinal)
		log.Debugf("sync pregenesis: syncing L1InfoTree from blocks [%d - %d] remains: %d", i, toBlock, toBlockFinal-toBlock)
		blocks, order, err := s.etherMan.GetRollupInfoByBlockRangePreviousRollupGenesis(s.ctx, i, &toBlock)
		if err != nil {
			log.Error("sync pregenesis: error getting rollupInfoByBlockRange before rollup genesis: ", err)
			return err
		}
		log.Debugf("sync pregenesis: syncing L1InfoTree from blocks [%d - %d] -> num_block:%d num_order:%d", i, toBlock, len(blocks), len(order))
		err = s.ProcessBlockRange(blocks, order)
		if err != nil {
			log.Error("sync pregenesis: error processing blocks before the genesis: ", err)
			return err
		}
		if toBlock == toBlockFinal {
			break
		}
	}
	elapsedTime := time.Since(startTime)
	log.Infof("sync pregenesis: sync L1InfoTree finish: from %d to %d total_block %d done in %s", fromBlock, toBlockFinal, toBlockFinal-fromBlock+1, &elapsedTime)
	return nil
}

func (s *ClientSynchronizer) processGenesis() (*state.Block, error) {
	log.Info("State is empty, verifying genesis block")
	valid, err := s.etherMan.VerifyGenBlockNumber(s.ctx, s.genesis.RollupBlockNumber)
	if err != nil {
		log.Error("error checking genesis block number. Error: ", err)
		return nil, err
	} else if !valid {
		log.Error("genesis Block number configured is not valid. It is required the block number where the PolygonZkEVM smc was deployed")
		return nil, fmt.Errorf("genesis Block number configured is not valid. It is required the block number where the PolygonZkEVM smc was deployed")
	}
	err = s.synchronizePreGenesisRollupEvents(s.cfg.SyncChunkSize, s.ctx)
	if err != nil {
		log.Error("error synchronizing pre genesis events: ", err)
		return nil, err
	}

	header, err := s.etherMan.HeaderByNumber(s.ctx, big.NewInt(0).SetUint64(s.genesis.RollupBlockNumber))
	if err != nil {
		log.Errorf("error getting l1 block header for block %d. Error: %v", s.genesis.RollupBlockNumber, err)
		return nil, err
	}
	log.Info("synchronizing rollup creation block")
	lastEthBlockSynced := &state.Block{
		BlockNumber: header.Number.Uint64(),
		BlockHash:   header.Hash(),
		ParentHash:  header.ParentHash,
		ReceivedAt:  time.Unix(int64(header.Time), 0),
	}
	dbTx, err := s.state.BeginStateTransaction(s.ctx)
	if err != nil {
		log.Errorf("error creating db transaction to get latest block. Error: %v", err)
		return nil, err
	}
	// This add the genesis block and set values on tree
	genesisRoot, err := s.state.SetGenesis(s.ctx, *lastEthBlockSynced, s.genesis, stateMetrics.SynchronizerCallerLabel, dbTx)
	if err != nil {
		log.Error("error setting genesis: ", err)
		return nil, rollback(s.ctx, dbTx, err)
	}
	err = s.RequestAndProcessRollupGenesisBlock(dbTx, lastEthBlockSynced)
	if err != nil {
		log.Error("error processing Rollup genesis block: ", err)
		return nil, rollback(s.ctx, dbTx, err)
	}

	if genesisRoot != s.genesis.Root {
		log.Errorf("Calculated newRoot should be %s instead of %s", s.genesis.Root.String(), genesisRoot.String())
		return nil, rollback(s.ctx, dbTx, fmt.Errorf("calculated newRoot should be %s instead of %s", s.genesis.Root.String(), genesisRoot.String()))
	}
	// Waiting for the flushID to be stored
	err = s.checkFlushID(dbTx)
	if err != nil {
		log.Error("error checking genesis flushID: ", err)
		return nil, rollback(s.ctx, dbTx, err)
	}
	if err := dbTx.Commit(s.ctx); err != nil {
		log.Errorf("error genesis committing dbTx, err: %v", err)
		return nil, rollback(s.ctx, dbTx, err)
	}
	log.Info("Genesis root matches! Stored genesis blocks.")
	return lastEthBlockSynced, nil
}

// Sync function will read the last state synced and will continue from that point.
// Sync() will read blockchain events to detect rollup updates
func (s *ClientSynchronizer) Sync() error {
	startInitialization := time.Now()
	// If there is no lastEthereumBlock means that sync from the beginning is necessary. If not, it continues from the retrieved ethereum block
	// Get the latest synced block. If there is no block on db, use genesis block
	log.Info("Sync started")

	genesisDone, lastEthBlockSynced, err := s.isGenesisProcessed(s.ctx, nil)
	if err != nil {
		log.Errorf("error checking if genesis is processed. Error: %v", err)
		return err
	}
	if !genesisDone {
		lastEthBlockSynced, err = s.processGenesis()
		if err != nil {
			log.Errorf("error processing genesis. Error: %v", err)
			return err
		}
	}

	dbTx, err := s.state.BeginStateTransaction(s.ctx)
	if err != nil {
		log.Errorf("error creating db transaction to get latest block. Error: %v", err)
		return err
	}

	initBatchNumber, err := s.state.GetLastBatchNumber(s.ctx, dbTx)
	if err != nil {
		log.Error("error getting latest batchNumber synced. Error: ", err)
		rollbackErr := dbTx.Rollback(s.ctx)
		if rollbackErr != nil {
			log.Errorf("error rolling back state. RollbackErr: %v, err: %s", rollbackErr, err.Error())
			return rollbackErr
		}
		return err
	}
	err = s.state.SetInitSyncBatch(s.ctx, initBatchNumber, dbTx)
	if err != nil {
		log.Error("error setting initial batch number. Error: ", err)
		rollbackErr := dbTx.Rollback(s.ctx)
		if rollbackErr != nil {
			log.Errorf("error rolling back state. RollbackErr: %v, err: %s", rollbackErr, err.Error())
			return rollbackErr
		}
		return err
	}
	if err := dbTx.Commit(s.ctx); err != nil {
		log.Errorf("error committing dbTx, err: %v", err)
		rollbackErr := dbTx.Rollback(s.ctx)
		if rollbackErr != nil {
			log.Errorf("error rolling back state. RollbackErr: %v, err: %s", rollbackErr, err.Error())
			return rollbackErr
		}
		return err
	}
	metrics.InitializationTime(time.Since(startInitialization))

	for {
		select {
		case <-s.ctx.Done():
			return nil
		case <-time.After(waitDuration):
			start := time.Now()
			latestSequencedBatchNumber, err := s.etherMan.GetLatestBatchNumber()
			if err != nil {
				log.Warn("error getting latest sequenced batch in the rollup. Error: ", err)
				continue
			}
			latestSyncedBatch, err := s.state.GetLastBatchNumber(s.ctx, nil)
			metrics.LastSyncedBatchNumber(float64(latestSyncedBatch))
			if err != nil {
				log.Warn("error getting latest batch synced in the db. Error: ", err)
				continue
			}
			// Check the latest verified Batch number in the smc
			lastVerifiedBatchNumber, err := s.etherMan.GetLatestVerifiedBatchNum()
			if err != nil {
				log.Warn("error getting last verified batch in the rollup. Error: ", err)
				continue
			}
			err = s.state.SetLastBatchInfoSeenOnEthereum(s.ctx, latestSequencedBatchNumber, lastVerifiedBatchNumber, nil)
			if err != nil {
				log.Warn("error setting latest batch info into db. Error: ", err)
				continue
			}
			log.Infof("latestSequencedBatchNumber: %d, latestSyncedBatch: %d, lastVerifiedBatchNumber: %d", latestSequencedBatchNumber, latestSyncedBatch, lastVerifiedBatchNumber)
			// Sync trusted state
			// latestSyncedBatch -> Last batch on DB
			// latestSequencedBatchNumber -> last batch on SMC
			if latestSyncedBatch >= latestSequencedBatchNumber {
				startTrusted := time.Now()
				if s.syncTrustedStateExecutor != nil && !s.isTrustedSequencer {
					log.Info("Syncing trusted state (permissionless)")
					err = s.syncTrustedState(latestSyncedBatch)
					metrics.FullTrustedSyncTime(time.Since(startTrusted))
					if err != nil {
						log.Warn("error syncing trusted state. Error: ", err)
						s.CleanTrustedState()
						if errors.Is(err, syncinterfaces.ErrFatalDesyncFromL1) {
							l1BlockNumber := err.(*l2_shared.DeSyncPermissionlessAndTrustedNodeError).L1BlockNumber
							log.Error("Trusted and permissionless desync! reseting to last common point: L1Block (%d-1)", l1BlockNumber)
							err = s.resetState(l1BlockNumber - 1)
							if err != nil {
								log.Errorf("error resetting the state to a discrepancy block. Retrying... Err: %v", err)
								continue
							}
						} else if errors.Is(err, syncinterfaces.ErrMissingSyncFromL1) {
							log.Info("Syncing from trusted node need data from L1")
						} else if errors.Is(err, syncinterfaces.ErrCantSyncFromL2) {
							log.Info("Can't sync from L2, going to sync from L1")
						} else {
							// We break for resync from Trusted
							log.Debug("Sleeping for 1 second to avoid respawn too fast, error: ", err)
							time.Sleep(time.Second)
							continue
						}
					}
				}
				waitDuration = s.cfg.SyncInterval.Duration
			}
			//Sync L1Blocks
			startL1 := time.Now()
			if s.l1SyncOrchestration != nil && (latestSyncedBatch < latestSequencedBatchNumber || !s.cfg.L1ParallelSynchronization.FallbackToSequentialModeOnSynchronized) {
				log.Infof("Syncing L1 blocks in parallel lastEthBlockSynced=%d", lastEthBlockSynced.BlockNumber)
				lastEthBlockSynced, err = s.syncBlocksParallel(lastEthBlockSynced)
			} else {
				if s.l1SyncOrchestration != nil {
					log.Infof("Switching to sequential mode, stopping parallel sync and deleting object")
					s.l1SyncOrchestration.Abort()
					s.l1SyncOrchestration = nil
				}
				log.Infof("Syncing L1 blocks sequentially lastEthBlockSynced=%d", lastEthBlockSynced.BlockNumber)
				lastEthBlockSynced, err = s.syncBlocksSequential(lastEthBlockSynced)
			}
			metrics.FullL1SyncTime(time.Since(startL1))
			if err != nil {
				log.Warn("error syncing blocks: ", err)
				s.CleanTrustedState()
				lastEthBlockSynced, err = s.state.GetLastBlock(s.ctx, nil)
				if err != nil {
					log.Fatal("error getting lastEthBlockSynced to resume the synchronization... Error: ", err)
				}
				if s.l1SyncOrchestration != nil {
					// If have failed execution and get starting point from DB, we must reset parallel sync to this point
					// producer must start requesting this block
					s.l1SyncOrchestration.Reset(lastEthBlockSynced.BlockNumber)
				}
				if s.ctx.Err() != nil {
					continue
				}
			}
			metrics.FullSyncIterationTime(time.Since(start))
			log.Info("L1 state fully synchronized")
		}
	}
}

// RequestAndProcessRollupGenesisBlock it requests the rollup genesis block and processes it
//
//	and execute it
func (s *ClientSynchronizer) RequestAndProcessRollupGenesisBlock(dbTx pgx.Tx, lastEthBlockSynced *state.Block) error {
	blocks, order, err := s.etherMan.GetRollupInfoByBlockRange(s.ctx, lastEthBlockSynced.BlockNumber, &lastEthBlockSynced.BlockNumber)
	if err != nil {
		log.Error("error getting rollupInfoByBlockRange after set the genesis: ", err)
		return err
	}
	// Check that the response is the expected. It should be 1 block with ForkID event
	log.Debugf("SanityCheck for genesis block (%d) events: %+v", lastEthBlockSynced.BlockNumber, order)
	err = sanityCheckForGenesisBlockRollupInfo(blocks, order)
	if err != nil {
		return err
	}
	log.Infof("Processing genesis block %d orders: %+v", lastEthBlockSynced.BlockNumber, order)
	err = s.internalProcessBlock(blocks[0], order[blocks[0].BlockHash], false, dbTx)
	if err != nil {
		log.Errorf("error processinge events on genesis block %d:  err:%w", lastEthBlockSynced.BlockNumber, err)
	}

	return nil
}

func sanityCheckForGenesisBlockRollupInfo(blocks []etherman.Block, order map[common.Hash][]etherman.Order) error {
	if len(blocks) != 1 || len(order) < 1 || len(order[blocks[0].BlockHash]) < 1 {
		log.Errorf("error getting rollupInfoByBlockRange after set the genesis. Expected 1 block with minimum 2 orders")
		return fmt.Errorf("error getting rollupInfoByBlockRange after set the genesis. Expected 1 block with minimum 2 orders")
	}
	// The genesis block implies 1 ForkID event
	for _, value := range order[blocks[0].BlockHash] {
		if value.Name == etherman.ForkIDsOrder {
			return nil
		}
	}
	err := fmt.Errorf("events on genesis block (%d) need a ForkIDsOrder event but this block got %+v",
		blocks[0].BlockNumber, order[blocks[0].BlockHash])
	log.Error(err.Error())
	return err
}

// This function syncs the node from a specific block to the latest
// lastEthBlockSynced -> last block synced in the db
func (s *ClientSynchronizer) syncBlocksParallel(lastEthBlockSynced *state.Block) (*state.Block, error) {
	// This function will read events fromBlockNum to latestEthBlock. Check reorg to be sure that everything is ok.
	block, err := s.checkReorg(lastEthBlockSynced)
	if err != nil {
		log.Errorf("error checking reorgs. Retrying... Err: %v", err)
		return lastEthBlockSynced, fmt.Errorf("error checking reorgs")
	}
	if block != nil {
		log.Infof("reorg detected. Resetting the state from block %v to block %v", lastEthBlockSynced.BlockNumber, block.BlockNumber)
		err = s.resetState(block.BlockNumber)
		if err != nil {
			log.Errorf("error resetting the state to a previous block. Retrying... Err: %v", err)
			s.l1SyncOrchestration.Reset(lastEthBlockSynced.BlockNumber)
			return lastEthBlockSynced, fmt.Errorf("error resetting the state to a previous block")
		}
		return block, nil
	}
	log.Infof("Starting L1 sync orchestrator in parallel block: %d", lastEthBlockSynced.BlockNumber)
	return s.l1SyncOrchestration.Start(lastEthBlockSynced)
}

// This function syncs the node from a specific block to the latest
func (s *ClientSynchronizer) syncBlocksSequential(lastEthBlockSynced *state.Block) (*state.Block, error) {
	// This function will read events fromBlockNum to latestEthBlock. Check reorg to be sure that everything is ok.
	block, err := s.checkReorg(lastEthBlockSynced)
	if err != nil {
		log.Errorf("error checking reorgs. Retrying... Err: %v", err)
		return lastEthBlockSynced, fmt.Errorf("error checking reorgs")
	}
	if block != nil {
		err = s.resetState(block.BlockNumber)
		if err != nil {
			log.Errorf("error resetting the state to a previous block. Retrying... Err: %v", err)
			return lastEthBlockSynced, fmt.Errorf("error resetting the state to a previous block")
		}
		return block, nil
	}

	// Call the blockchain to retrieve data
	header, err := s.etherMan.HeaderByNumber(s.ctx, big.NewInt(s.syncBlockProtection.Int64()))
	if err != nil {
		return lastEthBlockSynced, err
	}
	lastKnownBlock := header.Number

	var fromBlock uint64
	if lastEthBlockSynced.BlockNumber > 0 {
		fromBlock = lastEthBlockSynced.BlockNumber + 1
	}

	for {
		toBlock := fromBlock + s.cfg.SyncChunkSize
		if toBlock > lastKnownBlock.Uint64() {
			toBlock = lastKnownBlock.Uint64()
		}
		log.Infof("Syncing block %d of %d", fromBlock, lastKnownBlock.Uint64())
		log.Infof("Getting rollup info from block %d to block %d", fromBlock, toBlock)
		// This function returns the rollup information contained in the ethereum blocks and an extra param called order.
		// Order param is a map that contains the event order to allow the synchronizer store the info in the same order that is readed.
		// Name can be different in the order struct. For instance: Batches or Name:NewSequencers. This name is an identifier to check
		// if the next info that must be stored in the db is a new sequencer or a batch. The value pos (position) tells what is the
		// array index where this value is.
		start := time.Now()
		blocks, order, err := s.etherMan.GetRollupInfoByBlockRange(s.ctx, fromBlock, &toBlock)
		metrics.ReadL1DataTime(time.Since(start))
		if err != nil {
			return lastEthBlockSynced, err
		}
		start = time.Now()
		err = s.ProcessBlockRange(blocks, order)
		metrics.ProcessL1DataTime(time.Since(start))
		if err != nil {
			return lastEthBlockSynced, err
		}
		if len(blocks) > 0 {
			lastEthBlockSynced = &state.Block{
				BlockNumber: blocks[len(blocks)-1].BlockNumber,
				BlockHash:   blocks[len(blocks)-1].BlockHash,
				ParentHash:  blocks[len(blocks)-1].ParentHash,
				ReceivedAt:  blocks[len(blocks)-1].ReceivedAt,
			}
			for i := range blocks {
				log.Debug("Position: ", i, ". BlockNumber: ", blocks[i].BlockNumber, ". BlockHash: ", blocks[i].BlockHash)
			}
		}
		fromBlock = toBlock + 1

		if lastKnownBlock.Cmp(new(big.Int).SetUint64(toBlock)) < 1 {
			waitDuration = s.cfg.SyncInterval.Duration
			break
		}
		if len(blocks) == 0 { // If there is no events in the checked blocks range and lastKnownBlock > fromBlock.
			// Store the latest block of the block range. Get block info and process the block
			fb, err := s.etherMan.EthBlockByNumber(s.ctx, toBlock)
			if err != nil {
				return lastEthBlockSynced, err
			}
			b := etherman.Block{
				BlockNumber: fb.NumberU64(),
				BlockHash:   fb.Hash(),
				ParentHash:  fb.ParentHash(),
				ReceivedAt:  time.Unix(int64(fb.Time()), 0),
			}
			err = s.ProcessBlockRange([]etherman.Block{b}, order)
			if err != nil {
				return lastEthBlockSynced, err
			}
			block := state.Block{
				BlockNumber: fb.NumberU64(),
				BlockHash:   fb.Hash(),
				ParentHash:  fb.ParentHash(),
				ReceivedAt:  time.Unix(int64(fb.Time()), 0),
			}
			lastEthBlockSynced = &block
			log.Debug("Storing empty block. BlockNumber: ", b.BlockNumber, ". BlockHash: ", b.BlockHash)
		}
	}

	return lastEthBlockSynced, nil
}

// ProcessBlockRange process the L1 events and stores the information in the db
func (s *ClientSynchronizer) ProcessBlockRange(blocks []etherman.Block, order map[common.Hash][]etherman.Order) error {
	for i := range blocks {
		// Begin db transaction
		dbTx, err := s.state.BeginStateTransaction(s.ctx)
		if err != nil {
			log.Errorf("error creating db transaction to store block. BlockNumber: %d, error: %v", blocks[i].BlockNumber, err)
			return err
		}
		err = s.internalProcessBlock(blocks[i], order[blocks[i].BlockHash], true, dbTx)
		if err != nil {
			log.Error("rollingback BlockNumber: %d, because error internalProcessBlock: ", blocks[i].BlockNumber, err)
			// If any goes wrong we ensure that the state is rollbacked
			rollbackErr := dbTx.Rollback(s.ctx)
			if rollbackErr != nil && !errors.Is(rollbackErr, pgx.ErrTxClosed) {
				log.Warnf("error rolling back state to store block. BlockNumber: %d, rollbackErr: %s, error : %v", blocks[i].BlockNumber, rollbackErr.Error(), err)
				return fmt.Errorf("error rollback BlockNumber: %d: err:%w original error:%w", blocks[i].BlockNumber, rollbackErr, err)
			}
			return err
		}

		err = dbTx.Commit(s.ctx)
		if err != nil {
			log.Errorf("error committing state to store block. BlockNumber: %d, err: %v", blocks[i].BlockNumber, err)
			rollbackErr := dbTx.Rollback(s.ctx)
			if rollbackErr != nil {
				log.Errorf("error rolling back state to store block. BlockNumber: %d, rollbackErr: %s, error : %v", blocks[i].BlockNumber, rollbackErr.Error(), err)
				return rollbackErr
			}
			return err
		}
	}
	return nil
}

// ProcessBlockRange process the L1 events and stores the information in the db
func (s *ClientSynchronizer) internalProcessBlock(blocks etherman.Block, order []etherman.Order, addBlock bool, dbTx pgx.Tx) error {
	var err error
	// New info has to be included into the db using the state
	if addBlock {
		b := state.Block{
			BlockNumber: blocks.BlockNumber,
			BlockHash:   blocks.BlockHash,
			ParentHash:  blocks.ParentHash,
			ReceivedAt:  blocks.ReceivedAt,
		}
		// Add block information
		log.Debugf("Storing block. Block: %s", b.String())
		err = s.state.AddBlock(s.ctx, &b, dbTx)
		if err != nil {
			log.Errorf("error storing block. BlockNumber: %d, error: %v", blocks.BlockNumber, err)
			return err
		}
	}

	for _, element := range order {
		batchSequence := l1event_orders.GetSequenceFromL1EventOrder(element.Name, &blocks, element.Pos)
		var forkId uint64
		if batchSequence != nil {
			forkId = s.state.GetForkIDByBatchNumber(batchSequence.FromBatchNumber)
			log.Debug("EventOrder: ", element.Name, ". Batch Sequence: ", batchSequence, "forkId: ", forkId)
		} else {
			forkId = s.state.GetForkIDByBlockNumber(blocks.BlockNumber)
			log.Debug("EventOrder: ", element.Name, ". BlockNumber: ", blocks.BlockNumber, "forkId: ", forkId)
		}
		forkIdTyped := actions.ForkIdType(forkId)
		// Process event received from l1
		err := s.l1EventProcessors.Process(s.ctx, forkIdTyped, element, &blocks, dbTx)
		if err != nil {
			log.Error("1EventProcessors.Process error: ", err)
			return err
		}
	}
	log.Debug("Checking FlushID to commit L1 data to db")
	err = s.checkFlushID(dbTx)
	if err != nil {
		log.Errorf("error checking flushID. Error: %v", err)
		return err
	}

	return nil
}

func (s *ClientSynchronizer) syncTrustedState(latestSyncedBatch uint64) error {
	if s.syncTrustedStateExecutor == nil || s.isTrustedSequencer {
		return nil
	}

	return s.syncTrustedStateExecutor.SyncTrustedState(s.ctx, latestSyncedBatch, maxBatchNumber)
}

// This function allows reset the state until an specific ethereum block
func (s *ClientSynchronizer) resetState(blockNumber uint64) error {
	log.Info("Reverting synchronization to block: ", blockNumber)
	dbTx, err := s.state.BeginStateTransaction(s.ctx)
	if err != nil {
		log.Error("error starting a db transaction to reset the state. Error: ", err)
		return err
	}
	err = s.state.Reset(s.ctx, blockNumber, dbTx)
	if err != nil {
		rollbackErr := dbTx.Rollback(s.ctx)
		if rollbackErr != nil {
			log.Errorf("error rolling back state to store block. BlockNumber: %d, rollbackErr: %s, error : %v", blockNumber, rollbackErr.Error(), err)
			return rollbackErr
		}
		log.Error("error resetting the state. Error: ", err)
		return err
	}
	err = s.ethTxManager.Reorg(s.ctx, blockNumber+1, dbTx)
	if err != nil {
		rollbackErr := dbTx.Rollback(s.ctx)
		if rollbackErr != nil {
			log.Errorf("error rolling back eth tx manager when reorg detected. BlockNumber: %d, rollbackErr: %s, error : %v", blockNumber, rollbackErr.Error(), err)
			return rollbackErr
		}
		log.Error("error processing reorg on eth tx manager. Error: ", err)
		return err
	}
	err = dbTx.Commit(s.ctx)
	if err != nil {
		rollbackErr := dbTx.Rollback(s.ctx)
		if rollbackErr != nil {
			log.Errorf("error rolling back state to store block. BlockNumber: %d, rollbackErr: %s, error : %v", blockNumber, rollbackErr.Error(), err)
			return rollbackErr
		}
		log.Error("error committing the resetted state. Error: ", err)
		return err
	}
	if s.l1SyncOrchestration != nil {
		s.l1SyncOrchestration.Reset(blockNumber)
	}
	return nil
}

/*
This function will check if there is a reorg.
As input param needs the last ethereum block synced. Retrieve the block info from the blockchain
to compare it with the stored info. If hash and hash parent matches, then no reorg is detected and return a nil.
If hash or hash parent don't match, reorg detected and the function will return the block until the sync process
must be reverted. Then, check the previous ethereum block synced, get block info from the blockchain and check
hash and has parent. This operation has to be done until a match is found.
*/
func (s *ClientSynchronizer) checkReorg(latestBlock *state.Block) (*state.Block, error) {
	// This function only needs to worry about reorgs if some of the reorganized blocks contained rollup info.
	latestEthBlockSynced := *latestBlock
	var depth uint64
	for {
		block, err := s.etherMan.EthBlockByNumber(s.ctx, latestBlock.BlockNumber)
		if err != nil {
			log.Errorf("error getting latest block synced from blockchain. Block: %d, error: %v", latestBlock.BlockNumber, err)
			return nil, err
		}
		log.Infof("[checkReorg function] BlockNumber: %d BlockHash got from L1 provider: %s", block.Number().Uint64(), block.Hash().String())
		log.Infof("[checkReorg function] latestBlockNumber: %d latestBlockHash already synced: %s", latestBlock.BlockNumber, latestBlock.BlockHash.String())
		if block.NumberU64() != latestBlock.BlockNumber {
			err = fmt.Errorf("wrong ethereum block retrieved from blockchain. Block numbers don't match. BlockNumber stored: %d. BlockNumber retrieved: %d",
				latestBlock.BlockNumber, block.NumberU64())
			log.Error("error: ", err)
			return nil, err
		}
		// Compare hashes
		if (block.Hash() != latestBlock.BlockHash || block.ParentHash() != latestBlock.ParentHash) && latestBlock.BlockNumber > s.genesis.RollupBlockNumber {
			log.Infof("checkReorg: Bad block %d hashOk %t parentHashOk %t", latestBlock.BlockNumber, block.Hash() == latestBlock.BlockHash, block.ParentHash() == latestBlock.ParentHash)
			log.Debug("[checkReorg function] => latestBlockNumber: ", latestBlock.BlockNumber)
			log.Debug("[checkReorg function] => latestBlockHash: ", latestBlock.BlockHash)
			log.Debug("[checkReorg function] => latestBlockHashParent: ", latestBlock.ParentHash)
			log.Debug("[checkReorg function] => BlockNumber: ", latestBlock.BlockNumber, block.NumberU64())
			log.Debug("[checkReorg function] => BlockHash: ", block.Hash())
			log.Debug("[checkReorg function] => BlockHashParent: ", block.ParentHash())
			depth++
			log.Debug("REORG: Looking for the latest correct ethereum block. Depth: ", depth)
			// Reorg detected. Getting previous block
			dbTx, err := s.state.BeginStateTransaction(s.ctx)
			if err != nil {
				log.Errorf("error creating db transaction to get prevoius blocks")
				return nil, err
			}
			latestBlock, err = s.state.GetPreviousBlock(s.ctx, depth, dbTx)
			errC := dbTx.Commit(s.ctx)
			if errC != nil {
				log.Errorf("error committing dbTx, err: %v", errC)
				rollbackErr := dbTx.Rollback(s.ctx)
				if rollbackErr != nil {
					log.Errorf("error rolling back state. RollbackErr: %v", rollbackErr)
					return nil, rollbackErr
				}
				log.Errorf("error committing dbTx, err: %v", errC)
				return nil, errC
			}
			if errors.Is(err, state.ErrNotFound) {
				log.Warn("error checking reorg: previous block not found in db: ", err)
				return &state.Block{}, nil
			} else if err != nil {
				return nil, err
			}
		} else {
			break
		}
	}
	if latestEthBlockSynced.BlockHash != latestBlock.BlockHash {
		log.Info("Reorg detected in block: ", latestEthBlockSynced.BlockNumber, " last block OK: ", latestBlock.BlockNumber)
		return latestBlock, nil
	}
	return nil, nil
}

// Stop function stops the synchronizer
func (s *ClientSynchronizer) Stop() {
	s.cancelCtx()
}

// PendingFlushID is called when a flushID is pending to be stored in the db
func (s *ClientSynchronizer) PendingFlushID(flushID uint64, proverID string) {
	log.Infof("pending flushID: %d", flushID)
	if flushID == 0 {
		log.Fatal("flushID is 0. Please check that prover/executor config parameter dbReadOnly is false")
	}
	s.latestFlushID = flushID
	s.latestFlushIDIsFulfilled = false
	s.updateAndCheckProverID(proverID)
}

// deprecated: use PendingFlushID instead
//
//nolint:unused
func (s *ClientSynchronizer) pendingFlushID(flushID uint64, proverID string) {
	s.PendingFlushID(flushID, proverID)
}

func (s *ClientSynchronizer) updateAndCheckProverID(proverID string) {
	if s.proverID == "" {
		log.Infof("Current proverID is %s", proverID)
		s.proverID = proverID
		return
	}
	if s.proverID != proverID {
		event := &event.Event{
			ReceivedAt:  time.Now(),
			Source:      event.Source_Node,
			Component:   event.Component_Synchronizer,
			Level:       event.Level_Critical,
			EventID:     event.EventID_SynchronizerRestart,
			Description: fmt.Sprintf("proverID changed from %s to %s, restarting Synchronizer ", s.proverID, proverID),
		}

		err := s.eventLog.LogEvent(context.Background(), event)
		if err != nil {
			log.Errorf("error storing event payload: %v", err)
		}

		log.Fatal("restarting synchronizer because executor has been restarted (old=%s, new=%s)", s.proverID, proverID)
	}
}

// CheckFlushID is called when a flushID is pending to be stored in the db
func (s *ClientSynchronizer) CheckFlushID(dbTx pgx.Tx) error {
	return s.checkFlushID(dbTx)
}

func (s *ClientSynchronizer) checkFlushID(dbTx pgx.Tx) error {
	if s.latestFlushIDIsFulfilled {
		log.Debugf("no pending flushID, nothing to do. Last pending fulfilled flushID: %d, last executor flushId received: %d", s.latestFlushID, s.latestFlushID)
		return nil
	}
	storedFlushID, proverID, err := s.state.GetStoredFlushID(s.ctx)
	if err != nil {
		log.Error("error getting stored flushID. Error: ", err)
		return err
	}
	if s.previousExecutorFlushID != storedFlushID || s.proverID != proverID {
		log.Infof("executor vs local: flushid=%d/%d, proverID=%s/%s", storedFlushID,
			s.latestFlushID, proverID, s.proverID)
	} else {
		log.Debugf("executor vs local: flushid=%d/%d, proverID=%s/%s", storedFlushID,
			s.latestFlushID, proverID, s.proverID)
	}
	s.updateAndCheckProverID(proverID)
	log.Debugf("storedFlushID (executor reported): %d, latestFlushID (pending): %d", storedFlushID, s.latestFlushID)
	if storedFlushID < s.latestFlushID {
		log.Infof("Synchronized BLOCKED!: Wating for the flushID to be stored. FlushID to be stored: %d. Latest flushID stored: %d", s.latestFlushID, storedFlushID)
		iteration := 0
		start := time.Now()
		for storedFlushID < s.latestFlushID {
			log.Debugf("Waiting for the flushID to be stored. FlushID to be stored: %d. Latest flushID stored: %d iteration:%d elpased:%s",
				s.latestFlushID, storedFlushID, iteration, time.Since(start))
			time.Sleep(100 * time.Millisecond) //nolint:gomnd
			storedFlushID, _, err = s.state.GetStoredFlushID(s.ctx)
			if err != nil {
				log.Error("error getting stored flushID. Error: ", err)
				return err
			}
			iteration++
		}
		log.Infof("Synchronizer resumed, flushID stored: %d", s.latestFlushID)
	}
	log.Infof("Pending Flushid fullfiled: %d, executor have write %d", s.latestFlushID, storedFlushID)
	s.latestFlushIDIsFulfilled = true
	s.previousExecutorFlushID = storedFlushID
	return nil
}

const (
	//L2BlockHeaderForGenesis = "0b73e6af6f00000000"
	L2BlockHeaderForGenesis = "0b0000000000000000"
)
