package zg

import (
	"context"
	"errors"
	"time"

	pb "github.com/0glabs/0g-da-client/api/grpc/disperser"
	"github.com/0xPolygonHermez/zkevm-node/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const retryDuration = time.Duration(3 * time.Second)

type ZgConfig struct {
	Enable      bool   `mapstructure:"Enable"`
	Address     string `mapstructure:"Address"`
	MaxBlobSize int    `mapstructure:"MaxBlobSize"`
}

type ZgDA struct {
	Client pb.DisperserClient
	Cfg    ZgConfig
}

type BlobRequestParams struct {
	DataRoot []byte
	Epoch    uint64
	QuorumId uint64
}

func NewZgDA(cfg ZgConfig) (*ZgDA, error) {
	conn, err := grpc.Dial(cfg.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, err
	}

	c := pb.NewDisperserClient(conn)

	return &ZgDA{
		Client: c,
		Cfg:    cfg,
	}, nil
}

func (d *ZgDA) Init() error {
	return nil
}

// GetSequence gets backend data one hash at a time. This should be optimized on the DAC side to get them all at once.
func (d *ZgDA) GetSequence(ctx context.Context, hashes []common.Hash, dataAvailabilityMessage []byte) ([][]byte, error) {
	var blobRequestParams [][]BlobRequestParams
	err := rlp.DecodeBytes(dataAvailabilityMessage, &blobRequestParams)
	if err != nil {
		return nil, err
	}

	var batchData [][]byte
	for _, msg := range blobRequestParams {
		data, err := d.GetBatchL2Data(ctx, msg)

		if err != nil {
			return nil, err
		}
		batchData = append(batchData, data)
	}
	return batchData, nil
}

func (d *ZgDA) GetBatchL2Data(ctx context.Context, blobRequests []BlobRequestParams) ([]byte, error) {
	var blobData = make([]byte, 0)

	for _, requestParam := range blobRequests {
		log.Info("Requesting data from zgDA", "param", requestParam)

		retrieveBlobReply, err := d.Client.RetrieveBlob(ctx, &pb.RetrieveBlobRequest{
			StorageRoot: requestParam.DataRoot,
			Epoch:       requestParam.Epoch,
			QuorumId:    requestParam.QuorumId,
		})

		if err != nil {
			return nil, err
		}

		blobData = append(blobData, retrieveBlobReply.GetData()...)
	}

	return blobData, nil
}

func (s *ZgDA) PostSequence(ctx context.Context, batchesData [][]byte) ([]byte, error) {
	message := make([][]BlobRequestParams, 0)

	for _, seq := range batchesData {
		totalBlobSize := len(seq)

		requestParams := make([]BlobRequestParams, 0)
		if totalBlobSize > 0 {
			log.Infof("BatchL2Data %s, len %d", seq, totalBlobSize)

			for idx := 0; idx < totalBlobSize; idx += s.Cfg.MaxBlobSize {
				var endIdx int
				if totalBlobSize <= idx+s.Cfg.MaxBlobSize {
					endIdx = totalBlobSize
				} else {
					endIdx = idx + s.Cfg.MaxBlobSize
				}

				blob := pb.DisperseBlobRequest{
					Data: seq[idx:endIdx],
				}

				log.Infof("Disperse blob range %d %d", idx, endIdx)
				blobReply, err := s.Client.DisperseBlob(ctx, &blob)
				if err != nil {
					log.Warn("Disperse blob error", "err", err)
					return nil, err
				}

				requestId := blobReply.GetRequestId()
				log.Infof("Disperse request id %s", requestId)
				for {
					statusReply, err := s.Client.GetBlobStatus(ctx, &pb.BlobStatusRequest{RequestId: requestId})

					if err != nil {
						log.Warn("Get blob status error", "err", err)
						return nil, err
					}
					log.Infof("status reply %s", statusReply)

					if statusReply.GetStatus() == pb.BlobStatus_CONFIRMED || statusReply.GetStatus() == pb.BlobStatus_FINALIZED {
						blobInfo := statusReply.GetInfo()

						dataRoot := blobInfo.BlobHeader.GetStorageRoot()
						epoch := blobInfo.BlobHeader.GetEpoch()
						quorumId := blobInfo.BlobHeader.GetQuorumId()

						requestParams = append(requestParams, BlobRequestParams{
							DataRoot: dataRoot,
							Epoch:    epoch,
							QuorumId: quorumId,
						})

						break
					}

					if statusReply.GetStatus() == pb.BlobStatus_FAILED {
						return nil, errors.New("store blob failed")
					}

					time.Sleep(retryDuration)
				}
			}

			message = append(message, requestParams)
		}
	}

	rlpEncode, err := rlp.EncodeToBytes(&message)
	if err != nil {
		return nil, err
	}

	return rlpEncode, nil
}
