package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// SubModuleName is the error codespace
const SubModuleName string = "commitment"

// IBC connection sentinel errors
var (
	ErrInvalidProof       = sdkerrors.Register(SubModuleName, 50, "invalid proof")
	ErrInvalidPrefix      = sdkerrors.Register(SubModuleName, 51, "invalid prefix")
	ErrInvalidMerkleProof = sdkerrors.Register(SubModuleName, 52, "invalid merkle proof")
)
