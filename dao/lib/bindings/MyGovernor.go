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

// MyGovernorMetaData contains all meta data concerning the MyGovernor contract.
var MyGovernorMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIVotes\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"contractTimelockController\",\"name\":\"_timelock\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CheckpointUnorderedInsertion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"GovernorAlreadyCastVote\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"GovernorAlreadyQueuedProposal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernorDisabledDeposit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"GovernorInsufficientProposerVotes\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"targets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calldatas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"values\",\"type\":\"uint256\"}],\"name\":\"GovernorInvalidProposalLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quorumNumerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quorumDenominator\",\"type\":\"uint256\"}],\"name\":\"GovernorInvalidQuorumFraction\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"GovernorInvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernorInvalidVoteParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernorInvalidVoteType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"votingPeriod\",\"type\":\"uint256\"}],\"name\":\"GovernorInvalidVotingPeriod\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"GovernorNonexistentProposal\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"GovernorNotQueuedProposal\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorOnlyExecutor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernorQueueNotImplemented\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"}],\"name\":\"GovernorRestrictedProposer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorUnableToCancel\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"enumIGovernor.ProposalState\",\"name\":\"current\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"expectedStates\",\"type\":\"bytes32\"}],\"name\":\"GovernorUnexpectedProposalState\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"signatures\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voteStart\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voteEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"etaSeconds\",\"type\":\"uint256\"}],\"name\":\"ProposalQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldQuorumNumerator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newQuorumNumerator\",\"type\":\"uint256\"}],\"name\":\"QuorumNumeratorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldTimelock\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTimelock\",\"type\":\"address\"}],\"name\":\"TimelockChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"VoteCast\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"VoteCastWithParams\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BALLOT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLOCK_MODE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"COUNTING_MODE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXTENDED_BALLOT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"}],\"name\":\"castVote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"castVoteBySig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"castVoteWithReason\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"castVoteWithReasonAndParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"castVoteWithReasonAndParamsBySig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clock\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"getProposalId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timepoint\",\"type\":\"uint256\"}],\"name\":\"getVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timepoint\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"getVotesWithParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"hashProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalDeadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalEta\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalNeedsQueuing\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalProposer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"againstVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"abstainVotes\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"propose\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"queue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timepoint\",\"type\":\"uint256\"}],\"name\":\"quorum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quorumDenominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timepoint\",\"type\":\"uint256\"}],\"name\":\"quorumNumerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quorumNumerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"relay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumIGovernor.ProposalState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timelock\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC5805\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newQuorumNumerator\",\"type\":\"uint256\"}],\"name\":\"updateQuorumNumerator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractTimelockController\",\"name\":\"newTimelock\",\"type\":\"address\"}],\"name\":\"updateTimelock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	ID:  "MyGovernor",
	Bin: "0x610180604052348015610010575f5ffd5b50604051617e3e380380617e3e83398181016040528101906100329190610a61565b806004836040518060400160405280600a81526020017f4d79476f7665726e6f72000000000000000000000000000000000000000000008152508061007b61019260201b60201c565b61008e5f836101cf60201b90919060201c565b61012081815250506100aa6001826101cf60201b90919060201c565b6101408181525050818051906020012060e08181525050808051906020012061010081815250504660a081815250506100e761021c60201b60201c565b608081815250503073ffffffffffffffffffffffffffffffffffffffff1660c08173ffffffffffffffffffffffffffffffffffffffff1681525050505080600390816101339190610cdc565b50508073ffffffffffffffffffffffffffffffffffffffff166101608173ffffffffffffffffffffffffffffffffffffffff16815250505061017a8161027660201b60201c565b5061018a8161035060201b60201c565b50505061112b565b60606040518060400160405280600181526020017f3100000000000000000000000000000000000000000000000000000000000000815250905090565b5f6020835110156101f0576101e9836103ed60201b60201c565b9050610216565b826102008361045260201b60201c565b5f01908161020e9190610cdc565b5060ff5f1b90505b92915050565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60e05161010051463060405160200161025b959493929190610de1565b60405160208183030381529060405280519060200120905090565b5f61028561045b60201b60201c565b9050808211156102ce5781816040517f243e54450000000000000000000000000000000000000000000000000000000081526004016102c5929190610e32565b60405180910390fd5b5f6102dd61046360201b60201c565b90506103106102f061049560201b60201c565b6102ff8561052c60201b60201c565b600861059960201b9092919060201c565b50507f0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b46339978184604051610343929190610e32565b60405180910390a1505050565b7f08f74ea46ef7894f65eabfb5e6e695de773a000b47c529ab559178069b22640160095f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16826040516103a2929190610e59565b60405180910390a18060095f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f5f829050601f8151111561043957826040517f305a27a90000000000000000000000000000000000000000000000000000000081526004016104309190610ee6565b60405180910390fd5b80518161044590610f33565b5f1c175f1b915050919050565b5f819050919050565b5f6064905090565b5f61047460086105ba60201b60201c565b79ffffffffffffffffffffffffffffffffffffffffffffffffffff16905090565b5f6104a461062260201b60201c565b73ffffffffffffffffffffffffffffffffffffffff166391ddadf46040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801561050b57506040513d601f19601f820116820180604052508101906105089190610fd4565b60015b6105245761051d61062c60201b60201c565b9050610529565b809150505b90565b5f79ffffffffffffffffffffffffffffffffffffffffffffffffffff80168211156105915760d0826040517f6dfcc650000000000000000000000000000000000000000000000000000000008152600401610588929190611044565b60405180910390fd5b819050919050565b5f5f6105ae855f01858561064160201b60201c565b91509150935093915050565b5f5f825f018054905090505f8114610618576105ea835f016001836105df9190611098565b61093b60201b60201c565b5f0160069054906101000a900479ffffffffffffffffffffffffffffffffffffffffffffffffffff1661061a565b5f5b915050919050565b5f61016051905090565b5f61063c4361094d60201b60201c565b905090565b5f5f5f858054905090505f811115610853575f610670876001846106659190611098565b61093b60201b60201c565b90505f815f015f9054906101000a900465ffffffffffff1690505f825f0160069054906101000a900479ffffffffffffffffffffffffffffffffffffffffffffffffffff1690508765ffffffffffff168265ffffffffffff161115610701576040517f2520601d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8765ffffffffffff168265ffffffffffff160361076b5786835f0160066101000a81548179ffffffffffffffffffffffffffffffffffffffffffffffffffff021916908379ffffffffffffffffffffffffffffffffffffffffffffffffffff160217905550610844565b8860405180604001604052808a65ffffffffffff1681526020018979ffffffffffffffffffffffffffffffffffffffffffffffffffff16815250908060018154018082558091505060019003905f5260205f20015f909190919091505f820151815f015f6101000a81548165ffffffffffff021916908365ffffffffffff1602179055506020820151815f0160066101000a81548179ffffffffffffffffffffffffffffffffffffffffffffffffffff021916908379ffffffffffffffffffffffffffffffffffffffffffffffffffff16021790555050505b80879550955050505050610933565b8560405180604001604052808765ffffffffffff1681526020018679ffffffffffffffffffffffffffffffffffffffffffffffffffff16815250908060018154018082558091505060019003905f5260205f20015f909190919091505f820151815f015f6101000a81548165ffffffffffff021916908365ffffffffffff1602179055506020820151815f0160066101000a81548179ffffffffffffffffffffffffffffffffffffffffffffffffffff021916908379ffffffffffffffffffffffffffffffffffffffffffffffffffff16021790555050505f8492509250505b935093915050565b5f825f528160205f2001905092915050565b5f65ffffffffffff801682111561099e576030826040517f6dfcc650000000000000000000000000000000000000000000000000000000008152600401610995929190611104565b60405180910390fd5b819050919050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6109d3826109aa565b9050919050565b5f6109e4826109c9565b9050919050565b6109f4816109da565b81146109fe575f5ffd5b50565b5f81519050610a0f816109eb565b92915050565b5f610a1f826109aa565b9050919050565b5f610a3082610a15565b9050919050565b610a4081610a26565b8114610a4a575f5ffd5b50565b5f81519050610a5b81610a37565b92915050565b5f5f60408385031215610a7757610a766109a6565b5b5f610a8485828601610a01565b9250506020610a9585828601610a4d565b9150509250929050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610b1a57607f821691505b602082108103610b2d57610b2c610ad6565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302610b8f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610b54565b610b998683610b54565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f610bdd610bd8610bd384610bb1565b610bba565b610bb1565b9050919050565b5f819050919050565b610bf683610bc3565b610c0a610c0282610be4565b848454610b60565b825550505050565b5f5f905090565b610c21610c12565b610c2c818484610bed565b505050565b5b81811015610c4f57610c445f82610c19565b600181019050610c32565b5050565b601f821115610c9457610c6581610b33565b610c6e84610b45565b81016020851015610c7d578190505b610c91610c8985610b45565b830182610c31565b50505b505050565b5f82821c905092915050565b5f610cb45f1984600802610c99565b1980831691505092915050565b5f610ccc8383610ca5565b9150826002028217905092915050565b610ce582610a9f565b67ffffffffffffffff811115610cfe57610cfd610aa9565b5b610d088254610b03565b610d13828285610c53565b5f60209050601f831160018114610d44575f8415610d32578287015190505b610d3c8582610cc1565b865550610da3565b601f198416610d5286610b33565b5f5b82811015610d7957848901518255600182019150602085019450602081019050610d54565b86831015610d965784890151610d92601f891682610ca5565b8355505b6001600288020188555050505b505050505050565b5f819050919050565b610dbd81610dab565b82525050565b610dcc81610bb1565b82525050565b610ddb816109c9565b82525050565b5f60a082019050610df45f830188610db4565b610e016020830187610db4565b610e0e6040830186610db4565b610e1b6060830185610dc3565b610e286080830184610dd2565b9695505050505050565b5f604082019050610e455f830185610dc3565b610e526020830184610dc3565b9392505050565b5f604082019050610e6c5f830185610dd2565b610e796020830184610dd2565b9392505050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610eb882610a9f565b610ec28185610e80565b9350610ed2818560208601610e90565b610edb81610e9e565b840191505092915050565b5f6020820190508181035f830152610efe8184610eae565b905092915050565b5f81519050919050565b5f819050602082019050919050565b5f610f2a8251610dab565b80915050919050565b5f610f3d82610f06565b82610f4784610f10565b9050610f5281610f1f565b92506020821015610f9257610f8d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff83602003600802610b54565b831692505b5050919050565b5f65ffffffffffff82169050919050565b610fb381610f99565b8114610fbd575f5ffd5b50565b5f81519050610fce81610faa565b92915050565b5f60208284031215610fe957610fe86109a6565b5b5f610ff684828501610fc0565b91505092915050565b5f819050919050565b5f60ff82169050919050565b5f61102e61102961102484610fff565b610bba565b611008565b9050919050565b61103e81611014565b82525050565b5f6040820190506110575f830185611035565b6110646020830184610dc3565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6110a282610bb1565b91506110ad83610bb1565b92508282039050818111156110c5576110c461106b565b5b92915050565b5f819050919050565b5f6110ee6110e96110e4846110cb565b610bba565b611008565b9050919050565b6110fe816110d4565b82525050565b5f6040820190506111175f8301856110f5565b6111246020830184610dc3565b9392505050565b60805160a05160c05160e05161010051610120516101405161016051616cb86111865f395f611d9701525f612aab01525f612a7001525f61473701525f61471601525f6140c101525f61411701525f6141400152616cb85ff3fe608060405260043610610280575f3560e01c80637ecebe001161014e578063b58131b0116100c0578063dd4e2ba511610079578063dd4e2ba514610b33578063deaaa7cc14610b5d578063eb9019d414610b87578063f23a6e6114610bc3578063f8ce560a14610bff578063fc0c546a14610c3b576102f3565b8063b58131b014610a0f578063bc197c8114610a39578063c01f9e3714610a75578063c28bc2fa14610ab1578063c59057e414610acd578063d33219b414610b09576102f3565b80639a802a6d116101125780639a802a6d146108cd578063a7713a7014610909578063a890c91014610933578063a8f8a6681461095b578063a9a9529414610997578063ab58fb8e146109d3576102f3565b80637ecebe00146107d157806384b0196e1461080d5780638ff262e31461083d57806391ddadf41461087957806397c3d334146108a3576102f3565b80633e4f49e6116101f257806356781388116101ab57806356781388146106695780635b8d0e0d146106a55780635f398a14146106e157806360c4247f1461071d5780637b3c71d3146107595780637d5e81e214610795576102f3565b80633e4f49e614610523578063438596321461055f578063452115d61461059b5780634bf5d7e9146105d7578063544ffc9c1461060157806354fd4d501461063f576102f3565b8063150b7a0211610244578063150b7a02146103eb578063160cbed7146104275780632656227d146104635780632d63f693146104935780632fe3e261146104cf5780633932abb1146104f9576102f3565b806301ffc9a7146102f757806302a251a31461033357806306f3f9e61461035d57806306fdde0314610385578063143489d0146103af576102f3565b366102f3573073ffffffffffffffffffffffffffffffffffffffff166102a4610c65565b73ffffffffffffffffffffffffffffffffffffffff16146102f1576040517fe90a651e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b005b5f5ffd5b348015610302575f5ffd5b5061031d600480360381019061031891906149d7565b610c73565b60405161032a9190614a1c565b60405180910390f35b34801561033e575f5ffd5b50610347610dc5565b6040516103549190614a4d565b60405180910390f35b348015610368575f5ffd5b50610383600480360381019061037e9190614a90565b610dcd565b005b348015610390575f5ffd5b50610399610de1565b6040516103a69190614b2b565b60405180910390f35b3480156103ba575f5ffd5b506103d560048036038101906103d09190614a90565b610e71565b6040516103e29190614b8a565b60405180910390f35b3480156103f6575f5ffd5b50610411600480360381019061040c9190614cf9565b610eac565b60405161041e9190614d88565b60405180910390f35b348015610432575f5ffd5b5061044d60048036038101906104489190615036565b610f2b565b60405161045a9190614a4d565b60405180910390f35b61047d60048036038101906104789190615036565b611021565b60405161048a9190614a4d565b60405180910390f35b34801561049e575f5ffd5b506104b960048036038101906104b49190614a90565b6111fe565b6040516104c69190614a4d565b60405180910390f35b3480156104da575f5ffd5b506104e3611234565b6040516104f091906150fd565b60405180910390f35b348015610504575f5ffd5b5061050d611258565b60405161051a9190614a4d565b60405180910390f35b34801561052e575f5ffd5b5061054960048036038101906105449190614a90565b611260565b6040516105569190615189565b60405180910390f35b34801561056a575f5ffd5b50610585600480360381019061058091906151a2565b611271565b6040516105929190614a1c565b60405180910390f35b3480156105a6575f5ffd5b506105c160048036038101906105bc9190615036565b6112d6565b6040516105ce9190614a4d565b60405180910390f35b3480156105e2575f5ffd5b506105eb611356565b6040516105f89190614b2b565b60405180910390f35b34801561060c575f5ffd5b5061062760048036038101906106229190614a90565b611413565b604051610636939291906151e0565b60405180910390f35b34801561064a575f5ffd5b50610653611446565b6040516106609190614b2b565b60405180910390f35b348015610674575f5ffd5b5061068f600480360381019061068a919061524b565b611483565b60405161069c9190614a4d565b60405180910390f35b3480156106b0575f5ffd5b506106cb60048036038101906106c691906152e2565b6114b2565b6040516106d89190614a4d565b60405180910390f35b3480156106ec575f5ffd5b50610707600480360381019061070291906153c4565b6115a2565b6040516107149190614a4d565b60405180910390f35b348015610728575f5ffd5b50610743600480360381019061073e9190614a90565b61160a565b6040516107509190614a4d565b60405180910390f35b348015610764575f5ffd5b5061077f600480360381019061077a9190615464565b61161d565b60405161078c9190614a4d565b60405180910390f35b3480156107a0575f5ffd5b506107bb60048036038101906107b69190615573565b611683565b6040516107c89190614a4d565b60405180910390f35b3480156107dc575f5ffd5b506107f760048036038101906107f29190615647565b611779565b6040516108049190614a4d565b60405180910390f35b348015610818575f5ffd5b506108216117bf565b6040516108349796959493929190615763565b60405180910390f35b348015610848575f5ffd5b50610863600480360381019061085e91906157e5565b611864565b6040516108709190614a4d565b60405180910390f35b348015610884575f5ffd5b5061088d6118d6565b60405161089a9190615885565b60405180910390f35b3480156108ae575f5ffd5b506108b7611961565b6040516108c49190614a4d565b60405180910390f35b3480156108d8575f5ffd5b506108f360048036038101906108ee919061589e565b611969565b6040516109009190614a4d565b60405180910390f35b348015610914575f5ffd5b5061091d61197e565b60405161092a9190614a4d565b60405180910390f35b34801561093e575f5ffd5b5061095960048036038101906109549190615956565b6119aa565b005b348015610966575f5ffd5b50610981600480360381019061097c9190615036565b6119be565b60405161098e9190614a4d565b60405180910390f35b3480156109a2575f5ffd5b506109bd60048036038101906109b89190614a90565b6119d5565b6040516109ca9190614a1c565b60405180910390f35b3480156109de575f5ffd5b506109f960048036038101906109f49190614a90565b6119e6565b604051610a069190614a4d565b60405180910390f35b348015610a1a575f5ffd5b50610a23611a1c565b604051610a309190614a4d565b60405180910390f35b348015610a44575f5ffd5b50610a5f6004803603810190610a5a9190615981565b611a23565b604051610a6c9190614d88565b60405180910390f35b348015610a80575f5ffd5b50610a9b6004803603810190610a969190614a90565b611aa3565b604051610aa89190614a4d565b60405180910390f35b610acb6004803603810190610ac69190615aa1565b611b0d565b005b348015610ad8575f5ffd5b50610af36004803603810190610aee9190615036565b611b96565b604051610b009190614a4d565b60405180910390f35b348015610b14575f5ffd5b50610b1d611bd0565b604051610b2a9190614b8a565b60405180910390f35b348015610b3e575f5ffd5b50610b47611bf8565b604051610b549190614b2b565b60405180910390f35b348015610b68575f5ffd5b50610b71611c35565b604051610b7e91906150fd565b60405180910390f35b348015610b92575f5ffd5b50610bad6004803603810190610ba89190615b12565b611c59565b604051610bba9190614a4d565b60405180910390f35b348015610bce575f5ffd5b50610be96004803603810190610be49190615b50565b611c74565b604051610bf69190614d88565b60405180910390f35b348015610c0a575f5ffd5b50610c256004803603810190610c209190614a90565b611cf4565b604051610c329190614a4d565b60405180910390f35b348015610c46575f5ffd5b50610c4f611d94565b604051610c5c9190615c3e565b60405180910390f35b5f610c6e611dbb565b905090565b5f7fcdbdfcee000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480610d46575063a8f8a66860e01b7fcdbdfcee00000000000000000000000000000000000000000000000000000000187bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b80610dae57507f4e2312e0000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b80610dbe5750610dbd82611de3565b5b9050919050565b5f6064905090565b610dd5611e4c565b610dde81611f41565b50565b606060038054610df090615c84565b80601f0160208091040260200160405190810160405280929190818152602001828054610e1c90615c84565b8015610e675780601f10610e3e57610100808354040283529160200191610e67565b820191905f5260205f20905b815481529060010190602001808311610e4a57829003601f168201915b5050505050905090565b5f60045f8381526020019081526020015f205f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b5f3073ffffffffffffffffffffffffffffffffffffffff16610ecc610c65565b73ffffffffffffffffffffffffffffffffffffffff1614610f19576040517fe90a651e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63150b7a0260e01b9050949350505050565b5f5f610f39868686866119be565b9050610f4e81610f496004612003565b612027565b505f610f5d8288888888612092565b90505f8165ffffffffffff1614610fe2578060045f8481526020019081526020015f206001015f6101000a81548165ffffffffffff021916908365ffffffffffff1602179055507f9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda28928282604051610fd5929190615ce4565b60405180910390a1611014565b6040517f90884a4600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8192505050949350505050565b5f5f61102f868686866119be565b905061104f8161103f6005612003565b6110496004612003565b17612027565b50600160045f8381526020019081526020015f205f01601e6101000a81548160ff0219169083151502179055503073ffffffffffffffffffffffffffffffffffffffff1661109b610c65565b73ffffffffffffffffffffffffffffffffffffffff1614611154575f5f90505b8651811015611152573073ffffffffffffffffffffffffffffffffffffffff168782815181106110ee576110ed615d0b565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff16036111475761114685828151811061112757611126615d0b565b5b60200260200101518051906020012060056120ab90919063ffffffff16565b5b8060010190506110bb565b505b61116181878787876121a3565b3073ffffffffffffffffffffffffffffffffffffffff16611180610c65565b73ffffffffffffffffffffffffffffffffffffffff16141580156111ab57506111a960056121b7565b155b156111bb576111ba6005612223565b5b7f712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f816040516111ea9190614a4d565b60405180910390a180915050949350505050565b5f60045f8381526020019081526020015f205f0160149054906101000a900465ffffffffffff1665ffffffffffff169050919050565b7f3e83946653575f9a39005e1545185629e92736b7528ab20ca3816f315424a81181565b5f6064905090565b5f61126a82612299565b9050919050565b5f60075f8481526020019081526020015f206003015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b5f5f6112e4868686866119be565b90505f6112ef61244e565b90506112fb8282612455565b61133e5781816040517f8fe5d8a9000000000000000000000000000000000000000000000000000000008152600401611335929190615d38565b60405180910390fd5b61134a878787876124cc565b92505050949350505050565b6060611360611d94565b73ffffffffffffffffffffffffffffffffffffffff16634bf5d7e96040518163ffffffff1660e01b81526004015f60405180830381865afa9250505080156113ca57506040513d5f823e3d601f19601f820116820180604052508101906113c79190615dcd565b60015b61140b576040518060400160405280601d81526020017f6d6f64653d626c6f636b6e756d6265722666726f6d3d64656661756c740000008152509050611410565b809150505b90565b5f5f5f5f60075f8681526020019081526020015f209050805f015481600101548260020154935093509350509193909250565b60606040518060400160405280600181526020017f3100000000000000000000000000000000000000000000000000000000000000815250905090565b5f5f61148d61244e565b90506114a984828560405180602001604052805f8152506124e3565b91505092915050565b5f61150488888888888080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f820116905080830192505050505050508787612502565b61154557856040517f94ab6c0700000000000000000000000000000000000000000000000000000000815260040161153c9190614b8a565b60405180910390fd5b61159588878988888080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f820116905080830192505050505050508761258a565b9050979650505050505050565b5f5f6115ac61244e565b90506115fe87828888888080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f820116905080830192505050505050508761258a565b91505095945050505050565b5f611616600883612694565b9050919050565b5f5f61162761244e565b905061167886828787878080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f820116905080830192505050505050506124e3565b915050949350505050565b5f5f61168d61244e565b905061169981846126fd565b6116da57806040517fd9b395570000000000000000000000000000000000000000000000000000000081526004016116d19190614b8a565b60405180910390fd5b5f6116e3611a1c565b90505f811115611760575f6117138360016116fc6118d6565b6117069190615e41565b65ffffffffffff16611c59565b90508181101561175e578281836040517fc242ee1600000000000000000000000000000000000000000000000000000000815260040161175593929190615e7a565b60405180910390fd5b505b61176d87878787866127e3565b92505050949350505050565b5f60025f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b5f6060805f5f5f60606117d0612a68565b6117d8612aa2565b46305f5f1b5f67ffffffffffffffff8111156117f7576117f6614bd5565b5b6040519080825280602002602001820160405280156118255781602001602082028036833780820191505090505b507f0f00000000000000000000000000000000000000000000000000000000000000959493929190965096509650965096509650965090919293949596565b5f61187185858585612add565b6118b257826040517f94ab6c070000000000000000000000000000000000000000000000000000000081526004016118a99190614b8a565b60405180910390fd5b6118cc85848660405180602001604052805f8152506124e3565b9050949350505050565b5f6118df611d94565b73ffffffffffffffffffffffffffffffffffffffff166391ddadf46040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801561194657506040513d601f19601f820116820180604052508101906119439190615ed9565b60015b61195957611952612b51565b905061195e565b809150505b90565b5f6064905090565b5f611975848484612b60565b90509392505050565b5f6119896008612beb565b79ffffffffffffffffffffffffffffffffffffffffffffffffffff16905090565b6119b2611e4c565b6119bb81612c4d565b50565b5f6119cb85858585611b96565b9050949350505050565b5f6119df82612cea565b9050919050565b5f60045f8381526020019081526020015f206001015f9054906101000a900465ffffffffffff1665ffffffffffff169050919050565b5f5f905090565b5f3073ffffffffffffffffffffffffffffffffffffffff16611a43610c65565b73ffffffffffffffffffffffffffffffffffffffff1614611a90576040517fe90a651e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63bc197c8160e01b905095945050505050565b5f60045f8381526020019081526020015f205f01601a9054906101000a900463ffffffff1663ffffffff1660045f8481526020019081526020015f205f0160149054906101000a900465ffffffffffff16611afe9190615f04565b65ffffffffffff169050919050565b611b15611e4c565b5f5f8573ffffffffffffffffffffffffffffffffffffffff16858585604051611b3f929190615f6b565b5f6040518083038185875af1925050503d805f8114611b79576040519150601f19603f3d011682016040523d82523d5f602084013e611b7e565b606091505b5091509150611b8d8282612cf4565b50505050505050565b5f84848484604051602001611bae9493929190616147565b604051602081830303815290604052805190602001205f1c9050949350505050565b5f60095f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60606040518060400160405280602081526020017f737570706f72743d627261766f2671756f72756d3d666f722c6162737461696e815250905090565b7ff2aad550cf55f045cb27e9c559f9889fdfb6e6cdaa032301d6ea397784ae51d781565b5f611c6c8383611c67612d18565b612b60565b905092915050565b5f3073ffffffffffffffffffffffffffffffffffffffff16611c94610c65565b73ffffffffffffffffffffffffffffffffffffffff1614611ce1576040517fe90a651e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63f23a6e6160e01b905095945050505050565b5f611d8d611d00611d94565b73ffffffffffffffffffffffffffffffffffffffff16638e539e8c846040518263ffffffff1660e01b8152600401611d389190614a4d565b602060405180830381865afa158015611d53573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611d7791906161b3565b611d808461160a565b611d88611961565b612d2e565b9050919050565b5f7f0000000000000000000000000000000000000000000000000000000000000000905090565b5f60095f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b611e5461244e565b73ffffffffffffffffffffffffffffffffffffffff16611e72610c65565b73ffffffffffffffffffffffffffffffffffffffff1614611ed157611e9561244e565b6040517f47096e47000000000000000000000000000000000000000000000000000000008152600401611ec89190614b8a565b60405180910390fd5b3073ffffffffffffffffffffffffffffffffffffffff16611ef0610c65565b73ffffffffffffffffffffffffffffffffffffffff1614611f3f575f611f14612e0d565b604051611f22929190615f6b565b604051809103902090505b80611f386005612e19565b03611f2d57505b565b5f611f4a611961565b905080821115611f935781816040517f243e5445000000000000000000000000000000000000000000000000000000008152600401611f8a9291906161de565b60405180910390fd5b5f611f9c61197e565b9050611fc3611fa96118d6565b611fb285612f47565b6008612fb49092919063ffffffff16565b50507f0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b46339978184604051611ff69291906161de565b60405180910390a1505050565b5f81600781111561201757612016615116565b5b60ff166001901b5f1b9050919050565b5f5f61203284611260565b90505f5f1b8361204183612003565b1603612088578381846040517f31b75e4d00000000000000000000000000000000000000000000000000000000815260040161207f93929190616205565b60405180910390fd5b8091505092915050565b5f6120a08686868686612fcf565b905095945050505050565b5f825f0160109054906101000a90046fffffffffffffffffffffffffffffffff169050825f015f9054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff16600182016fffffffffffffffffffffffffffffffff16036121255761212460416131d5565b5b81836001015f836fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff1681526020019081526020015f208190555060018101835f0160106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff160217905550505050565b6121b085858585856131e6565b5050505050565b5f815f015f9054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff16825f0160109054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff16149050919050565b5f815f015f6101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff1602179055505f815f0160106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555050565b5f5f6122a48361329a565b9050600560078111156122ba576122b9615116565b5b8160078111156122cd576122cc615116565b5b146122db5780915050612449565b5f600a5f8581526020019081526020015f2054905060095f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663584b153e826040518263ffffffff1660e01b815260040161234a91906150fd565b602060405180830381865afa158015612365573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906123899190616264565b1561239957600592505050612449565b60095f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632ab0f529826040518263ffffffff1660e01b81526004016123f391906150fd565b602060405180830381865afa15801561240e573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906124329190616264565b1561244257600792505050612449565b6002925050505b919050565b5f33905090565b5f5f600781111561246957612468615116565b5b61247284611260565b600781111561248457612483615116565b5b1480156124c4575061249583610e71565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16145b905092915050565b5f6124d9858585856133f1565b9050949350505050565b5f6124f8858585856124f3612d18565b61258a565b9050949350505050565b5f61257e856125787f3e83946653575f9a39005e1545185629e92736b7528ab20ca3816f315424a8118a8a8a6125378c6134c8565b8b805190602001208b8051906020012060405160200161255d979695949392919061629e565b6040516020818303038152906040528051906020012061351b565b84613534565b90509695505050505050565b5f61259e866125996001612003565b612027565b505f6125b3866125ad896111fe565b85612b60565b90505f6125c388888885886135e0565b90505f845103612626578673ffffffffffffffffffffffffffffffffffffffff167fb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda489888489604051612619949392919061630b565b60405180910390a261267d565b8673ffffffffffffffffffffffffffffffffffffffff167fe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb8712898884898960405161267495949392919061639d565b60405180910390a25b612686886137d5565b809250505095945050505050565b5f5f5f6126a0856137d8565b9250925050838265ffffffffffff1611156126d5576126d06126c185613867565b866138c090919063ffffffff16565b6126d7565b805b79ffffffffffffffffffffffffffffffffffffffffffffffffffff169250505092915050565b5f5f8251905060348110156127165760019150506127dd565b5f61272484603484036139b0565b90507f2370726f706f7365723d0000000000000000000000000000000000000000000075ffffffffffffffffffffffffffffffffffffffffffff19168175ffffffffffffffffffffffffffffffffffffffffffff19161461278a576001925050506127dd565b5f5f61279a86602a8603866139c0565b915091508115806127d657508673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b9450505050505b92915050565b5f6127f786868686805190602001206119be565b90508451865114158061280c57508351865114155b8061281757505f8651145b15612860578551845186516040517f447b05d0000000000000000000000000000000000000000000000000000000008152600401612857939291906151e0565b60405180910390fd5b5f60045f8381526020019081526020015f205f0160149054906101000a900465ffffffffffff1665ffffffffffff16146128df578061289e82611260565b5f5f1b6040517f31b75e4d0000000000000000000000000000000000000000000000000000000081526004016128d693929190616205565b60405180910390fd5b5f6128e8611258565b6128f06118d6565b65ffffffffffff1661290291906163fc565b90505f61290d610dc5565b90505f60045f8581526020019081526020015f20905084815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061296d83613867565b815f0160146101000a81548165ffffffffffff021916908365ffffffffffff16021790555061299b82613ad1565b815f01601a6101000a81548163ffffffff021916908363ffffffff1602179055507f7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e084868b8b8d5167ffffffffffffffff8111156129fc576129fb614bd5565b5b604051908082528060200260200182016040528015612a2f57816020015b6060815260200190600190039081612a1a5790505b508c89898b612a3e91906163fc565b8e604051612a5499989796959493929190616532565b60405180910390a150505095945050505050565b6060612a9d5f7f0000000000000000000000000000000000000000000000000000000000000000613b2890919063ffffffff16565b905090565b6060612ad860017f0000000000000000000000000000000000000000000000000000000000000000613b2890919063ffffffff16565b905090565b5f612b4783612b417ff2aad550cf55f045cb27e9c559f9889fdfb6e6cdaa032301d6ea397784ae51d7888888612b128a6134c8565b604051602001612b269594939291906165e0565b6040516020818303038152906040528051906020012061351b565b84613534565b9050949350505050565b5f612b5b43613867565b905090565b5f612b69611d94565b73ffffffffffffffffffffffffffffffffffffffff16633a46b1a885856040518363ffffffff1660e01b8152600401612ba3929190616631565b602060405180830381865afa158015612bbe573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612be291906161b3565b90509392505050565b5f5f825f018054905090505f8114612c4357612c15835f01600183612c109190616658565b613bd5565b5f0160069054906101000a900479ffffffffffffffffffffffffffffffffffffffffffffffffffff16612c45565b5f5b915050919050565b7f08f74ea46ef7894f65eabfb5e6e695de773a000b47c529ab559178069b22640160095f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1682604051612c9f92919061668b565b60405180910390a18060095f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f60019050919050565b606082612d0957612d0482613be7565b612d11565b819050612d12565b5b92915050565b606060405180602001604052805f815250905090565b5f5f5f612d3b8686613c2a565b915091505f8203612d6057838181612d5657612d556166b2565b5b0492505050612e06565b818411612d7f57612d7e612d795f861460126011613c47565b6131d5565b5b5f8486880990508181118303925080820391505f855f038616905080860495508083049250600181825f0304019050808402831792505f600287600302189050808702600203810290508087026002038102905080870260020381029050808702600203810290508087026002038102905080870260020381029050808402955050505050505b9392505050565b365f5f36915091509091565b5f5f825f015f9054906101000a90046fffffffffffffffffffffffffffffffff169050825f0160109054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff16816fffffffffffffffffffffffffffffffff1603612e9157612e9060316131d5565b5b826001015f826fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff1681526020019081526020015f20549150826001015f826fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff1681526020019081526020015f205f905560018101835f015f6101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555050919050565b5f79ffffffffffffffffffffffffffffffffffffffffffffffffffff8016821115612fac5760d0826040517f6dfcc650000000000000000000000000000000000000000000000000000000008152600401612fa3929190616718565b60405180910390fd5b819050919050565b5f5f612fc3855f018585613c60565b91509150935093915050565b5f5f60095f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f27a0c926040518163ffffffff1660e01b8152600401602060405180830381865afa15801561303b573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061305f91906161b3565b90505f61306b84613f54565b905060095f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b1c5f4278888885f866040518663ffffffff1660e01b81526004016130cf959493929190616783565b602060405180830381865afa1580156130ea573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061310e91906167fd565b600a5f8a81526020019081526020015f208190555060095f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638f2a0bb08888885f86886040518763ffffffff1660e01b815260040161318796959493929190616828565b5f604051808303815f87803b15801561319e575f5ffd5b505af11580156131b0573d5f5f3e3d5ffd5b505050506131c882426131c391906163fc565b613867565b9250505095945050505050565b634e487b715f52806020526024601cfd5b60095f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e38335e5348686865f61323188613f54565b6040518763ffffffff1660e01b8152600401613251959493929190616783565b5f604051808303818588803b158015613268575f5ffd5b505af115801561327a573d5f5f3e3d5ffd5b5050505050600a5f8681526020019081526020015f205f90555050505050565b5f5f60045f8481526020019081526020015f2090505f815f01601e9054906101000a900460ff1690505f825f01601f9054906101000a900460ff16905081156132e957600793505050506133ec565b80156132fb57600293505050506133ec565b5f613305866111fe565b90505f810361334b57856040517f6ad060750000000000000000000000000000000000000000000000000000000081526004016133429190614a4d565b60405180910390fd5b5f6133546118d6565b65ffffffffffff169050808210613372575f955050505050506133ec565b5f61337c88611aa3565b905081811061339457600196505050505050506133ec565b61339d88613f71565b15806133af57506133ad88613fb5565b155b156133c357600396505050505050506133ec565b5f6133cd896119e6565b036133e157600496505050505050506133ec565b600596505050505050505b919050565b5f5f6133ff86868686613fdc565b90505f600a5f8381526020019081526020015f205490505f5f1b81146134bb5760095f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c4d252f5826040518263ffffffff1660e01b815260040161347991906150fd565b5f604051808303815f87803b158015613490575f5ffd5b505af11580156134a2573d5f5f3e3d5ffd5b50505050600a5f8381526020019081526020015f205f90555b8192505050949350505050565b5f60025f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f815480929190600101919050559050919050565b5f61352d6135276140be565b83614174565b9050919050565b5f5f8473ffffffffffffffffffffffffffffffffffffffff163b036135cb575f5f61355f85856141b4565b50915091505f600381111561357757613576615116565b5b81600381111561358a57613589615116565b5b1480156135c257508573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16145b925050506135d9565b6135d6848484614209565b90505b9392505050565b5f5f60075f8881526020019081526020015f209050806003015f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff161561368357856040517f71c6af4900000000000000000000000000000000000000000000000000000000815260040161367a9190614b8a565b60405180910390fd5b6001816003015f8873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055505f60028111156136ed576136ec615116565b5b60ff168560ff16036137175783815f015f82825461370b91906163fc565b925050819055506137c8565b6001600281111561372b5761372a615116565b5b60ff168560ff16036137565783816001015f82825461374a91906163fc565b925050819055506137c7565b60028081111561376957613768615116565b5b60ff168560ff16036137945783816002015f82825461378891906163fc565b925050819055506137c6565b6040517f06b337c200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5b5b8391505095945050505050565b50565b5f5f5f5f845f018054905090505f81036137fb575f5f5f93509350935050613860565b5f613814865f0160018461380f9190616658565b613bd5565b90506001815f015f9054906101000a900465ffffffffffff16825f0160069054906101000a900479ffffffffffffffffffffffffffffffffffffffffffffffffffff1694509450945050505b9193909250565b5f65ffffffffffff80168211156138b8576030826040517f6dfcc6500000000000000000000000000000000000000000000000000000000081526004016138af9291906168d5565b60405180910390fd5b819050919050565b5f5f835f018054905090505f5f90505f8290506005831115613944575f6138e684614328565b846138f19190616658565b90506138ff875f0182613bd5565b5f015f9054906101000a900465ffffffffffff1665ffffffffffff168665ffffffffffff16101561393257809150613942565b60018161393f91906163fc565b92505b505b5f613953875f018785856144c2565b90505f81146139a257613974875f0160018361396f9190616658565b613bd5565b5f0160069054906101000a900479ffffffffffffffffffffffffffffffffffffffffffffffffffff166139a4565b5f5b94505050505092915050565b5f81602084010151905092915050565b5f5f84518311806139d057508284115b156139e0575f5f91509150613ac9565b5f6001856139ee91906163fc565b84118015613a6657507f30780000000000000000000000000000000000000000000000000000000000007dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916613a438787614537565b7dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b90505f6002613a76831515614547565b613a8091906168fc565b6028613a8c91906163fc565b9050808686613a9b9190616658565b03613ac0575f5f613aad898989614552565b9150915081819550955050505050613ac9565b5f5f9350935050505b935093915050565b5f63ffffffff8016821115613b20576020826040517f6dfcc650000000000000000000000000000000000000000000000000000000008152600401613b17929190616976565b60405180910390fd5b819050919050565b606060ff5f1b8314613b4457613b3d83614680565b9050613bcf565b818054613b5090615c84565b80601f0160208091040260200160405190810160405280929190818152602001828054613b7c90615c84565b8015613bc75780601f10613b9e57610100808354040283529160200191613bc7565b820191905f5260205f20905b815481529060010190602001808311613baa57829003601f168201915b505050505090505b92915050565b5f825f528160205f2001905092915050565b5f81511115613bf857805160208201fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5f5f198385098385029150818110828203039250509250929050565b5f613c5184614547565b82841802821890509392505050565b5f5f5f858054905090505f811115613e6c575f613c8987600184613c849190616658565b613bd5565b90505f815f015f9054906101000a900465ffffffffffff1690505f825f0160069054906101000a900479ffffffffffffffffffffffffffffffffffffffffffffffffffff1690508765ffffffffffff168265ffffffffffff161115613d1a576040517f2520601d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8765ffffffffffff168265ffffffffffff1603613d845786835f0160066101000a81548179ffffffffffffffffffffffffffffffffffffffffffffffffffff021916908379ffffffffffffffffffffffffffffffffffffffffffffffffffff160217905550613e5d565b8860405180604001604052808a65ffffffffffff1681526020018979ffffffffffffffffffffffffffffffffffffffffffffffffffff16815250908060018154018082558091505060019003905f5260205f20015f909190919091505f820151815f015f6101000a81548165ffffffffffff021916908365ffffffffffff1602179055506020820151815f0160066101000a81548179ffffffffffffffffffffffffffffffffffffffffffffffffffff021916908379ffffffffffffffffffffffffffffffffffffffffffffffffffff16021790555050505b80879550955050505050613f4c565b8560405180604001604052808765ffffffffffff1681526020018679ffffffffffffffffffffffffffffffffffffffffffffffffffff16815250908060018154018082558091505060019003905f5260205f20015f909190919091505f820151815f015f6101000a81548165ffffffffffff021916908365ffffffffffff1602179055506020820151815f0160066101000a81548179ffffffffffffffffffffffffffffffffffffffffffffffffffff021916908379ffffffffffffffffffffffffffffffffffffffffffffffffffff16021790555050505f8492509250505b935093915050565b5f813060601b6bffffffffffffffffffffffff1916189050919050565b5f5f60075f8481526020019081526020015f20905080600201548160010154613f9a91906163fc565b613fab613fa6856111fe565b611cf4565b1115915050919050565b5f5f60075f8481526020019081526020015f209050805f0154816001015411915050919050565b5f5f613fea868686866119be565b905061404e81613ffa6007612003565b6140046006612003565b61400e6002612003565b60018060078081111561402457614023615116565b5b61402e919061699d565b600261403a9190616b00565b6140449190616658565b5f1b181818612027565b50600160045f8381526020019081526020015f205f01601f6101000a81548160ff0219169083151502179055507f789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c816040516140aa9190614a4d565b60405180910390a180915050949350505050565b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1614801561413957507f000000000000000000000000000000000000000000000000000000000000000046145b15614166577f00000000000000000000000000000000000000000000000000000000000000009050614171565b61416e6146f2565b90505b90565b5f6040517f190100000000000000000000000000000000000000000000000000000000000081528360028201528260228201526042812091505092915050565b5f5f5f60418451036141f4575f5f5f602087015192506040870151915060608701515f1a90506141e688828585614787565b955095509550505050614202565b5f600285515f1b9250925092505b9250925092565b5f5f5f8573ffffffffffffffffffffffffffffffffffffffff168585604051602401614236929190616b4a565b604051602081830303815290604052631626ba7e60e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516142889190616ba8565b5f60405180830381855afa9150503d805f81146142c0576040519150601f19603f3d011682016040523d82523d5f602084013e6142c5565b606091505b50915091508180156142d957506020815110155b801561431d5750631626ba7e60e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168180602001905181019061431b91906167fd565b145b925050509392505050565b5f60018211614339578190506144bd565b5f8290505f60019050700100000000000000000000000000000000821061436957608082901c9150604081901b90505b68010000000000000000821061438857604082901c9150602081901b90505b64010000000082106143a357602082901c9150601081901b90505b6201000082106143bc57601082901c9150600881901b90505b61010082106143d457600882901c9150600481901b90505b601082106143eb57600482901c9150600281901b90505b600482106143fb57600181901b90505b600181600302901c90506001818581614417576144166166b2565b5b048201901c905060018185816144305761442f6166b2565b5b048201901c90506001818581614449576144486166b2565b5b048201901c90506001818581614462576144616166b2565b5b048201901c9050600181858161447b5761447a6166b2565b5b048201901c90506001818581614494576144936166b2565b5b048201901c90506144b68185816144ae576144ad6166b2565b5b048211614547565b8103925050505b919050565b5f5b8183101561452c575f6144d7848461486e565b90508465ffffffffffff166144ec8783613bd5565b5f015f9054906101000a900465ffffffffffff1665ffffffffffff16111561451657809250614526565b60018161452391906163fc565b93505b506144c4565b819050949350505050565b5f81602084010151905092915050565b5f8115159050919050565b5f5f5f8590505f60018661456691906163fc565b851180156145de57507f30780000000000000000000000000000000000000000000000000000000000007dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166145bb8388614537565b7dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b90505f60026145ee831515614547565b6145f891906168fc565b90505f5f90505f828961460b91906163fc565b90505b8781101561466b575f6146296146248784614537565b614893565b9050600f8160ff161115614648575f5f97509750505050505050614678565b60108361465591906168fc565b92508060ff16830192505080600101905061460e565b5060018195509550505050505b935093915050565b60605f61468c83614923565b90505f602067ffffffffffffffff8111156146aa576146a9614bd5565b5b6040519080825280601f01601f1916602001820160405280156146dc5781602001600182028036833780820191505090505b5090508181528360208201528092505050919050565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f7f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000463060405160200161476c959493929190616bbe565b60405160208183030381529060405280519060200120905090565b5f5f5f7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0845f1c11156147c3575f600385925092509250614864565b5f6001888888886040515f81526020016040526040516147e69493929190616c0f565b6020604051602081039080840390855afa158015614806573d5f5f3e3d5ffd5b5050506020604051035190505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603614857575f60015f5f1b93509350935050614864565b805f5f5f1b935093509350505b9450945094915050565b5f600282841861487e9190616c52565b82841661488b91906163fc565b905092915050565b5f5f8260f81c9050602f8160ff161180156148b15750603a8160ff16105b156148c157603081039050614919565b60608160ff161180156148d7575060678160ff16105b156148e757605781039050614918565b60408160ff161180156148fd575060478160ff16105b1561490d57603781039050614917565b60ff91505061491e565b5b5b809150505b919050565b5f5f60ff835f1c169050601f811115614968576040517fb3512b0c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80915050919050565b5f604051905090565b5f5ffd5b5f5ffd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6149b681614982565b81146149c0575f5ffd5b50565b5f813590506149d1816149ad565b92915050565b5f602082840312156149ec576149eb61497a565b5b5f6149f9848285016149c3565b91505092915050565b5f8115159050919050565b614a1681614a02565b82525050565b5f602082019050614a2f5f830184614a0d565b92915050565b5f819050919050565b614a4781614a35565b82525050565b5f602082019050614a605f830184614a3e565b92915050565b614a6f81614a35565b8114614a79575f5ffd5b50565b5f81359050614a8a81614a66565b92915050565b5f60208284031215614aa557614aa461497a565b5b5f614ab284828501614a7c565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f614afd82614abb565b614b078185614ac5565b9350614b17818560208601614ad5565b614b2081614ae3565b840191505092915050565b5f6020820190508181035f830152614b438184614af3565b905092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f614b7482614b4b565b9050919050565b614b8481614b6a565b82525050565b5f602082019050614b9d5f830184614b7b565b92915050565b614bac81614b6a565b8114614bb6575f5ffd5b50565b5f81359050614bc781614ba3565b92915050565b5f5ffd5b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b614c0b82614ae3565b810181811067ffffffffffffffff82111715614c2a57614c29614bd5565b5b80604052505050565b5f614c3c614971565b9050614c488282614c02565b919050565b5f67ffffffffffffffff821115614c6757614c66614bd5565b5b614c7082614ae3565b9050602081019050919050565b828183375f83830152505050565b5f614c9d614c9884614c4d565b614c33565b905082815260208101848484011115614cb957614cb8614bd1565b5b614cc4848285614c7d565b509392505050565b5f82601f830112614ce057614cdf614bcd565b5b8135614cf0848260208601614c8b565b91505092915050565b5f5f5f5f60808587031215614d1157614d1061497a565b5b5f614d1e87828801614bb9565b9450506020614d2f87828801614bb9565b9350506040614d4087828801614a7c565b925050606085013567ffffffffffffffff811115614d6157614d6061497e565b5b614d6d87828801614ccc565b91505092959194509250565b614d8281614982565b82525050565b5f602082019050614d9b5f830184614d79565b92915050565b5f67ffffffffffffffff821115614dbb57614dba614bd5565b5b602082029050602081019050919050565b5f5ffd5b5f614de2614ddd84614da1565b614c33565b90508083825260208201905060208402830185811115614e0557614e04614dcc565b5b835b81811015614e2e5780614e1a8882614bb9565b845260208401935050602081019050614e07565b5050509392505050565b5f82601f830112614e4c57614e4b614bcd565b5b8135614e5c848260208601614dd0565b91505092915050565b5f67ffffffffffffffff821115614e7f57614e7e614bd5565b5b602082029050602081019050919050565b5f614ea2614e9d84614e65565b614c33565b90508083825260208201905060208402830185811115614ec557614ec4614dcc565b5b835b81811015614eee5780614eda8882614a7c565b845260208401935050602081019050614ec7565b5050509392505050565b5f82601f830112614f0c57614f0b614bcd565b5b8135614f1c848260208601614e90565b91505092915050565b5f67ffffffffffffffff821115614f3f57614f3e614bd5565b5b602082029050602081019050919050565b5f614f62614f5d84614f25565b614c33565b90508083825260208201905060208402830185811115614f8557614f84614dcc565b5b835b81811015614fcc57803567ffffffffffffffff811115614faa57614fa9614bcd565b5b808601614fb78982614ccc565b85526020850194505050602081019050614f87565b5050509392505050565b5f82601f830112614fea57614fe9614bcd565b5b8135614ffa848260208601614f50565b91505092915050565b5f819050919050565b61501581615003565b811461501f575f5ffd5b50565b5f813590506150308161500c565b92915050565b5f5f5f5f6080858703121561504e5761504d61497a565b5b5f85013567ffffffffffffffff81111561506b5761506a61497e565b5b61507787828801614e38565b945050602085013567ffffffffffffffff8111156150985761509761497e565b5b6150a487828801614ef8565b935050604085013567ffffffffffffffff8111156150c5576150c461497e565b5b6150d187828801614fd6565b92505060606150e287828801615022565b91505092959194509250565b6150f781615003565b82525050565b5f6020820190506151105f8301846150ee565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b6008811061515457615153615116565b5b50565b5f81905061516482615143565b919050565b5f61517382615157565b9050919050565b61518381615169565b82525050565b5f60208201905061519c5f83018461517a565b92915050565b5f5f604083850312156151b8576151b761497a565b5b5f6151c585828601614a7c565b92505060206151d685828601614bb9565b9150509250929050565b5f6060820190506151f35f830186614a3e565b6152006020830185614a3e565b61520d6040830184614a3e565b949350505050565b5f60ff82169050919050565b61522a81615215565b8114615234575f5ffd5b50565b5f8135905061524581615221565b92915050565b5f5f604083850312156152615761526061497a565b5b5f61526e85828601614a7c565b925050602061527f85828601615237565b9150509250929050565b5f5ffd5b5f5f83601f8401126152a2576152a1614bcd565b5b8235905067ffffffffffffffff8111156152bf576152be615289565b5b6020830191508360018202830111156152db576152da614dcc565b5b9250929050565b5f5f5f5f5f5f5f60c0888a0312156152fd576152fc61497a565b5b5f61530a8a828b01614a7c565b975050602061531b8a828b01615237565b965050604061532c8a828b01614bb9565b955050606088013567ffffffffffffffff81111561534d5761534c61497e565b5b6153598a828b0161528d565b9450945050608088013567ffffffffffffffff81111561537c5761537b61497e565b5b6153888a828b01614ccc565b92505060a088013567ffffffffffffffff8111156153a9576153a861497e565b5b6153b58a828b01614ccc565b91505092959891949750929550565b5f5f5f5f5f608086880312156153dd576153dc61497a565b5b5f6153ea88828901614a7c565b95505060206153fb88828901615237565b945050604086013567ffffffffffffffff81111561541c5761541b61497e565b5b6154288882890161528d565b9350935050606086013567ffffffffffffffff81111561544b5761544a61497e565b5b61545788828901614ccc565b9150509295509295909350565b5f5f5f5f6060858703121561547c5761547b61497a565b5b5f61548987828801614a7c565b945050602061549a87828801615237565b935050604085013567ffffffffffffffff8111156154bb576154ba61497e565b5b6154c78782880161528d565b925092505092959194509250565b5f67ffffffffffffffff8211156154ef576154ee614bd5565b5b6154f882614ae3565b9050602081019050919050565b5f615517615512846154d5565b614c33565b90508281526020810184848401111561553357615532614bd1565b5b61553e848285614c7d565b509392505050565b5f82601f83011261555a57615559614bcd565b5b813561556a848260208601615505565b91505092915050565b5f5f5f5f6080858703121561558b5761558a61497a565b5b5f85013567ffffffffffffffff8111156155a8576155a761497e565b5b6155b487828801614e38565b945050602085013567ffffffffffffffff8111156155d5576155d461497e565b5b6155e187828801614ef8565b935050604085013567ffffffffffffffff8111156156025761560161497e565b5b61560e87828801614fd6565b925050606085013567ffffffffffffffff81111561562f5761562e61497e565b5b61563b87828801615546565b91505092959194509250565b5f6020828403121561565c5761565b61497a565b5b5f61566984828501614bb9565b91505092915050565b5f7fff0000000000000000000000000000000000000000000000000000000000000082169050919050565b6156a681615672565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b6156de81614a35565b82525050565b5f6156ef83836156d5565b60208301905092915050565b5f602082019050919050565b5f615711826156ac565b61571b81856156b6565b9350615726836156c6565b805f5b8381101561575657815161573d88826156e4565b9750615748836156fb565b925050600181019050615729565b5085935050505092915050565b5f60e0820190506157765f83018a61569d565b81810360208301526157888189614af3565b9050818103604083015261579c8188614af3565b90506157ab6060830187614a3e565b6157b86080830186614b7b565b6157c560a08301856150ee565b81810360c08301526157d78184615707565b905098975050505050505050565b5f5f5f5f608085870312156157fd576157fc61497a565b5b5f61580a87828801614a7c565b945050602061581b87828801615237565b935050604061582c87828801614bb9565b925050606085013567ffffffffffffffff81111561584d5761584c61497e565b5b61585987828801614ccc565b91505092959194509250565b5f65ffffffffffff82169050919050565b61587f81615865565b82525050565b5f6020820190506158985f830184615876565b92915050565b5f5f5f606084860312156158b5576158b461497a565b5b5f6158c286828701614bb9565b93505060206158d386828701614a7c565b925050604084013567ffffffffffffffff8111156158f4576158f361497e565b5b61590086828701614ccc565b9150509250925092565b5f61591482614b4b565b9050919050565b5f6159258261590a565b9050919050565b6159358161591b565b811461593f575f5ffd5b50565b5f813590506159508161592c565b92915050565b5f6020828403121561596b5761596a61497a565b5b5f61597884828501615942565b91505092915050565b5f5f5f5f5f60a0868803121561599a5761599961497a565b5b5f6159a788828901614bb9565b95505060206159b888828901614bb9565b945050604086013567ffffffffffffffff8111156159d9576159d861497e565b5b6159e588828901614ef8565b935050606086013567ffffffffffffffff811115615a0657615a0561497e565b5b615a1288828901614ef8565b925050608086013567ffffffffffffffff811115615a3357615a3261497e565b5b615a3f88828901614ccc565b9150509295509295909350565b5f5f83601f840112615a6157615a60614bcd565b5b8235905067ffffffffffffffff811115615a7e57615a7d615289565b5b602083019150836001820283011115615a9a57615a99614dcc565b5b9250929050565b5f5f5f5f60608587031215615ab957615ab861497a565b5b5f615ac687828801614bb9565b9450506020615ad787828801614a7c565b935050604085013567ffffffffffffffff811115615af857615af761497e565b5b615b0487828801615a4c565b925092505092959194509250565b5f5f60408385031215615b2857615b2761497a565b5b5f615b3585828601614bb9565b9250506020615b4685828601614a7c565b9150509250929050565b5f5f5f5f5f60a08688031215615b6957615b6861497a565b5b5f615b7688828901614bb9565b9550506020615b8788828901614bb9565b9450506040615b9888828901614a7c565b9350506060615ba988828901614a7c565b925050608086013567ffffffffffffffff811115615bca57615bc961497e565b5b615bd688828901614ccc565b9150509295509295909350565b5f819050919050565b5f615c06615c01615bfc84614b4b565b615be3565b614b4b565b9050919050565b5f615c1782615bec565b9050919050565b5f615c2882615c0d565b9050919050565b615c3881615c1e565b82525050565b5f602082019050615c515f830184615c2f565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680615c9b57607f821691505b602082108103615cae57615cad615c57565b5b50919050565b5f615cce615cc9615cc484615865565b615be3565b614a35565b9050919050565b615cde81615cb4565b82525050565b5f604082019050615cf75f830185614a3e565b615d046020830184615cd5565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f604082019050615d4b5f830185614a3e565b615d586020830184614b7b565b9392505050565b5f615d71615d6c846154d5565b614c33565b905082815260208101848484011115615d8d57615d8c614bd1565b5b615d98848285614ad5565b509392505050565b5f82601f830112615db457615db3614bcd565b5b8151615dc4848260208601615d5f565b91505092915050565b5f60208284031215615de257615de161497a565b5b5f82015167ffffffffffffffff811115615dff57615dfe61497e565b5b615e0b84828501615da0565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f615e4b82615865565b9150615e5683615865565b9250828203905065ffffffffffff811115615e7457615e73615e14565b5b92915050565b5f606082019050615e8d5f830186614b7b565b615e9a6020830185614a3e565b615ea76040830184614a3e565b949350505050565b615eb881615865565b8114615ec2575f5ffd5b50565b5f81519050615ed381615eaf565b92915050565b5f60208284031215615eee57615eed61497a565b5b5f615efb84828501615ec5565b91505092915050565b5f615f0e82615865565b9150615f1983615865565b9250828201905065ffffffffffff811115615f3757615f36615e14565b5b92915050565b5f81905092915050565b5f615f528385615f3d565b9350615f5f838584614c7d565b82840190509392505050565b5f615f77828486615f47565b91508190509392505050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b615fb581614b6a565b82525050565b5f615fc68383615fac565b60208301905092915050565b5f602082019050919050565b5f615fe882615f83565b615ff28185615f8d565b9350615ffd83615f9d565b805f5b8381101561602d5781516160148882615fbb565b975061601f83615fd2565b925050600181019050616000565b5085935050505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f81519050919050565b5f82825260208201905092915050565b5f61608782616063565b616091818561606d565b93506160a1818560208601614ad5565b6160aa81614ae3565b840191505092915050565b5f6160c0838361607d565b905092915050565b5f602082019050919050565b5f6160de8261603a565b6160e88185616044565b9350836020820285016160fa85616054565b805f5b85811015616135578484038952815161611685826160b5565b9450616121836160c8565b925060208a019950506001810190506160fd565b50829750879550505050505092915050565b5f6080820190508181035f83015261615f8187615fde565b905081810360208301526161738186615707565b9050818103604083015261618781856160d4565b905061619660608301846150ee565b95945050505050565b5f815190506161ad81614a66565b92915050565b5f602082840312156161c8576161c761497a565b5b5f6161d58482850161619f565b91505092915050565b5f6040820190506161f15f830185614a3e565b6161fe6020830184614a3e565b9392505050565b5f6060820190506162185f830186614a3e565b616225602083018561517a565b61623260408301846150ee565b949350505050565b61624381614a02565b811461624d575f5ffd5b50565b5f8151905061625e8161623a565b92915050565b5f602082840312156162795761627861497a565b5b5f61628684828501616250565b91505092915050565b61629881615215565b82525050565b5f60e0820190506162b15f83018a6150ee565b6162be6020830189614a3e565b6162cb604083018861628f565b6162d86060830187614b7b565b6162e56080830186614a3e565b6162f260a08301856150ee565b6162ff60c08301846150ee565b98975050505050505050565b5f60808201905061631e5f830187614a3e565b61632b602083018661628f565b6163386040830185614a3e565b818103606083015261634a8184614af3565b905095945050505050565b5f82825260208201905092915050565b5f61636f82616063565b6163798185616355565b9350616389818560208601614ad5565b61639281614ae3565b840191505092915050565b5f60a0820190506163b05f830188614a3e565b6163bd602083018761628f565b6163ca6040830186614a3e565b81810360608301526163dc8185614af3565b905081810360808301526163f08184616365565b90509695505050505050565b5f61640682614a35565b915061641183614a35565b925082820190508082111561642957616428615e14565b5b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f82825260208201905092915050565b5f61647282614abb565b61647c8185616458565b935061648c818560208601614ad5565b61649581614ae3565b840191505092915050565b5f6164ab8383616468565b905092915050565b5f602082019050919050565b5f6164c98261642f565b6164d38185616439565b9350836020820285016164e585616449565b805f5b85811015616520578484038952815161650185826164a0565b945061650c836164b3565b925060208a019950506001810190506164e8565b50829750879550505050505092915050565b5f610120820190506165465f83018c614a3e565b616553602083018b614b7b565b8181036040830152616565818a615fde565b905081810360608301526165798189615707565b9050818103608083015261658d81886164bf565b905081810360a08301526165a181876160d4565b90506165b060c0830186614a3e565b6165bd60e0830185614a3e565b8181036101008301526165d08184614af3565b90509a9950505050505050505050565b5f60a0820190506165f35f8301886150ee565b6166006020830187614a3e565b61660d604083018661628f565b61661a6060830185614b7b565b6166276080830184614a3e565b9695505050505050565b5f6040820190506166445f830185614b7b565b6166516020830184614a3e565b9392505050565b5f61666282614a35565b915061666d83614a35565b925082820390508181111561668557616684615e14565b5b92915050565b5f60408201905061669e5f830185614b7b565b6166ab6020830184614b7b565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f819050919050565b5f6167026166fd6166f8846166df565b615be3565b615215565b9050919050565b616712816166e8565b82525050565b5f60408201905061672b5f830185616709565b6167386020830184614a3e565b9392505050565b5f819050919050565b5f815f1b9050919050565b5f61676d6167686167638461673f565b616748565b615003565b9050919050565b61677d81616753565b82525050565b5f60a0820190508181035f83015261679b8188615fde565b905081810360208301526167af8187615707565b905081810360408301526167c381866160d4565b90506167d26060830185616774565b6167df60808301846150ee565b9695505050505050565b5f815190506167f78161500c565b92915050565b5f602082840312156168125761681161497a565b5b5f61681f848285016167e9565b91505092915050565b5f60c0820190508181035f8301526168408189615fde565b905081810360208301526168548188615707565b9050818103604083015261686881876160d4565b90506168776060830186616774565b61688460808301856150ee565b61689160a0830184614a3e565b979650505050505050565b5f819050919050565b5f6168bf6168ba6168b58461689c565b615be3565b615215565b9050919050565b6168cf816168a5565b82525050565b5f6040820190506168e85f8301856168c6565b6168f56020830184614a3e565b9392505050565b5f61690682614a35565b915061691183614a35565b925082820261691f81614a35565b9150828204841483151761693657616935615e14565b5b5092915050565b5f819050919050565b5f61696061695b6169568461693d565b615be3565b615215565b9050919050565b61697081616946565b82525050565b5f6040820190506169895f830185616967565b6169966020830184614a3e565b9392505050565b5f6169a782615215565b91506169b283615215565b9250828201905060ff8111156169cb576169ca615e14565b5b92915050565b5f8160011c9050919050565b5f5f8291508390505b6001851115616a2657808604811115616a0257616a01615e14565b5b6001851615616a115780820291505b8081029050616a1f856169d1565b94506169e6565b94509492505050565b5f82616a3e5760019050616af9565b81616a4b575f9050616af9565b8160018114616a615760028114616a6b57616a9a565b6001915050616af9565b60ff841115616a7d57616a7c615e14565b5b8360020a915084821115616a9457616a93615e14565b5b50616af9565b5060208310610133831016604e8410600b8410161715616acf5782820a905083811115616aca57616ac9615e14565b5b616af9565b616adc84848460016169dd565b92509050818404811115616af357616af2615e14565b5b81810290505b9392505050565b5f616b0a82614a35565b9150616b1583615215565b9250616b427fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484616a2f565b905092915050565b5f604082019050616b5d5f8301856150ee565b8181036020830152616b6f8184616365565b90509392505050565b5f616b8282616063565b616b8c8185615f3d565b9350616b9c818560208601614ad5565b80840191505092915050565b5f616bb38284616b78565b915081905092915050565b5f60a082019050616bd15f8301886150ee565b616bde60208301876150ee565b616beb60408301866150ee565b616bf86060830185614a3e565b616c056080830184614b7b565b9695505050505050565b5f608082019050616c225f8301876150ee565b616c2f602083018661628f565b616c3c60408301856150ee565b616c4960608301846150ee565b95945050505050565b5f616c5c82614a35565b9150616c6783614a35565b925082616c7757616c766166b2565b5b82820490509291505056fea26469706673582212205d453077b532bacaff914579122104e6bf63a62edc70a6054736de4339ad62e964736f6c634300081e0033",
}

