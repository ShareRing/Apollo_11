package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// You can see how they are constructed below:
var (
	ErrInvalidCollection = sdkerrors.Register(ModuleName, 1, "invalid NFT collection")
	ErrUnknownCollection = sdkerrors.Register(ModuleName, 2, "unknown NFT collection")
	ErrInvalidNFT        = sdkerrors.Register(ModuleName, 3, "invalid NFT")
	ErrUnknownNFT        = sdkerrors.Register(ModuleName, 4, "unknown NFT")
	ErrNFTAlreadyExists  = sdkerrors.Register(ModuleName, 5, "NFT already exists")
	ErrEmptyMetadata     = sdkerrors.Register(ModuleName, 6, "NFT metadata can't be empty")
	ErrWrongURI          = sdkerrors.Register(ModuleName, 7, "wrong URI")
	ErrWrongDigitalHash  = sdkerrors.Register(ModuleName, 8, "wrong digital hash")
)
