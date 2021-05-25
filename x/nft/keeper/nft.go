package keeper

import (
	"fmt"

	"github.com/ShareRing/nft-module/x/nft/exported"
	"github.com/ShareRing/nft-module/x/nft/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// IsNFT returns whether an NFT exists
func (k Keeper) IsNFT(ctx sdk.Context, denom, id string) (exists bool) {
	_, err := k.GetNFT(ctx, denom, id)
	return err == nil
}

// GetNFT gets the entire NFT metadata struct for a uint64
func (k Keeper) GetNFT(ctx sdk.Context, denom, id string) (nft exported.NFT, err error) {
	collection, found := k.GetCollection(ctx, denom)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrUnknownCollection, fmt.Sprintf("collection of %s doesn't exist", denom))
	}
	nft, err = collection.GetNFT(id)

	if err != nil {
		return nil, err
	}
	return nft, err
}

func (k Keeper) GetNFTByDigitalHash(ctx sdk.Context, digitalHash string) (nft *types.ExtraBaseNFT, err error) {
	nftLink, found := k.GetDigitalHashKVStoreValue(ctx, digitalHash)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrUnknownCollection, fmt.Sprintf("no one linked token with this DigitalHash %s", digitalHash))
	}

	collection, found := k.GetCollection(ctx, nftLink.Denom)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrUnknownCollection, fmt.Sprintf("dont found any collection by Denom %s", nftLink.Denom))
	}

	basenft, err := collection.GetNFT(nftLink.TokenID)
	nft = types.NewExtraBaseNFT(basenft.GetID(), basenft.GetOwner(), basenft.GetTokenURI(), basenft.GetDigitalHash(), nftLink.Denom)
	if err != nil {
		return nil, err
	}

	return nft, nil

}

// UpdateNFT updates an already existing NFTs
func (k Keeper) UpdateNFT(ctx sdk.Context, denom string, nft exported.NFT) (err error) {
	collection, found := k.GetCollection(ctx, denom)
	if !found {
		return sdkerrors.Wrap(types.ErrUnknownCollection, fmt.Sprintf("collection #%s doesn't exist", denom))
	}
	oldNFT, err := collection.GetNFT(nft.GetID())
	if err != nil {
		return err
	}
	// if the owner changed then update the owners KVStore too
	if !oldNFT.GetOwner().Equals(nft.GetOwner()) {
		err = k.SwapOwners(ctx, denom, nft.GetID(), oldNFT.GetOwner(), nft.GetOwner())
		if err != nil {
			return err
		}
	}

	oldTokenURI := oldNFT.GetTokenURI()
	tokenURI := nft.GetTokenURI()
	if oldTokenURI != tokenURI {
		err = k.IsURIExists(ctx, tokenURI)
		if err != nil {
			return err
		}
	}

	oldDigitalHash := oldNFT.GetDigitalHash()
	digitalHash := nft.GetDigitalHash()
	if oldDigitalHash != digitalHash {
		err = k.IsDigitalHashExists(ctx, digitalHash)
		if err != nil {
			return err
		}
	}

	collection, err = collection.UpdateNFT(nft)

	if err != nil {
		return err
	}

	k.RemoveURIKVStoreValue(ctx, oldTokenURI)
	k.RemoveDigitalHashKVStoreValue(ctx, oldDigitalHash)
	k.SetCollection(ctx, denom, collection)

	return nil
}

// MintNFT mints an NFT and manages that NFTs existence within Collections and Owners
func (k Keeper) MintNFT(ctx sdk.Context, denom string, nft exported.NFT) (err error) {
	collection, found := k.GetCollection(ctx, denom)
	if found {
		collection, err = collection.AddNFT(nft)
		if err != nil {
			return err
		}
	} else {
		collection = types.NewCollection(denom, types.NewNFTs(nft))
	}

	tokenURI := nft.GetTokenURI()
	digitalHash := nft.GetDigitalHash()

	err = k.IsURIExists(ctx, tokenURI)
	if err != nil {
		return err
	}

	err = k.IsDigitalHashExists(ctx, digitalHash)
	if err != nil {
		return err
	}

	k.SetCollection(ctx, denom, collection)

	tokenID := nft.GetID()
	if tokenURI != "" {
		k.SetURIKVstoreValue(ctx, tokenURI, tokenID, denom)
	}

	if digitalHash != "" {
		k.SetDigitalHashKVStoreValue(ctx, digitalHash, tokenID, denom)
	}

	ownerIDCollection, _ := k.GetOwnerByDenom(ctx, nft.GetOwner(), denom)
	ownerIDCollection = ownerIDCollection.AddID(tokenID)
	k.SetOwnerByDenom(ctx, nft.GetOwner(), denom, ownerIDCollection.IDs)
	return
}

// DeleteNFT deletes an existing NFT from store
func (k Keeper) DeleteNFT(ctx sdk.Context, denom, id string) (err error) {
	collection, found := k.GetCollection(ctx, denom)
	if !found {
		return sdkerrors.Wrap(types.ErrUnknownCollection, fmt.Sprintf("collection of %s doesn't exist", denom))
	}
	nft, err := collection.GetNFT(id)
	if err != nil {
		return err
	}
	ownerIDCollection, found := k.GetOwnerByDenom(ctx, nft.GetOwner(), denom)
	if !found {
		return sdkerrors.Wrap(types.ErrUnknownCollection,
			fmt.Sprintf("id collection #%s doesn't exist for owner %s", denom, nft.GetOwner()),
		)
	}
	ownerIDCollection, err = ownerIDCollection.DeleteID(nft.GetID())
	if err != nil {
		return err
	}
	k.SetOwnerByDenom(ctx, nft.GetOwner(), denom, ownerIDCollection.IDs)

	collection, err = collection.DeleteNFT(nft)
	if err != nil {
		return err
	}

	k.RemoveURIKVStoreValue(ctx, nft.GetTokenURI())
	k.RemoveDigitalHashKVStoreValue(ctx, nft.GetDigitalHash())
	k.SetCollection(ctx, denom, collection)

	return
}
