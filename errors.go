package checkers

import "cosmossdk.io/errors"

var (
	ErrIndexTooLong     = errors.Register(ModuleName, 2, "index too long")
	ErrDuplicateAddress = errors.Register(ModuleName, 3, "duplicate address")
	ErrInvalidBlack     = errors.Register(ModuleName, 4, "black address is invalid: %s")
	ErrInvalidRed       = errors.Register(ModuleName, 5, "red address is invalid: %s")
	ErrGameNotParseable = errors.Register(ModuleName, 6, "game cannot be parsed")
)
