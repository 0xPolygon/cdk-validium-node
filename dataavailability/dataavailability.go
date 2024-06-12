package dataavailability

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/0xPolygonHermez/zkevm-node/etherman/types"
	jsontypes "github.com/0xPolygonHermez/zkevm-node/jsonrpc/types"
	"github.com/0xPolygonHermez/zkevm-node/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	unexpectedHashTemplate      = "mismatch on transaction data for batch num %d. Expected hash %s, actual hash: %s"
	failedDataRetrievalTemplate = "failed to retrieve local data for batches %v: %s"
	invalidBatchRetrievalArgs   = "invalid L2 batch data retrieval arguments, %d != %d"
)

// DataSourcePriority defines where data is retrieved from
type DataSourcePriority string

const (
	// Local indicates data stored in this nodes DB
	Local DataSourcePriority = "local"
	// Trusted indicates data stored in the Trusted Sequencer
	Trusted DataSourcePriority = "trusted"
	// External indicates data stored in the Data Availability layer
	External DataSourcePriority = "external"
)

// DefaultPriority is the default order in which data is retrieved
var DefaultPriority = []DataSourcePriority{Local, Trusted, External}

// DataAvailability implements an abstract data availability integration
type DataAvailability struct {
	isTrustedSequencer bool
	state              stateInterface
	zkEVMClient        ZKEVMClientTrustedBatchesGetter
	backend            DABackender
	dataSourcePriority []DataSourcePriority
	ctx                context.Context
}

// New creates a DataAvailability instance
func New(
	isTrustedSequencer bool,
	backend DABackender,
	state stateInterface,
	zkEVMClient ZKEVMClientTrustedBatchesGetter,
	priority []DataSourcePriority,
) (*DataAvailability, error) {
	da := &DataAvailability{
		isTrustedSequencer: isTrustedSequencer,
		backend:            backend,
		state:              state,
		zkEVMClient:        zkEVMClient,
		ctx:                context.Background(),
		dataSourcePriority: priority,
	}
	if len(da.dataSourcePriority) == 0 {
		da.dataSourcePriority = DefaultPriority
	}
	err := da.backend.Init()
	return da, err
}

// PostSequence sends the sequence data to the data availability backend, and returns the dataAvailabilityMessage
// as expected by the contract
func (d *DataAvailability) PostSequence(ctx context.Context, sequences []types.Sequence) ([]byte, error) {
	batchesData := [][]byte{}
	for _, batch := range sequences {
		// Do not send to the DA backend data that will be stored to L1
		if batch.ForcedBatchTimestamp == 0 {
			batchesData = append(batchesData, batch.BatchL2Data)
		}
	}
	return d.backend.PostSequence(ctx, batchesData)
}

// GetBatchL2Data tries to return the data from a batch, in the following priorities. batchNums should not include forced batches.
// 1. From local DB
// 2. From Trusted Sequencer (if not self)
// 3. From DA backend
func (d *DataAvailability) GetBatchL2Data(batchNums []uint64, batchHashes []common.Hash, dataAvailabilityMessage []byte) ([][]byte, error) {
	if len(batchNums) != len(batchHashes) {
		return nil, fmt.Errorf(invalidBatchRetrievalArgs, len(batchNums), len(batchHashes))
	}

	for _, p := range d.dataSourcePriority {
		switch p {
		case Local:
			localData, err := d.state.GetBatchL2DataByNumbers(d.ctx, batchNums, nil)
			if err != nil {
				return nil, err
			}

			data, err := checkBatches(batchNums, batchHashes, localData)
			if err != nil {
				log.Warnf(failedDataRetrievalTemplate, batchNums, err.Error())
			} else {
				return data, nil
			}
		case Trusted:
			if !d.isTrustedSequencer {
				data, err := d.rpcData(batchNums, batchHashes, d.zkEVMClient.BatchesByNumbers)
				if err != nil {
					log.Warnf(failedDataRetrievalTemplate, batchNums, err.Error())
				} else {
					return data, nil
				}
			}
		case External:
			return d.backend.GetSequence(d.ctx, batchHashes, dataAvailabilityMessage)
		default:
			log.Warnf("invalid data retrieval priority: %s", p)
		}
	}

	return nil, errors.New("failed to retrieve l2 batch data")
}

func checkBatches(batchNumbers []uint64, expectedHashes []common.Hash, batchData map[uint64][]byte) ([][]byte, error) {
	if len(batchNumbers) != len(expectedHashes) {
		return nil, fmt.Errorf("invalid batch parameters")
	}
	result := make([][]byte, len(batchNumbers))
	for i := 0; i < len(batchNumbers); i++ {
		batchNumber := batchNumbers[i]
		expectedHash := expectedHashes[i]
		bd, ok := batchData[batchNumber]
		if !ok {
			return nil, fmt.Errorf("missing batch data: [%d] %s", batchNumber, expectedHash.Hex())
		}
		actualHash := crypto.Keccak256Hash(bd)
		if actualHash != expectedHash {
			err := fmt.Errorf(unexpectedHashTemplate, batchNumber, expectedHash, actualHash)
			log.Warnf("wrong local data for hash: %s", err.Error())
			return nil, err
		}
		result[i] = bd
	}
	return result, nil
}

type rpcBatchDataFunc func(ctx context.Context, numbers []*big.Int) ([]*jsontypes.BatchData, error)

// rpcData retrieves batch data from rpcBatchDataFunc, returns an error unless all are found and correct
func (d *DataAvailability) rpcData(batchNums []uint64, expectedHashes []common.Hash, rpcFunc rpcBatchDataFunc) ([][]byte, error) {
	if len(batchNums) != len(expectedHashes) {
		return nil, fmt.Errorf("invalid arguments, len of batch numbers does not equal length of expected hashes: %d != %d",
			len(batchNums), len(expectedHashes))
	}
	nums := make([]*big.Int, 0, len(batchNums))
	for _, n := range batchNums {
		nums = append(nums, new(big.Int).SetUint64(n))
	}
	batchData, err := rpcFunc(d.ctx, nums)
	if err != nil {
		return nil, err
	}
	if len(batchData) != len(batchNums) {
		return nil, fmt.Errorf("missing batch data, expected %d, got %d", len(batchNums), len(batchData))
	}
	result := make(map[uint64][]byte)
	for i := 0; i < len(batchNums); i++ {
		number := batchNums[i]
		batch := batchData[i]
		expectedTransactionsHash := expectedHashes[i]
		actualTransactionsHash := crypto.Keccak256Hash(batch.BatchL2Data)
		if expectedTransactionsHash != actualTransactionsHash {
			return nil, fmt.Errorf(unexpectedHashTemplate, number, expectedTransactionsHash, actualTransactionsHash)
		}
		result[number] = batch.BatchL2Data
	}
	checked, err := checkBatches(batchNums, expectedHashes, result)
	if err != nil {
		return nil, err
	}
	return checked, nil
}
