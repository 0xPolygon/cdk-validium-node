// Tests

package near

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/near/rollup-data-availability/gopkg/sidecar"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/sha3"
)

// stub 32 byte tx id
const tx string = "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"

var random = GenerateRandomBatchesData(1, 10)
var sequence = Sequence{Batches: random}

func mocks(t *testing.T, sequenceMap map[string]*Sequence, refMap map[common.Hash]*string) (*NearProtocolBackend, *httptest.Server) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/health":
			w.WriteHeader(http.StatusOK)
		case "/blob":
			if r.Method == http.MethodPost {
				var blob sidecar.Blob
				err := json.NewDecoder(r.Body).Decode(&blob)
				require.NoError(t, err)

				txId := tx

				maybeTx := refMap[common.BytesToHash(hash(blob.Data))]
				if maybeTx != nil {
					println("found tx: ", *maybeTx)
					txId = *maybeTx
				}

				blobRef, err := sidecar.NewBlobRef(common.FromHex(txId))
				require.NoError(t, err)

				jsonData, err := blobRef.MarshalJSON()
				require.NoError(t, err)

				w.Write(jsonData)
			} else if r.Method == http.MethodGet {
				txId := r.URL.Query().Get("transaction_id")

				var seq = sequence

				maybeSequence := sequenceMap[txId]
				if maybeSequence != nil {
					seq = *maybeSequence
				}

				blob, err := seq.encode()
				require.NoError(t, err)

				jsonData, err := blob.MarshalJSON()
				require.NoError(t, err)
				w.Write(jsonData)
			}
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	backend, err := New(mockServer.URL, nil)
	backend.Init()
	require.NoError(t, err)

	return backend, mockServer
}

func TestPostSequence(t *testing.T) {
	backend, mockServer := mocks(t, nil, nil)
	defer backend.Client.Close()
	defer mockServer.Close()

	// Test PostSequence
	batchesData := GenerateRandomBatchesData(3, 10)
	println("test batches data: ", hexEncode(batchesData))
	transactionID, err := backend.PostSequence(context.Background(), batchesData)
	require.NoError(t, err)
	expectedTransactionID, err := hex.DecodeString(tx)
	require.NoError(t, err)

	assert.Equal(t, expectedTransactionID, transactionID)

	// TODO: test overloadpostssequence
}

func TestGetSequence(t *testing.T) {

	backend, mockServer := mocks(t, nil, nil)
	defer backend.Client.Close()
	defer mockServer.Close()

	batchHashes := []common.Hash{}
	for _, v := range sequence.Batches {
		batchHashes = append(batchHashes, common.BytesToHash(v))
	}

	dataAvailabilityMessage := common.Hex2Bytes(tx)
	batchData, err := backend.GetSequence(context.Background(), batchHashes, dataAvailabilityMessage)
	require.NoError(t, err)

	assert.Equal(t, 1, len(batchData))
	assert.Equal(t, sequence.Batches, batchData)

}

func TestLargeE2e(t *testing.T) {
	numBatches := 10

	batches := GenerateChunks(MaxBatchSize * numBatches)
	sequences := ChunkedSequences(batches)
	println("Expected sequences: ", len(sequences))

	seqMap := map[string]*Sequence{}
	refMap := map[common.Hash]*string{}

	var index byte = 0
	for _, v := range sequences {
		// Populate blob ref
		blobRef := hash([]byte{index})
		txId := hex.EncodeToString(blobRef[:])

		// Encode the blob
		encoded, err := v.encode()
		require.NoError(t, err)

		seqMap[txId] = &v
		refMap[common.BytesToHash(hash(encoded.Data))] = &txId

		index++
	}
	backend, mockServer := mocks(t, seqMap, refMap)
	defer backend.Client.Close()
	defer mockServer.Close()

	// Test PostSequence
	transactionIds, err := backend.PostSequence(context.Background(), batches)
	require.NoError(t, err)

	// Assert that all the transaction ids are the ones we know about
	for _, v := range Chunks(transactionIds, sidecar.EncodedBlobRefSize) {
		seq := seqMap[common.Bytes2Hex(v)]
		require.NotNil(t, seq)
	}

	batchHashes := []common.Hash{}
	for _, v := range sequences {
		var appended []byte
		for _, b := range v.Batches {
			appended = append(appended, b...)
		}
		batchHashes = append(batchHashes, common.BytesToHash(hash(appended)))
	}

	batchData, err := backend.GetSequence(context.Background(), batchHashes, transactionIds)
	require.NoError(t, err)
	assert.Equal(t, len(batches), len(batchData))
}

func TestChunkedSequences(t *testing.T) {
	batches := GenerateRandomBatchesData(4, MaxBatchSize+1)
	sequences := ChunkedSequences(batches)

	// should be 5 sequences: 4 normals + overflow
	assert.Equal(t, 5, len(sequences))
}

func hash(buf []byte) []byte {
	sha := sha3.NewLegacyKeccak256()
	sha.Write(buf[:])
	sha256 := sha.Sum(nil)
	return sha256
}

// GenerateRandomBlobData generates random blob data of the specified size.
func GenerateRandomBlobData(size int) []byte {
	blob := make([]byte, size)
	_, err := rand.Read(blob)
	if err != nil {
		panic(err)
	}
	return blob
}

// GenerateRandomBatchesData generates a slice of random batch data.
func GenerateRandomBatchesData(numBatches, batchSize int) [][]byte {
	println("generate random batches data: ", numBatches, batchSize)
	batchesData := make([][]byte, numBatches)
	for i := 0; i < numBatches; i++ {
		batchesData[i] = GenerateRandomBlobData(batchSize)
	}
	return batchesData
}

func GenerateChunks(size int) [][]byte {
	data := GenerateRandomBlobData(size)
	return Chunks(data, MaxBatchSize)
}
