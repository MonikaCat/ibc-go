package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// IBC connection sentinel errors
var (
	ErrConnectionExists              = sdkerrors.Register(SubModuleName, 50, "connection already exists")
	ErrConnectionNotFound            = sdkerrors.Register(SubModuleName, 51, "connection not found")
	ErrClientConnectionPathsNotFound = sdkerrors.Register(SubModuleName, 52, "light client connection paths not found")
	ErrConnectionPath                = sdkerrors.Register(SubModuleName, 53, "connection path is not associated to the given light client")
	ErrInvalidConnectionState        = sdkerrors.Register(SubModuleName, 54, "invalid connection state")
	ErrInvalidCounterparty           = sdkerrors.Register(SubModuleName, 55, "invalid counterparty connection")
	ErrInvalidConnection             = sdkerrors.Register(SubModuleName, 56, "invalid connection")
	ErrInvalidVersion                = sdkerrors.Register(SubModuleName, 57, "invalid connection version")
	ErrVersionNegotiationFailed      = sdkerrors.Register(SubModuleName, 58, "connection version negotiation failed")
	ErrInvalidConnectionIdentifier   = sdkerrors.Register(SubModuleName, 59, "invalid connection identifier")
)
