package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// IBC client sentinel errors
var (
	ErrClientExists                           = sdkerrors.Register(SubModuleName, 50, "light client already exists")
	ErrInvalidClient                          = sdkerrors.Register(SubModuleName, 51, "light client is invalid")
	ErrClientNotFound                         = sdkerrors.Register(SubModuleName, 52, "light client not found")
	ErrClientFrozen                           = sdkerrors.Register(SubModuleName, 53, "light client is frozen due to misbehaviour")
	ErrInvalidClientMetadata                  = sdkerrors.Register(SubModuleName, 54, "invalid client metadata")
	ErrConsensusStateNotFound                 = sdkerrors.Register(SubModuleName, 55, "consensus state not found")
	ErrInvalidConsensus                       = sdkerrors.Register(SubModuleName, 56, "invalid consensus state")
	ErrClientTypeNotFound                     = sdkerrors.Register(SubModuleName, 57, "client type not found")
	ErrInvalidClientType                      = sdkerrors.Register(SubModuleName, 58, "invalid client type")
	ErrRootNotFound                           = sdkerrors.Register(SubModuleName, 59, "commitment root not found")
	ErrInvalidHeader                          = sdkerrors.Register(SubModuleName, 60, "invalid client header")
	ErrInvalidMisbehaviour                    = sdkerrors.Register(SubModuleName, 61, "invalid light client misbehaviour")
	ErrFailedClientStateVerification          = sdkerrors.Register(SubModuleName, 62, "client state verification failed")
	ErrFailedClientConsensusStateVerification = sdkerrors.Register(SubModuleName, 63, "client consensus state verification failed")
	ErrFailedConnectionStateVerification      = sdkerrors.Register(SubModuleName, 64, "connection state verification failed")
	ErrFailedChannelStateVerification         = sdkerrors.Register(SubModuleName, 65, "channel state verification failed")
	ErrFailedPacketCommitmentVerification     = sdkerrors.Register(SubModuleName, 66, "packet commitment verification failed")
	ErrFailedPacketAckVerification            = sdkerrors.Register(SubModuleName, 67, "packet acknowledgement verification failed")
	ErrFailedPacketReceiptVerification        = sdkerrors.Register(SubModuleName, 68, "packet receipt verification failed")
	ErrFailedNextSeqRecvVerification          = sdkerrors.Register(SubModuleName, 69, "next sequence receive verification failed")
	ErrSelfConsensusStateNotFound             = sdkerrors.Register(SubModuleName, 70, "self consensus state not found")
	ErrUpdateClientFailed                     = sdkerrors.Register(SubModuleName, 71, "unable to update light client")
	ErrInvalidUpdateClientProposal            = sdkerrors.Register(SubModuleName, 72, "invalid update client proposal")
	ErrInvalidUpgradeClient                   = sdkerrors.Register(SubModuleName, 73, "invalid client upgrade")
	ErrInvalidHeight                          = sdkerrors.Register(SubModuleName, 74, "invalid height")
	ErrInvalidSubstitute                      = sdkerrors.Register(SubModuleName, 75, "invalid client state substitute")
	ErrInvalidUpgradeProposal                 = sdkerrors.Register(SubModuleName, 76, "invalid upgrade proposal")
	ErrClientNotActive                        = sdkerrors.Register(SubModuleName, 77, "client is not active")
)
