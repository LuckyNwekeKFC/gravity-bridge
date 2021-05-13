package keeper

import (
	"context"
	"encoding/hex"
	"fmt"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/cosmos/gravity-bridge/module/x/gravity/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the gov MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) DelegateKeys(c context.Context, msg *types.MsgDelegateKeys) (*types.MsgDelegateKeysResponse, error) {
	// ensure that this passes validation
	err := msg.ValidateBasic()
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(c)
	val, _ := sdk.ValAddressFromBech32(msg.Validator)
	orch, _ := sdk.AccAddressFromBech32(msg.Orchestrator)

	// ensure that the validator exists
	if k.Keeper.StakingKeeper.Validator(ctx, val) == nil {
		return nil, sdkerrors.Wrap(stakingtypes.ErrNoValidatorFound, val.String())
	}

	// TODO consider impact of maliciously setting duplicate delegate
	// addresses since no signatures from the private keys of these addresses
	// are required for this message it could be sent in a hostile way.

	// set the orchestrator address
	k.SetOrchestratorValidator(ctx, val, orch)
	// set the ethereum address
	k.SetEthAddressForValidator(ctx, val, msg.EthAddress)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Type()),
			sdk.NewAttribute(types.AttributeKeySetOperatorAddr, orch.String()),
		),
	)

	return &types.MsgDelegateKeysResponse{}, nil

}

// TODO: check msgSignerSetTxSignature to have an Orchestrator field instead of a Validator field
func (k msgServer) SignerSetTxSignature(c context.Context, msg *types.MsgSignerSetTxSignature) (*types.MsgSignerSetTxSignatureResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	signerSetTx := k.GetSignerSetTx(ctx, msg.Nonce)
	if signerSetTx == nil {
		return nil, sdkerrors.Wrap(types.ErrInvalid, "couldn't find signerSetTx")
	}

	gravityID := k.GetGravityID(ctx)
	checkpoint := signerSetTx.GetCheckpoint(gravityID)

	sigBytes, err := hex.DecodeString(msg.Signature)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalid, "signature decoding")
	}

	orchaddr, _ := sdk.AccAddressFromBech32(msg.Orchestrator)
	validator := k.GetOrchestratorValidator(ctx, orchaddr)
	if validator == nil {
		return nil, sdkerrors.Wrap(types.ErrUnknown, "validator")
	}

	ethAddress := k.GetEthAddressByValidator(ctx, validator)
	if ethAddress == "" {
		return nil, sdkerrors.Wrap(types.ErrEmpty, "eth address")
	}

	if err = types.ValidateEthereumSignature(checkpoint, sigBytes, ethAddress); err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalid, fmt.Sprintf("signature verification failed expected sig by %s with gravity-id %s with checkpoint %s found %s", ethAddress, gravityID, hex.EncodeToString(checkpoint), msg.Signature))
	}

	// persist signature
	if k.GetSignerSetTxSignature(ctx, msg.Nonce, orchaddr) != nil {
		return nil, sdkerrors.Wrap(types.ErrDuplicate, "signature duplicate")
	}
	key := k.SetSignerSetTxSignature(ctx, *msg)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Type()),
			sdk.NewAttribute(types.AttributeKeySignerSetTxSignatureKey, string(key)),
		),
	)

	return &types.MsgSignerSetTxSignatureResponse{}, nil
}

func (k msgServer) SendToEthereum(c context.Context, msg *types.MsgSendToEthereum) (*types.MsgSendToEthereumResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}
	txID, err := k.AddToSendToEthereumPool(ctx, sender, msg.EthDest, msg.Amount, msg.BridgeFee)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Type()),
			sdk.NewAttribute(types.AttributeKeySendToEthereumID, fmt.Sprint(txID)),
		),
	)

	return &types.MsgSendToEthereumResponse{}, nil
}

func (k msgServer) RequestBatchTx(c context.Context, msg *types.MsgRequestBatchTx) (*types.MsgRequestBatchTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	// Check if the denom is a gravity coin, if not, check if there is a deployed ERC20 representing it.
	// If not, error out
	_, tokenContract, err := k.DenomToERC20Lookup(ctx, msg.Denom)
	if err != nil {
		return nil, err
	}

	batch, err := k.BuildBatchTx(ctx, tokenContract, BatchTxSize)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Type()),
			sdk.NewAttribute(types.AttributeKeyBatchNonce, fmt.Sprint(batch.BatchNonce)),
		),
	)

	return &types.MsgRequestBatchTxResponse{}, nil
}