// MyGovernor is an auto generated Go binding around an Ethereum contract.
type MyGovernor struct {
	abi abi.ABI
}

// NewMyGovernor creates a new instance of MyGovernor.
func NewMyGovernor() *MyGovernor {
	parsed, err := MyGovernorMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &MyGovernor{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *MyGovernor) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address _token, address _timelock) returns()
func (myGovernor *MyGovernor) PackConstructor(_token common.Address, _timelock common.Address) []byte {
	enc, err := myGovernor.abi.Pack("", _token, _timelock)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBALLOTTYPEHASH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdeaaa7cc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (myGovernor *MyGovernor) PackBALLOTTYPEHASH() []byte {
	enc, err := myGovernor.abi.Pack("BALLOT_TYPEHASH")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBALLOTTYPEHASH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdeaaa7cc.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (myGovernor *MyGovernor) TryPackBALLOTTYPEHASH() ([]byte, error) {
	return myGovernor.abi.Pack("BALLOT_TYPEHASH")
}

// UnpackBALLOTTYPEHASH is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdeaaa7cc.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (myGovernor *MyGovernor) UnpackBALLOTTYPEHASH(data []byte) ([32]byte, error) {
	out, err := myGovernor.abi.Unpack("BALLOT_TYPEHASH", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackCLOCKMODE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4bf5d7e9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (myGovernor *MyGovernor) PackCLOCKMODE() []byte {
	enc, err := myGovernor.abi.Pack("CLOCK_MODE")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCLOCKMODE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4bf5d7e9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (myGovernor *MyGovernor) TryPackCLOCKMODE() ([]byte, error) {
	return myGovernor.abi.Pack("CLOCK_MODE")
}

// UnpackCLOCKMODE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (myGovernor *MyGovernor) UnpackCLOCKMODE(data []byte) (string, error) {
	out, err := myGovernor.abi.Unpack("CLOCK_MODE", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackCOUNTINGMODE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdd4e2ba5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function COUNTING_MODE() pure returns(string)
func (myGovernor *MyGovernor) PackCOUNTINGMODE() []byte {
	enc, err := myGovernor.abi.Pack("COUNTING_MODE")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCOUNTINGMODE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdd4e2ba5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function COUNTING_MODE() pure returns(string)
func (myGovernor *MyGovernor) TryPackCOUNTINGMODE() ([]byte, error) {
	return myGovernor.abi.Pack("COUNTING_MODE")
}

// UnpackCOUNTINGMODE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdd4e2ba5.
//
// Solidity: function COUNTING_MODE() pure returns(string)
func (myGovernor *MyGovernor) UnpackCOUNTINGMODE(data []byte) (string, error) {
	out, err := myGovernor.abi.Unpack("COUNTING_MODE", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackEXTENDEDBALLOTTYPEHASH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2fe3e261.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function EXTENDED_BALLOT_TYPEHASH() view returns(bytes32)
func (myGovernor *MyGovernor) PackEXTENDEDBALLOTTYPEHASH() []byte {
	enc, err := myGovernor.abi.Pack("EXTENDED_BALLOT_TYPEHASH")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackEXTENDEDBALLOTTYPEHASH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2fe3e261.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function EXTENDED_BALLOT_TYPEHASH() view returns(bytes32)
func (myGovernor *MyGovernor) TryPackEXTENDEDBALLOTTYPEHASH() ([]byte, error) {
	return myGovernor.abi.Pack("EXTENDED_BALLOT_TYPEHASH")
}

// UnpackEXTENDEDBALLOTTYPEHASH is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2fe3e261.
//
// Solidity: function EXTENDED_BALLOT_TYPEHASH() view returns(bytes32)
func (myGovernor *MyGovernor) UnpackEXTENDEDBALLOTTYPEHASH(data []byte) ([32]byte, error) {
	out, err := myGovernor.abi.Unpack("EXTENDED_BALLOT_TYPEHASH", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackCancel is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x452115d6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function cancel(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (myGovernor *MyGovernor) PackCancel(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) []byte {
	enc, err := myGovernor.abi.Pack("cancel", targets, values, calldatas, descriptionHash)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCancel is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x452115d6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function cancel(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (myGovernor *MyGovernor) TryPackCancel(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) ([]byte, error) {
	return myGovernor.abi.Pack("cancel", targets, values, calldatas, descriptionHash)
}

// UnpackCancel is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x452115d6.
//
// Solidity: function cancel(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (myGovernor *MyGovernor) UnpackCancel(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("cancel", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackCastVote is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x56781388.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function castVote(uint256 proposalId, uint8 support) returns(uint256)
func (myGovernor *MyGovernor) PackCastVote(proposalId *big.Int, support uint8) []byte {
	enc, err := myGovernor.abi.Pack("castVote", proposalId, support)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCastVote is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x56781388.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function castVote(uint256 proposalId, uint8 support) returns(uint256)
func (myGovernor *MyGovernor) TryPackCastVote(proposalId *big.Int, support uint8) ([]byte, error) {
	return myGovernor.abi.Pack("castVote", proposalId, support)
}

// UnpackCastVote is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x56781388.
//
// Solidity: function castVote(uint256 proposalId, uint8 support) returns(uint256)
func (myGovernor *MyGovernor) UnpackCastVote(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("castVote", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackCastVoteBySig is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8ff262e3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function castVoteBySig(uint256 proposalId, uint8 support, address voter, bytes signature) returns(uint256)
func (myGovernor *MyGovernor) PackCastVoteBySig(proposalId *big.Int, support uint8, voter common.Address, signature []byte) []byte {
	enc, err := myGovernor.abi.Pack("castVoteBySig", proposalId, support, voter, signature)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCastVoteBySig is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8ff262e3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function castVoteBySig(uint256 proposalId, uint8 support, address voter, bytes signature) returns(uint256)
func (myGovernor *MyGovernor) TryPackCastVoteBySig(proposalId *big.Int, support uint8, voter common.Address, signature []byte) ([]byte, error) {
	return myGovernor.abi.Pack("castVoteBySig", proposalId, support, voter, signature)
}

// UnpackCastVoteBySig is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8ff262e3.
//
// Solidity: function castVoteBySig(uint256 proposalId, uint8 support, address voter, bytes signature) returns(uint256)
func (myGovernor *MyGovernor) UnpackCastVoteBySig(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("castVoteBySig", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackCastVoteWithReason is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7b3c71d3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason) returns(uint256)
func (myGovernor *MyGovernor) PackCastVoteWithReason(proposalId *big.Int, support uint8, reason string) []byte {
	enc, err := myGovernor.abi.Pack("castVoteWithReason", proposalId, support, reason)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCastVoteWithReason is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7b3c71d3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason) returns(uint256)
func (myGovernor *MyGovernor) TryPackCastVoteWithReason(proposalId *big.Int, support uint8, reason string) ([]byte, error) {
	return myGovernor.abi.Pack("castVoteWithReason", proposalId, support, reason)
}

// UnpackCastVoteWithReason is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7b3c71d3.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason) returns(uint256)
func (myGovernor *MyGovernor) UnpackCastVoteWithReason(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("castVoteWithReason", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackCastVoteWithReasonAndParams is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5f398a14.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function castVoteWithReasonAndParams(uint256 proposalId, uint8 support, string reason, bytes params) returns(uint256)
func (myGovernor *MyGovernor) PackCastVoteWithReasonAndParams(proposalId *big.Int, support uint8, reason string, params []byte) []byte {
	enc, err := myGovernor.abi.Pack("castVoteWithReasonAndParams", proposalId, support, reason, params)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCastVoteWithReasonAndParams is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5f398a14.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function castVoteWithReasonAndParams(uint256 proposalId, uint8 support, string reason, bytes params) returns(uint256)
func (myGovernor *MyGovernor) TryPackCastVoteWithReasonAndParams(proposalId *big.Int, support uint8, reason string, params []byte) ([]byte, error) {
	return myGovernor.abi.Pack("castVoteWithReasonAndParams", proposalId, support, reason, params)
}

// UnpackCastVoteWithReasonAndParams is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5f398a14.
//
// Solidity: function castVoteWithReasonAndParams(uint256 proposalId, uint8 support, string reason, bytes params) returns(uint256)
func (myGovernor *MyGovernor) UnpackCastVoteWithReasonAndParams(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("castVoteWithReasonAndParams", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackCastVoteWithReasonAndParamsBySig is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5b8d0e0d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function castVoteWithReasonAndParamsBySig(uint256 proposalId, uint8 support, address voter, string reason, bytes params, bytes signature) returns(uint256)
func (myGovernor *MyGovernor) PackCastVoteWithReasonAndParamsBySig(proposalId *big.Int, support uint8, voter common.Address, reason string, params []byte, signature []byte) []byte {
	enc, err := myGovernor.abi.Pack("castVoteWithReasonAndParamsBySig", proposalId, support, voter, reason, params, signature)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCastVoteWithReasonAndParamsBySig is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5b8d0e0d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function castVoteWithReasonAndParamsBySig(uint256 proposalId, uint8 support, address voter, string reason, bytes params, bytes signature) returns(uint256)
func (myGovernor *MyGovernor) TryPackCastVoteWithReasonAndParamsBySig(proposalId *big.Int, support uint8, voter common.Address, reason string, params []byte, signature []byte) ([]byte, error) {
	return myGovernor.abi.Pack("castVoteWithReasonAndParamsBySig", proposalId, support, voter, reason, params, signature)
}

// UnpackCastVoteWithReasonAndParamsBySig is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5b8d0e0d.
//
// Solidity: function castVoteWithReasonAndParamsBySig(uint256 proposalId, uint8 support, address voter, string reason, bytes params, bytes signature) returns(uint256)
func (myGovernor *MyGovernor) UnpackCastVoteWithReasonAndParamsBySig(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("castVoteWithReasonAndParamsBySig", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackClock is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91ddadf4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function clock() view returns(uint48)
func (myGovernor *MyGovernor) PackClock() []byte {
	enc, err := myGovernor.abi.Pack("clock")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackClock is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91ddadf4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function clock() view returns(uint48)
func (myGovernor *MyGovernor) TryPackClock() ([]byte, error) {
	return myGovernor.abi.Pack("clock")
}

// UnpackClock is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (myGovernor *MyGovernor) UnpackClock(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("clock", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackEip712Domain is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x84b0196e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (myGovernor *MyGovernor) PackEip712Domain() []byte {
	enc, err := myGovernor.abi.Pack("eip712Domain")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackEip712Domain is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x84b0196e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (myGovernor *MyGovernor) TryPackEip712Domain() ([]byte, error) {
	return myGovernor.abi.Pack("eip712Domain")
}

// Eip712DomainOutput serves as a container for the return parameters of contract
// method Eip712Domain.
// type Eip712DomainOutput struct {
// 	Fields            [1]byte
// 	Name              string
// 	Version           string
// 	ChainId           *big.Int
// 	VerifyingContract common.Address
// 	Salt              [32]byte
// 	Extensions        []*big.Int
// }

// UnpackEip712Domain is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (myGovernor *MyGovernor) UnpackEip712Domain(data []byte) (Eip712DomainOutput, error) {
	out, err := myGovernor.abi.Unpack("eip712Domain", data)
	outstruct := new(Eip712DomainOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = abi.ConvertType(out[3], new(big.Int)).(*big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)
	return *outstruct, nil
}

// PackExecute is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2656227d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function execute(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) payable returns(uint256)
func (myGovernor *MyGovernor) PackExecute(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) []byte {
	enc, err := myGovernor.abi.Pack("execute", targets, values, calldatas, descriptionHash)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackExecute is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2656227d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function execute(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) payable returns(uint256)
func (myGovernor *MyGovernor) TryPackExecute(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) ([]byte, error) {
	return myGovernor.abi.Pack("execute", targets, values, calldatas, descriptionHash)
}

// UnpackExecute is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2656227d.
//
// Solidity: function execute(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) payable returns(uint256)
func (myGovernor *MyGovernor) UnpackExecute(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("execute", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetProposalId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa8f8a668.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getProposalId(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) view returns(uint256)
func (myGovernor *MyGovernor) PackGetProposalId(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) []byte {
	enc, err := myGovernor.abi.Pack("getProposalId", targets, values, calldatas, descriptionHash)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetProposalId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa8f8a668.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getProposalId(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) view returns(uint256)
func (myGovernor *MyGovernor) TryPackGetProposalId(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) ([]byte, error) {
	return myGovernor.abi.Pack("getProposalId", targets, values, calldatas, descriptionHash)
}

// UnpackGetProposalId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa8f8a668.
//
// Solidity: function getProposalId(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) view returns(uint256)
func (myGovernor *MyGovernor) UnpackGetProposalId(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("getProposalId", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetVotes is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xeb9019d4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getVotes(address account, uint256 timepoint) view returns(uint256)
func (myGovernor *MyGovernor) PackGetVotes(account common.Address, timepoint *big.Int) []byte {
	enc, err := myGovernor.abi.Pack("getVotes", account, timepoint)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetVotes is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xeb9019d4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getVotes(address account, uint256 timepoint) view returns(uint256)
func (myGovernor *MyGovernor) TryPackGetVotes(account common.Address, timepoint *big.Int) ([]byte, error) {
	return myGovernor.abi.Pack("getVotes", account, timepoint)
}

// UnpackGetVotes is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xeb9019d4.
//
// Solidity: function getVotes(address account, uint256 timepoint) view returns(uint256)
func (myGovernor *MyGovernor) UnpackGetVotes(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("getVotes", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetVotesWithParams is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9a802a6d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getVotesWithParams(address account, uint256 timepoint, bytes params) view returns(uint256)
func (myGovernor *MyGovernor) PackGetVotesWithParams(account common.Address, timepoint *big.Int, params []byte) []byte {
	enc, err := myGovernor.abi.Pack("getVotesWithParams", account, timepoint, params)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetVotesWithParams is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9a802a6d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getVotesWithParams(address account, uint256 timepoint, bytes params) view returns(uint256)
func (myGovernor *MyGovernor) TryPackGetVotesWithParams(account common.Address, timepoint *big.Int, params []byte) ([]byte, error) {
	return myGovernor.abi.Pack("getVotesWithParams", account, timepoint, params)
}

// UnpackGetVotesWithParams is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9a802a6d.
//
// Solidity: function getVotesWithParams(address account, uint256 timepoint, bytes params) view returns(uint256)
func (myGovernor *MyGovernor) UnpackGetVotesWithParams(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("getVotesWithParams", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackHasVoted is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x43859632.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function hasVoted(uint256 proposalId, address account) view returns(bool)
func (myGovernor *MyGovernor) PackHasVoted(proposalId *big.Int, account common.Address) []byte {
	enc, err := myGovernor.abi.Pack("hasVoted", proposalId, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackHasVoted is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x43859632.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function hasVoted(uint256 proposalId, address account) view returns(bool)
func (myGovernor *MyGovernor) TryPackHasVoted(proposalId *big.Int, account common.Address) ([]byte, error) {
	return myGovernor.abi.Pack("hasVoted", proposalId, account)
}

// UnpackHasVoted is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x43859632.
//
// Solidity: function hasVoted(uint256 proposalId, address account) view returns(bool)
func (myGovernor *MyGovernor) UnpackHasVoted(data []byte) (bool, error) {
	out, err := myGovernor.abi.Unpack("hasVoted", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackHashProposal is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc59057e4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function hashProposal(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) pure returns(uint256)
func (myGovernor *MyGovernor) PackHashProposal(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) []byte {
	enc, err := myGovernor.abi.Pack("hashProposal", targets, values, calldatas, descriptionHash)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackHashProposal is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc59057e4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function hashProposal(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) pure returns(uint256)
func (myGovernor *MyGovernor) TryPackHashProposal(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) ([]byte, error) {
	return myGovernor.abi.Pack("hashProposal", targets, values, calldatas, descriptionHash)
}

// UnpackHashProposal is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc59057e4.
//
// Solidity: function hashProposal(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) pure returns(uint256)
func (myGovernor *MyGovernor) UnpackHashProposal(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("hashProposal", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function name() view returns(string)
func (myGovernor *MyGovernor) PackName() []byte {
	enc, err := myGovernor.abi.Pack("name")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function name() view returns(string)
func (myGovernor *MyGovernor) TryPackName() ([]byte, error) {
	return myGovernor.abi.Pack("name")
}

// UnpackName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (myGovernor *MyGovernor) UnpackName(data []byte) (string, error) {
	out, err := myGovernor.abi.Unpack("name", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackNonces is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7ecebe00.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (myGovernor *MyGovernor) PackNonces(owner common.Address) []byte {
	enc, err := myGovernor.abi.Pack("nonces", owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackNonces is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7ecebe00.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (myGovernor *MyGovernor) TryPackNonces(owner common.Address) ([]byte, error) {
	return myGovernor.abi.Pack("nonces", owner)
}

// UnpackNonces is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (myGovernor *MyGovernor) UnpackNonces(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("nonces", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackOnERC1155BatchReceived is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbc197c81.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (myGovernor *MyGovernor) PackOnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) []byte {
	enc, err := myGovernor.abi.Pack("onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOnERC1155BatchReceived is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbc197c81.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (myGovernor *MyGovernor) TryPackOnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([]byte, error) {
	return myGovernor.abi.Pack("onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// UnpackOnERC1155BatchReceived is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (myGovernor *MyGovernor) UnpackOnERC1155BatchReceived(data []byte) ([4]byte, error) {
	out, err := myGovernor.abi.Unpack("onERC1155BatchReceived", data)
	if err != nil {
		return *new([4]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	return out0, nil
}

// PackOnERC1155Received is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf23a6e61.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (myGovernor *MyGovernor) PackOnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) []byte {
	enc, err := myGovernor.abi.Pack("onERC1155Received", arg0, arg1, arg2, arg3, arg4)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOnERC1155Received is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf23a6e61.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (myGovernor *MyGovernor) TryPackOnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([]byte, error) {
	return myGovernor.abi.Pack("onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// UnpackOnERC1155Received is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (myGovernor *MyGovernor) UnpackOnERC1155Received(data []byte) ([4]byte, error) {
	out, err := myGovernor.abi.Unpack("onERC1155Received", data)
	if err != nil {
		return *new([4]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	return out0, nil
}

// PackOnERC721Received is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x150b7a02.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (myGovernor *MyGovernor) PackOnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) []byte {
	enc, err := myGovernor.abi.Pack("onERC721Received", arg0, arg1, arg2, arg3)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOnERC721Received is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x150b7a02.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (myGovernor *MyGovernor) TryPackOnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([]byte, error) {
	return myGovernor.abi.Pack("onERC721Received", arg0, arg1, arg2, arg3)
}

// UnpackOnERC721Received is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (myGovernor *MyGovernor) UnpackOnERC721Received(data []byte) ([4]byte, error) {
	out, err := myGovernor.abi.Unpack("onERC721Received", data)
	if err != nil {
		return *new([4]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	return out0, nil
}

// PackProposalDeadline is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc01f9e37.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function proposalDeadline(uint256 proposalId) view returns(uint256)
func (myGovernor *MyGovernor) PackProposalDeadline(proposalId *big.Int) []byte {
	enc, err := myGovernor.abi.Pack("proposalDeadline", proposalId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackProposalDeadline is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc01f9e37.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function proposalDeadline(uint256 proposalId) view returns(uint256)
func (myGovernor *MyGovernor) TryPackProposalDeadline(proposalId *big.Int) ([]byte, error) {
	return myGovernor.abi.Pack("proposalDeadline", proposalId)
}

// UnpackProposalDeadline is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc01f9e37.
//
// Solidity: function proposalDeadline(uint256 proposalId) view returns(uint256)
func (myGovernor *MyGovernor) UnpackProposalDeadline(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("proposalDeadline", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackProposalEta is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xab58fb8e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function proposalEta(uint256 proposalId) view returns(uint256)
func (myGovernor *MyGovernor) PackProposalEta(proposalId *big.Int) []byte {
	enc, err := myGovernor.abi.Pack("proposalEta", proposalId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackProposalEta is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xab58fb8e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function proposalEta(uint256 proposalId) view returns(uint256)
func (myGovernor *MyGovernor) TryPackProposalEta(proposalId *big.Int) ([]byte, error) {
	return myGovernor.abi.Pack("proposalEta", proposalId)
}

// UnpackProposalEta is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xab58fb8e.
//
// Solidity: function proposalEta(uint256 proposalId) view returns(uint256)
func (myGovernor *MyGovernor) UnpackProposalEta(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("proposalEta", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackProposalNeedsQueuing is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa9a95294.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function proposalNeedsQueuing(uint256 proposalId) view returns(bool)
func (myGovernor *MyGovernor) PackProposalNeedsQueuing(proposalId *big.Int) []byte {
	enc, err := myGovernor.abi.Pack("proposalNeedsQueuing", proposalId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackProposalNeedsQueuing is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa9a95294.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function proposalNeedsQueuing(uint256 proposalId) view returns(bool)
func (myGovernor *MyGovernor) TryPackProposalNeedsQueuing(proposalId *big.Int) ([]byte, error) {
	return myGovernor.abi.Pack("proposalNeedsQueuing", proposalId)
}

// UnpackProposalNeedsQueuing is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa9a95294.
//
// Solidity: function proposalNeedsQueuing(uint256 proposalId) view returns(bool)
func (myGovernor *MyGovernor) UnpackProposalNeedsQueuing(data []byte) (bool, error) {
	out, err := myGovernor.abi.Unpack("proposalNeedsQueuing", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackProposalProposer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x143489d0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function proposalProposer(uint256 proposalId) view returns(address)
func (myGovernor *MyGovernor) PackProposalProposer(proposalId *big.Int) []byte {
	enc, err := myGovernor.abi.Pack("proposalProposer", proposalId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackProposalProposer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x143489d0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function proposalProposer(uint256 proposalId) view returns(address)
func (myGovernor *MyGovernor) TryPackProposalProposer(proposalId *big.Int) ([]byte, error) {
	return myGovernor.abi.Pack("proposalProposer", proposalId)
}

// UnpackProposalProposer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x143489d0.
//
// Solidity: function proposalProposer(uint256 proposalId) view returns(address)
func (myGovernor *MyGovernor) UnpackProposalProposer(data []byte) (common.Address, error) {
	out, err := myGovernor.abi.Unpack("proposalProposer", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackProposalSnapshot is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2d63f693.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function proposalSnapshot(uint256 proposalId) view returns(uint256)
func (myGovernor *MyGovernor) PackProposalSnapshot(proposalId *big.Int) []byte {
	enc, err := myGovernor.abi.Pack("proposalSnapshot", proposalId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackProposalSnapshot is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2d63f693.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function proposalSnapshot(uint256 proposalId) view returns(uint256)
func (myGovernor *MyGovernor) TryPackProposalSnapshot(proposalId *big.Int) ([]byte, error) {
	return myGovernor.abi.Pack("proposalSnapshot", proposalId)
}

// UnpackProposalSnapshot is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2d63f693.
//
// Solidity: function proposalSnapshot(uint256 proposalId) view returns(uint256)
func (myGovernor *MyGovernor) UnpackProposalSnapshot(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("proposalSnapshot", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackProposalThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb58131b0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function proposalThreshold() pure returns(uint256)
func (myGovernor *MyGovernor) PackProposalThreshold() []byte {
	enc, err := myGovernor.abi.Pack("proposalThreshold")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackProposalThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb58131b0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function proposalThreshold() pure returns(uint256)
func (myGovernor *MyGovernor) TryPackProposalThreshold() ([]byte, error) {
	return myGovernor.abi.Pack("proposalThreshold")
}

// UnpackProposalThreshold is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb58131b0.
//
// Solidity: function proposalThreshold() pure returns(uint256)
func (myGovernor *MyGovernor) UnpackProposalThreshold(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("proposalThreshold", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackProposalVotes is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x544ffc9c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function proposalVotes(uint256 proposalId) view returns(uint256 againstVotes, uint256 forVotes, uint256 abstainVotes)
func (myGovernor *MyGovernor) PackProposalVotes(proposalId *big.Int) []byte {
	enc, err := myGovernor.abi.Pack("proposalVotes", proposalId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackProposalVotes is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x544ffc9c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function proposalVotes(uint256 proposalId) view returns(uint256 againstVotes, uint256 forVotes, uint256 abstainVotes)
func (myGovernor *MyGovernor) TryPackProposalVotes(proposalId *big.Int) ([]byte, error) {
	return myGovernor.abi.Pack("proposalVotes", proposalId)
}

// ProposalVotesOutput serves as a container for the return parameters of contract
// method ProposalVotes.
type ProposalVotesOutput struct {
	AgainstVotes *big.Int
	ForVotes     *big.Int
	AbstainVotes *big.Int
}

// UnpackProposalVotes is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x544ffc9c.
//
// Solidity: function proposalVotes(uint256 proposalId) view returns(uint256 againstVotes, uint256 forVotes, uint256 abstainVotes)
func (myGovernor *MyGovernor) UnpackProposalVotes(data []byte) (ProposalVotesOutput, error) {
	out, err := myGovernor.abi.Unpack("proposalVotes", data)
	outstruct := new(ProposalVotesOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.AgainstVotes = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.ForVotes = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	outstruct.AbstainVotes = abi.ConvertType(out[2], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackPropose is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7d5e81e2.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function propose(address[] targets, uint256[] values, bytes[] calldatas, string description) returns(uint256)
func (myGovernor *MyGovernor) PackPropose(targets []common.Address, values []*big.Int, calldatas [][]byte, description string) []byte {
	enc, err := myGovernor.abi.Pack("propose", targets, values, calldatas, description)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPropose is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7d5e81e2.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function propose(address[] targets, uint256[] values, bytes[] calldatas, string description) returns(uint256)
func (myGovernor *MyGovernor) TryPackPropose(targets []common.Address, values []*big.Int, calldatas [][]byte, description string) ([]byte, error) {
	return myGovernor.abi.Pack("propose", targets, values, calldatas, description)
}

// UnpackPropose is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7d5e81e2.
//
// Solidity: function propose(address[] targets, uint256[] values, bytes[] calldatas, string description) returns(uint256)
func (myGovernor *MyGovernor) UnpackPropose(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("propose", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackQueue is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x160cbed7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function queue(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (myGovernor *MyGovernor) PackQueue(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) []byte {
	enc, err := myGovernor.abi.Pack("queue", targets, values, calldatas, descriptionHash)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackQueue is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x160cbed7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function queue(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (myGovernor *MyGovernor) TryPackQueue(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) ([]byte, error) {
	return myGovernor.abi.Pack("queue", targets, values, calldatas, descriptionHash)
}

// UnpackQueue is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x160cbed7.
//
// Solidity: function queue(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (myGovernor *MyGovernor) UnpackQueue(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("queue", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackQuorum is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf8ce560a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function quorum(uint256 timepoint) view returns(uint256)
func (myGovernor *MyGovernor) PackQuorum(timepoint *big.Int) []byte {
	enc, err := myGovernor.abi.Pack("quorum", timepoint)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackQuorum is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf8ce560a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function quorum(uint256 timepoint) view returns(uint256)
func (myGovernor *MyGovernor) TryPackQuorum(timepoint *big.Int) ([]byte, error) {
	return myGovernor.abi.Pack("quorum", timepoint)
}

// UnpackQuorum is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf8ce560a.
//
// Solidity: function quorum(uint256 timepoint) view returns(uint256)
func (myGovernor *MyGovernor) UnpackQuorum(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("quorum", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackQuorumDenominator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x97c3d334.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function quorumDenominator() view returns(uint256)
func (myGovernor *MyGovernor) PackQuorumDenominator() []byte {
	enc, err := myGovernor.abi.Pack("quorumDenominator")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackQuorumDenominator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x97c3d334.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function quorumDenominator() view returns(uint256)
func (myGovernor *MyGovernor) TryPackQuorumDenominator() ([]byte, error) {
	return myGovernor.abi.Pack("quorumDenominator")
}

// UnpackQuorumDenominator is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x97c3d334.
//
// Solidity: function quorumDenominator() view returns(uint256)
func (myGovernor *MyGovernor) UnpackQuorumDenominator(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("quorumDenominator", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackQuorumNumerator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x60c4247f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function quorumNumerator(uint256 timepoint) view returns(uint256)
func (myGovernor *MyGovernor) PackQuorumNumerator(timepoint *big.Int) []byte {
	enc, err := myGovernor.abi.Pack("quorumNumerator", timepoint)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackQuorumNumerator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x60c4247f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function quorumNumerator(uint256 timepoint) view returns(uint256)
func (myGovernor *MyGovernor) TryPackQuorumNumerator(timepoint *big.Int) ([]byte, error) {
	return myGovernor.abi.Pack("quorumNumerator", timepoint)
}

// UnpackQuorumNumerator is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x60c4247f.
//
// Solidity: function quorumNumerator(uint256 timepoint) view returns(uint256)
func (myGovernor *MyGovernor) UnpackQuorumNumerator(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("quorumNumerator", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackQuorumNumerator0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa7713a70.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function quorumNumerator() view returns(uint256)
func (myGovernor *MyGovernor) PackQuorumNumerator0() []byte {
	enc, err := myGovernor.abi.Pack("quorumNumerator0")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackQuorumNumerator0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa7713a70.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function quorumNumerator() view returns(uint256)
func (myGovernor *MyGovernor) TryPackQuorumNumerator0() ([]byte, error) {
	return myGovernor.abi.Pack("quorumNumerator0")
}

// UnpackQuorumNumerator0 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa7713a70.
//
// Solidity: function quorumNumerator() view returns(uint256)
func (myGovernor *MyGovernor) UnpackQuorumNumerator0(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("quorumNumerator0", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackRelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc28bc2fa.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function relay(address target, uint256 value, bytes data) payable returns()
func (myGovernor *MyGovernor) PackRelay(target common.Address, value *big.Int, data []byte) []byte {
	enc, err := myGovernor.abi.Pack("relay", target, value, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc28bc2fa.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function relay(address target, uint256 value, bytes data) payable returns()
func (myGovernor *MyGovernor) TryPackRelay(target common.Address, value *big.Int, data []byte) ([]byte, error) {
	return myGovernor.abi.Pack("relay", target, value, data)
}

// PackState is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3e4f49e6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (myGovernor *MyGovernor) PackState(proposalId *big.Int) []byte {
	enc, err := myGovernor.abi.Pack("state", proposalId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackState is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3e4f49e6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (myGovernor *MyGovernor) TryPackState(proposalId *big.Int) ([]byte, error) {
	return myGovernor.abi.Pack("state", proposalId)
}

// UnpackState is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (myGovernor *MyGovernor) UnpackState(data []byte) (uint8, error) {
	out, err := myGovernor.abi.Unpack("state", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, nil
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (myGovernor *MyGovernor) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := myGovernor.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (myGovernor *MyGovernor) TryPackSupportsInterface(interfaceId [4]byte) ([]byte, error) {
	return myGovernor.abi.Pack("supportsInterface", interfaceId)
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (myGovernor *MyGovernor) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := myGovernor.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackTimelock is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd33219b4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function timelock() view returns(address)
func (myGovernor *MyGovernor) PackTimelock() []byte {
	enc, err := myGovernor.abi.Pack("timelock")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTimelock is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd33219b4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function timelock() view returns(address)
func (myGovernor *MyGovernor) TryPackTimelock() ([]byte, error) {
	return myGovernor.abi.Pack("timelock")
}

// UnpackTimelock is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd33219b4.
//
// Solidity: function timelock() view returns(address)
func (myGovernor *MyGovernor) UnpackTimelock(data []byte) (common.Address, error) {
	out, err := myGovernor.abi.Unpack("timelock", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfc0c546a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function token() view returns(address)
func (myGovernor *MyGovernor) PackToken() []byte {
	enc, err := myGovernor.abi.Pack("token")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfc0c546a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function token() view returns(address)
func (myGovernor *MyGovernor) TryPackToken() ([]byte, error) {
	return myGovernor.abi.Pack("token")
}

// UnpackToken is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (myGovernor *MyGovernor) UnpackToken(data []byte) (common.Address, error) {
	out, err := myGovernor.abi.Unpack("token", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackUpdateQuorumNumerator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06f3f9e6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function updateQuorumNumerator(uint256 newQuorumNumerator) returns()
func (myGovernor *MyGovernor) PackUpdateQuorumNumerator(newQuorumNumerator *big.Int) []byte {
	enc, err := myGovernor.abi.Pack("updateQuorumNumerator", newQuorumNumerator)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUpdateQuorumNumerator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06f3f9e6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function updateQuorumNumerator(uint256 newQuorumNumerator) returns()
func (myGovernor *MyGovernor) TryPackUpdateQuorumNumerator(newQuorumNumerator *big.Int) ([]byte, error) {
	return myGovernor.abi.Pack("updateQuorumNumerator", newQuorumNumerator)
}

// PackUpdateTimelock is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa890c910.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function updateTimelock(address newTimelock) returns()
func (myGovernor *MyGovernor) PackUpdateTimelock(newTimelock common.Address) []byte {
	enc, err := myGovernor.abi.Pack("updateTimelock", newTimelock)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUpdateTimelock is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa890c910.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function updateTimelock(address newTimelock) returns()
func (myGovernor *MyGovernor) TryPackUpdateTimelock(newTimelock common.Address) ([]byte, error) {
	return myGovernor.abi.Pack("updateTimelock", newTimelock)
}

// PackVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x54fd4d50.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function version() view returns(string)
func (myGovernor *MyGovernor) PackVersion() []byte {
	enc, err := myGovernor.abi.Pack("version")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x54fd4d50.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function version() view returns(string)
func (myGovernor *MyGovernor) TryPackVersion() ([]byte, error) {
	return myGovernor.abi.Pack("version")
}

// UnpackVersion is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (myGovernor *MyGovernor) UnpackVersion(data []byte) (string, error) {
	out, err := myGovernor.abi.Unpack("version", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackVotingDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3932abb1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function votingDelay() pure returns(uint256)
func (myGovernor *MyGovernor) PackVotingDelay() []byte {
	enc, err := myGovernor.abi.Pack("votingDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackVotingDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3932abb1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function votingDelay() pure returns(uint256)
func (myGovernor *MyGovernor) TryPackVotingDelay() ([]byte, error) {
	return myGovernor.abi.Pack("votingDelay")
}

// UnpackVotingDelay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3932abb1.
//
// Solidity: function votingDelay() pure returns(uint256)
func (myGovernor *MyGovernor) UnpackVotingDelay(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("votingDelay", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackVotingPeriod is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x02a251a3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function votingPeriod() pure returns(uint256)
func (myGovernor *MyGovernor) PackVotingPeriod() []byte {
	enc, err := myGovernor.abi.Pack("votingPeriod")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackVotingPeriod is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x02a251a3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function votingPeriod() pure returns(uint256)
func (myGovernor *MyGovernor) TryPackVotingPeriod() ([]byte, error) {
	return myGovernor.abi.Pack("votingPeriod")
}

// UnpackVotingPeriod is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x02a251a3.
//
// Solidity: function votingPeriod() pure returns(uint256)
func (myGovernor *MyGovernor) UnpackVotingPeriod(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("votingPeriod", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// MyGovernorEIP712DomainChanged represents a EIP712DomainChanged event raised by the MyGovernor contract.
type MyGovernorEIP712DomainChanged struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const MyGovernorEIP712DomainChangedEventName = "EIP712DomainChanged"

// ContractEventName returns the user-defined event name.
func (MyGovernorEIP712DomainChanged) ContractEventName() string {
	return MyGovernorEIP712DomainChangedEventName
}

// UnpackEIP712DomainChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EIP712DomainChanged()
func (myGovernor *MyGovernor) UnpackEIP712DomainChangedEvent(log *types.Log) (*MyGovernorEIP712DomainChanged, error) {
	event := "EIP712DomainChanged"
	if log.Topics[0] != myGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MyGovernorEIP712DomainChanged)
	if len(log.Data) > 0 {
		if err := myGovernor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range myGovernor.abi.Events[event].Inputs {
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

// MyGovernorProposalCanceled represents a ProposalCanceled event raised by the MyGovernor contract.
type MyGovernorProposalCanceled struct {
	ProposalId *big.Int
	Raw        *types.Log // Blockchain specific contextual infos
}

const MyGovernorProposalCanceledEventName = "ProposalCanceled"

// ContractEventName returns the user-defined event name.
func (MyGovernorProposalCanceled) ContractEventName() string {
	return MyGovernorProposalCanceledEventName
}

// UnpackProposalCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ProposalCanceled(uint256 proposalId)
func (myGovernor *MyGovernor) UnpackProposalCanceledEvent(log *types.Log) (*MyGovernorProposalCanceled, error) {
	event := "ProposalCanceled"
	if log.Topics[0] != myGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MyGovernorProposalCanceled)
	if len(log.Data) > 0 {
		if err := myGovernor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range myGovernor.abi.Events[event].Inputs {
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

// MyGovernorProposalCreated represents a ProposalCreated event raised by the MyGovernor contract.
type MyGovernorProposalCreated struct {
	ProposalId  *big.Int
	Proposer    common.Address
	Targets     []common.Address
	Values      []*big.Int
	Signatures  []string
	Calldatas   [][]byte
	VoteStart   *big.Int
	VoteEnd     *big.Int
	Description string
	Raw         *types.Log // Blockchain specific contextual infos
}

const MyGovernorProposalCreatedEventName = "ProposalCreated"

// ContractEventName returns the user-defined event name.
func (MyGovernorProposalCreated) ContractEventName() string {
	return MyGovernorProposalCreatedEventName
}

// UnpackProposalCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ProposalCreated(uint256 proposalId, address proposer, address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, uint256 voteStart, uint256 voteEnd, string description)
func (myGovernor *MyGovernor) UnpackProposalCreatedEvent(log *types.Log) (*MyGovernorProposalCreated, error) {
	event := "ProposalCreated"
	if log.Topics[0] != myGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MyGovernorProposalCreated)
	if len(log.Data) > 0 {
		if err := myGovernor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range myGovernor.abi.Events[event].Inputs {
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

// MyGovernorProposalExecuted represents a ProposalExecuted event raised by the MyGovernor contract.
type MyGovernorProposalExecuted struct {
	ProposalId *big.Int
	Raw        *types.Log // Blockchain specific contextual infos
}

const MyGovernorProposalExecutedEventName = "ProposalExecuted"

// ContractEventName returns the user-defined event name.
func (MyGovernorProposalExecuted) ContractEventName() string {
	return MyGovernorProposalExecutedEventName
}

// UnpackProposalExecutedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (myGovernor *MyGovernor) UnpackProposalExecutedEvent(log *types.Log) (*MyGovernorProposalExecuted, error) {
	event := "ProposalExecuted"
	if log.Topics[0] != myGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MyGovernorProposalExecuted)
	if len(log.Data) > 0 {
		if err := myGovernor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range myGovernor.abi.Events[event].Inputs {
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

// MyGovernorProposalQueued represents a ProposalQueued event raised by the MyGovernor contract.
type MyGovernorProposalQueued struct {
	ProposalId *big.Int
	EtaSeconds *big.Int
	Raw        *types.Log // Blockchain specific contextual infos
}

const MyGovernorProposalQueuedEventName = "ProposalQueued"

// ContractEventName returns the user-defined event name.
func (MyGovernorProposalQueued) ContractEventName() string {
	return MyGovernorProposalQueuedEventName
}

// UnpackProposalQueuedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ProposalQueued(uint256 proposalId, uint256 etaSeconds)
func (myGovernor *MyGovernor) UnpackProposalQueuedEvent(log *types.Log) (*MyGovernorProposalQueued, error) {
	event := "ProposalQueued"
	if log.Topics[0] != myGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MyGovernorProposalQueued)
	if len(log.Data) > 0 {
		if err := myGovernor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range myGovernor.abi.Events[event].Inputs {
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

// MyGovernorQuorumNumeratorUpdated represents a QuorumNumeratorUpdated event raised by the MyGovernor contract.
type MyGovernorQuorumNumeratorUpdated struct {
	OldQuorumNumerator *big.Int
	NewQuorumNumerator *big.Int
	Raw                *types.Log // Blockchain specific contextual infos
}

const MyGovernorQuorumNumeratorUpdatedEventName = "QuorumNumeratorUpdated"

// ContractEventName returns the user-defined event name.
func (MyGovernorQuorumNumeratorUpdated) ContractEventName() string {
	return MyGovernorQuorumNumeratorUpdatedEventName
}

// UnpackQuorumNumeratorUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event QuorumNumeratorUpdated(uint256 oldQuorumNumerator, uint256 newQuorumNumerator)
func (myGovernor *MyGovernor) UnpackQuorumNumeratorUpdatedEvent(log *types.Log) (*MyGovernorQuorumNumeratorUpdated, error) {
	event := "QuorumNumeratorUpdated"
	if log.Topics[0] != myGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MyGovernorQuorumNumeratorUpdated)
	if len(log.Data) > 0 {
		if err := myGovernor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range myGovernor.abi.Events[event].Inputs {
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

// MyGovernorTimelockChange represents a TimelockChange event raised by the MyGovernor contract.
type MyGovernorTimelockChange struct {
	OldTimelock common.Address
	NewTimelock common.Address
	Raw         *types.Log // Blockchain specific contextual infos
}

const MyGovernorTimelockChangeEventName = "TimelockChange"

// ContractEventName returns the user-defined event name.
func (MyGovernorTimelockChange) ContractEventName() string {
	return MyGovernorTimelockChangeEventName
}

// UnpackTimelockChangeEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TimelockChange(address oldTimelock, address newTimelock)
func (myGovernor *MyGovernor) UnpackTimelockChangeEvent(log *types.Log) (*MyGovernorTimelockChange, error) {
	event := "TimelockChange"
	if log.Topics[0] != myGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MyGovernorTimelockChange)
	if len(log.Data) > 0 {
		if err := myGovernor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range myGovernor.abi.Events[event].Inputs {
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

// MyGovernorVoteCast represents a VoteCast event raised by the MyGovernor contract.
type MyGovernorVoteCast struct {
	Voter      common.Address
	ProposalId *big.Int
	Support    uint8
	Weight     *big.Int
	Reason     string
	Raw        *types.Log // Blockchain specific contextual infos
}

const MyGovernorVoteCastEventName = "VoteCast"

// ContractEventName returns the user-defined event name.
func (MyGovernorVoteCast) ContractEventName() string {
	return MyGovernorVoteCastEventName
}

// UnpackVoteCastEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event VoteCast(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason)
func (myGovernor *MyGovernor) UnpackVoteCastEvent(log *types.Log) (*MyGovernorVoteCast, error) {
	event := "VoteCast"
	if log.Topics[0] != myGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MyGovernorVoteCast)
	if len(log.Data) > 0 {
		if err := myGovernor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range myGovernor.abi.Events[event].Inputs {
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

// MyGovernorVoteCastWithParams represents a VoteCastWithParams event raised by the MyGovernor contract.
type MyGovernorVoteCastWithParams struct {
	Voter      common.Address
	ProposalId *big.Int
	Support    uint8
	Weight     *big.Int
	Reason     string
	Params     []byte
	Raw        *types.Log // Blockchain specific contextual infos
}

const MyGovernorVoteCastWithParamsEventName = "VoteCastWithParams"

// ContractEventName returns the user-defined event name.
func (MyGovernorVoteCastWithParams) ContractEventName() string {
	return MyGovernorVoteCastWithParamsEventName
}

// UnpackVoteCastWithParamsEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event VoteCastWithParams(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason, bytes params)
func (myGovernor *MyGovernor) UnpackVoteCastWithParamsEvent(log *types.Log) (*MyGovernorVoteCastWithParams, error) {
	event := "VoteCastWithParams"
	if log.Topics[0] != myGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MyGovernorVoteCastWithParams)
	if len(log.Data) > 0 {
		if err := myGovernor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range myGovernor.abi.Events[event].Inputs {
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
func (myGovernor *MyGovernor) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["CheckpointUnorderedInsertion"].ID.Bytes()[:4]) {
		return myGovernor.UnpackCheckpointUnorderedInsertionError(raw[4:])
	}
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["FailedCall"].ID.Bytes()[:4]) {
		return myGovernor.UnpackFailedCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["GovernorAlreadyCastVote"].ID.Bytes()[:4]) {
		return myGovernor.UnpackGovernorAlreadyCastVoteError(raw[4:])
	}
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["GovernorAlreadyQueuedProposal"].ID.Bytes()[:4]) {
		return myGovernor.UnpackGovernorAlreadyQueuedProposalError(raw[4:])
	}
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["GovernorDisabledDeposit"].ID.Bytes()[:4]) {
		return myGovernor.UnpackGovernorDisabledDepositError(raw[4:])
	}
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["GovernorInsufficientProposerVotes"].ID.Bytes()[:4]) {
		return myGovernor.UnpackGovernorInsufficientProposerVotesError(raw[4:])
	}
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["GovernorInvalidProposalLength"].ID.Bytes()[:4]) {
		return myGovernor.UnpackGovernorInvalidProposalLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["GovernorInvalidQuorumFraction"].ID.Bytes()[:4]) {
		return myGovernor.UnpackGovernorInvalidQuorumFractionError(raw[4:])
	}
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["GovernorInvalidSignature"].ID.Bytes()[:4]) {
		return myGovernor.UnpackGovernorInvalidSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["GovernorInvalidVoteParams"].ID.Bytes()[:4]) {
		return myGovernor.UnpackGovernorInvalidVoteParamsError(raw[4:])
	}
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["GovernorInvalidVoteType"].ID.Bytes()[:4]) {
		return myGovernor.UnpackGovernorInvalidVoteTypeError(raw[4:])
	}
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["GovernorInvalidVotingPeriod"].ID.Bytes()[:4]) {
		return myGovernor.UnpackGovernorInvalidVotingPeriodError(raw[4:])
	}
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["GovernorNonexistentProposal"].ID.Bytes()[:4]) {
		return myGovernor.UnpackGovernorNonexistentProposalError(raw[4:])
	}
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["GovernorNotQueuedProposal"].ID.Bytes()[:4]) {
		return myGovernor.UnpackGovernorNotQueuedProposalError(raw[4:])
	}
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["GovernorOnlyExecutor"].ID.Bytes()[:4]) {
		return myGovernor.UnpackGovernorOnlyExecutorError(raw[4:])
	}
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["GovernorQueueNotImplemented"].ID.Bytes()[:4]) {
		return myGovernor.UnpackGovernorQueueNotImplementedError(raw[4:])
	}
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["GovernorRestrictedProposer"].ID.Bytes()[:4]) {
		return myGovernor.UnpackGovernorRestrictedProposerError(raw[4:])
	}
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["GovernorUnableToCancel"].ID.Bytes()[:4]) {
		return myGovernor.UnpackGovernorUnableToCancelError(raw[4:])
	}
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["GovernorUnexpectedProposalState"].ID.Bytes()[:4]) {
		return myGovernor.UnpackGovernorUnexpectedProposalStateError(raw[4:])
	}
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["InvalidAccountNonce"].ID.Bytes()[:4]) {
		return myGovernor.UnpackInvalidAccountNonceError(raw[4:])
	}
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["InvalidShortString"].ID.Bytes()[:4]) {
		return myGovernor.UnpackInvalidShortStringError(raw[4:])
	}
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["SafeCastOverflowedUintDowncast"].ID.Bytes()[:4]) {
		return myGovernor.UnpackSafeCastOverflowedUintDowncastError(raw[4:])
	}
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["StringTooLong"].ID.Bytes()[:4]) {
		return myGovernor.UnpackStringTooLongError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// MyGovernorCheckpointUnorderedInsertion represents a CheckpointUnorderedInsertion error raised by the MyGovernor contract.
type MyGovernorCheckpointUnorderedInsertion struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CheckpointUnorderedInsertion()
func MyGovernorCheckpointUnorderedInsertionErrorID() common.Hash {
	return common.HexToHash("0x2520601d9d60b717c34a36ad270857824c5a1ebbfd08ff39aba6930089495cfa")
}

// UnpackCheckpointUnorderedInsertionError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CheckpointUnorderedInsertion()
func (myGovernor *MyGovernor) UnpackCheckpointUnorderedInsertionError(raw []byte) (*MyGovernorCheckpointUnorderedInsertion, error) {
	out := new(MyGovernorCheckpointUnorderedInsertion)
	if err := myGovernor.abi.UnpackIntoInterface(out, "CheckpointUnorderedInsertion", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyGovernorFailedCall represents a FailedCall error raised by the MyGovernor contract.
type MyGovernorFailedCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedCall()
func MyGovernorFailedCallErrorID() common.Hash {
	return common.HexToHash("0xd6bda27508c0fb6d8a39b4b122878dab26f731a7d4e4abe711dd3731899052a4")
}

// UnpackFailedCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedCall()
func (myGovernor *MyGovernor) UnpackFailedCallError(raw []byte) (*MyGovernorFailedCall, error) {
	out := new(MyGovernorFailedCall)
	if err := myGovernor.abi.UnpackIntoInterface(out, "FailedCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyGovernorGovernorAlreadyCastVote represents a GovernorAlreadyCastVote error raised by the MyGovernor contract.
type MyGovernorGovernorAlreadyCastVote struct {
	Voter common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorAlreadyCastVote(address voter)
func MyGovernorGovernorAlreadyCastVoteErrorID() common.Hash {
	return common.HexToHash("0x71c6af4932ed543cdb181fcbb800f4b9940a2ccceeaee5e6e141de8c50856ede")
}

// UnpackGovernorAlreadyCastVoteError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorAlreadyCastVote(address voter)
func (myGovernor *MyGovernor) UnpackGovernorAlreadyCastVoteError(raw []byte) (*MyGovernorGovernorAlreadyCastVote, error) {
	out := new(MyGovernorGovernorAlreadyCastVote)
	if err := myGovernor.abi.UnpackIntoInterface(out, "GovernorAlreadyCastVote", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyGovernorGovernorAlreadyQueuedProposal represents a GovernorAlreadyQueuedProposal error raised by the MyGovernor contract.
type MyGovernorGovernorAlreadyQueuedProposal struct {
	ProposalId *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorAlreadyQueuedProposal(uint256 proposalId)
func MyGovernorGovernorAlreadyQueuedProposalErrorID() common.Hash {
	return common.HexToHash("0xf20e7d374e58691196b2e081c7753efc94203ab3520c58efe153076e279fde0a")
}

// UnpackGovernorAlreadyQueuedProposalError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorAlreadyQueuedProposal(uint256 proposalId)
func (myGovernor *MyGovernor) UnpackGovernorAlreadyQueuedProposalError(raw []byte) (*MyGovernorGovernorAlreadyQueuedProposal, error) {
	out := new(MyGovernorGovernorAlreadyQueuedProposal)
	if err := myGovernor.abi.UnpackIntoInterface(out, "GovernorAlreadyQueuedProposal", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyGovernorGovernorDisabledDeposit represents a GovernorDisabledDeposit error raised by the MyGovernor contract.
type MyGovernorGovernorDisabledDeposit struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorDisabledDeposit()
func MyGovernorGovernorDisabledDepositErrorID() common.Hash {
	return common.HexToHash("0xe90a651e5fdea7022846d10b5f36994c22e1f46db4b5021aa3e26ad83b24bfd8")
}

// UnpackGovernorDisabledDepositError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorDisabledDeposit()
func (myGovernor *MyGovernor) UnpackGovernorDisabledDepositError(raw []byte) (*MyGovernorGovernorDisabledDeposit, error) {
	out := new(MyGovernorGovernorDisabledDeposit)
	if err := myGovernor.abi.UnpackIntoInterface(out, "GovernorDisabledDeposit", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyGovernorGovernorInsufficientProposerVotes represents a GovernorInsufficientProposerVotes error raised by the MyGovernor contract.
type MyGovernorGovernorInsufficientProposerVotes struct {
	Proposer  common.Address
	Votes     *big.Int
	Threshold *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorInsufficientProposerVotes(address proposer, uint256 votes, uint256 threshold)
func MyGovernorGovernorInsufficientProposerVotesErrorID() common.Hash {
	return common.HexToHash("0xc242ee16ab08d11dbce60e744efdbd91b4e07ac4c074d993992519795a6324d0")
}

// UnpackGovernorInsufficientProposerVotesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorInsufficientProposerVotes(address proposer, uint256 votes, uint256 threshold)
func (myGovernor *MyGovernor) UnpackGovernorInsufficientProposerVotesError(raw []byte) (*MyGovernorGovernorInsufficientProposerVotes, error) {
	out := new(MyGovernorGovernorInsufficientProposerVotes)
	if err := myGovernor.abi.UnpackIntoInterface(out, "GovernorInsufficientProposerVotes", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyGovernorGovernorInvalidProposalLength represents a GovernorInvalidProposalLength error raised by the MyGovernor contract.
type MyGovernorGovernorInvalidProposalLength struct {
	Targets   *big.Int
	Calldatas *big.Int
	Values    *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorInvalidProposalLength(uint256 targets, uint256 calldatas, uint256 values)
func MyGovernorGovernorInvalidProposalLengthErrorID() common.Hash {
	return common.HexToHash("0x447b05d0c41e339e22932be48ca2091a47f3c39df25e2152ad11a8729753b2b4")
}

// UnpackGovernorInvalidProposalLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorInvalidProposalLength(uint256 targets, uint256 calldatas, uint256 values)
func (myGovernor *MyGovernor) UnpackGovernorInvalidProposalLengthError(raw []byte) (*MyGovernorGovernorInvalidProposalLength, error) {
	out := new(MyGovernorGovernorInvalidProposalLength)
	if err := myGovernor.abi.UnpackIntoInterface(out, "GovernorInvalidProposalLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyGovernorGovernorInvalidQuorumFraction represents a GovernorInvalidQuorumFraction error raised by the MyGovernor contract.
type MyGovernorGovernorInvalidQuorumFraction struct {
	QuorumNumerator   *big.Int
	QuorumDenominator *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorInvalidQuorumFraction(uint256 quorumNumerator, uint256 quorumDenominator)
func MyGovernorGovernorInvalidQuorumFractionErrorID() common.Hash {
	return common.HexToHash("0x243e5445050913bb1c3de7a2f82eba0c3b0b8a55c2aacf660392fa35087a1919")
}

// UnpackGovernorInvalidQuorumFractionError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorInvalidQuorumFraction(uint256 quorumNumerator, uint256 quorumDenominator)
func (myGovernor *MyGovernor) UnpackGovernorInvalidQuorumFractionError(raw []byte) (*MyGovernorGovernorInvalidQuorumFraction, error) {
	out := new(MyGovernorGovernorInvalidQuorumFraction)
	if err := myGovernor.abi.UnpackIntoInterface(out, "GovernorInvalidQuorumFraction", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyGovernorGovernorInvalidSignature represents a GovernorInvalidSignature error raised by the MyGovernor contract.
type MyGovernorGovernorInvalidSignature struct {
	Voter common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorInvalidSignature(address voter)
func MyGovernorGovernorInvalidSignatureErrorID() common.Hash {
	return common.HexToHash("0x94ab6c07905fb25046d2809e4563b09690f891c9495bfe19551d602e7eddbb1b")
}

// UnpackGovernorInvalidSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorInvalidSignature(address voter)
func (myGovernor *MyGovernor) UnpackGovernorInvalidSignatureError(raw []byte) (*MyGovernorGovernorInvalidSignature, error) {
	out := new(MyGovernorGovernorInvalidSignature)
	if err := myGovernor.abi.UnpackIntoInterface(out, "GovernorInvalidSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyGovernorGovernorInvalidVoteParams represents a GovernorInvalidVoteParams error raised by the MyGovernor contract.
type MyGovernorGovernorInvalidVoteParams struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorInvalidVoteParams()
func MyGovernorGovernorInvalidVoteParamsErrorID() common.Hash {
	return common.HexToHash("0x867db7717d0cc803be5726127d33cc17ae07776705d8ba6659af8aaa5b418dd8")
}

// UnpackGovernorInvalidVoteParamsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorInvalidVoteParams()
func (myGovernor *MyGovernor) UnpackGovernorInvalidVoteParamsError(raw []byte) (*MyGovernorGovernorInvalidVoteParams, error) {
	out := new(MyGovernorGovernorInvalidVoteParams)
	if err := myGovernor.abi.UnpackIntoInterface(out, "GovernorInvalidVoteParams", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyGovernorGovernorInvalidVoteType represents a GovernorInvalidVoteType error raised by the MyGovernor contract.
type MyGovernorGovernorInvalidVoteType struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorInvalidVoteType()
func MyGovernorGovernorInvalidVoteTypeErrorID() common.Hash {
	return common.HexToHash("0x06b337c26289d63178b4d4ed5fd513f38a1d8832d0edd309ef07bfc9ba5caf49")
}

// UnpackGovernorInvalidVoteTypeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorInvalidVoteType()
func (myGovernor *MyGovernor) UnpackGovernorInvalidVoteTypeError(raw []byte) (*MyGovernorGovernorInvalidVoteType, error) {
	out := new(MyGovernorGovernorInvalidVoteType)
	if err := myGovernor.abi.UnpackIntoInterface(out, "GovernorInvalidVoteType", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyGovernorGovernorInvalidVotingPeriod represents a GovernorInvalidVotingPeriod error raised by the MyGovernor contract.
type MyGovernorGovernorInvalidVotingPeriod struct {
	VotingPeriod *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorInvalidVotingPeriod(uint256 votingPeriod)
func MyGovernorGovernorInvalidVotingPeriodErrorID() common.Hash {
	return common.HexToHash("0xf1cfbf057db43f9730bc42a52728d66da9151a5c6929758ee824e299f82f5689")
}

// UnpackGovernorInvalidVotingPeriodError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorInvalidVotingPeriod(uint256 votingPeriod)
func (myGovernor *MyGovernor) UnpackGovernorInvalidVotingPeriodError(raw []byte) (*MyGovernorGovernorInvalidVotingPeriod, error) {
	out := new(MyGovernorGovernorInvalidVotingPeriod)
	if err := myGovernor.abi.UnpackIntoInterface(out, "GovernorInvalidVotingPeriod", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyGovernorGovernorNonexistentProposal represents a GovernorNonexistentProposal error raised by the MyGovernor contract.
type MyGovernorGovernorNonexistentProposal struct {
	ProposalId *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorNonexistentProposal(uint256 proposalId)
func MyGovernorGovernorNonexistentProposalErrorID() common.Hash {
	return common.HexToHash("0x6ad06075316ea071ccae80931b756598be5aad3433b2c47b38607a8eec344a70")
}

// UnpackGovernorNonexistentProposalError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorNonexistentProposal(uint256 proposalId)
func (myGovernor *MyGovernor) UnpackGovernorNonexistentProposalError(raw []byte) (*MyGovernorGovernorNonexistentProposal, error) {
	out := new(MyGovernorGovernorNonexistentProposal)
	if err := myGovernor.abi.UnpackIntoInterface(out, "GovernorNonexistentProposal", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyGovernorGovernorNotQueuedProposal represents a GovernorNotQueuedProposal error raised by the MyGovernor contract.
type MyGovernorGovernorNotQueuedProposal struct {
	ProposalId *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorNotQueuedProposal(uint256 proposalId)
func MyGovernorGovernorNotQueuedProposalErrorID() common.Hash {
	return common.HexToHash("0xd5ddb8255ec3d5fb4ee2dd5d919eb1db6a2f1e4420bb74c3741830500cfb6a4f")
}

// UnpackGovernorNotQueuedProposalError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorNotQueuedProposal(uint256 proposalId)
func (myGovernor *MyGovernor) UnpackGovernorNotQueuedProposalError(raw []byte) (*MyGovernorGovernorNotQueuedProposal, error) {
	out := new(MyGovernorGovernorNotQueuedProposal)
	if err := myGovernor.abi.UnpackIntoInterface(out, "GovernorNotQueuedProposal", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyGovernorGovernorOnlyExecutor represents a GovernorOnlyExecutor error raised by the MyGovernor contract.
type MyGovernorGovernorOnlyExecutor struct {
	Account common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorOnlyExecutor(address account)
func MyGovernorGovernorOnlyExecutorErrorID() common.Hash {
	return common.HexToHash("0x47096e4749c231af946d5efc74a7fd817371e756031e98787f18bf70aaa7753c")
}

// UnpackGovernorOnlyExecutorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorOnlyExecutor(address account)
func (myGovernor *MyGovernor) UnpackGovernorOnlyExecutorError(raw []byte) (*MyGovernorGovernorOnlyExecutor, error) {
	out := new(MyGovernorGovernorOnlyExecutor)
	if err := myGovernor.abi.UnpackIntoInterface(out, "GovernorOnlyExecutor", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyGovernorGovernorQueueNotImplemented represents a GovernorQueueNotImplemented error raised by the MyGovernor contract.
type MyGovernorGovernorQueueNotImplemented struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorQueueNotImplemented()
func MyGovernorGovernorQueueNotImplementedErrorID() common.Hash {
	return common.HexToHash("0x90884a46490684db0a73766419939e5ace793ae0f80195a286e159115c6628a0")
}

// UnpackGovernorQueueNotImplementedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorQueueNotImplemented()
func (myGovernor *MyGovernor) UnpackGovernorQueueNotImplementedError(raw []byte) (*MyGovernorGovernorQueueNotImplemented, error) {
	out := new(MyGovernorGovernorQueueNotImplemented)
	if err := myGovernor.abi.UnpackIntoInterface(out, "GovernorQueueNotImplemented", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyGovernorGovernorRestrictedProposer represents a GovernorRestrictedProposer error raised by the MyGovernor contract.
type MyGovernorGovernorRestrictedProposer struct {
	Proposer common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorRestrictedProposer(address proposer)
func MyGovernorGovernorRestrictedProposerErrorID() common.Hash {
	return common.HexToHash("0xd9b395579c6f1566cc7608394c53136f366f7fa719d01a941bee075ef8c704f4")
}

// UnpackGovernorRestrictedProposerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorRestrictedProposer(address proposer)
func (myGovernor *MyGovernor) UnpackGovernorRestrictedProposerError(raw []byte) (*MyGovernorGovernorRestrictedProposer, error) {
	out := new(MyGovernorGovernorRestrictedProposer)
	if err := myGovernor.abi.UnpackIntoInterface(out, "GovernorRestrictedProposer", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyGovernorGovernorUnableToCancel represents a GovernorUnableToCancel error raised by the MyGovernor contract.
type MyGovernorGovernorUnableToCancel struct {
	ProposalId *big.Int
	Account    common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorUnableToCancel(uint256 proposalId, address account)
func MyGovernorGovernorUnableToCancelErrorID() common.Hash {
	return common.HexToHash("0x8fe5d8a9809b4b1a3569a8d98ce25e21fb956a89b6b9a2741d15bc699f46ef07")
}

// UnpackGovernorUnableToCancelError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorUnableToCancel(uint256 proposalId, address account)
func (myGovernor *MyGovernor) UnpackGovernorUnableToCancelError(raw []byte) (*MyGovernorGovernorUnableToCancel, error) {
	out := new(MyGovernorGovernorUnableToCancel)
	if err := myGovernor.abi.UnpackIntoInterface(out, "GovernorUnableToCancel", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyGovernorGovernorUnexpectedProposalState represents a GovernorUnexpectedProposalState error raised by the MyGovernor contract.
type MyGovernorGovernorUnexpectedProposalState struct {
	ProposalId     *big.Int
	Current        uint8
	ExpectedStates [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorUnexpectedProposalState(uint256 proposalId, uint8 current, bytes32 expectedStates)
func MyGovernorGovernorUnexpectedProposalStateErrorID() common.Hash {
	return common.HexToHash("0x31b75e4d4f8317c390cf01cbc79dfe4f67ce2d27f65a099074fdc67f00f76908")
}

// UnpackGovernorUnexpectedProposalStateError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorUnexpectedProposalState(uint256 proposalId, uint8 current, bytes32 expectedStates)
func (myGovernor *MyGovernor) UnpackGovernorUnexpectedProposalStateError(raw []byte) (*MyGovernorGovernorUnexpectedProposalState, error) {
	out := new(MyGovernorGovernorUnexpectedProposalState)
	if err := myGovernor.abi.UnpackIntoInterface(out, "GovernorUnexpectedProposalState", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyGovernorInvalidAccountNonce represents a InvalidAccountNonce error raised by the MyGovernor contract.
type MyGovernorInvalidAccountNonce struct {
	Account      common.Address
	CurrentNonce *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func MyGovernorInvalidAccountNonceErrorID() common.Hash {
	return common.HexToHash("0x752d88c0de02638abf10e8e31861e4c68dc1f3a1e7d840e580683f2c282bfc7a")
}

// UnpackInvalidAccountNonceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func (myGovernor *MyGovernor) UnpackInvalidAccountNonceError(raw []byte) (*MyGovernorInvalidAccountNonce, error) {
	out := new(MyGovernorInvalidAccountNonce)
	if err := myGovernor.abi.UnpackIntoInterface(out, "InvalidAccountNonce", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyGovernorInvalidShortString represents a InvalidShortString error raised by the MyGovernor contract.
type MyGovernorInvalidShortString struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidShortString()
func MyGovernorInvalidShortStringErrorID() common.Hash {
	return common.HexToHash("0xb3512b0c6163e5f0bafab72bb631b9d58cd7a731b082f910338aa21c83d5c274")
}

// UnpackInvalidShortStringError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidShortString()
func (myGovernor *MyGovernor) UnpackInvalidShortStringError(raw []byte) (*MyGovernorInvalidShortString, error) {
	out := new(MyGovernorInvalidShortString)
	if err := myGovernor.abi.UnpackIntoInterface(out, "InvalidShortString", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyGovernorSafeCastOverflowedUintDowncast represents a SafeCastOverflowedUintDowncast error raised by the MyGovernor contract.
type MyGovernorSafeCastOverflowedUintDowncast struct {
	Bits  uint8
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func MyGovernorSafeCastOverflowedUintDowncastErrorID() common.Hash {
	return common.HexToHash("0x6dfcc6503a32754ce7a89698e18201fc5294fd4aad43edefee786f88423b1a12")
}

// UnpackSafeCastOverflowedUintDowncastError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func (myGovernor *MyGovernor) UnpackSafeCastOverflowedUintDowncastError(raw []byte) (*MyGovernorSafeCastOverflowedUintDowncast, error) {
	out := new(MyGovernorSafeCastOverflowedUintDowncast)
	if err := myGovernor.abi.UnpackIntoInterface(out, "SafeCastOverflowedUintDowncast", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyGovernorStringTooLong represents a StringTooLong error raised by the MyGovernor contract.
type MyGovernorStringTooLong struct {
	Str string
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error StringTooLong(string str)
func MyGovernorStringTooLongErrorID() common.Hash {
	return common.HexToHash("0x305a27a93f8e33b7392df0a0f91d6fc63847395853c45991eec52dbf24d72381")
}

// UnpackStringTooLongError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error StringTooLong(string str)
func (myGovernor *MyGovernor) UnpackStringTooLongError(raw []byte) (*MyGovernorStringTooLong, error) {
	out := new(MyGovernorStringTooLong)
	if err := myGovernor.abi.UnpackIntoInterface(out, "StringTooLong", raw); err != nil {
		return nil, err
	}
	return out, nil
}
