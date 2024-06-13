package etherman

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/jackc/pgx/v4"
)

type dataAvailabilityProvider interface {
	GetBatchL2Data(batchNum []uint64, hash []common.Hash, dataAvailabilityMessage []byte) ([][]byte, error)
}

type stateProvider interface {
	GetForcedBatchDataByNumbers(ctx context.Context, batchNumbers []uint64, dbTx pgx.Tx) (map[uint64][]byte, error)
}
