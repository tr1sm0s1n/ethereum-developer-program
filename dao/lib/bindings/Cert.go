// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = bytes.Equal
	_ = errors.New
	_ = big.NewInt
	_ = common.Big1
	_ = types.BloomLookup
	_ = abi.ConvertType
)

// CertMetaData contains all meta data concerning the Cert contract.
var CertMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"date\",\"type\":\"string\"}],\"name\":\"Issued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Certificates\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"course\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"grade\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"date\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_course\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_grade\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_date\",\"type\":\"string\"}],\"name\":\"issue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "Cert",
	Bin: "0x608060405234801561000f575f5ffd5b50604051610f84380380610f84833981810160405281019061003191906101d7565b805f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036100a2575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016100999190610211565b60405180910390fd5b6100b1816100b860201b60201c565b505061022a565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6101a68261017d565b9050919050565b6101b68161019c565b81146101c0575f5ffd5b50565b5f815190506101d1816101ad565b92915050565b5f602082840312156101ec576101eb610179565b5b5f6101f9848285016101c3565b91505092915050565b61020b8161019c565b82525050565b5f6020820190506102245f830184610202565b92915050565b610d4d806102375f395ff3fe608060405234801561000f575f5ffd5b5060043610610055575f3560e01c8063715018a6146100595780638da5cb5b146100635780639622c83614610081578063d1e7be26146100b4578063f2fde38b146100d0575b5f5ffd5b6100616100ec565b005b61006b6100ff565b604051610078919061064f565b60405180910390f35b61009b600480360381019061009691906106ac565b610126565b6040516100ab9493929190610747565b60405180910390f35b6100ce60048036038101906100c991906108d2565b61036a565b005b6100ea60048036038101906100e591906109e3565b61043d565b005b6100f46104c1565b6100fd5f610548565b565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6001602052805f5260405f205f91509050805f01805461014590610a3b565b80601f016020809104026020016040519081016040528092919081815260200182805461017190610a3b565b80156101bc5780601f10610193576101008083540402835291602001916101bc565b820191905f5260205f20905b81548152906001019060200180831161019f57829003601f168201915b5050505050908060010180546101d190610a3b565b80601f01602080910402602001604051908101604052809291908181526020018280546101fd90610a3b565b80156102485780601f1061021f57610100808354040283529160200191610248565b820191905f5260205f20905b81548152906001019060200180831161022b57829003601f168201915b50505050509080600201805461025d90610a3b565b80601f016020809104026020016040519081016040528092919081815260200182805461028990610a3b565b80156102d45780601f106102ab576101008083540402835291602001916102d4565b820191905f5260205f20905b8154815290600101906020018083116102b757829003601f168201915b5050505050908060030180546102e990610a3b565b80601f016020809104026020016040519081016040528092919081815260200182805461031590610a3b565b80156103605780601f1061033757610100808354040283529160200191610360565b820191905f5260205f20905b81548152906001019060200180831161034357829003601f168201915b5050505050905084565b6103726104c1565b60405180608001604052808581526020018481526020018381526020018281525060015f8781526020019081526020015f205f820151815f0190816103b79190610c0b565b5060208201518160010190816103cd9190610c0b565b5060408201518160020190816103e39190610c0b565b5060608201518160030190816103f99190610c0b565b509050507f6f0d4ed09360dcc2710057310fc932744b9cfddf8c87f0ed6cf244b7abd2f9a0858260405161042e929190610ce9565b60405180910390a15050505050565b6104456104c1565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036104b5575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016104ac919061064f565b60405180910390fd5b6104be81610548565b50565b6104c9610609565b73ffffffffffffffffffffffffffffffffffffffff166104e76100ff565b73ffffffffffffffffffffffffffffffffffffffff16146105465761050a610609565b6040517f118cdaa700000000000000000000000000000000000000000000000000000000815260040161053d919061064f565b60405180910390fd5b565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f33905090565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61063982610610565b9050919050565b6106498161062f565b82525050565b5f6020820190506106625f830184610640565b92915050565b5f604051905090565b5f5ffd5b5f5ffd5b5f819050919050565b61068b81610679565b8114610695575f5ffd5b50565b5f813590506106a681610682565b92915050565b5f602082840312156106c1576106c0610671565b5b5f6106ce84828501610698565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610719826106d7565b61072381856106e1565b93506107338185602086016106f1565b61073c816106ff565b840191505092915050565b5f6080820190508181035f83015261075f818761070f565b90508181036020830152610773818661070f565b90508181036040830152610787818561070f565b9050818103606083015261079b818461070f565b905095945050505050565b5f5ffd5b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6107e4826106ff565b810181811067ffffffffffffffff82111715610803576108026107ae565b5b80604052505050565b5f610815610668565b905061082182826107db565b919050565b5f67ffffffffffffffff8211156108405761083f6107ae565b5b610849826106ff565b9050602081019050919050565b828183375f83830152505050565b5f61087661087184610826565b61080c565b905082815260208101848484011115610892576108916107aa565b5b61089d848285610856565b509392505050565b5f82601f8301126108b9576108b86107a6565b5b81356108c9848260208601610864565b91505092915050565b5f5f5f5f5f60a086880312156108eb576108ea610671565b5b5f6108f888828901610698565b955050602086013567ffffffffffffffff81111561091957610918610675565b5b610925888289016108a5565b945050604086013567ffffffffffffffff81111561094657610945610675565b5b610952888289016108a5565b935050606086013567ffffffffffffffff81111561097357610972610675565b5b61097f888289016108a5565b925050608086013567ffffffffffffffff8111156109a05761099f610675565b5b6109ac888289016108a5565b9150509295509295909350565b6109c28161062f565b81146109cc575f5ffd5b50565b5f813590506109dd816109b9565b92915050565b5f602082840312156109f8576109f7610671565b5b5f610a05848285016109cf565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610a5257607f821691505b602082108103610a6557610a64610a0e565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302610ac77fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610a8c565b610ad18683610a8c565b95508019841693508086168417925050509392505050565b5f819050919050565b5f610b0c610b07610b0284610679565b610ae9565b610679565b9050919050565b5f819050919050565b610b2583610af2565b610b39610b3182610b13565b848454610a98565b825550505050565b5f5f905090565b610b50610b41565b610b5b818484610b1c565b505050565b5b81811015610b7e57610b735f82610b48565b600181019050610b61565b5050565b601f821115610bc357610b9481610a6b565b610b9d84610a7d565b81016020851015610bac578190505b610bc0610bb885610a7d565b830182610b60565b50505b505050565b5f82821c905092915050565b5f610be35f1984600802610bc8565b1980831691505092915050565b5f610bfb8383610bd4565b9150826002028217905092915050565b610c14826106d7565b67ffffffffffffffff811115610c2d57610c2c6107ae565b5b610c378254610a3b565b610c42828285610b82565b5f60209050601f831160018114610c73575f8415610c61578287015190505b610c6b8582610bf0565b865550610cd2565b601f198416610c8186610a6b565b5f5b82811015610ca857848901518255600182019150602085019450602081019050610c83565b86831015610cc55784890151610cc1601f891682610bd4565b8355505b6001600288020188555050505b505050505050565b610ce381610679565b82525050565b5f604082019050610cfc5f830185610cda565b8181036020830152610d0e818461070f565b9050939250505056fea26469706673582212208002dbe165641de474bd06b0caaf8b37b2915473cff5ed6159a234eec9f20f1c64736f6c634300081e0033",
}

