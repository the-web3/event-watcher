// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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
	_ = abi.ConvertType
)

// TreasureManagerMetaData contains all meta data concerning the TreasureManager contract.
var TreasureManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"a\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimAllTokens\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimToken\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositERC20\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositETH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"ethAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenWhiteList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRewards\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"granter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_treasureManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_withdrawManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"queryReward\",\"inputs\":[{\"name\":\"_tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTokenWhiteList\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setWithdrawManager\",\"inputs\":[{\"name\":\"_withdrawManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenBalances\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenWhiteList\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"treasureManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"userRewardAmounts\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawERC20\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"withdrawAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawETH\",\"inputs\":[{\"name\":\"withdrawAddress\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdrawManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DepositToken\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GrantRewardTokenAmount\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"granter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawManagerUpdate\",\"inputs\":[{\"name\":\"withdrawManager\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawToken\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"withdrawAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IsZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x6080806040523460b4577ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009081549060ff8260401c1660a557506001600160401b036002600160401b0319828216016061575b60405161162290816100b98239f35b6001600160401b031990911681179091556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80806052565b63f92ee8a960e01b8152600490fd5b5f80fdfe6080604090808252600480361015610029575b505050361561001f575f80fd5b610027611222565b005b5f3560e01c91826301ffc9a714610e945750816303186d2214610e6d5781630dbe671f14610e4f5781630ec08b0014610e0f57816317e3e2e814610da95781631e4bd42c14610c9457816323ecdbee14610bd2578163248a9ca314610b9c578163254c5d8714610b525781632f2ff15d14610b2a57816332f289cf1461096e57816336568abe1461092a578163410c9f9a1461080a57816341398b15146107dc57816344004cc1146106a25781634782f77914610664578163523fba7f1461062d578163715018a6146105c65781638da5cb5b1461059257816391d148541461054257816397feb92614610492578163a217fddf14610478578163a2672ad7146103de578163c0c53b8b14610272578163c37220ea14610230578163d547741f146101ea57508063ec3e9da5146101c2578063f2fde38b146101995763f6326fb314610176578080610012565b5f3660031901126101955760209061018c611222565b90519015158152f35b5f80fd5b34610195576020366003190112610195576100276101b5610f2f565b6101bd6112c5565b6111b1565b5034610195575f3660031901126101955760015490516001600160a01b039091168152602090f35b8234610195578060031936011261019557610027913561022b600161020d610f45565b93835f525f805160206115cd8339815191526020525f200154611340565b61140f565b82346101955760203660031901126101955760209161024d610f2f565b335f9081529184528282206001600160a01b0390911682528352819020549051908152f35b9050346101955760603660031901126101955761028d610f2f565b90610296610f45565b6044356001600160a01b038181169291839003610195577ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0094855460ff81891c16159567ffffffffffffffff8216801590816103d6575b60011490816103cc575b1590816103c3575b506103b5575067ffffffffffffffff1981166001178755610332919086610396575b5061032a611508565b6101bd611508565b6bffffffffffffffffffffffff60a01b9116815f5416175f55600154161760015561035957005b805468ff00000000000000001916905551600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1005b68ffffffffffffffffff1916680100000000000000011787555f610321565b885163f92ee8a960e01b8152fd5b9050155f6102ff565b303b1591506102f7565b8891506102ed565b8234610195576020366003190112610195576103f8610f2f565b5f546001600160a01b03939184916104139083163314610ff1565b1691821561046b5750600254906801000000000000000082101561045857508060016104429201600255610ee6565b909283549160031b92831b921b19161790555f80f35b604190634e487b7160e01b5f525260245ffd5b51631326d6d560e01b8152fd5b8234610195575f36600319011261019557602090515f8152f35b82346101955780600319360112610195576020906104ae610f2f565b81516323b872dd60e01b8482015233602480830191909152306044830152356064808301829052825291906104ee906104e8608482610f7c565b8261148e565b6001600160a01b03165f81815260038552839020805461050f908490611049565b905582519182527f4b3f81827ede20c81afbf1bb77b954afcdcae24d391d99042310cb1d9210dd57843393a35160018152f35b823461019557806003193601126101955760209161055e610f45565b90355f525f805160206115cd8339815191528352815f209060018060a01b03165f52825260ff815f20541690519015158152f35b8234610195575f366003190112610195575f805160206115ad8339815191525490516001600160a01b039091168152602090f35b34610195575f366003190112610195576105de6112c5565b5f805160206115ad83398151915280546001600160a01b031981169091555f906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b8234610195576020366003190112610195576020906001600160a01b03610652610f2f565b165f5260038252805f20549051908152f35b9050816003193601126101955735906001600160a01b0380831683036101955760209261069961018c92600154163314611056565b602435906110a2565b8234610195576060366003190112610195576106bc610f2f565b6106c4610f45565b6044359060018060a01b036106de81600154163314611056565b831692835f52600360205282855f20541061076557602095509061073c83827f9ca7c1e047552a8048d924a5a8d3c150eb861086a72a9100e5f19d1176c1b7469594875f5260038a52885f20610735848254610f5b565b90556112fd565b84513381526001600160a01b039190911660208201526040810191909152606090a25160018152f35b845162461bcd60e51b8152602081880152604560248201527f54726561737572654d616e6167657220776974686472617745524332303a204960448201527f6e73756666696369656e7420746f6b656e2062616c616e636520696e20636f6e6064820152641d1c9858dd60da1b608482015260a490fd5b8234610195575f366003190112610195576020905173eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee8152f35b90503461019557606036600319011261019557610825610f2f565b9061082e610f45565b6044359160018060a01b038094610849825f54163314610ff1565b16938415158061091f575b156108c657946108c1917f10621458f3ad2a9cfcb87c216122570629e44079d6af4d717035eb55d2c60424959684165f52602052805f20865f52602052805f2061089f858254611049565b9055516001600160a01b03909216825260208201929092529081906040820190565b0390a2005b855162461bcd60e51b8152602081840152602d60248201527f54726561737572654d616e61676572206772616e74526577617264733a20696e60448201526c76616c6964206164647265737360981b6064820152608490fd5b508083161515610854565b8234610195578060031936011261019557610943610f45565b90336001600160a01b0383160361095f5750610027913561140f565b5163334bd91960e11b81529050fd5b90503461019557602080600319360112610195576001600160a01b03610992610f2f565b168015610acf57335f52828252835f20815f528252835f2054908115610a7557335f52838352845f20815f5283525f8581205560038352845f206109d7838254610f5b565b905573eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee8103610a6557505f80808093335af1610a05610fb2565b5015610a0d57005b608492519162461bcd60e51b8352820152602f60248201527f54726561737572654d616e6167657220636c61696d546f6b656e3a204554482060448201526e1d1c985b9cd9995c8819985a5b1959608a1b6064820152fd5b91509150610027925033906112fd565b5050608492519162461bcd60e51b8352820152602f60248201527f54726561737572654d616e6167657220636c61696d546f6b656e3a206e6f207260448201526e657761726420617661696c61626c6560881b6064820152fd5b50608492519162461bcd60e51b8352820152603160248201527f54726561737572654d616e6167657220636c61696d546f6b656e3a20696e76616044820152706c696420746f6b656e206164647265737360781b6064820152fd5b82346101955780600319360112610195576100279135610b4d600161020d610f45565b61138c565b8234610195578060031936011261019557602091610b6e610f2f565b610b76610f45565b6001600160a01b039182165f908152928552838320911682528352819020549051908152f35b823461019557602036600319011261019557602091355f525f805160206115cd83398151915282526001815f2001549051908152f35b8234610195575f3660031901126101955780519081600254908181526020809101809260025f527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace905f5b818110610c775750505084610c33910385610f7c565b825181815293518185018190528493840192915f5b828110610c5757505050500390f35b83516001600160a01b031685528695509381019392810192600101610c48565b82546001600160a01b031684529284019260019283019201610c1d565b905034610195575f366003190112610195575f5b60025481101561002757610cbb81610ee6565b90549060039160018060a01b0391831b1c1690335f52602091848352855f20815f528352855f20549182610cf6575b50505050600101610ca8565b335f52858452865f20825f5284525f878120558352855f20610d19838254610f5b565b905573eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee8103610d9357505f80808093335af1610d47610fb2565b5015610d5b57506001905b905f8080610cea565b9050606492519162461bcd60e51b83528201526013602482015272115512081d1c985b9cd9995c8819985a5b1959606a1b6044820152fd5b600193925090610da49133906112fd565b610d52565b3461019557602036600319011261019557610dc2610f2f565b610dca6112c5565b600180546001600160a01b0319166001600160a01b039290921691821790557f799e16a314d482c87bc47fd219011aaf33f4f9c7e302be5d7a0af286a52998b95f80a2005b90503461019557602036600319011261019557359060025482101561019557610e39602092610ee6565b905491519160018060a01b039160031b1c168152f35b8234610195575f366003190112610195576020906005549051908152f35b8234610195575f366003190112610195575f5490516001600160a01b039091168152602090f35b903461019557602036600319011261019557359063ffffffff60e01b821680920361019557602091637965db0b60e01b8114908115610ed5575b5015158152f35b6301ffc9a760e01b14905083610ece565b600254811015610f1b5760025f527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace01905f90565b634e487b7160e01b5f52603260045260245ffd5b600435906001600160a01b038216820361019557565b602435906001600160a01b038216820361019557565b91908203918211610f6857565b634e487b7160e01b5f52601160045260245ffd5b90601f8019910116810190811067ffffffffffffffff821117610f9e57604052565b634e487b7160e01b5f52604160045260245ffd5b3d15610fec573d9067ffffffffffffffff8211610f9e5760405191610fe1601f8201601f191660200184610f7c565b82523d5f602084013e565b606090565b15610ff857565b60405162461bcd60e51b815260206004820152602360248201527f54726561737572654d616e616765722e6f6e6c7954726561737572654d616e6160448201526233b2b960e91b6064820152608490fd5b91908201809211610f6857565b1561105d57565b60405162461bcd60e51b815260206004820152601e60248201527f54726561737572654d616e616765722e6f6e6c795769746864726177657200006044820152606490fd5b9080471061113c5773eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee91825f52600360205260405f206110d7838254610f5b565b90556001600160a01b0316905f80808084865af16110f3610fb2565b5015611135577f9ca7c1e047552a8048d924a5a8d3c150eb861086a72a9100e5f19d1176c1b746916060916040519133835260208301526040820152a2600190565b5050505f90565b60405162461bcd60e51b815260206004820152604160248201527f54726561737572654d616e616765722077697468647261774554483a20696e7360448201527f756666696369656e74204554482062616c616e636520696e20636f6e747261636064820152601d60fa1b608482015260a490fd5b6001600160a01b0390811690811561120a575f805160206115ad83398151915280546001600160a01b031981168417909155167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3565b604051631e4fbdf760e01b81525f6004820152602490fd5b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f009060028254146112b3576002825573eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee805f52600360205260405f2061127d348254611049565b9055604051903482527f4b3f81827ede20c81afbf1bb77b954afcdcae24d391d99042310cb1d9210dd5760203393a36001809255565b604051633ee5aeb560e01b8152600490fd5b5f805160206115ad833981519152546001600160a01b031633036112e557565b60405163118cdaa760e01b8152336004820152602490fd5b60405163a9059cbb60e01b60208201526001600160a01b0392909216602483015260448083019390935291815261133e91611339606483610f7c565b61148e565b565b805f525f805160206115cd83398151915260205260405f20335f5260205260ff60405f2054161561136e5750565b6044906040519063e2517d3f60e01b82523360048301526024820152fd5b90815f525f805160206115cd8339815191528060205260405f209160018060a01b031691825f5260205260ff60405f205416155f1461113557825f5260205260405f20815f5260205260405f20600160ff1982541617905533917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d5f80a4600190565b90815f525f805160206115cd8339815191528060205260405f209160018060a01b031691825f5260205260ff60405f2054165f1461113557825f5260205260405f20815f5260205260405f2060ff19815416905533917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b5f80a4600190565b5f806114b69260018060a01b03169360208151910182865af16114af610fb2565b9083611549565b80519081151591826114e4575b50506114cc5750565b60249060405190635274afe760e01b82526004820152fd5b81925090602091810103126101955760200151801590811503610195575f806114c3565b60ff7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005460401c161561153757565b604051631afcd79f60e31b8152600490fd5b90611570575080511561155e57805190602001fd5b604051630a12f52160e11b8152600490fd5b815115806115a3575b611581575090565b604051639996b31560e01b81526001600160a01b039091166004820152602490fd5b50803b1561157956fe9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930002dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800a26469706673582212201270c9d2354aae0057b80e6aac9c2f97f6c286603aa19cf123218a51afbf7eef64736f6c63430008190033",
}

// TreasureManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use TreasureManagerMetaData.ABI instead.
var TreasureManagerABI = TreasureManagerMetaData.ABI

// TreasureManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TreasureManagerMetaData.Bin instead.
var TreasureManagerBin = TreasureManagerMetaData.Bin

// DeployTreasureManager deploys a new Ethereum contract, binding an instance of TreasureManager to it.
func DeployTreasureManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TreasureManager, error) {
	parsed, err := TreasureManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TreasureManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TreasureManager{TreasureManagerCaller: TreasureManagerCaller{contract: contract}, TreasureManagerTransactor: TreasureManagerTransactor{contract: contract}, TreasureManagerFilterer: TreasureManagerFilterer{contract: contract}}, nil
}

// TreasureManager is an auto generated Go binding around an Ethereum contract.
type TreasureManager struct {
	TreasureManagerCaller     // Read-only binding to the contract
	TreasureManagerTransactor // Write-only binding to the contract
	TreasureManagerFilterer   // Log filterer for contract events
}

// TreasureManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TreasureManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TreasureManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TreasureManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TreasureManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TreasureManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TreasureManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TreasureManagerSession struct {
	Contract     *TreasureManager  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TreasureManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TreasureManagerCallerSession struct {
	Contract *TreasureManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// TreasureManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TreasureManagerTransactorSession struct {
	Contract     *TreasureManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// TreasureManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TreasureManagerRaw struct {
	Contract *TreasureManager // Generic contract binding to access the raw methods on
}

// TreasureManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TreasureManagerCallerRaw struct {
	Contract *TreasureManagerCaller // Generic read-only contract binding to access the raw methods on
}

// TreasureManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TreasureManagerTransactorRaw struct {
	Contract *TreasureManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTreasureManager creates a new instance of TreasureManager, bound to a specific deployed contract.
func NewTreasureManager(address common.Address, backend bind.ContractBackend) (*TreasureManager, error) {
	contract, err := bindTreasureManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TreasureManager{TreasureManagerCaller: TreasureManagerCaller{contract: contract}, TreasureManagerTransactor: TreasureManagerTransactor{contract: contract}, TreasureManagerFilterer: TreasureManagerFilterer{contract: contract}}, nil
}

// NewTreasureManagerCaller creates a new read-only instance of TreasureManager, bound to a specific deployed contract.
func NewTreasureManagerCaller(address common.Address, caller bind.ContractCaller) (*TreasureManagerCaller, error) {
	contract, err := bindTreasureManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TreasureManagerCaller{contract: contract}, nil
}

// NewTreasureManagerTransactor creates a new write-only instance of TreasureManager, bound to a specific deployed contract.
func NewTreasureManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*TreasureManagerTransactor, error) {
	contract, err := bindTreasureManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TreasureManagerTransactor{contract: contract}, nil
}

// NewTreasureManagerFilterer creates a new log filterer instance of TreasureManager, bound to a specific deployed contract.
func NewTreasureManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*TreasureManagerFilterer, error) {
	contract, err := bindTreasureManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TreasureManagerFilterer{contract: contract}, nil
}

// bindTreasureManager binds a generic wrapper to an already deployed contract.
func bindTreasureManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TreasureManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TreasureManager *TreasureManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TreasureManager.Contract.TreasureManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TreasureManager *TreasureManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TreasureManager.Contract.TreasureManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TreasureManager *TreasureManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TreasureManager.Contract.TreasureManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TreasureManager *TreasureManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TreasureManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TreasureManager *TreasureManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TreasureManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TreasureManager *TreasureManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TreasureManager.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TreasureManager *TreasureManagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TreasureManager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TreasureManager *TreasureManagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TreasureManager.Contract.DEFAULTADMINROLE(&_TreasureManager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TreasureManager *TreasureManagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TreasureManager.Contract.DEFAULTADMINROLE(&_TreasureManager.CallOpts)
}