// BatchTxSignature handles MsgBatchTxSignature
func (k msgServer) BatchTxSignature(c context.Context, msg *types.MsgBatchTxSignature) (*types.MsgBatchTxSignatureResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	// fetch the batch tx the nonce
	batch := k.GetBatchTx(ctx, msg.TokenContract, msg.Nonce)
	if batch == nil {
		return nil, sdkerrors.Wrap(types.ErrInvalid, "couldn't find batch")
	}

	gravityID := k.GetGravityID(ctx)
	checkpoint := batch.GetCheckpoint(gravityID)

	sigBytes, err := hex.DecodeString(msg.Signature)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalid, "signature decoding")
	}

	orchaddr, _ := sdk.AccAddressFromBech32(msg.Orchestrator)
	validator := k.GetOrchestratorValidator(ctx, orchaddr)
	if validator == nil {
		return nil, sdkerrors.Wrap(types.ErrUnknown, "validator")
	}

	ethAddress := k.GetEthAddressByValidator(ctx, validator)
	if ethAddress == "" {
		return nil, sdkerrors.Wrap(types.ErrEmpty, "eth address")
	}

	err = types.ValidateEthereumSignature(checkpoint, sigBytes, ethAddress)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalid, fmt.Sprintf("signature verification failed expected sig by %s with gravity-id %s with checkpoint %s found %s", ethAddress, gravityID, hex.EncodeToString(checkpoint), msg.Signature))
	}

	// check if we already have this signature
	if k.GetBatchTxSignature(ctx, msg.Nonce, msg.TokenContract, orchaddr) != nil {
		return nil, sdkerrors.Wrap(types.ErrDuplicate, "duplicate signature")
	}
	key := k.SetBatchTxSignature(ctx, msg)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Type()),
			sdk.NewAttribute(types.AttributeKeyBatchTxSignatureKey, string(key)),
		),
	)

	return nil, nil
}

// ContractCallTxSignature handles MsgContractCallTxSignature
func (k msgServer) ContractCallTxSignature(c context.Context, msg *types.MsgContractCallTxSignature) (*types.MsgContractCallTxSignatureResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	invalidationIdBytes, err := hex.DecodeString(msg.InvalidationId)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalid, "invalidation id encoding")
	}

	// fetch the contract call given the nonce
	contractCall := k.GetContractCallTx(ctx, invalidationIdBytes, msg.InvalidationNonce)
	if contractCall == nil {
		return nil, sdkerrors.Wrap(types.ErrInvalid, "couldn't find contractCall")
	}

	gravityID := k.GetGravityID(ctx)
	checkpoint := contractCall.GetCheckpoint(gravityID)

	sigBytes, err := hex.DecodeString(msg.Signature)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalid, "signature decoding")
	}

	orchaddr, _ := sdk.AccAddressFromBech32(msg.Orchestrator)
	validator := k.GetOrchestratorValidator(ctx, orchaddr)
	if validator == nil {
		return nil, sdkerrors.Wrap(types.ErrUnknown, "validator")
	}

	ethAddress := k.GetEthAddressByValidator(ctx, validator)
	if ethAddress == "" {
		return nil, sdkerrors.Wrap(types.ErrEmpty, "eth address")
	}

	err = types.ValidateEthereumSignature(checkpoint, sigBytes, ethAddress)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalid, fmt.Sprintf("signature verification failed expected sig by %s with gravity-id %s with checkpoint %s found %s", ethAddress, gravityID, hex.EncodeToString(checkpoint), msg.Signature))
	}

	// check if we already have this signature
	if k.GetContractCallTxSignature(ctx, invalidationIdBytes, msg.InvalidationNonce, orchaddr) != nil {
		return nil, sdkerrors.Wrap(types.ErrDuplicate, "duplicate signature")
	}

	k.SetContractCallTxSignature(ctx, msg)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Type()),
		),
	)

	return nil, nil
}

