package exported

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NFT non fungible token interface
type NFT interface {
	GetID() string
	GetOwner() sdk.AccAddress
	SetOwner(address sdk.AccAddress)
	GetTokenURI() string
	GetDigitalHash() string
	EditMetadata(tokenURI string)
	EditDigitalHash(digitalHash string)
	String() string
}
