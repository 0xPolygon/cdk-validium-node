package sequencesender

import (
	"github.com/0xPolygon/cdk-validium-node/config/types"
	"github.com/ethereum/go-ethereum/common"
)

// Config represents the configuration of a sequence sender
type Config struct {
	// WaitPeriodSendSequence is the time the sequencer waits until
	// trying to send a sequence to L1
	WaitPeriodSendSequence types.Duration `mapstructure:"WaitPeriodSendSequence"`
	// LastBatchVirtualizationTimeMaxWaitPeriod is time since sequences should be sent
	LastBatchVirtualizationTimeMaxWaitPeriod types.Duration `mapstructure:"LastBatchVirtualizationTimeMaxWaitPeriod"`
	// MaxBatchesForL1 is the maximum amount of batches to be sequenced in a single L1 tx
	MaxBatchesForL1 uint64 `mapstructure:"MaxBatchesForL1"`
	// SenderAddress defines which private key the eth tx manager needs to use
	// to sign the L1 txs
	SenderAddress common.Address
	// L2Coinbase defines which addess is going to receive the fees
	L2Coinbase common.Address `mapstructure:"L2Coinbase"`
	// PrivateKey defines all the key store files that are going
	// to be read in order to provide the private keys to sign the L1 txs
	PrivateKey types.KeystoreFileConfig `mapstructure:"PrivateKey"`
	// Batch number where there is a forkid change (fork upgrade)
	ForkUpgradeBatchNumber uint64
}
