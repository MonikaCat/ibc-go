package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// IBC channel sentinel errors
var (
	ErrChannelExists             = sdkerrors.Register(SubModuleName, 50, "channel already exists")
	ErrChannelNotFound           = sdkerrors.Register(SubModuleName, 51, "channel not found")
	ErrInvalidChannel            = sdkerrors.Register(SubModuleName, 52, "invalid channel")
	ErrInvalidChannelState       = sdkerrors.Register(SubModuleName, 53, "invalid channel state")
	ErrInvalidChannelOrdering    = sdkerrors.Register(SubModuleName, 54, "invalid channel ordering")
	ErrInvalidCounterparty       = sdkerrors.Register(SubModuleName, 55, "invalid counterparty channel")
	ErrInvalidChannelCapability  = sdkerrors.Register(SubModuleName, 56, "invalid channel capability")
	ErrChannelCapabilityNotFound = sdkerrors.Register(SubModuleName, 57, "channel capability not found")
	ErrSequenceSendNotFound      = sdkerrors.Register(SubModuleName, 58, "sequence send not found")
	ErrSequenceReceiveNotFound   = sdkerrors.Register(SubModuleName, 59, "sequence receive not found")
	ErrSequenceAckNotFound       = sdkerrors.Register(SubModuleName, 60, "sequence acknowledgement not found")
	ErrInvalidPacket             = sdkerrors.Register(SubModuleName, 61, "invalid packet")
	ErrPacketTimeout             = sdkerrors.Register(SubModuleName, 62, "packet timeout")
	ErrTooManyConnectionHops     = sdkerrors.Register(SubModuleName, 63, "too many connection hops")
	ErrInvalidAcknowledgement    = sdkerrors.Register(SubModuleName, 64, "invalid acknowledgement")
	ErrAcknowledgementExists     = sdkerrors.Register(SubModuleName, 65, "acknowledgement for packet already exists")
	ErrInvalidChannelIdentifier  = sdkerrors.Register(SubModuleName, 66, "invalid channel identifier")

	// packets already relayed errors
	ErrPacketReceived           = sdkerrors.Register(SubModuleName, 67, "packet already received")
	ErrPacketCommitmentNotFound = sdkerrors.Register(SubModuleName, 68, "packet commitment not found") // may occur for already received acknowledgements or timeouts and in rare cases for packets never sent

	// ORDERED channel error
	ErrPacketSequenceOutOfOrder = sdkerrors.Register(SubModuleName, 69, "packet sequence is out of order")

	// Antehandler error
	ErrRedundantTx = sdkerrors.Register(SubModuleName, 70, "packet messages are redundant")

	// Perform a no-op on the current Msg
	ErrNoOpMsg = sdkerrors.Register(SubModuleName, 71, "message is redundant, no-op will be performed")

	ErrInvalidChannelVersion = sdkerrors.Register(SubModuleName, 72, "invalid channel version")
)
