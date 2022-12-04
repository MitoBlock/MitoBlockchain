package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/mitoblockchaindev module sentinel errors
var (
	ErrTokenStatusOld = sdkerrors.Register(ModuleName, 1300, "")
	ErrID             = sdkerrors.Register(ModuleName, 1400, "")
)