// A is a free data retrieval call binding the contract method 0x0dbe671f.
//
// Solidity: function a() view returns(uint256)
func (_TreasureManager *TreasureManagerCaller) A(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TreasureManager.contract.Call(opts, &out, "a")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// A is a free data retrieval call binding the contract method 0x0dbe671f.
//
// Solidity: function a() view returns(uint256)
func (_TreasureManager *TreasureManagerSession) A() (*big.Int, error) {
	return _TreasureManager.Contract.A(&_TreasureManager.CallOpts)
}

// A is a free data retrieval call binding the contract method 0x0dbe671f.
//
// Solidity: function a() view returns(uint256)
func (_TreasureManager *TreasureManagerCallerSession) A() (*big.Int, error) {
	return _TreasureManager.Contract.A(&_TreasureManager.CallOpts)
}

// EthAddress is a free data retrieval call binding the contract method 0x41398b15.
//
// Solidity: function ethAddress() view returns(address)
func (_TreasureManager *TreasureManagerCaller) EthAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TreasureManager.contract.Call(opts, &out, "ethAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EthAddress is a free data retrieval call binding the contract method 0x41398b15.
//
// Solidity: function ethAddress() view returns(address)
func (_TreasureManager *TreasureManagerSession) EthAddress() (common.Address, error) {
	return _TreasureManager.Contract.EthAddress(&_TreasureManager.CallOpts)
}

// EthAddress is a free data retrieval call binding the contract method 0x41398b15.
//
// Solidity: function ethAddress() view returns(address)
func (_TreasureManager *TreasureManagerCallerSession) EthAddress() (common.Address, error) {
	return _TreasureManager.Contract.EthAddress(&_TreasureManager.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TreasureManager *TreasureManagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TreasureManager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TreasureManager *TreasureManagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TreasureManager.Contract.GetRoleAdmin(&_TreasureManager.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TreasureManager *TreasureManagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TreasureManager.Contract.GetRoleAdmin(&_TreasureManager.CallOpts, role)
}

// GetTokenWhiteList is a free data retrieval call binding the contract method 0x23ecdbee.
//
// Solidity: function getTokenWhiteList() view returns(address[])
func (_TreasureManager *TreasureManagerCaller) GetTokenWhiteList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _TreasureManager.contract.Call(opts, &out, "getTokenWhiteList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTokenWhiteList is a free data retrieval call binding the contract method 0x23ecdbee.
//
// Solidity: function getTokenWhiteList() view returns(address[])
func (_TreasureManager *TreasureManagerSession) GetTokenWhiteList() ([]common.Address, error) {
	return _TreasureManager.Contract.GetTokenWhiteList(&_TreasureManager.CallOpts)
}

// GetTokenWhiteList is a free data retrieval call binding the contract method 0x23ecdbee.
//
// Solidity: function getTokenWhiteList() view returns(address[])
func (_TreasureManager *TreasureManagerCallerSession) GetTokenWhiteList() ([]common.Address, error) {
	return _TreasureManager.Contract.GetTokenWhiteList(&_TreasureManager.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TreasureManager *TreasureManagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _TreasureManager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TreasureManager *TreasureManagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TreasureManager.Contract.HasRole(&_TreasureManager.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TreasureManager *TreasureManagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TreasureManager.Contract.HasRole(&_TreasureManager.CallOpts, role, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TreasureManager *TreasureManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TreasureManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TreasureManager *TreasureManagerSession) Owner() (common.Address, error) {
	return _TreasureManager.Contract.Owner(&_TreasureManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TreasureManager *TreasureManagerCallerSession) Owner() (common.Address, error) {
	return _TreasureManager.Contract.Owner(&_TreasureManager.CallOpts)
}

// QueryReward is a free data retrieval call binding the contract method 0xc37220ea.
//
// Solidity: function queryReward(address _tokenAddress) view returns(uint256)
func (_TreasureManager *TreasureManagerCaller) QueryReward(opts *bind.CallOpts, _tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TreasureManager.contract.Call(opts, &out, "queryReward", _tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueryReward is a free data retrieval call binding the contract method 0xc37220ea.
//
// Solidity: function queryReward(address _tokenAddress) view returns(uint256)
func (_TreasureManager *TreasureManagerSession) QueryReward(_tokenAddress common.Address) (*big.Int, error) {
	return _TreasureManager.Contract.QueryReward(&_TreasureManager.CallOpts, _tokenAddress)
}

// QueryReward is a free data retrieval call binding the contract method 0xc37220ea.
//
// Solidity: function queryReward(address _tokenAddress) view returns(uint256)
func (_TreasureManager *TreasureManagerCallerSession) QueryReward(_tokenAddress common.Address) (*big.Int, error) {
	return _TreasureManager.Contract.QueryReward(&_TreasureManager.CallOpts, _tokenAddress)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TreasureManager *TreasureManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TreasureManager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TreasureManager *TreasureManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TreasureManager.Contract.SupportsInterface(&_TreasureManager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TreasureManager *TreasureManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TreasureManager.Contract.SupportsInterface(&_TreasureManager.CallOpts, interfaceId)
}

// TokenBalances is a free data retrieval call binding the contract method 0x523fba7f.
//
// Solidity: function tokenBalances(address ) view returns(uint256)
func (_TreasureManager *TreasureManagerCaller) TokenBalances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TreasureManager.contract.Call(opts, &out, "tokenBalances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenBalances is a free data retrieval call binding the contract method 0x523fba7f.
//
// Solidity: function tokenBalances(address ) view returns(uint256)
func (_TreasureManager *TreasureManagerSession) TokenBalances(arg0 common.Address) (*big.Int, error) {
	return _TreasureManager.Contract.TokenBalances(&_TreasureManager.CallOpts, arg0)
}

// TokenBalances is a free data retrieval call binding the contract method 0x523fba7f.
//
// Solidity: function tokenBalances(address ) view returns(uint256)
func (_TreasureManager *TreasureManagerCallerSession) TokenBalances(arg0 common.Address) (*big.Int, error) {
	return _TreasureManager.Contract.TokenBalances(&_TreasureManager.CallOpts, arg0)
}

// TokenWhiteList is a free data retrieval call binding the contract method 0x0ec08b00.
//
// Solidity: function tokenWhiteList(uint256 ) view returns(address)
func (_TreasureManager *TreasureManagerCaller) TokenWhiteList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TreasureManager.contract.Call(opts, &out, "tokenWhiteList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenWhiteList is a free data retrieval call binding the contract method 0x0ec08b00.
//
// Solidity: function tokenWhiteList(uint256 ) view returns(address)
func (_TreasureManager *TreasureManagerSession) TokenWhiteList(arg0 *big.Int) (common.Address, error) {
	return _TreasureManager.Contract.TokenWhiteList(&_TreasureManager.CallOpts, arg0)
}

// TokenWhiteList is a free data retrieval call binding the contract method 0x0ec08b00.
//
// Solidity: function tokenWhiteList(uint256 ) view returns(address)
func (_TreasureManager *TreasureManagerCallerSession) TokenWhiteList(arg0 *big.Int) (common.Address, error) {
	return _TreasureManager.Contract.TokenWhiteList(&_TreasureManager.CallOpts, arg0)
}

// TreasureManager is a free data retrieval call binding the contract method 0x03186d22.
//
// Solidity: function treasureManager() view returns(address)
func (_TreasureManager *TreasureManagerCaller) TreasureManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TreasureManager.contract.Call(opts, &out, "treasureManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TreasureManager is a free data retrieval call binding the contract method 0x03186d22.
//
// Solidity: function treasureManager() view returns(address)
func (_TreasureManager *TreasureManagerSession) TreasureManager() (common.Address, error) {
	return _TreasureManager.Contract.TreasureManager(&_TreasureManager.CallOpts)
}

// TreasureManager is a free data retrieval call binding the contract method 0x03186d22.
//
// Solidity: function treasureManager() view returns(address)
func (_TreasureManager *TreasureManagerCallerSession) TreasureManager() (common.Address, error) {
	return _TreasureManager.Contract.TreasureManager(&_TreasureManager.CallOpts)
}

// UserRewardAmounts is a free data retrieval call binding the contract method 0x254c5d87.
//
// Solidity: function userRewardAmounts(address , address ) view returns(uint256)
func (_TreasureManager *TreasureManagerCaller) UserRewardAmounts(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TreasureManager.contract.Call(opts, &out, "userRewardAmounts", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRewardAmounts is a free data retrieval call binding the contract method 0x254c5d87.
//
// Solidity: function userRewardAmounts(address , address ) view returns(uint256)
func (_TreasureManager *TreasureManagerSession) UserRewardAmounts(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TreasureManager.Contract.UserRewardAmounts(&_TreasureManager.CallOpts, arg0, arg1)
}

// UserRewardAmounts is a free data retrieval call binding the contract method 0x254c5d87.
//
// Solidity: function userRewardAmounts(address , address ) view returns(uint256)
func (_TreasureManager *TreasureManagerCallerSession) UserRewardAmounts(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TreasureManager.Contract.UserRewardAmounts(&_TreasureManager.CallOpts, arg0, arg1)
}

// WithdrawManager is a free data retrieval call binding the contract method 0xec3e9da5.
//
// Solidity: function withdrawManager() view returns(address)
func (_TreasureManager *TreasureManagerCaller) WithdrawManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TreasureManager.contract.Call(opts, &out, "withdrawManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WithdrawManager is a free data retrieval call binding the contract method 0xec3e9da5.
//
// Solidity: function withdrawManager() view returns(address)
func (_TreasureManager *TreasureManagerSession) WithdrawManager() (common.Address, error) {
	return _TreasureManager.Contract.WithdrawManager(&_TreasureManager.CallOpts)
}

// WithdrawManager is a free data retrieval call binding the contract method 0xec3e9da5.
//
// Solidity: function withdrawManager() view returns(address)
func (_TreasureManager *TreasureManagerCallerSession) WithdrawManager() (common.Address, error) {
	return _TreasureManager.Contract.WithdrawManager(&_TreasureManager.CallOpts)
}

// ClaimAllTokens is a paid mutator transaction binding the contract method 0x1e4bd42c.
//
// Solidity: function claimAllTokens() returns()
func (_TreasureManager *TreasureManagerTransactor) ClaimAllTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "claimAllTokens")
}

// ClaimAllTokens is a paid mutator transaction binding the contract method 0x1e4bd42c.
//
// Solidity: function claimAllTokens() returns()
func (_TreasureManager *TreasureManagerSession) ClaimAllTokens() (*types.Transaction, error) {
	return _TreasureManager.Contract.ClaimAllTokens(&_TreasureManager.TransactOpts)
}

// ClaimAllTokens is a paid mutator transaction binding the contract method 0x1e4bd42c.
//
// Solidity: function claimAllTokens() returns()
func (_TreasureManager *TreasureManagerTransactorSession) ClaimAllTokens() (*types.Transaction, error) {
	return _TreasureManager.Contract.ClaimAllTokens(&_TreasureManager.TransactOpts)
}

// ClaimToken is a paid mutator transaction binding the contract method 0x32f289cf.
//
// Solidity: function claimToken(address tokenAddress) returns()
func (_TreasureManager *TreasureManagerTransactor) ClaimToken(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "claimToken", tokenAddress)
}

// ClaimToken is a paid mutator transaction binding the contract method 0x32f289cf.
//
// Solidity: function claimToken(address tokenAddress) returns()
func (_TreasureManager *TreasureManagerSession) ClaimToken(tokenAddress common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.ClaimToken(&_TreasureManager.TransactOpts, tokenAddress)
}

// ClaimToken is a paid mutator transaction binding the contract method 0x32f289cf.
//
// Solidity: function claimToken(address tokenAddress) returns()
func (_TreasureManager *TreasureManagerTransactorSession) ClaimToken(tokenAddress common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.ClaimToken(&_TreasureManager.TransactOpts, tokenAddress)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x97feb926.
//
// Solidity: function depositERC20(address tokenAddress, uint256 amount) returns(bool)
func (_TreasureManager *TreasureManagerTransactor) DepositERC20(opts *bind.TransactOpts, tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "depositERC20", tokenAddress, amount)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x97feb926.
//
// Solidity: function depositERC20(address tokenAddress, uint256 amount) returns(bool)
func (_TreasureManager *TreasureManagerSession) DepositERC20(tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TreasureManager.Contract.DepositERC20(&_TreasureManager.TransactOpts, tokenAddress, amount)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x97feb926.
//
// Solidity: function depositERC20(address tokenAddress, uint256 amount) returns(bool)
func (_TreasureManager *TreasureManagerTransactorSession) DepositERC20(tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TreasureManager.Contract.DepositERC20(&_TreasureManager.TransactOpts, tokenAddress, amount)
}

// DepositETH is a paid mutator transaction binding the contract method 0xf6326fb3.
//
// Solidity: function depositETH() payable returns(bool)
func (_TreasureManager *TreasureManagerTransactor) DepositETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "depositETH")
}

// DepositETH is a paid mutator transaction binding the contract method 0xf6326fb3.
//
// Solidity: function depositETH() payable returns(bool)
func (_TreasureManager *TreasureManagerSession) DepositETH() (*types.Transaction, error) {
	return _TreasureManager.Contract.DepositETH(&_TreasureManager.TransactOpts)
}

// DepositETH is a paid mutator transaction binding the contract method 0xf6326fb3.
//
// Solidity: function depositETH() payable returns(bool)
func (_TreasureManager *TreasureManagerTransactorSession) DepositETH() (*types.Transaction, error) {
	return _TreasureManager.Contract.DepositETH(&_TreasureManager.TransactOpts)
}

// GrantRewards is a paid mutator transaction binding the contract method 0x410c9f9a.
//
// Solidity: function grantRewards(address tokenAddress, address granter, uint256 amount) returns()
func (_TreasureManager *TreasureManagerTransactor) GrantRewards(opts *bind.TransactOpts, tokenAddress common.Address, granter common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "grantRewards", tokenAddress, granter, amount)
}

// GrantRewards is a paid mutator transaction binding the contract method 0x410c9f9a.
//
// Solidity: function grantRewards(address tokenAddress, address granter, uint256 amount) returns()
func (_TreasureManager *TreasureManagerSession) GrantRewards(tokenAddress common.Address, granter common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TreasureManager.Contract.GrantRewards(&_TreasureManager.TransactOpts, tokenAddress, granter, amount)
}

// GrantRewards is a paid mutator transaction binding the contract method 0x410c9f9a.
//
// Solidity: function grantRewards(address tokenAddress, address granter, uint256 amount) returns()
func (_TreasureManager *TreasureManagerTransactorSession) GrantRewards(tokenAddress common.Address, granter common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TreasureManager.Contract.GrantRewards(&_TreasureManager.TransactOpts, tokenAddress, granter, amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TreasureManager *TreasureManagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TreasureManager *TreasureManagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.GrantRole(&_TreasureManager.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TreasureManager *TreasureManagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.GrantRole(&_TreasureManager.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address initialOwner, address _treasureManager, address _withdrawManager) returns()
func (_TreasureManager *TreasureManagerTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, _treasureManager common.Address, _withdrawManager common.Address) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "initialize", initialOwner, _treasureManager, _withdrawManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address initialOwner, address _treasureManager, address _withdrawManager) returns()
func (_TreasureManager *TreasureManagerSession) Initialize(initialOwner common.Address, _treasureManager common.Address, _withdrawManager common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.Initialize(&_TreasureManager.TransactOpts, initialOwner, _treasureManager, _withdrawManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address initialOwner, address _treasureManager, address _withdrawManager) returns()
func (_TreasureManager *TreasureManagerTransactorSession) Initialize(initialOwner common.Address, _treasureManager common.Address, _withdrawManager common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.Initialize(&_TreasureManager.TransactOpts, initialOwner, _treasureManager, _withdrawManager)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TreasureManager *TreasureManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TreasureManager *TreasureManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _TreasureManager.Contract.RenounceOwnership(&_TreasureManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TreasureManager *TreasureManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TreasureManager.Contract.RenounceOwnership(&_TreasureManager.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_TreasureManager *TreasureManagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_TreasureManager *TreasureManagerSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.RenounceRole(&_TreasureManager.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_TreasureManager *TreasureManagerTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.RenounceRole(&_TreasureManager.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TreasureManager *TreasureManagerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TreasureManager *TreasureManagerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.RevokeRole(&_TreasureManager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TreasureManager *TreasureManagerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.RevokeRole(&_TreasureManager.TransactOpts, role, account)
}

// SetTokenWhiteList is a paid mutator transaction binding the contract method 0xa2672ad7.
//
// Solidity: function setTokenWhiteList(address tokenAddress) returns()
func (_TreasureManager *TreasureManagerTransactor) SetTokenWhiteList(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "setTokenWhiteList", tokenAddress)
}

// SetTokenWhiteList is a paid mutator transaction binding the contract method 0xa2672ad7.
//
// Solidity: function setTokenWhiteList(address tokenAddress) returns()
func (_TreasureManager *TreasureManagerSession) SetTokenWhiteList(tokenAddress common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.SetTokenWhiteList(&_TreasureManager.TransactOpts, tokenAddress)
}

// SetTokenWhiteList is a paid mutator transaction binding the contract method 0xa2672ad7.
//
// Solidity: function setTokenWhiteList(address tokenAddress) returns()
func (_TreasureManager *TreasureManagerTransactorSession) SetTokenWhiteList(tokenAddress common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.SetTokenWhiteList(&_TreasureManager.TransactOpts, tokenAddress)
}

// SetWithdrawManager is a paid mutator transaction binding the contract method 0x17e3e2e8.
//
// Solidity: function setWithdrawManager(address _withdrawManager) returns()
func (_TreasureManager *TreasureManagerTransactor) SetWithdrawManager(opts *bind.TransactOpts, _withdrawManager common.Address) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "setWithdrawManager", _withdrawManager)
}

// SetWithdrawManager is a paid mutator transaction binding the contract method 0x17e3e2e8.
//
// Solidity: function setWithdrawManager(address _withdrawManager) returns()
func (_TreasureManager *TreasureManagerSession) SetWithdrawManager(_withdrawManager common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.SetWithdrawManager(&_TreasureManager.TransactOpts, _withdrawManager)
}

// SetWithdrawManager is a paid mutator transaction binding the contract method 0x17e3e2e8.
//
// Solidity: function setWithdrawManager(address _withdrawManager) returns()
func (_TreasureManager *TreasureManagerTransactorSession) SetWithdrawManager(_withdrawManager common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.SetWithdrawManager(&_TreasureManager.TransactOpts, _withdrawManager)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TreasureManager *TreasureManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TreasureManager *TreasureManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.TransferOwnership(&_TreasureManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TreasureManager *TreasureManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.TransferOwnership(&_TreasureManager.TransactOpts, newOwner)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x44004cc1.
//
// Solidity: function withdrawERC20(address tokenAddress, address withdrawAddress, uint256 amount) returns(bool)
func (_TreasureManager *TreasureManagerTransactor) WithdrawERC20(opts *bind.TransactOpts, tokenAddress common.Address, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "withdrawERC20", tokenAddress, withdrawAddress, amount)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x44004cc1.
//
// Solidity: function withdrawERC20(address tokenAddress, address withdrawAddress, uint256 amount) returns(bool)
func (_TreasureManager *TreasureManagerSession) WithdrawERC20(tokenAddress common.Address, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TreasureManager.Contract.WithdrawERC20(&_TreasureManager.TransactOpts, tokenAddress, withdrawAddress, amount)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x44004cc1.
//
// Solidity: function withdrawERC20(address tokenAddress, address withdrawAddress, uint256 amount) returns(bool)
func (_TreasureManager *TreasureManagerTransactorSession) WithdrawERC20(tokenAddress common.Address, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TreasureManager.Contract.WithdrawERC20(&_TreasureManager.TransactOpts, tokenAddress, withdrawAddress, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address withdrawAddress, uint256 amount) payable returns(bool)
func (_TreasureManager *TreasureManagerTransactor) WithdrawETH(opts *bind.TransactOpts, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "withdrawETH", withdrawAddress, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address withdrawAddress, uint256 amount) payable returns(bool)
func (_TreasureManager *TreasureManagerSession) WithdrawETH(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TreasureManager.Contract.WithdrawETH(&_TreasureManager.TransactOpts, withdrawAddress, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address withdrawAddress, uint256 amount) payable returns(bool)
func (_TreasureManager *TreasureManagerTransactorSession) WithdrawETH(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TreasureManager.Contract.WithdrawETH(&_TreasureManager.TransactOpts, withdrawAddress, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TreasureManager *TreasureManagerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TreasureManager.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TreasureManager *TreasureManagerSession) Receive() (*types.Transaction, error) {
	return _TreasureManager.Contract.Receive(&_TreasureManager.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TreasureManager *TreasureManagerTransactorSession) Receive() (*types.Transaction, error) {
	return _TreasureManager.Contract.Receive(&_TreasureManager.TransactOpts)
}

// TreasureManagerDepositTokenIterator is returned from FilterDepositToken and is used to iterate over the raw logs and unpacked data for DepositToken events raised by the TreasureManager contract.
type TreasureManagerDepositTokenIterator struct {
	Event *TreasureManagerDepositToken // Event containing the contract specifics and raw log

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
func (it *TreasureManagerDepositTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasureManagerDepositToken)
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
		it.Event = new(TreasureManagerDepositToken)
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
func (it *TreasureManagerDepositTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasureManagerDepositTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasureManagerDepositToken represents a DepositToken event raised by the TreasureManager contract.
type TreasureManagerDepositToken struct {
	TokenAddress common.Address
	Sender       common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDepositToken is a free log retrieval operation binding the contract event 0x4b3f81827ede20c81afbf1bb77b954afcdcae24d391d99042310cb1d9210dd57.
//
// Solidity: event DepositToken(address indexed tokenAddress, address indexed sender, uint256 amount)
func (_TreasureManager *TreasureManagerFilterer) FilterDepositToken(opts *bind.FilterOpts, tokenAddress []common.Address, sender []common.Address) (*TreasureManagerDepositTokenIterator, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TreasureManager.contract.FilterLogs(opts, "DepositToken", tokenAddressRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TreasureManagerDepositTokenIterator{contract: _TreasureManager.contract, event: "DepositToken", logs: logs, sub: sub}, nil
}

// WatchDepositToken is a free log subscription operation binding the contract event 0x4b3f81827ede20c81afbf1bb77b954afcdcae24d391d99042310cb1d9210dd57.
//
// Solidity: event DepositToken(address indexed tokenAddress, address indexed sender, uint256 amount)
func (_TreasureManager *TreasureManagerFilterer) WatchDepositToken(opts *bind.WatchOpts, sink chan<- *TreasureManagerDepositToken, tokenAddress []common.Address, sender []common.Address) (event.Subscription, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TreasureManager.contract.WatchLogs(opts, "DepositToken", tokenAddressRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasureManagerDepositToken)
				if err := _TreasureManager.contract.UnpackLog(event, "DepositToken", log); err != nil {
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

// ParseDepositToken is a log parse operation binding the contract event 0x4b3f81827ede20c81afbf1bb77b954afcdcae24d391d99042310cb1d9210dd57.
//
// Solidity: event DepositToken(address indexed tokenAddress, address indexed sender, uint256 amount)
func (_TreasureManager *TreasureManagerFilterer) ParseDepositToken(log types.Log) (*TreasureManagerDepositToken, error) {
	event := new(TreasureManagerDepositToken)
	if err := _TreasureManager.contract.UnpackLog(event, "DepositToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasureManagerGrantRewardTokenAmountIterator is returned from FilterGrantRewardTokenAmount and is used to iterate over the raw logs and unpacked data for GrantRewardTokenAmount events raised by the TreasureManager contract.
type TreasureManagerGrantRewardTokenAmountIterator struct {
	Event *TreasureManagerGrantRewardTokenAmount // Event containing the contract specifics and raw log

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
func (it *TreasureManagerGrantRewardTokenAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasureManagerGrantRewardTokenAmount)
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
		it.Event = new(TreasureManagerGrantRewardTokenAmount)
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
func (it *TreasureManagerGrantRewardTokenAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasureManagerGrantRewardTokenAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasureManagerGrantRewardTokenAmount represents a GrantRewardTokenAmount event raised by the TreasureManager contract.
type TreasureManagerGrantRewardTokenAmount struct {
	TokenAddress common.Address
	Granter      common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterGrantRewardTokenAmount is a free log retrieval operation binding the contract event 0x10621458f3ad2a9cfcb87c216122570629e44079d6af4d717035eb55d2c60424.
//
// Solidity: event GrantRewardTokenAmount(address indexed tokenAddress, address granter, uint256 amount)
func (_TreasureManager *TreasureManagerFilterer) FilterGrantRewardTokenAmount(opts *bind.FilterOpts, tokenAddress []common.Address) (*TreasureManagerGrantRewardTokenAmountIterator, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _TreasureManager.contract.FilterLogs(opts, "GrantRewardTokenAmount", tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &TreasureManagerGrantRewardTokenAmountIterator{contract: _TreasureManager.contract, event: "GrantRewardTokenAmount", logs: logs, sub: sub}, nil
}

// WatchGrantRewardTokenAmount is a free log subscription operation binding the contract event 0x10621458f3ad2a9cfcb87c216122570629e44079d6af4d717035eb55d2c60424.
//
// Solidity: event GrantRewardTokenAmount(address indexed tokenAddress, address granter, uint256 amount)
func (_TreasureManager *TreasureManagerFilterer) WatchGrantRewardTokenAmount(opts *bind.WatchOpts, sink chan<- *TreasureManagerGrantRewardTokenAmount, tokenAddress []common.Address) (event.Subscription, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _TreasureManager.contract.WatchLogs(opts, "GrantRewardTokenAmount", tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasureManagerGrantRewardTokenAmount)
				if err := _TreasureManager.contract.UnpackLog(event, "GrantRewardTokenAmount", log); err != nil {
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

// ParseGrantRewardTokenAmount is a log parse operation binding the contract event 0x10621458f3ad2a9cfcb87c216122570629e44079d6af4d717035eb55d2c60424.
//
// Solidity: event GrantRewardTokenAmount(address indexed tokenAddress, address granter, uint256 amount)
func (_TreasureManager *TreasureManagerFilterer) ParseGrantRewardTokenAmount(log types.Log) (*TreasureManagerGrantRewardTokenAmount, error) {
	event := new(TreasureManagerGrantRewardTokenAmount)
	if err := _TreasureManager.contract.UnpackLog(event, "GrantRewardTokenAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasureManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TreasureManager contract.
type TreasureManagerInitializedIterator struct {
	Event *TreasureManagerInitialized // Event containing the contract specifics and raw log

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
func (it *TreasureManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasureManagerInitialized)
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
		it.Event = new(TreasureManagerInitialized)
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
func (it *TreasureManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasureManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasureManagerInitialized represents a Initialized event raised by the TreasureManager contract.
type TreasureManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TreasureManager *TreasureManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*TreasureManagerInitializedIterator, error) {

	logs, sub, err := _TreasureManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TreasureManagerInitializedIterator{contract: _TreasureManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TreasureManager *TreasureManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TreasureManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _TreasureManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasureManagerInitialized)
				if err := _TreasureManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_TreasureManager *TreasureManagerFilterer) ParseInitialized(log types.Log) (*TreasureManagerInitialized, error) {
	event := new(TreasureManagerInitialized)
	if err := _TreasureManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasureManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TreasureManager contract.
type TreasureManagerOwnershipTransferredIterator struct {
	Event *TreasureManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TreasureManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasureManagerOwnershipTransferred)
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
		it.Event = new(TreasureManagerOwnershipTransferred)
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
func (it *TreasureManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasureManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasureManagerOwnershipTransferred represents a OwnershipTransferred event raised by the TreasureManager contract.
type TreasureManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TreasureManager *TreasureManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TreasureManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TreasureManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TreasureManagerOwnershipTransferredIterator{contract: _TreasureManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TreasureManager *TreasureManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TreasureManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TreasureManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasureManagerOwnershipTransferred)
				if err := _TreasureManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TreasureManager *TreasureManagerFilterer) ParseOwnershipTransferred(log types.Log) (*TreasureManagerOwnershipTransferred, error) {
	event := new(TreasureManagerOwnershipTransferred)
	if err := _TreasureManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasureManagerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the TreasureManager contract.
type TreasureManagerRoleAdminChangedIterator struct {
	Event *TreasureManagerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *TreasureManagerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasureManagerRoleAdminChanged)
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
		it.Event = new(TreasureManagerRoleAdminChanged)
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
func (it *TreasureManagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasureManagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasureManagerRoleAdminChanged represents a RoleAdminChanged event raised by the TreasureManager contract.
type TreasureManagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TreasureManager *TreasureManagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*TreasureManagerRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _TreasureManager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &TreasureManagerRoleAdminChangedIterator{contract: _TreasureManager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TreasureManager *TreasureManagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *TreasureManagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _TreasureManager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasureManagerRoleAdminChanged)
				if err := _TreasureManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TreasureManager *TreasureManagerFilterer) ParseRoleAdminChanged(log types.Log) (*TreasureManagerRoleAdminChanged, error) {
	event := new(TreasureManagerRoleAdminChanged)
	if err := _TreasureManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasureManagerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the TreasureManager contract.
type TreasureManagerRoleGrantedIterator struct {
	Event *TreasureManagerRoleGranted // Event containing the contract specifics and raw log

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
func (it *TreasureManagerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasureManagerRoleGranted)
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
		it.Event = new(TreasureManagerRoleGranted)
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
func (it *TreasureManagerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasureManagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasureManagerRoleGranted represents a RoleGranted event raised by the TreasureManager contract.
type TreasureManagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TreasureManager *TreasureManagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TreasureManagerRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TreasureManager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TreasureManagerRoleGrantedIterator{contract: _TreasureManager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TreasureManager *TreasureManagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *TreasureManagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TreasureManager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasureManagerRoleGranted)
				if err := _TreasureManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TreasureManager *TreasureManagerFilterer) ParseRoleGranted(log types.Log) (*TreasureManagerRoleGranted, error) {
	event := new(TreasureManagerRoleGranted)
	if err := _TreasureManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasureManagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the TreasureManager contract.
type TreasureManagerRoleRevokedIterator struct {
	Event *TreasureManagerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *TreasureManagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasureManagerRoleRevoked)
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
		it.Event = new(TreasureManagerRoleRevoked)
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
func (it *TreasureManagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasureManagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasureManagerRoleRevoked represents a RoleRevoked event raised by the TreasureManager contract.
type TreasureManagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TreasureManager *TreasureManagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TreasureManagerRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TreasureManager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TreasureManagerRoleRevokedIterator{contract: _TreasureManager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TreasureManager *TreasureManagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *TreasureManagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TreasureManager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasureManagerRoleRevoked)
				if err := _TreasureManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TreasureManager *TreasureManagerFilterer) ParseRoleRevoked(log types.Log) (*TreasureManagerRoleRevoked, error) {
	event := new(TreasureManagerRoleRevoked)
	if err := _TreasureManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasureManagerWithdrawManagerUpdateIterator is returned from FilterWithdrawManagerUpdate and is used to iterate over the raw logs and unpacked data for WithdrawManagerUpdate events raised by the TreasureManager contract.
type TreasureManagerWithdrawManagerUpdateIterator struct {
	Event *TreasureManagerWithdrawManagerUpdate // Event containing the contract specifics and raw log

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
func (it *TreasureManagerWithdrawManagerUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasureManagerWithdrawManagerUpdate)
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
		it.Event = new(TreasureManagerWithdrawManagerUpdate)
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
func (it *TreasureManagerWithdrawManagerUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasureManagerWithdrawManagerUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasureManagerWithdrawManagerUpdate represents a WithdrawManagerUpdate event raised by the TreasureManager contract.
type TreasureManagerWithdrawManagerUpdate struct {
	WithdrawManager common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawManagerUpdate is a free log retrieval operation binding the contract event 0x799e16a314d482c87bc47fd219011aaf33f4f9c7e302be5d7a0af286a52998b9.
//
// Solidity: event WithdrawManagerUpdate(address indexed withdrawManager)
func (_TreasureManager *TreasureManagerFilterer) FilterWithdrawManagerUpdate(opts *bind.FilterOpts, withdrawManager []common.Address) (*TreasureManagerWithdrawManagerUpdateIterator, error) {

	var withdrawManagerRule []interface{}
	for _, withdrawManagerItem := range withdrawManager {
		withdrawManagerRule = append(withdrawManagerRule, withdrawManagerItem)
	}

	logs, sub, err := _TreasureManager.contract.FilterLogs(opts, "WithdrawManagerUpdate", withdrawManagerRule)
	if err != nil {
		return nil, err
	}
	return &TreasureManagerWithdrawManagerUpdateIterator{contract: _TreasureManager.contract, event: "WithdrawManagerUpdate", logs: logs, sub: sub}, nil
}

// WatchWithdrawManagerUpdate is a free log subscription operation binding the contract event 0x799e16a314d482c87bc47fd219011aaf33f4f9c7e302be5d7a0af286a52998b9.
//
// Solidity: event WithdrawManagerUpdate(address indexed withdrawManager)
func (_TreasureManager *TreasureManagerFilterer) WatchWithdrawManagerUpdate(opts *bind.WatchOpts, sink chan<- *TreasureManagerWithdrawManagerUpdate, withdrawManager []common.Address) (event.Subscription, error) {

	var withdrawManagerRule []interface{}
	for _, withdrawManagerItem := range withdrawManager {
		withdrawManagerRule = append(withdrawManagerRule, withdrawManagerItem)
	}

	logs, sub, err := _TreasureManager.contract.WatchLogs(opts, "WithdrawManagerUpdate", withdrawManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasureManagerWithdrawManagerUpdate)
				if err := _TreasureManager.contract.UnpackLog(event, "WithdrawManagerUpdate", log); err != nil {
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

// ParseWithdrawManagerUpdate is a log parse operation binding the contract event 0x799e16a314d482c87bc47fd219011aaf33f4f9c7e302be5d7a0af286a52998b9.
//
// Solidity: event WithdrawManagerUpdate(address indexed withdrawManager)
func (_TreasureManager *TreasureManagerFilterer) ParseWithdrawManagerUpdate(log types.Log) (*TreasureManagerWithdrawManagerUpdate, error) {
	event := new(TreasureManagerWithdrawManagerUpdate)
	if err := _TreasureManager.contract.UnpackLog(event, "WithdrawManagerUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasureManagerWithdrawTokenIterator is returned from FilterWithdrawToken and is used to iterate over the raw logs and unpacked data for WithdrawToken events raised by the TreasureManager contract.
type TreasureManagerWithdrawTokenIterator struct {
	Event *TreasureManagerWithdrawToken // Event containing the contract specifics and raw log

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
func (it *TreasureManagerWithdrawTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasureManagerWithdrawToken)
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
		it.Event = new(TreasureManagerWithdrawToken)
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
func (it *TreasureManagerWithdrawTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasureManagerWithdrawTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasureManagerWithdrawToken represents a WithdrawToken event raised by the TreasureManager contract.
type TreasureManagerWithdrawToken struct {
	TokenAddress    common.Address
	Sender          common.Address
	WithdrawAddress common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawToken is a free log retrieval operation binding the contract event 0x9ca7c1e047552a8048d924a5a8d3c150eb861086a72a9100e5f19d1176c1b746.
//
// Solidity: event WithdrawToken(address indexed tokenAddress, address sender, address withdrawAddress, uint256 amount)
func (_TreasureManager *TreasureManagerFilterer) FilterWithdrawToken(opts *bind.FilterOpts, tokenAddress []common.Address) (*TreasureManagerWithdrawTokenIterator, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _TreasureManager.contract.FilterLogs(opts, "WithdrawToken", tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &TreasureManagerWithdrawTokenIterator{contract: _TreasureManager.contract, event: "WithdrawToken", logs: logs, sub: sub}, nil
}

// WatchWithdrawToken is a free log subscription operation binding the contract event 0x9ca7c1e047552a8048d924a5a8d3c150eb861086a72a9100e5f19d1176c1b746.
//
// Solidity: event WithdrawToken(address indexed tokenAddress, address sender, address withdrawAddress, uint256 amount)
func (_TreasureManager *TreasureManagerFilterer) WatchWithdrawToken(opts *bind.WatchOpts, sink chan<- *TreasureManagerWithdrawToken, tokenAddress []common.Address) (event.Subscription, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _TreasureManager.contract.WatchLogs(opts, "WithdrawToken", tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasureManagerWithdrawToken)
				if err := _TreasureManager.contract.UnpackLog(event, "WithdrawToken", log); err != nil {
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

// ParseWithdrawToken is a log parse operation binding the contract event 0x9ca7c1e047552a8048d924a5a8d3c150eb861086a72a9100e5f19d1176c1b746.
//
// Solidity: event WithdrawToken(address indexed tokenAddress, address sender, address withdrawAddress, uint256 amount)
func (_TreasureManager *TreasureManagerFilterer) ParseWithdrawToken(log types.Log) (*TreasureManagerWithdrawToken, error) {
	event := new(TreasureManagerWithdrawToken)
	if err := _TreasureManager.contract.UnpackLog(event, "WithdrawToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
