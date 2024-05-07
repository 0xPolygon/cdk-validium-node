package main

import (
	"context"
	"fmt"

	etrog "github.com/0xPolygonHermez/zkevm-node/etherman/smartcontracts/etrogpolygonzkevm"
	nearda "github.com/0xPolygonHermez/zkevm-node/etherman/smartcontracts/neardataavailability"

	"math/big"
	"strings"
	"time"

	"github.com/0xPolygonHermez/zkevm-node/log"
	"github.com/0xPolygonHermez/zkevm-node/test/operations"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const ()

var (
	executedTransctionsCount uint64 = 0
)

const (
	txTimeout = 60 * time.Second
	// if you want to test using goerli network
	// replace this by your goerli infura url
	//networkURL = "http://localhost:8123"
	networkURL = "http://localhost:8545"
	//pk = "0xdfd01798f92667dbf91df722434e8fbe96af0211d4d1b82bbbbc8f1def7a814f"
	pk        = operations.DefaultSequencerPrivateKey
	zkEvmAddr = operations.DefaultL1ZkEVMSmartContract
)

type Deployments struct {
	NearDA     *nearda.Neardataavailability
	NearDAAddr common.Address
}

func main() {
	ctx := context.Background()
	log.Infof("connecting to %v", networkURL)
	client, err := ethclient.Dial(networkURL)
	ChkErr(err)

	log.Infof("connected")
	chainID, err := client.ChainID(ctx)
	ChkErr(err)

	log.Infof("chainID: %v", chainID)
	auth := GetAuth(ctx, client, pk)
	fmt.Println()
	deployments := DeployContractsAndSetDap(client, auth)
	log.Infof("deployed: %v", deployments.NearDAAddr.Hex())
}

func DeployContractsAndSetDap(client *ethclient.Client, auth *bind.TransactOpts) Deployments {
	ctx := context.Background()
	fmt.Println()
	addr, tx, neardataavailability, err := nearda.DeployNeardataavailability(auth, client)
	fmt.Println()
	err = WaitForTransactionAndIncrementNonce(client, auth, err, ctx, tx)
	log.Debugf("Deploy tx: %v", tx.Hash().Hex())
	log.Debugf("Addr: %v", addr.Hex())
	fmt.Println()
	ChkErr(err)
	// Set DAP
	log.Debugf("Setting DAP on %v to %v", zkEvmAddr, addr)
	etrog, err := etrog.NewEtrogpolygonzkevm(common.HexToAddress(zkEvmAddr), client)
	etrog.SetDataAvailabilityProtocol(auth, addr)
	fmt.Println()
	ChkErr(err)

	return Deployments{
		NearDA:     neardataavailability,
		NearDAAddr: addr,
	}
}

func WaitForTransactionAndIncrementNonce(l2Client *ethclient.Client, auth *bind.TransactOpts, err error, ctx context.Context, tx *types.Transaction) error {
	ChkErr(err)
	err = operations.WaitTxToBeMined(ctx, l2Client, tx, txTimeout)
	ChkErr(err)
	executedTransctionsCount++
	auth.Nonce = nil
	auth.Value = nil

	return err
}

func GetAuth(ctx context.Context, client *ethclient.Client, pkHex string) *bind.TransactOpts {
	chainID, err := client.ChainID(ctx)
	ChkErr(err)
	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(pkHex, "0x"))
	ChkErr(err)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	ChkErr(err)
	senderNonce, err := client.PendingNonceAt(ctx, auth.From)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(senderNonce))
	return auth
}

func ChkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetExecutedTransactionsCount() uint64 {
	return executedTransctionsCount
}

func getDeadline() *big.Int {
	const deadLinelimit = 5 * time.Minute
	return big.NewInt(time.Now().UTC().Add(deadLinelimit).Unix())
}
