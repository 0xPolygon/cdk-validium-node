// Package near provides a backend implementation for interacting with the NEAR DA.
//
// The NEAR DA backend allows posting and retrieving sequences of batches using the Near blockchain.
// It utilizes the NEAR DA Sidecar service for submitting and retrieving data blobs.
//
// Usage:
//
//	backend, err := near.New("http://localhost:5888", &sidecar.ConfigureClientRequest{...})
//	if err != nil {
//	    // Handle error
//	}
//
//	// Post a sequence of batches
//	batchesData := [][]byte{...}
//	transactionID, err := backend.PostSequence(context.Background(), batchesData)
//	if err != nil {
//	    // Handle error
//	}
//
//	// Get a sequence of batches
//	batchHashes := []common.Hash{...}
//	dataAvailabilityMessage := []byte{...}
//	batchData, err := backend.GetSequence(context.Background(), batchHashes, dataAvailabilityMessage)
//	if err != nil {
//	    // Handle error
//	}
//
// The package also provides logging functionality using the "github.com/0xPolygonHermez/zkevm-node/log" package.
//
// Dependencies:
//
//	github.com/near/rollup-data-availability/gopkg/sidecar
package near

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"

	"github.com/0xPolygonHermez/zkevm-node/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/near/rollup-data-availability/gopkg/sidecar"
)

// NearProtocolBackend represents a backend for interacting with NEAR DA.
type NearProtocolBackend struct {
	host   string
	Client *sidecar.Client
	Config *sidecar.ConfigureClientRequest
}

// New creates a new instance of the NearProtocolBackend.
// It takes the host and configuration as parameters and returns a pointer to the backend.
func New(host string, config *sidecar.ConfigureClientRequest) (*NearProtocolBackend, error) {
	if config == nil {
		log.Debug("no near config, assuming sidecar setup outside of CDK")
	}
	return &NearProtocolBackend{
		host:   host,
		Config: config,
	}, nil
}

// Init loads the DAC to be cached when needed
func (d *NearProtocolBackend) Init() error {
	client, err := sidecar.NewClient(d.host, d.Config)
	if err != nil {
		return fmt.Errorf("error connecting to NEAR: %s", err)
	}

	d.Client = client
	return nil
}

func ParseConfig(accountId, secretKey, contractId, network string) sidecar.ConfigureClientRequest {
	sidecarConfig := sidecar.ConfigureClientRequest{
		AccountID:  accountId,
		SecretKey:  secretKey,
		ContractID: contractId,
		Network:    sidecar.Network(network),
	}
	return sidecarConfig
}

type Sequence struct {
	Batches [][]byte
}

func (s *Sequence) encode() (*sidecar.Blob, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(s)
	if err != nil {
		return nil, fmt.Errorf("error encoding sequence for NEAR: %s", err)
	}
	blob := sidecar.Blob{
		Data: buf.Bytes(),
	}

	return &blob, err
}

func decode(b *sidecar.Blob) (*Sequence, error) {
	var s Sequence

	buf := bytes.NewReader(b.Data)
	codec := gob.NewDecoder(buf)
	err := codec.Decode(&s)
	if len(s.Batches) == 0 {
		return nil, fmt.Errorf("sequence is empty")
	}
	return &s, err
}

func hexEncode(b [][]byte) string {
	var bytes []byte
	for _, batch := range b {
		bytes = append(bytes, batch...)
	}
	return common.Bytes2Hex(bytes)
}

func hexEncode2(b []byte) string {
	return common.Bytes2Hex(b)
}

const MaxBatchSize = 4 * 1024 * 1024 // Max batch size is 4MB

// PostSequence posts a sequence of batches to the Near blockchain.
// It takes a context and a slice of byte slices representing the batches data.
// It returns the transaction ID of the submitted sequence and any error encountered.
func (n *NearProtocolBackend) PostSequence(ctx context.Context, batches [][]byte) ([]byte, error) {
	// count the size of all batches, overflowing batches into multiple sequences
	sequences := ChunkedSequences(batches)
	log.Debugf("Submitting %d sequences", len(sequences))

	var blobRefs []byte
	// For 	each sequence weno
	for _, seq := range sequences {
		blob, err := seq.encode()
		if err != nil {
			return nil, err
		}

		blobRef, err := n.Client.SubmitBlob(*blob)
		if err != nil {
			return nil, fmt.Errorf("error submitting data to NEAR: %s", err)
		}
		log.Debugf("Blob ref: %s", hexEncode2(blobRef.Deref()))
		blobRefs = append(blobRefs, blobRef.Deref()...)
	}
	return blobRefs, nil
}

func ChunkedSequences(batches [][]byte) []Sequence {
	// count the size of all batches, overflowing batches into multiple sequences
	var sequences []Sequence
	sequences = append(sequences, Sequence{})

	size := 0
	seqIndex := 0
	for _, batch := range batches {
		size += len(batch)

		if size > MaxBatchSize {
			seqIndex++
			size = len(batch)

			var seq Sequence
			seq.Batches = append(seq.Batches, batch)
			sequences = append(sequences, seq)
		} else {
			sequences[seqIndex].Batches = append(sequences[seqIndex].Batches, batch)
		}
	}
	return sequences
}

// GetSequence retrieves a sequence of batches from the Near blockchain.
// It takes a context, a slice of batch hashes, and a data availability message.
// It returns a slice of byte slices representing the retrieved batches data and any error encountered.
func (n *NearProtocolBackend) GetSequence(ctx context.Context, batchHashes []common.Hash, dataAvailabilityMessage []byte) ([][]byte, error) {
	var batchData [][]byte

	log.Debugf("Retrieving %d batches from dataAvailabilityMessage %s: %d", len(batchHashes), hexEncode2(dataAvailabilityMessage), len(dataAvailabilityMessage))

	// Chunk the da message into references
	for _, blobRef := range Chunks(dataAvailabilityMessage, sidecar.EncodedBlobRefSize) {
		blobRef, err := sidecar.NewBlobRef(blobRef)
		if err != nil {
			return nil, fmt.Errorf("error reading blob reference: %s", err)
		}

		blob, err := n.Client.GetBlob(*blobRef)
		if err != nil {
			return nil, fmt.Errorf("error fetching blob %s:  %s", hexEncode2(blobRef.Deref()), err)
		}

		seq, err := decode(blob)
		if err != nil {
			return nil, fmt.Errorf("error decoding sequence from NEAR: %s", err)
		}

		batchData = append(batchData, seq.Batches...)
	}

	return batchData, nil
}

func Chunks(message []byte, chunkSize int) [][]byte {
	var chunks [][]byte
	for {
		if len(message) == 0 {
			break
		}

		// necessary check to avoid slicing beyond
		// slice capacity
		if len(message) < chunkSize {
			chunkSize = len(message)
		}

		chunks = append(chunks, message[0:chunkSize])
		message = message[chunkSize:]
	}

	return chunks
}
