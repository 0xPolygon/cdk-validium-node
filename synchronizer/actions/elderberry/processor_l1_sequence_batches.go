package elderberry

import (
	"context"
	"errors"
	"time"

	"github.com/0xPolygonHermez/zkevm-node/etherman"
	"github.com/0xPolygonHermez/zkevm-node/log"
	"github.com/0xPolygonHermez/zkevm-node/state"
	"github.com/0xPolygonHermez/zkevm-node/synchronizer/actions"
	"github.com/jackc/pgx/v4"
)

var (
	// ErrInvalidInitialBatchNumber is returned when the initial batch number is not the expected one
	ErrInvalidInitialBatchNumber = errors.New("invalid initial batch number")
)

// PreviousProcessor is the interface that the previous processor (Etrog)
type PreviousProcessor interface {
	Process(ctx context.Context, order etherman.Order, l1Block *etherman.Block, dbTx pgx.Tx) error
	ProcessSequenceBatches(ctx context.Context, sequencedBatches []etherman.SequencedBatch, blockNumber uint64, l1BlockTimestamp time.Time, dbTx pgx.Tx) error
}

// StateL1SequenceBatchesElderberry state interface
type StateL1SequenceBatchesElderberry interface {
	GetLastVirtualBatchNum(ctx context.Context, dbTx pgx.Tx) (uint64, error)
	GetLastL2BlockByBatchNumber(ctx context.Context, batchNumber uint64, dbTx pgx.Tx) (*state.L2Block, error)
}

// ProcessorL1SequenceBatchesElderberry is the processor for SequenceBatches for Elderberry
type ProcessorL1SequenceBatchesElderberry struct {
	actions.ProcessorBase[ProcessorL1SequenceBatchesElderberry]
	previousProcessor PreviousProcessor
	state             StateL1SequenceBatchesElderberry
}

// NewProcessorL1SequenceBatchesElderberry returns instance of a processor for SequenceBatchesOrder
func NewProcessorL1SequenceBatchesElderberry(previousProcessor PreviousProcessor, state StateL1SequenceBatchesElderberry) *ProcessorL1SequenceBatchesElderberry {
	return &ProcessorL1SequenceBatchesElderberry{
		ProcessorBase: actions.ProcessorBase[ProcessorL1SequenceBatchesElderberry]{
			SupportedEvent:    []etherman.EventOrder{etherman.SequenceBatchesOrder},
			SupportedForkdIds: &actions.ForksIdOnlyElderberry},
		previousProcessor: previousProcessor,
		state:             state,
	}
}

// Process process event
func (g *ProcessorL1SequenceBatchesElderberry) Process(ctx context.Context, order etherman.Order, l1Block *etherman.Block, dbTx pgx.Tx) error {
	if l1Block == nil || len(l1Block.SequencedBatches) <= order.Pos {
		return actions.ErrInvalidParams
	}
	if len(l1Block.SequencedBatches[order.Pos]) == 0 {
		log.Warnf("No sequenced batches for position")
		return nil
	}

	sbatch := l1Block.SequencedBatches[order.Pos][0]

	executionTime := l1Block.ReceivedAt
	if sbatch.SequencedBatchElderberryData == nil {
		log.Warnf("No elderberry sequenced batch data for batch %d", sbatch.BatchNumber)
	} else {
		executionTime = time.Unix(int64(sbatch.SequencedBatchElderberryData.MaxSequenceTimestamp), 0)
	}

	return g.previousProcessor.ProcessSequenceBatches(ctx, l1Block.SequencedBatches[order.Pos], l1Block.BlockNumber, executionTime, dbTx)
}