// SendToCosmosEvent handles MsgSendToCosmosEvent
// TODO it is possible to submit an old msgSendToCosmosEvent (old defined as covering an event nonce that has already been
// executed aka 'accepted' and had it's slashing window expire) that will never be cleaned up in the endblocker. This
// should not be a security risk as 'old' events can never execute but it does store spam in the chain.
func (k msgServer) SendToCosmosEvent(c context.Context, msg *types.MsgSendToCosmosEvent) (*types.MsgSendToCosmosEventResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	orchaddr, _ := sdk.AccAddressFromBech32(msg.Orchestrator)
	validator := k.GetOrchestratorValidator(ctx, orchaddr)
	if validator == nil {
		return nil, sdkerrors.Wrap(types.ErrUnknown, "validator")
	}

	// return an error if the validator isn't in the active set
	val := k.StakingKeeper.Validator(ctx, validator)
	if val == nil || !val.IsBonded() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrorInvalidSigner, "validator not in active set")
	}

	any, err := codectypes.NewAnyWithValue(msg)
	if err != nil {
		return nil, err
	}

	// Add the event to the store
	_, err = k.Vote(ctx, msg, any)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "create ethereumEventVoteRecord")
	}

	// Emit the handle message event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Type()),
			// TODO: maybe return something better here? is this the right string representation?
			sdk.NewAttribute(types.AttributeKeyEthereumEventVoteRecordID, string(types.GetEthereumEventVoteRecordKey(msg.EventNonce, msg.EventHash()))),
		),
	)

	return &types.MsgSendToCosmosEventResponse{}, nil
}

// BatchExecutedEvent handles MsgBatchExecutedEvent
// TODO it is possible to submit an old msgBatchExecutedEvent (old defined as covering an event nonce that has already been
// executed aka 'accepted' and had it's slashing window expire) that will never be cleaned up in the endblocker. This
// should not be a security risk as 'old' events can never execute but it does store spam in the chain.
func (k msgServer) BatchExecutedEvent(c context.Context, msg *types.MsgBatchExecutedEvent) (*types.MsgBatchExecutedEventResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	orchaddr, _ := sdk.AccAddressFromBech32(msg.Orchestrator)
	validator := k.GetOrchestratorValidator(ctx, orchaddr)
	if validator == nil {
		return nil, sdkerrors.Wrap(types.ErrUnknown, "validator")
	}

	// return an error if the validator isn't in the active set
	val := k.StakingKeeper.Validator(ctx, validator)
	if val == nil || !val.IsBonded() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrorInvalidSigner, "validator not in acitve set")
	}

	any, err := codectypes.NewAnyWithValue(msg)
	if err != nil {
		return nil, err
	}

	// Add the event to the store
	_, err = k.Vote(ctx, msg, any)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "create ethereumEventVoteRecord")
	}

	// Emit the handle message event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Type()),
			// TODO: maybe return something better here? is this the right string representation?
			sdk.NewAttribute(types.AttributeKeyEthereumEventVoteRecordID, string(types.GetEthereumEventVoteRecordKey(msg.EventNonce, msg.EventHash()))),
		),
	)

	return &types.MsgBatchExecutedEventResponse{}, nil
}

// ERC20Deployed handles MsgERC20Deployed
func (k msgServer) ERC20DeployedEvent(c context.Context, msg *types.MsgERC20DeployedEvent) (*types.MsgERC20DeployedEventResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	orchaddr, _ := sdk.AccAddressFromBech32(msg.Orchestrator)
	validator := k.GetOrchestratorValidator(ctx, orchaddr)
	if validator == nil {
		return nil, sdkerrors.Wrap(types.ErrUnknown, "validator")
	}

	// return an error if the validator isn't in the active set
	val := k.StakingKeeper.Validator(ctx, validator)
	if val == nil || !val.IsBonded() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrorInvalidSigner, "validator not in acitve set")
	}

	any, err := codectypes.NewAnyWithValue(msg)
	if err != nil {
		return nil, err
	}

	// Add the event to the store
	_, err = k.Vote(ctx, msg, any)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "create ethereumEventVoteRecord")
	}

	// Emit the handle message event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Type()),
			// TODO: maybe return something better here? is this the right string representation?
			sdk.NewAttribute(types.AttributeKeyEthereumEventVoteRecordID, string(types.GetEthereumEventVoteRecordKey(msg.EventNonce, msg.EventHash()))),
		),
	)

	return &types.MsgERC20DeployedEventResponse{}, nil
}

