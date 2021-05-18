package types

// nft module event types
const (
	EventTypeTransfer           = "transfer_nft"
	EventTypeEditNFTMetadata    = "edit_nft_metadata"
	EventTypeEditNFTDigitalHash = "edit_nft_digital_hash"
	EventTypeMintNFT            = "mint_nft"
	EventTypeBurnNFT            = "burn_nft"

	AttributeValueCategory = ModuleName

	AttributeKeySender           = "sender"
	AttributeKeyRecipient        = "recipient"
	AttributeKeyOwner            = "owner"
	AttributeKeyNFTID            = "nft-id"
	AttributeKeyNFTDByigitalHash = "nft-digital-hash"
	AttributeKeyNFTTokenURI      = "token-uri"
	AttributeKeyNFTDigitalHash   = "digital-hash"
	AttributeKeyDenom            = "denom"
)
