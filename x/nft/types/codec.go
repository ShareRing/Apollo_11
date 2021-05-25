package types

import (
	"github.com/ShareRing/nft-module/x/nft/exported"
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterInterface((*exported.NFT)(nil), nil)
	cdc.RegisterConcrete(&BaseNFT{}, "cosmos-sdk/BaseNFT", nil)
	cdc.RegisterConcrete(&ExtraBaseNFT{}, "cosmos-sdk/ExtraBaseNFT", nil)
	cdc.RegisterConcrete(&IDCollection{}, "cosmos-sdk/IDCollection", nil)
	cdc.RegisterConcrete(&Collection{}, "cosmos-sdk/Collection", nil)
	cdc.RegisterConcrete(&Owner{}, "cosmos-sdk/Owner", nil)
	cdc.RegisterConcrete(MsgTransferNFT{}, "cosmos-sdk/MsgTransferNFT", nil)
	cdc.RegisterConcrete(MsgEditNFTMetadata{}, "cosmos-sdk/MsgEditNFTMetadata", nil)
	cdc.RegisterConcrete(MsgEditNFTDigitalHash{}, "cosmos-sdk/MsgEditNFTDigitalHash", nil)
	cdc.RegisterConcrete(MsgMintNFT{}, "cosmos-sdk/MsgMintNFT", nil)
	cdc.RegisterConcrete(MsgBurnNFT{}, "cosmos-sdk/MsgBurnNFT", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
