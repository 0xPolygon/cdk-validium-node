package synchronizer

import (
	"github.com/0xPolygon/cdk-validium-node/config/types"
)

// Config represents the configuration of the synchronizer
type Config struct {
	// SyncInterval is the delay interval between reading new rollup information
	SyncInterval types.Duration `mapstructure:"SyncInterval"`
	// SyncChunkSize is the number of blocks to sync on each chunk
	SyncChunkSize uint64 `mapstructure:"SyncChunkSize"`
	// TrustedSequencerURL is the rpc url to connect and sync the trusted state
	TrustedSequencerURL string `mapstructure:"TrustedSequencerURL"`
	// NEAR DA related info
	DaAccount           string `mapstructure:"DaAccount"`
	DaContract          string `mapstructure:"DaContract"`
	DaKey               string `mapstructure:"DaKey"`
	DaNamespaceId       uint32 `mapstructure:"DaNamespaceId"`
}
