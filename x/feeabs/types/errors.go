package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidExchangeRate    = sdkerrors.Register(ModuleName, 1, "invalid exchange rate")
	ErrInvalidIBCFees         = sdkerrors.Register(ModuleName, 2, "invalid ibc fees")
	ErrHostZoneConfigNotFound = sdkerrors.Register(ModuleName, 3, "host chain config not found")
)
