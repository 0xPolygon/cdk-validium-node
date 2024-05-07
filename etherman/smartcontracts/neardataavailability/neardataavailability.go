// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package neardataavailability

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// VerifiedBatch is an auto generated low-level Go binding around an user-defined struct.
type VerifiedBatch struct {
	Id           [32]byte
	VerifyTxHash [32]byte
	SubmitTxId   [32]byte
}

// NeardataavailabilityMetaData contains all meta data concerning the Neardataavailability contract.
var NeardataavailabilityMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"_NOTIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_STORED_BATCH_AMT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_SUBMITTED_BATCH_AMT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"batchInfo\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"verifyTxHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"submitTxId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"completeOwnershipHandover\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getProcotolName\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"grantRoles\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"hasAllRoles\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasAnyRole\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"notifyAvailable\",\"inputs\":[{\"name\":\"verifiedBatch\",\"type\":\"tuple\",\"internalType\":\"structVerifiedBatch\",\"components\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"verifyTxHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"submitTxId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"notifySubmitted\",\"inputs\":[{\"name\":\"batches\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownershipHandoverExpiresAt\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"renounceRoles\",\"inputs\":[{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"requestOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"revokeRoles\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"rolesOf\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"submittedBatches\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"verifyMessage\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"dataAvailabilityBatch\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IsAvailable\",\"inputs\":[{\"name\":\"bucketIdx\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"batch\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structVerifiedBatch\",\"components\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"verifyTxHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"submitTxId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverCanceled\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverRequested\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"oldOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RolesUpdated\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Submitted\",\"inputs\":[{\"name\":\"bucketIdx\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"submitTxId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewOwnerIsZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoHandoverRequest\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
	Bin: "0x6080604052348015600f57600080fd5b506016601a565b60ca565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161560695760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b039081161460c75780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b610f15806100d96000396000f3fe60806040526004361061018b5760003560e01c8063715018a6116100d6578063c4d66de81161007f578063f04e283e11610059578063f04e283e14610461578063f2fde38b14610474578063fee81cf41461048757600080fd5b8063c4d66de8146103d5578063d0262817146103f5578063e4f171201461041557600080fd5b8063adc31e01116100b0578063adc31e011461038a578063b62aa7f91461039f578063ba0d0ebc146103bf57600080fd5b8063715018a6146102f3578063815bda47146102fb5780638da5cb5b1461033657600080fd5b80633b51be4b11610138578063514e62fc11610112578063514e62fc1461029f57806354d1f13d146102d657806355731347146102de57600080fd5b80633b51be4b146102565780634a4ee7b114610276578063510586421461028957600080fd5b8063256929621161016957806325692962146101ed57806329afa6a3146101f55780632de948071461021557600080fd5b8063183a4f6e146101905780631c10893f146101a55780631cd64df4146101b8575b600080fd5b6101a361019e366004610baa565b6104ba565b005b6101a36101b3366004610bec565b6104c7565b3480156101c457600080fd5b506101d86101d3366004610bec565b6104dd565b60405190151581526020015b60405180910390f35b6101a36104fc565b34801561020157600080fd5b506101a3610210366004610c16565b61054c565b34801561022157600080fd5b50610248610230366004610c99565b638b78c6d8600c908152600091909152602090205490565b6040519081526020016101e4565b34801561026257600080fd5b506101a3610271366004610d04565b6105ba565b6101a3610284366004610bec565b610613565b34801561029557600080fd5b5061024861080081565b3480156102ab57600080fd5b506101d86102ba366004610bec565b638b78c6d8600c90815260009290925260209091205416151590565b6101a3610625565b3480156102ea57600080fd5b50610248608081565b6101a3610661565b34801561030757600080fd5b5061031b610316366004610baa565b610675565b604080519384526020840192909252908201526060016101e4565b34801561034257600080fd5b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffff748739275460405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101e4565b34801561039657600080fd5b50610248602081565b3480156103ab57600080fd5b506101a36103ba366004610d50565b61069c565b3480156103cb57600080fd5b5061024861040081565b3480156103e157600080fd5b506101a36103f0366004610c99565b610755565b34801561040157600080fd5b50610248610410366004610baa565b6108d4565b34801561042157600080fd5b50604080518082018252600c81527f4e65617250726f746f636f6c0000000000000000000000000000000000000000602082015290516101e49190610d92565b6101a361046f366004610c99565b6108eb565b6101a3610482366004610c99565b610928565b34801561049357600080fd5b506102486104a2366004610c99565b63389a75e1600c908152600091909152602090205490565b6104c4338261094f565b50565b6104cf61095b565b6104d98282610991565b5050565b638b78c6d8600c90815260008390526020902054811681145b92915050565b60006202a30067ffffffffffffffff164201905063389a75e1600c5233600052806020600c2055337fdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d600080a250565b6108006105588161099d565b6000610563836109eb565b604080518281528551602080830191909152860151818301529085015160608201529091507f8985609fb3668d8b8e01db2e596eab8479f1e2447af72dab7baefdc1fd99d1bc9060800160405180910390a1505050565b6000806105c983850185610baa565b905060005b602081101561060b57600081602081106105ea576105ea610dff565b6003020192508183600001540361060357505050505050565b6001016105ce565b505050505050565b61061b61095b565b6104d9828261094f565b63389a75e1600c523360005260006020600c2055337ffa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92600080a2565b61066961095b565b6106736000610a4b565b565b6000816020811061068557600080fd5b600302018054600182015460029092015490925083565b6104006106a88161099d565b602060006106b68285610e8c565b905060005b8181101561060b57600086866106d18685610ea0565b90866106de866001610eb7565b6106e89190610ea0565b926106f593929190610eca565b8101906107029190610baa565b9050600061070f82610ab1565b60408051828152602081018590529192507f465ddfeee4b3d0fb72ba16ca1e8d4ecf86373b44e51658af98043162de372b4c910160405180910390a150506001016106bb565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff166000811580156107a05750825b905060008267ffffffffffffffff1660011480156107bd5750303b155b9050811580156107cb575080155b15610802576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117855583156108635784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b61086c86610aed565b831561060b5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a1505050505050565b606181608081106108e457600080fd5b0154905081565b6108f361095b565b63389a75e1600c52806000526020600c20805442111561091b57636f5e88186000526004601cfd5b600090556104c481610a4b565b61093061095b565b8060601b61094657637448fbae6000526004601cfd5b6104c481610a4b565b6104d982826000610b51565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffff74873927543314610673576382b429006000526004601cfd5b6104d982826001610b51565b638b78c6d8600c5233600052806020600c2054166104c4577fffffffffffffffffffffffffffffffffffffffffffffffffffffffff748739275433146104c4576382b429006000526004601cfd5b60605460009082828260208110610a0457610a04610dff565b600302016000820151816000015560208201518160010155604082015181600201559050506020816001610a389190610eb7565b610a429190610ef4565b60605592915050565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffff74873927805473ffffffffffffffffffffffffffffffffffffffff9092169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a355565b60e1546000908260618260808110610acb57610acb610dff565b01556080610ada826001610eb7565b610ae49190610ef4565b60e15592915050565b73ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff748739278190558060007f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08180a350565b638b78c6d8600c52826000526020600c20805483811783610b73575080841681185b80835580600c5160601c7f715ad5ce61fc9595c7b415289d59cf203f23a94fa06f04af7e489a0a76e1fe26600080a3505050505050565b600060208284031215610bbc57600080fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610be757600080fd5b919050565b60008060408385031215610bff57600080fd5b610c0883610bc3565b946020939093013593505050565b600060608284031215610c2857600080fd5b6040516060810181811067ffffffffffffffff82111715610c72577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b80604052508235815260208301356020820152604083013560408201528091505092915050565b600060208284031215610cab57600080fd5b610cb482610bc3565b9392505050565b60008083601f840112610ccd57600080fd5b50813567ffffffffffffffff811115610ce557600080fd5b602083019150836020828501011115610cfd57600080fd5b9250929050565b600080600060408486031215610d1957600080fd5b83359250602084013567ffffffffffffffff811115610d3757600080fd5b610d4386828701610cbb565b9497909650939450505050565b60008060208385031215610d6357600080fd5b823567ffffffffffffffff811115610d7a57600080fd5b610d8685828601610cbb565b90969095509350505050565b60006020808352835180602085015260005b81811015610dc057858101830151858201604001528201610da4565b5060006040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082610e9b57610e9b610e2e565b500490565b80820281158282048414176104f6576104f6610e5d565b808201808211156104f6576104f6610e5d565b60008085851115610eda57600080fd5b83861115610ee757600080fd5b5050820193919092039150565b600082610f0357610f03610e2e565b50069056fea164736f6c6343000819000a",
}