// ContractCallExecutedEvent handles events for executing a contract call on Ethereum
func (k msgServer) ContractCallExecutedEvent(c context.Context, msg *types.MsgContractCallExecutedEvent) (*types.MsgContractCallExecutedEventResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	orchaddr, _ := sdk.AccAddressFromBech32(msg.Orchestrator)
	validator := k.GetOrchestratorValidator(ctx, orchaddr)
	if validator == nil {
		return nil, sdkerrors.Wrap(types.ErrUnknown, "validator")
	}

	// return an error if the validator isn't in the active set
	val := k.StakingKeeper.Validator(ctx, validator)
	if val == nil || !val.IsBonded() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrorInvalidSigner, "validator not in acitve set")
	}

	any, err := codectypes.NewAnyWithValue(msg)
	if err != nil {
		return nil, err
	}

	// Add the event to the store
	_, err = k.Vote(ctx, msg, any)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "create ethereumEventVoteRecord")
	}

	// Emit the handle message event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Type()),
			// TODO: maybe return something better here? is this the right string representation?
			sdk.NewAttribute(types.AttributeKeyEthereumEventVoteRecordID, string(types.GetEthereumEventVoteRecordKey(msg.EventNonce, msg.EventHash()))),
		),
	)

	return &types.MsgContractCallExecutedEventResponse{}, nil
}

// SignerSetUpdatedEvent handles the event that is fired when a signer set is updated
func (k msgServer) SignerSetUpdatedEvent(c context.Context, msg *types.MsgSignerSetUpdatedEvent) (*types.MsgSignerSetUpdatedEventResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	orchaddr, _ := sdk.AccAddressFromBech32(msg.Orchestrator)
	validator := k.GetOrchestratorValidator(ctx, orchaddr)
	if validator == nil {
		return nil, sdkerrors.Wrap(types.ErrUnknown, "validator")
	}

	// return an error if the validator isn't in the active set
	val := k.StakingKeeper.Validator(ctx, validator)
	if val == nil || !val.IsBonded() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrorInvalidSigner, "validator not in acitve set")
	}

	any, err := codectypes.NewAnyWithValue(msg)
	if err != nil {
		return nil, err
	}

	// Add the event to the store
	_, err = k.Vote(ctx, msg, any)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "create ethereumEventVoteRecord")
	}

	// Emit the handle message event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Type()),
			// TODO: maybe return something better here? is this the right string representation?
			sdk.NewAttribute(types.AttributeKeyEthereumEventVoteRecordID, string(types.GetEthereumEventVoteRecordKey(msg.EventNonce, msg.EventHash()))),
		),
	)

	return &types.MsgSignerSetUpdatedEventResponse{}, nil
}

func (k msgServer) CancelSendToEthereum(c context.Context, msg *types.MsgCancelSendToEthereum) (*types.MsgCancelSendToEthereumResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}
	err = k.RemoveFromAddToSendToEthereumPoolAndRefund(ctx, msg.TransactionId, sender)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Type()),
			sdk.NewAttribute(types.AttributeKeySendToEthereumID, fmt.Sprint(msg.TransactionId)),
		),
	)

	return &types.MsgCancelSendToEthereumResponse{}, nil
}

func (k msgServer) SubmitBadEthereumSignatureEvidence(c context.Context, msg *types.MsgSubmitBadEthereumSignatureEvidence) (*types.MsgSubmitBadEthereumSignatureEvidenceResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	err := k.CheckBadSignatureEvidence(ctx, msg)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Type()),
			sdk.NewAttribute(types.AttributeKeyBadEthSignature, fmt.Sprint(msg.Signature)),
			sdk.NewAttribute(types.AttributeKeyBadEthSignatureSubject, fmt.Sprint(msg.Subject)),
		),
	)

	return &types.MsgSubmitBadEthereumSignatureEvidenceResponse{}, err
}
