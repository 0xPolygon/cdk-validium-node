package types

import "github.com/0xPolygon/cdk-validium-node/aggregator/prover"

// FinalProofInputs struct
type FinalProofInputs struct {
	FinalProof       *prover.FinalProof
	NewLocalExitRoot []byte
	NewStateRoot     []byte
}