// NeardataavailabilityABI is the input ABI used to generate the binding from.
// Deprecated: Use NeardataavailabilityMetaData.ABI instead.
var NeardataavailabilityABI = NeardataavailabilityMetaData.ABI

// NeardataavailabilityBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NeardataavailabilityMetaData.Bin instead.
var NeardataavailabilityBin = NeardataavailabilityMetaData.Bin

// DeployNeardataavailability deploys a new Ethereum contract, binding an instance of Neardataavailability to it.
func DeployNeardataavailability(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Neardataavailability, error) {
	parsed, err := NeardataavailabilityMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NeardataavailabilityBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Neardataavailability{NeardataavailabilityCaller: NeardataavailabilityCaller{contract: contract}, NeardataavailabilityTransactor: NeardataavailabilityTransactor{contract: contract}, NeardataavailabilityFilterer: NeardataavailabilityFilterer{contract: contract}}, nil
}

// Neardataavailability is an auto generated Go binding around an Ethereum contract.
type Neardataavailability struct {
	NeardataavailabilityCaller     // Read-only binding to the contract
	NeardataavailabilityTransactor // Write-only binding to the contract
	NeardataavailabilityFilterer   // Log filterer for contract events
}

// NeardataavailabilityCaller is an auto generated read-only Go binding around an Ethereum contract.
type NeardataavailabilityCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NeardataavailabilityTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NeardataavailabilityTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NeardataavailabilityFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NeardataavailabilityFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NeardataavailabilitySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NeardataavailabilitySession struct {
	Contract     *Neardataavailability // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// NeardataavailabilityCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NeardataavailabilityCallerSession struct {
	Contract *NeardataavailabilityCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// NeardataavailabilityTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NeardataavailabilityTransactorSession struct {
	Contract     *NeardataavailabilityTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// NeardataavailabilityRaw is an auto generated low-level Go binding around an Ethereum contract.
type NeardataavailabilityRaw struct {
	Contract *Neardataavailability // Generic contract binding to access the raw methods on
}

// NeardataavailabilityCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NeardataavailabilityCallerRaw struct {
	Contract *NeardataavailabilityCaller // Generic read-only contract binding to access the raw methods on
}

// NeardataavailabilityTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NeardataavailabilityTransactorRaw struct {
	Contract *NeardataavailabilityTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNeardataavailability creates a new instance of Neardataavailability, bound to a specific deployed contract.
