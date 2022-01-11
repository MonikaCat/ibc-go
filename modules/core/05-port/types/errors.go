package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// IBC port sentinel errors
var (
	ErrPortExists   = sdkerrors.Register(SubModuleName, 50, "port is already binded")
	ErrPortNotFound = sdkerrors.Register(SubModuleName, 51, "port not found")
	ErrInvalidPort  = sdkerrors.Register(SubModuleName, 52, "invalid port")
	ErrInvalidRoute = sdkerrors.Register(SubModuleName, 53, "route not found")
)