// Cert is an auto generated Go binding around an Ethereum contract.
type Cert struct {
	abi abi.ABI
}

// NewCert creates a new instance of Cert.
func NewCert() *Cert {
	parsed, err := CertMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Cert{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Cert) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address initialOwner) returns()
func (cert *Cert) PackConstructor(initialOwner common.Address) []byte {
	enc, err := cert.abi.Pack("", initialOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCertificates is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9622c836.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function Certificates(uint256 ) view returns(string name, string course, string grade, string date)
func (cert *Cert) PackCertificates(arg0 *big.Int) []byte {
	enc, err := cert.abi.Pack("Certificates", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCertificates is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9622c836.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function Certificates(uint256 ) view returns(string name, string course, string grade, string date)
func (cert *Cert) TryPackCertificates(arg0 *big.Int) ([]byte, error) {
	return cert.abi.Pack("Certificates", arg0)
}

// CertificatesOutput serves as a container for the return parameters of contract
// method Certificates.
type CertificatesOutput struct {
	Name   string
	Course string
	Grade  string
	Date   string
}

// UnpackCertificates is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9622c836.
//
// Solidity: function Certificates(uint256 ) view returns(string name, string course, string grade, string date)
func (cert *Cert) UnpackCertificates(data []byte) (CertificatesOutput, error) {
	out, err := cert.abi.Unpack("Certificates", data)
	outstruct := new(CertificatesOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Course = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Grade = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Date = *abi.ConvertType(out[3], new(string)).(*string)
	return *outstruct, nil
}

// PackIssue is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd1e7be26.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function issue(uint256 _id, string _name, string _course, string _grade, string _date) returns()
func (cert *Cert) PackIssue(id *big.Int, name string, course string, grade string, date string) []byte {
	enc, err := cert.abi.Pack("issue", id, name, course, grade, date)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIssue is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd1e7be26.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function issue(uint256 _id, string _name, string _course, string _grade, string _date) returns()
func (cert *Cert) TryPackIssue(id *big.Int, name string, course string, grade string, date string) ([]byte, error) {
	return cert.abi.Pack("issue", id, name, course, grade, date)
}

// PackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function owner() view returns(address)
func (cert *Cert) PackOwner() []byte {
	enc, err := cert.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function owner() view returns(address)
func (cert *Cert) TryPackOwner() ([]byte, error) {
	return cert.abi.Pack("owner")
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (cert *Cert) UnpackOwner(data []byte) (common.Address, error) {
	out, err := cert.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackRenounceOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x715018a6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function renounceOwnership() returns()
func (cert *Cert) PackRenounceOwnership() []byte {
	enc, err := cert.abi.Pack("renounceOwnership")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRenounceOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x715018a6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function renounceOwnership() returns()
func (cert *Cert) TryPackRenounceOwnership() ([]byte, error) {
	return cert.abi.Pack("renounceOwnership")
}

// PackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (cert *Cert) PackTransferOwnership(newOwner common.Address) []byte {
	enc, err := cert.abi.Pack("transferOwnership", newOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (cert *Cert) TryPackTransferOwnership(newOwner common.Address) ([]byte, error) {
	return cert.abi.Pack("transferOwnership", newOwner)
}

// CertIssued represents a Issued event raised by the Cert contract.
type CertIssued struct {
	Id   *big.Int
	Date string
	Raw  *types.Log // Blockchain specific contextual infos
}

const CertIssuedEventName = "Issued"

// ContractEventName returns the user-defined event name.
func (CertIssued) ContractEventName() string {
	return CertIssuedEventName
}

// UnpackIssuedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Issued(uint256 id, string date)
func (cert *Cert) UnpackIssuedEvent(log *types.Log) (*CertIssued, error) {
	event := "Issued"
	if log.Topics[0] != cert.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CertIssued)
	if len(log.Data) > 0 {
		if err := cert.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range cert.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// CertOwnershipTransferred represents a OwnershipTransferred event raised by the Cert contract.
type CertOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const CertOwnershipTransferredEventName = "OwnershipTransferred"

// ContractEventName returns the user-defined event name.
func (CertOwnershipTransferred) ContractEventName() string {
	return CertOwnershipTransferredEventName
}

// UnpackOwnershipTransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (cert *Cert) UnpackOwnershipTransferredEvent(log *types.Log) (*CertOwnershipTransferred, error) {
	event := "OwnershipTransferred"
	if log.Topics[0] != cert.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CertOwnershipTransferred)
	if len(log.Data) > 0 {
		if err := cert.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range cert.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (cert *Cert) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], cert.abi.Errors["OwnableInvalidOwner"].ID.Bytes()[:4]) {
		return cert.UnpackOwnableInvalidOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], cert.abi.Errors["OwnableUnauthorizedAccount"].ID.Bytes()[:4]) {
		return cert.UnpackOwnableUnauthorizedAccountError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// CertOwnableInvalidOwner represents a OwnableInvalidOwner error raised by the Cert contract.
type CertOwnableInvalidOwner struct {
	Owner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnableInvalidOwner(address owner)
func CertOwnableInvalidOwnerErrorID() common.Hash {
	return common.HexToHash("0x1e4fbdf7f3ef8bcaa855599e3abf48b232380f183f08f6f813d9ffa5bd585188")
}

// UnpackOwnableInvalidOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnableInvalidOwner(address owner)
func (cert *Cert) UnpackOwnableInvalidOwnerError(raw []byte) (*CertOwnableInvalidOwner, error) {
	out := new(CertOwnableInvalidOwner)
	if err := cert.abi.UnpackIntoInterface(out, "OwnableInvalidOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CertOwnableUnauthorizedAccount represents a OwnableUnauthorizedAccount error raised by the Cert contract.
type CertOwnableUnauthorizedAccount struct {
	Account common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnableUnauthorizedAccount(address account)
func CertOwnableUnauthorizedAccountErrorID() common.Hash {
	return common.HexToHash("0x118cdaa7a341953d1887a2245fd6665d741c67c8c50581daa59e1d03373fa188")
}

// UnpackOwnableUnauthorizedAccountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnableUnauthorizedAccount(address account)
func (cert *Cert) UnpackOwnableUnauthorizedAccountError(raw []byte) (*CertOwnableUnauthorizedAccount, error) {
	out := new(CertOwnableUnauthorizedAccount)
	if err := cert.abi.UnpackIntoInterface(out, "OwnableUnauthorizedAccount", raw); err != nil {
		return nil, err
	}
	return out, nil
}