func NewNeardataavailability(address common.Address, backend bind.ContractBackend) (*Neardataavailability, error) {
	contract, err := bindNeardataavailability(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Neardataavailability{NeardataavailabilityCaller: NeardataavailabilityCaller{contract: contract}, NeardataavailabilityTransactor: NeardataavailabilityTransactor{contract: contract}, NeardataavailabilityFilterer: NeardataavailabilityFilterer{contract: contract}}, nil
}

// NewNeardataavailabilityCaller creates a new read-only instance of Neardataavailability, bound to a specific deployed contract.
func NewNeardataavailabilityCaller(address common.Address, caller bind.ContractCaller) (*NeardataavailabilityCaller, error) {
	contract, err := bindNeardataavailability(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NeardataavailabilityCaller{contract: contract}, nil
}

// NewNeardataavailabilityTransactor creates a new write-only instance of Neardataavailability, bound to a specific deployed contract.
func NewNeardataavailabilityTransactor(address common.Address, transactor bind.ContractTransactor) (*NeardataavailabilityTransactor, error) {
	contract, err := bindNeardataavailability(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NeardataavailabilityTransactor{contract: contract}, nil
}

// NewNeardataavailabilityFilterer creates a new log filterer instance of Neardataavailability, bound to a specific deployed contract.
func NewNeardataavailabilityFilterer(address common.Address, filterer bind.ContractFilterer) (*NeardataavailabilityFilterer, error) {
	contract, err := bindNeardataavailability(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NeardataavailabilityFilterer{contract: contract}, nil
}

// bindNeardataavailability binds a generic wrapper to an already deployed contract.
func bindNeardataavailability(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NeardataavailabilityABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Neardataavailability *NeardataavailabilityRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Neardataavailability.Contract.NeardataavailabilityCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Neardataavailability *NeardataavailabilityRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Neardataavailability.Contract.NeardataavailabilityTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Neardataavailability *NeardataavailabilityRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Neardataavailability.Contract.NeardataavailabilityTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Neardataavailability *NeardataavailabilityCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Neardataavailability.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Neardataavailability *NeardataavailabilityTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Neardataavailability.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Neardataavailability *NeardataavailabilityTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Neardataavailability.Contract.contract.Transact(opts, method, params...)
}

// NOTIFIER is a free data retrieval call binding the contract method 0xba0d0ebc.
//
// Solidity: function _NOTIFIER() view returns(uint256)
func (_Neardataavailability *NeardataavailabilityCaller) NOTIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Neardataavailability.contract.Call(opts, &out, "_NOTIFIER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NOTIFIER is a free data retrieval call binding the contract method 0xba0d0ebc.
//
// Solidity: function _NOTIFIER() view returns(uint256)
func (_Neardataavailability *NeardataavailabilitySession) NOTIFIER() (*big.Int, error) {
	return _Neardataavailability.Contract.NOTIFIER(&_Neardataavailability.CallOpts)
}

// NOTIFIER is a free data retrieval call binding the contract method 0xba0d0ebc.
//
// Solidity: function _NOTIFIER() view returns(uint256)
func (_Neardataavailability *NeardataavailabilityCallerSession) NOTIFIER() (*big.Int, error) {
	return _Neardataavailability.Contract.NOTIFIER(&_Neardataavailability.CallOpts)
}

// STOREDBATCHAMT is a free data retrieval call binding the contract method 0xadc31e01.
//
// Solidity: function _STORED_BATCH_AMT() view returns(uint256)
func (_Neardataavailability *NeardataavailabilityCaller) STOREDBATCHAMT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Neardataavailability.contract.Call(opts, &out, "_STORED_BATCH_AMT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// STOREDBATCHAMT is a free data retrieval call binding the contract method 0xadc31e01.
//
// Solidity: function _STORED_BATCH_AMT() view returns(uint256)
func (_Neardataavailability *NeardataavailabilitySession) STOREDBATCHAMT() (*big.Int, error) {
	return _Neardataavailability.Contract.STOREDBATCHAMT(&_Neardataavailability.CallOpts)
}

// STOREDBATCHAMT is a free data retrieval call binding the contract method 0xadc31e01.
//
// Solidity: function _STORED_BATCH_AMT() view returns(uint256)
func (_Neardataavailability *NeardataavailabilityCallerSession) STOREDBATCHAMT() (*big.Int, error) {
	return _Neardataavailability.Contract.STOREDBATCHAMT(&_Neardataavailability.CallOpts)
}

// SUBMITTEDBATCHAMT is a free data retrieval call binding the contract method 0x55731347.
//
// Solidity: function _SUBMITTED_BATCH_AMT() view returns(uint256)
func (_Neardataavailability *NeardataavailabilityCaller) SUBMITTEDBATCHAMT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Neardataavailability.contract.Call(opts, &out, "_SUBMITTED_BATCH_AMT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SUBMITTEDBATCHAMT is a free data retrieval call binding the contract method 0x55731347.
//
// Solidity: function _SUBMITTED_BATCH_AMT() view returns(uint256)
func (_Neardataavailability *NeardataavailabilitySession) SUBMITTEDBATCHAMT() (*big.Int, error) {
	return _Neardataavailability.Contract.SUBMITTEDBATCHAMT(&_Neardataavailability.CallOpts)
}

// SUBMITTEDBATCHAMT is a free data retrieval call binding the contract method 0x55731347.
//
// Solidity: function _SUBMITTED_BATCH_AMT() view returns(uint256)
func (_Neardataavailability *NeardataavailabilityCallerSession) SUBMITTEDBATCHAMT() (*big.Int, error) {
	return _Neardataavailability.Contract.SUBMITTEDBATCHAMT(&_Neardataavailability.CallOpts)
}

// VERIFIER is a free data retrieval call binding the contract method 0x51058642.
//
// Solidity: function _VERIFIER() view returns(uint256)
func (_Neardataavailability *NeardataavailabilityCaller) VERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Neardataavailability.contract.Call(opts, &out, "_VERIFIER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VERIFIER is a free data retrieval call binding the contract method 0x51058642.
//
// Solidity: function _VERIFIER() view returns(uint256)
func (_Neardataavailability *NeardataavailabilitySession) VERIFIER() (*big.Int, error) {
	return _Neardataavailability.Contract.VERIFIER(&_Neardataavailability.CallOpts)
}

// VERIFIER is a free data retrieval call binding the contract method 0x51058642.
//
// Solidity: function _VERIFIER() view returns(uint256)
func (_Neardataavailability *NeardataavailabilityCallerSession) VERIFIER() (*big.Int, error) {
	return _Neardataavailability.Contract.VERIFIER(&_Neardataavailability.CallOpts)
}

// BatchInfo is a free data retrieval call binding the contract method 0x815bda47.
//
// Solidity: function batchInfo(uint256 ) view returns(bytes32 id, bytes32 verifyTxHash, bytes32 submitTxId)
func (_Neardataavailability *NeardataavailabilityCaller) BatchInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id           [32]byte
	VerifyTxHash [32]byte
	SubmitTxId   [32]byte
}, error) {
	var out []interface{}
	err := _Neardataavailability.contract.Call(opts, &out, "batchInfo", arg0)

	outstruct := new(struct {
		Id           [32]byte
		VerifyTxHash [32]byte
		SubmitTxId   [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.VerifyTxHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.SubmitTxId = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// BatchInfo is a free data retrieval call binding the contract method 0x815bda47.
//
// Solidity: function batchInfo(uint256 ) view returns(bytes32 id, bytes32 verifyTxHash, bytes32 submitTxId)
func (_Neardataavailability *NeardataavailabilitySession) BatchInfo(arg0 *big.Int) (struct {
	Id           [32]byte
	VerifyTxHash [32]byte
	SubmitTxId   [32]byte
}, error) {
	return _Neardataavailability.Contract.BatchInfo(&_Neardataavailability.CallOpts, arg0)
}

// BatchInfo is a free data retrieval call binding the contract method 0x815bda47.
//
// Solidity: function batchInfo(uint256 ) view returns(bytes32 id, bytes32 verifyTxHash, bytes32 submitTxId)
func (_Neardataavailability *NeardataavailabilityCallerSession) BatchInfo(arg0 *big.Int) (struct {
	Id           [32]byte
	VerifyTxHash [32]byte
	SubmitTxId   [32]byte
}, error) {
	return _Neardataavailability.Contract.BatchInfo(&_Neardataavailability.CallOpts, arg0)
}

// GetProcotolName is a free data retrieval call binding the contract method 0xe4f17120.
//
// Solidity: function getProcotolName() pure returns(string)
func (_Neardataavailability *NeardataavailabilityCaller) GetProcotolName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Neardataavailability.contract.Call(opts, &out, "getProcotolName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetProcotolName is a free data retrieval call binding the contract method 0xe4f17120.
//
// Solidity: function getProcotolName() pure returns(string)
func (_Neardataavailability *NeardataavailabilitySession) GetProcotolName() (string, error) {
	return _Neardataavailability.Contract.GetProcotolName(&_Neardataavailability.CallOpts)
}

// GetProcotolName is a free data retrieval call binding the contract method 0xe4f17120.
//
// Solidity: function getProcotolName() pure returns(string)
func (_Neardataavailability *NeardataavailabilityCallerSession) GetProcotolName() (string, error) {
	return _Neardataavailability.Contract.GetProcotolName(&_Neardataavailability.CallOpts)
}

// HasAllRoles is a free data retrieval call binding the contract method 0x1cd64df4.
//
// Solidity: function hasAllRoles(address user, uint256 roles) view returns(bool)
func (_Neardataavailability *NeardataavailabilityCaller) HasAllRoles(opts *bind.CallOpts, user common.Address, roles *big.Int) (bool, error) {
	var out []interface{}
	err := _Neardataavailability.contract.Call(opts, &out, "hasAllRoles", user, roles)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasAllRoles is a free data retrieval call binding the contract method 0x1cd64df4.
//
// Solidity: function hasAllRoles(address user, uint256 roles) view returns(bool)
func (_Neardataavailability *NeardataavailabilitySession) HasAllRoles(user common.Address, roles *big.Int) (bool, error) {
	return _Neardataavailability.Contract.HasAllRoles(&_Neardataavailability.CallOpts, user, roles)
}

// HasAllRoles is a free data retrieval call binding the contract method 0x1cd64df4.
//
// Solidity: function hasAllRoles(address user, uint256 roles) view returns(bool)
func (_Neardataavailability *NeardataavailabilityCallerSession) HasAllRoles(user common.Address, roles *big.Int) (bool, error) {
	return _Neardataavailability.Contract.HasAllRoles(&_Neardataavailability.CallOpts, user, roles)
}

// HasAnyRole is a free data retrieval call binding the contract method 0x514e62fc.
//
// Solidity: function hasAnyRole(address user, uint256 roles) view returns(bool)
func (_Neardataavailability *NeardataavailabilityCaller) HasAnyRole(opts *bind.CallOpts, user common.Address, roles *big.Int) (bool, error) {
	var out []interface{}
	err := _Neardataavailability.contract.Call(opts, &out, "hasAnyRole", user, roles)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasAnyRole is a free data retrieval call binding the contract method 0x514e62fc.
//
// Solidity: function hasAnyRole(address user, uint256 roles) view returns(bool)
func (_Neardataavailability *NeardataavailabilitySession) HasAnyRole(user common.Address, roles *big.Int) (bool, error) {
	return _Neardataavailability.Contract.HasAnyRole(&_Neardataavailability.CallOpts, user, roles)
}

// HasAnyRole is a free data retrieval call binding the contract method 0x514e62fc.
//
// Solidity: function hasAnyRole(address user, uint256 roles) view returns(bool)
func (_Neardataavailability *NeardataavailabilityCallerSession) HasAnyRole(user common.Address, roles *big.Int) (bool, error) {
	return _Neardataavailability.Contract.HasAnyRole(&_Neardataavailability.CallOpts, user, roles)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_Neardataavailability *NeardataavailabilityCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Neardataavailability.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_Neardataavailability *NeardataavailabilitySession) Owner() (common.Address, error) {
	return _Neardataavailability.Contract.Owner(&_Neardataavailability.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_Neardataavailability *NeardataavailabilityCallerSession) Owner() (common.Address, error) {
	return _Neardataavailability.Contract.Owner(&_Neardataavailability.CallOpts)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_Neardataavailability *NeardataavailabilityCaller) OwnershipHandoverExpiresAt(opts *bind.CallOpts, pendingOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Neardataavailability.contract.Call(opts, &out, "ownershipHandoverExpiresAt", pendingOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_Neardataavailability *NeardataavailabilitySession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _Neardataavailability.Contract.OwnershipHandoverExpiresAt(&_Neardataavailability.CallOpts, pendingOwner)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_Neardataavailability *NeardataavailabilityCallerSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _Neardataavailability.Contract.OwnershipHandoverExpiresAt(&_Neardataavailability.CallOpts, pendingOwner)
}

// RolesOf is a free data retrieval call binding the contract method 0x2de94807.
//
// Solidity: function rolesOf(address user) view returns(uint256 roles)
func (_Neardataavailability *NeardataavailabilityCaller) RolesOf(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Neardataavailability.contract.Call(opts, &out, "rolesOf", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RolesOf is a free data retrieval call binding the contract method 0x2de94807.
//
// Solidity: function rolesOf(address user) view returns(uint256 roles)
func (_Neardataavailability *NeardataavailabilitySession) RolesOf(user common.Address) (*big.Int, error) {
	return _Neardataavailability.Contract.RolesOf(&_Neardataavailability.CallOpts, user)
}

// RolesOf is a free data retrieval call binding the contract method 0x2de94807.
//
// Solidity: function rolesOf(address user) view returns(uint256 roles)
func (_Neardataavailability *NeardataavailabilityCallerSession) RolesOf(user common.Address) (*big.Int, error) {
	return _Neardataavailability.Contract.RolesOf(&_Neardataavailability.CallOpts, user)
}

// SubmittedBatches is a free data retrieval call binding the contract method 0xd0262817.
//
// Solidity: function submittedBatches(uint256 ) view returns(bytes32)
func (_Neardataavailability *NeardataavailabilityCaller) SubmittedBatches(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Neardataavailability.contract.Call(opts, &out, "submittedBatches", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SubmittedBatches is a free data retrieval call binding the contract method 0xd0262817.
//
// Solidity: function submittedBatches(uint256 ) view returns(bytes32)
func (_Neardataavailability *NeardataavailabilitySession) SubmittedBatches(arg0 *big.Int) ([32]byte, error) {
	return _Neardataavailability.Contract.SubmittedBatches(&_Neardataavailability.CallOpts, arg0)
}

// SubmittedBatches is a free data retrieval call binding the contract method 0xd0262817.
//
// Solidity: function submittedBatches(uint256 ) view returns(bytes32)
func (_Neardataavailability *NeardataavailabilityCallerSession) SubmittedBatches(arg0 *big.Int) ([32]byte, error) {
	return _Neardataavailability.Contract.SubmittedBatches(&_Neardataavailability.CallOpts, arg0)
}

// VerifyMessage is a free data retrieval call binding the contract method 0x3b51be4b.
//
// Solidity: function verifyMessage(bytes32 , bytes dataAvailabilityBatch) view returns()
func (_Neardataavailability *NeardataavailabilityCaller) VerifyMessage(opts *bind.CallOpts, arg0 [32]byte, dataAvailabilityBatch []byte) error {
	var out []interface{}
	err := _Neardataavailability.contract.Call(opts, &out, "verifyMessage", arg0, dataAvailabilityBatch)

	if err != nil {
		return err
	}

	return err

}

// VerifyMessage is a free data retrieval call binding the contract method 0x3b51be4b.
//
// Solidity: function verifyMessage(bytes32 , bytes dataAvailabilityBatch) view returns()
func (_Neardataavailability *NeardataavailabilitySession) VerifyMessage(arg0 [32]byte, dataAvailabilityBatch []byte) error {
	return _Neardataavailability.Contract.VerifyMessage(&_Neardataavailability.CallOpts, arg0, dataAvailabilityBatch)
}

// VerifyMessage is a free data retrieval call binding the contract method 0x3b51be4b.
//
// Solidity: function verifyMessage(bytes32 , bytes dataAvailabilityBatch) view returns()
func (_Neardataavailability *NeardataavailabilityCallerSession) VerifyMessage(arg0 [32]byte, dataAvailabilityBatch []byte) error {
	return _Neardataavailability.Contract.VerifyMessage(&_Neardataavailability.CallOpts, arg0, dataAvailabilityBatch)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_Neardataavailability *NeardataavailabilityTransactor) CancelOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Neardataavailability.contract.Transact(opts, "cancelOwnershipHandover")
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_Neardataavailability *NeardataavailabilitySession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _Neardataavailability.Contract.CancelOwnershipHandover(&_Neardataavailability.TransactOpts)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_Neardataavailability *NeardataavailabilityTransactorSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _Neardataavailability.Contract.CancelOwnershipHandover(&_Neardataavailability.TransactOpts)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_Neardataavailability *NeardataavailabilityTransactor) CompleteOwnershipHandover(opts *bind.TransactOpts, pendingOwner common.Address) (*types.Transaction, error) {
	return _Neardataavailability.contract.Transact(opts, "completeOwnershipHandover", pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_Neardataavailability *NeardataavailabilitySession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _Neardataavailability.Contract.CompleteOwnershipHandover(&_Neardataavailability.TransactOpts, pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_Neardataavailability *NeardataavailabilityTransactorSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _Neardataavailability.Contract.CompleteOwnershipHandover(&_Neardataavailability.TransactOpts, pendingOwner)
}

// GrantRoles is a paid mutator transaction binding the contract method 0x1c10893f.
//
// Solidity: function grantRoles(address user, uint256 roles) payable returns()
func (_Neardataavailability *NeardataavailabilityTransactor) GrantRoles(opts *bind.TransactOpts, user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _Neardataavailability.contract.Transact(opts, "grantRoles", user, roles)
}

// GrantRoles is a paid mutator transaction binding the contract method 0x1c10893f.
//
// Solidity: function grantRoles(address user, uint256 roles) payable returns()
func (_Neardataavailability *NeardataavailabilitySession) GrantRoles(user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _Neardataavailability.Contract.GrantRoles(&_Neardataavailability.TransactOpts, user, roles)
}

// GrantRoles is a paid mutator transaction binding the contract method 0x1c10893f.
//
// Solidity: function grantRoles(address user, uint256 roles) payable returns()
func (_Neardataavailability *NeardataavailabilityTransactorSession) GrantRoles(user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _Neardataavailability.Contract.GrantRoles(&_Neardataavailability.TransactOpts, user, roles)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_Neardataavailability *NeardataavailabilityTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address) (*types.Transaction, error) {
	return _Neardataavailability.contract.Transact(opts, "initialize", initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_Neardataavailability *NeardataavailabilitySession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _Neardataavailability.Contract.Initialize(&_Neardataavailability.TransactOpts, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_Neardataavailability *NeardataavailabilityTransactorSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _Neardataavailability.Contract.Initialize(&_Neardataavailability.TransactOpts, initialOwner)
}

// NotifyAvailable is a paid mutator transaction binding the contract method 0x29afa6a3.
//
// Solidity: function notifyAvailable((bytes32,bytes32,bytes32) verifiedBatch) returns()
func (_Neardataavailability *NeardataavailabilityTransactor) NotifyAvailable(opts *bind.TransactOpts, verifiedBatch VerifiedBatch) (*types.Transaction, error) {
	return _Neardataavailability.contract.Transact(opts, "notifyAvailable", verifiedBatch)
}

// NotifyAvailable is a paid mutator transaction binding the contract method 0x29afa6a3.
//
// Solidity: function notifyAvailable((bytes32,bytes32,bytes32) verifiedBatch) returns()
func (_Neardataavailability *NeardataavailabilitySession) NotifyAvailable(verifiedBatch VerifiedBatch) (*types.Transaction, error) {
	return _Neardataavailability.Contract.NotifyAvailable(&_Neardataavailability.TransactOpts, verifiedBatch)
}

// NotifyAvailable is a paid mutator transaction binding the contract method 0x29afa6a3.
//
// Solidity: function notifyAvailable((bytes32,bytes32,bytes32) verifiedBatch) returns()
func (_Neardataavailability *NeardataavailabilityTransactorSession) NotifyAvailable(verifiedBatch VerifiedBatch) (*types.Transaction, error) {
	return _Neardataavailability.Contract.NotifyAvailable(&_Neardataavailability.TransactOpts, verifiedBatch)
}

// NotifySubmitted is a paid mutator transaction binding the contract method 0xb62aa7f9.
//
// Solidity: function notifySubmitted(bytes batches) returns()
func (_Neardataavailability *NeardataavailabilityTransactor) NotifySubmitted(opts *bind.TransactOpts, batches []byte) (*types.Transaction, error) {
	return _Neardataavailability.contract.Transact(opts, "notifySubmitted", batches)
}

// NotifySubmitted is a paid mutator transaction binding the contract method 0xb62aa7f9.
//
// Solidity: function notifySubmitted(bytes batches) returns()
func (_Neardataavailability *NeardataavailabilitySession) NotifySubmitted(batches []byte) (*types.Transaction, error) {
	return _Neardataavailability.Contract.NotifySubmitted(&_Neardataavailability.TransactOpts, batches)
}

// NotifySubmitted is a paid mutator transaction binding the contract method 0xb62aa7f9.
//
// Solidity: function notifySubmitted(bytes batches) returns()
func (_Neardataavailability *NeardataavailabilityTransactorSession) NotifySubmitted(batches []byte) (*types.Transaction, error) {
	return _Neardataavailability.Contract.NotifySubmitted(&_Neardataavailability.TransactOpts, batches)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_Neardataavailability *NeardataavailabilityTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Neardataavailability.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_Neardataavailability *NeardataavailabilitySession) RenounceOwnership() (*types.Transaction, error) {
	return _Neardataavailability.Contract.RenounceOwnership(&_Neardataavailability.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_Neardataavailability *NeardataavailabilityTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Neardataavailability.Contract.RenounceOwnership(&_Neardataavailability.TransactOpts)
}

// RenounceRoles is a paid mutator transaction binding the contract method 0x183a4f6e.
//
// Solidity: function renounceRoles(uint256 roles) payable returns()
func (_Neardataavailability *NeardataavailabilityTransactor) RenounceRoles(opts *bind.TransactOpts, roles *big.Int) (*types.Transaction, error) {
	return _Neardataavailability.contract.Transact(opts, "renounceRoles", roles)
}

// RenounceRoles is a paid mutator transaction binding the contract method 0x183a4f6e.
//
// Solidity: function renounceRoles(uint256 roles) payable returns()
func (_Neardataavailability *NeardataavailabilitySession) RenounceRoles(roles *big.Int) (*types.Transaction, error) {
	return _Neardataavailability.Contract.RenounceRoles(&_Neardataavailability.TransactOpts, roles)
}

// RenounceRoles is a paid mutator transaction binding the contract method 0x183a4f6e.
//
// Solidity: function renounceRoles(uint256 roles) payable returns()
func (_Neardataavailability *NeardataavailabilityTransactorSession) RenounceRoles(roles *big.Int) (*types.Transaction, error) {
	return _Neardataavailability.Contract.RenounceRoles(&_Neardataavailability.TransactOpts, roles)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_Neardataavailability *NeardataavailabilityTransactor) RequestOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Neardataavailability.contract.Transact(opts, "requestOwnershipHandover")
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_Neardataavailability *NeardataavailabilitySession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _Neardataavailability.Contract.RequestOwnershipHandover(&_Neardataavailability.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_Neardataavailability *NeardataavailabilityTransactorSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _Neardataavailability.Contract.RequestOwnershipHandover(&_Neardataavailability.TransactOpts)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x4a4ee7b1.
//
// Solidity: function revokeRoles(address user, uint256 roles) payable returns()
func (_Neardataavailability *NeardataavailabilityTransactor) RevokeRoles(opts *bind.TransactOpts, user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _Neardataavailability.contract.Transact(opts, "revokeRoles", user, roles)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x4a4ee7b1.
//
// Solidity: function revokeRoles(address user, uint256 roles) payable returns()
func (_Neardataavailability *NeardataavailabilitySession) RevokeRoles(user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _Neardataavailability.Contract.RevokeRoles(&_Neardataavailability.TransactOpts, user, roles)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x4a4ee7b1.
//
// Solidity: function revokeRoles(address user, uint256 roles) payable returns()
func (_Neardataavailability *NeardataavailabilityTransactorSession) RevokeRoles(user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _Neardataavailability.Contract.RevokeRoles(&_Neardataavailability.TransactOpts, user, roles)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_Neardataavailability *NeardataavailabilityTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Neardataavailability.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_Neardataavailability *NeardataavailabilitySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Neardataavailability.Contract.TransferOwnership(&_Neardataavailability.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_Neardataavailability *NeardataavailabilityTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Neardataavailability.Contract.TransferOwnership(&_Neardataavailability.TransactOpts, newOwner)
}

// NeardataavailabilityInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Neardataavailability contract.
type NeardataavailabilityInitializedIterator struct {
	Event *NeardataavailabilityInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NeardataavailabilityInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NeardataavailabilityInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NeardataavailabilityInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NeardataavailabilityInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NeardataavailabilityInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NeardataavailabilityInitialized represents a Initialized event raised by the Neardataavailability contract.
type NeardataavailabilityInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Neardataavailability *NeardataavailabilityFilterer) FilterInitialized(opts *bind.FilterOpts) (*NeardataavailabilityInitializedIterator, error) {

	logs, sub, err := _Neardataavailability.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NeardataavailabilityInitializedIterator{contract: _Neardataavailability.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Neardataavailability *NeardataavailabilityFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NeardataavailabilityInitialized) (event.Subscription, error) {

	logs, sub, err := _Neardataavailability.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NeardataavailabilityInitialized)
				if err := _Neardataavailability.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Neardataavailability *NeardataavailabilityFilterer) ParseInitialized(log types.Log) (*NeardataavailabilityInitialized, error) {
	event := new(NeardataavailabilityInitialized)
	if err := _Neardataavailability.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NeardataavailabilityIsAvailableIterator is returned from FilterIsAvailable and is used to iterate over the raw logs and unpacked data for IsAvailable events raised by the Neardataavailability contract.
type NeardataavailabilityIsAvailableIterator struct {
	Event *NeardataavailabilityIsAvailable // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NeardataavailabilityIsAvailableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NeardataavailabilityIsAvailable)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NeardataavailabilityIsAvailable)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NeardataavailabilityIsAvailableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NeardataavailabilityIsAvailableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NeardataavailabilityIsAvailable represents a IsAvailable event raised by the Neardataavailability contract.
type NeardataavailabilityIsAvailable struct {
	BucketIdx *big.Int
	Batch     VerifiedBatch
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIsAvailable is a free log retrieval operation binding the contract event 0x8985609fb3668d8b8e01db2e596eab8479f1e2447af72dab7baefdc1fd99d1bc.
//
// Solidity: event IsAvailable(uint256 bucketIdx, (bytes32,bytes32,bytes32) batch)
func (_Neardataavailability *NeardataavailabilityFilterer) FilterIsAvailable(opts *bind.FilterOpts) (*NeardataavailabilityIsAvailableIterator, error) {

	logs, sub, err := _Neardataavailability.contract.FilterLogs(opts, "IsAvailable")
	if err != nil {
		return nil, err
	}
	return &NeardataavailabilityIsAvailableIterator{contract: _Neardataavailability.contract, event: "IsAvailable", logs: logs, sub: sub}, nil
}

// WatchIsAvailable is a free log subscription operation binding the contract event 0x8985609fb3668d8b8e01db2e596eab8479f1e2447af72dab7baefdc1fd99d1bc.
//
// Solidity: event IsAvailable(uint256 bucketIdx, (bytes32,bytes32,bytes32) batch)
func (_Neardataavailability *NeardataavailabilityFilterer) WatchIsAvailable(opts *bind.WatchOpts, sink chan<- *NeardataavailabilityIsAvailable) (event.Subscription, error) {

	logs, sub, err := _Neardataavailability.contract.WatchLogs(opts, "IsAvailable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NeardataavailabilityIsAvailable)
				if err := _Neardataavailability.contract.UnpackLog(event, "IsAvailable", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIsAvailable is a log parse operation binding the contract event 0x8985609fb3668d8b8e01db2e596eab8479f1e2447af72dab7baefdc1fd99d1bc.
//
// Solidity: event IsAvailable(uint256 bucketIdx, (bytes32,bytes32,bytes32) batch)
func (_Neardataavailability *NeardataavailabilityFilterer) ParseIsAvailable(log types.Log) (*NeardataavailabilityIsAvailable, error) {
	event := new(NeardataavailabilityIsAvailable)
	if err := _Neardataavailability.contract.UnpackLog(event, "IsAvailable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NeardataavailabilityOwnershipHandoverCanceledIterator is returned from FilterOwnershipHandoverCanceled and is used to iterate over the raw logs and unpacked data for OwnershipHandoverCanceled events raised by the Neardataavailability contract.
type NeardataavailabilityOwnershipHandoverCanceledIterator struct {
	Event *NeardataavailabilityOwnershipHandoverCanceled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NeardataavailabilityOwnershipHandoverCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NeardataavailabilityOwnershipHandoverCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NeardataavailabilityOwnershipHandoverCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NeardataavailabilityOwnershipHandoverCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NeardataavailabilityOwnershipHandoverCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NeardataavailabilityOwnershipHandoverCanceled represents a OwnershipHandoverCanceled event raised by the Neardataavailability contract.
type NeardataavailabilityOwnershipHandoverCanceled struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverCanceled is a free log retrieval operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_Neardataavailability *NeardataavailabilityFilterer) FilterOwnershipHandoverCanceled(opts *bind.FilterOpts, pendingOwner []common.Address) (*NeardataavailabilityOwnershipHandoverCanceledIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Neardataavailability.contract.FilterLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NeardataavailabilityOwnershipHandoverCanceledIterator{contract: _Neardataavailability.contract, event: "OwnershipHandoverCanceled", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverCanceled is a free log subscription operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_Neardataavailability *NeardataavailabilityFilterer) WatchOwnershipHandoverCanceled(opts *bind.WatchOpts, sink chan<- *NeardataavailabilityOwnershipHandoverCanceled, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Neardataavailability.contract.WatchLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NeardataavailabilityOwnershipHandoverCanceled)
				if err := _Neardataavailability.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipHandoverCanceled is a log parse operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_Neardataavailability *NeardataavailabilityFilterer) ParseOwnershipHandoverCanceled(log types.Log) (*NeardataavailabilityOwnershipHandoverCanceled, error) {
	event := new(NeardataavailabilityOwnershipHandoverCanceled)
	if err := _Neardataavailability.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NeardataavailabilityOwnershipHandoverRequestedIterator is returned from FilterOwnershipHandoverRequested and is used to iterate over the raw logs and unpacked data for OwnershipHandoverRequested events raised by the Neardataavailability contract.
type NeardataavailabilityOwnershipHandoverRequestedIterator struct {
	Event *NeardataavailabilityOwnershipHandoverRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NeardataavailabilityOwnershipHandoverRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NeardataavailabilityOwnershipHandoverRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NeardataavailabilityOwnershipHandoverRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NeardataavailabilityOwnershipHandoverRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NeardataavailabilityOwnershipHandoverRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NeardataavailabilityOwnershipHandoverRequested represents a OwnershipHandoverRequested event raised by the Neardataavailability contract.
type NeardataavailabilityOwnershipHandoverRequested struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverRequested is a free log retrieval operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_Neardataavailability *NeardataavailabilityFilterer) FilterOwnershipHandoverRequested(opts *bind.FilterOpts, pendingOwner []common.Address) (*NeardataavailabilityOwnershipHandoverRequestedIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Neardataavailability.contract.FilterLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NeardataavailabilityOwnershipHandoverRequestedIterator{contract: _Neardataavailability.contract, event: "OwnershipHandoverRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverRequested is a free log subscription operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_Neardataavailability *NeardataavailabilityFilterer) WatchOwnershipHandoverRequested(opts *bind.WatchOpts, sink chan<- *NeardataavailabilityOwnershipHandoverRequested, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Neardataavailability.contract.WatchLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NeardataavailabilityOwnershipHandoverRequested)
				if err := _Neardataavailability.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipHandoverRequested is a log parse operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_Neardataavailability *NeardataavailabilityFilterer) ParseOwnershipHandoverRequested(log types.Log) (*NeardataavailabilityOwnershipHandoverRequested, error) {
	event := new(NeardataavailabilityOwnershipHandoverRequested)
	if err := _Neardataavailability.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NeardataavailabilityOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Neardataavailability contract.
type NeardataavailabilityOwnershipTransferredIterator struct {
	Event *NeardataavailabilityOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NeardataavailabilityOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NeardataavailabilityOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NeardataavailabilityOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NeardataavailabilityOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NeardataavailabilityOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NeardataavailabilityOwnershipTransferred represents a OwnershipTransferred event raised by the Neardataavailability contract.
type NeardataavailabilityOwnershipTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_Neardataavailability *NeardataavailabilityFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*NeardataavailabilityOwnershipTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Neardataavailability.contract.FilterLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NeardataavailabilityOwnershipTransferredIterator{contract: _Neardataavailability.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_Neardataavailability *NeardataavailabilityFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NeardataavailabilityOwnershipTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Neardataavailability.contract.WatchLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NeardataavailabilityOwnershipTransferred)
				if err := _Neardataavailability.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_Neardataavailability *NeardataavailabilityFilterer) ParseOwnershipTransferred(log types.Log) (*NeardataavailabilityOwnershipTransferred, error) {
	event := new(NeardataavailabilityOwnershipTransferred)
	if err := _Neardataavailability.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NeardataavailabilityRolesUpdatedIterator is returned from FilterRolesUpdated and is used to iterate over the raw logs and unpacked data for RolesUpdated events raised by the Neardataavailability contract.
type NeardataavailabilityRolesUpdatedIterator struct {
	Event *NeardataavailabilityRolesUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NeardataavailabilityRolesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NeardataavailabilityRolesUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NeardataavailabilityRolesUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NeardataavailabilityRolesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NeardataavailabilityRolesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NeardataavailabilityRolesUpdated represents a RolesUpdated event raised by the Neardataavailability contract.
type NeardataavailabilityRolesUpdated struct {
	User  common.Address
	Roles *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRolesUpdated is a free log retrieval operation binding the contract event 0x715ad5ce61fc9595c7b415289d59cf203f23a94fa06f04af7e489a0a76e1fe26.
//
// Solidity: event RolesUpdated(address indexed user, uint256 indexed roles)
func (_Neardataavailability *NeardataavailabilityFilterer) FilterRolesUpdated(opts *bind.FilterOpts, user []common.Address, roles []*big.Int) (*NeardataavailabilityRolesUpdatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var rolesRule []interface{}
	for _, rolesItem := range roles {
		rolesRule = append(rolesRule, rolesItem)
	}

	logs, sub, err := _Neardataavailability.contract.FilterLogs(opts, "RolesUpdated", userRule, rolesRule)
	if err != nil {
		return nil, err
	}
	return &NeardataavailabilityRolesUpdatedIterator{contract: _Neardataavailability.contract, event: "RolesUpdated", logs: logs, sub: sub}, nil
}

// WatchRolesUpdated is a free log subscription operation binding the contract event 0x715ad5ce61fc9595c7b415289d59cf203f23a94fa06f04af7e489a0a76e1fe26.
//
// Solidity: event RolesUpdated(address indexed user, uint256 indexed roles)
func (_Neardataavailability *NeardataavailabilityFilterer) WatchRolesUpdated(opts *bind.WatchOpts, sink chan<- *NeardataavailabilityRolesUpdated, user []common.Address, roles []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var rolesRule []interface{}
	for _, rolesItem := range roles {
		rolesRule = append(rolesRule, rolesItem)
	}

	logs, sub, err := _Neardataavailability.contract.WatchLogs(opts, "RolesUpdated", userRule, rolesRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NeardataavailabilityRolesUpdated)
				if err := _Neardataavailability.contract.UnpackLog(event, "RolesUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRolesUpdated is a log parse operation binding the contract event 0x715ad5ce61fc9595c7b415289d59cf203f23a94fa06f04af7e489a0a76e1fe26.
//
// Solidity: event RolesUpdated(address indexed user, uint256 indexed roles)
func (_Neardataavailability *NeardataavailabilityFilterer) ParseRolesUpdated(log types.Log) (*NeardataavailabilityRolesUpdated, error) {
	event := new(NeardataavailabilityRolesUpdated)
	if err := _Neardataavailability.contract.UnpackLog(event, "RolesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NeardataavailabilitySubmittedIterator is returned from FilterSubmitted and is used to iterate over the raw logs and unpacked data for Submitted events raised by the Neardataavailability contract.
type NeardataavailabilitySubmittedIterator struct {
	Event *NeardataavailabilitySubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NeardataavailabilitySubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NeardataavailabilitySubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NeardataavailabilitySubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NeardataavailabilitySubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NeardataavailabilitySubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NeardataavailabilitySubmitted represents a Submitted event raised by the Neardataavailability contract.
type NeardataavailabilitySubmitted struct {
	BucketIdx  *big.Int
	SubmitTxId [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSubmitted is a free log retrieval operation binding the contract event 0x465ddfeee4b3d0fb72ba16ca1e8d4ecf86373b44e51658af98043162de372b4c.
//
// Solidity: event Submitted(uint256 bucketIdx, bytes32 submitTxId)
func (_Neardataavailability *NeardataavailabilityFilterer) FilterSubmitted(opts *bind.FilterOpts) (*NeardataavailabilitySubmittedIterator, error) {

	logs, sub, err := _Neardataavailability.contract.FilterLogs(opts, "Submitted")
	if err != nil {
		return nil, err
	}
	return &NeardataavailabilitySubmittedIterator{contract: _Neardataavailability.contract, event: "Submitted", logs: logs, sub: sub}, nil
}

// WatchSubmitted is a free log subscription operation binding the contract event 0x465ddfeee4b3d0fb72ba16ca1e8d4ecf86373b44e51658af98043162de372b4c.
//
// Solidity: event Submitted(uint256 bucketIdx, bytes32 submitTxId)
func (_Neardataavailability *NeardataavailabilityFilterer) WatchSubmitted(opts *bind.WatchOpts, sink chan<- *NeardataavailabilitySubmitted) (event.Subscription, error) {

	logs, sub, err := _Neardataavailability.contract.WatchLogs(opts, "Submitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NeardataavailabilitySubmitted)
				if err := _Neardataavailability.contract.UnpackLog(event, "Submitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSubmitted is a log parse operation binding the contract event 0x465ddfeee4b3d0fb72ba16ca1e8d4ecf86373b44e51658af98043162de372b4c.
//
// Solidity: event Submitted(uint256 bucketIdx, bytes32 submitTxId)
func (_Neardataavailability *NeardataavailabilityFilterer) ParseSubmitted(log types.Log) (*NeardataavailabilitySubmitted, error) {
	event := new(NeardataavailabilitySubmitted)
	if err := _Neardataavailability.contract.UnpackLog(event, "Submitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
