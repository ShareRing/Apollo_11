package keeper

import (
	"fmt"
	"github.com/ShareRing/nft-module/x/nft/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// SetURIKVstoreValue set uri's value to kvstore
func (k Keeper) SetURIKVstoreValue(ctx sdk.Context, uri, tokenID, denom string) {
	store := ctx.KVStore(k.storeKey)
	URIKey := types.GetURIKey(uri)
	nftLink := types.NFTLinkDenom{
		TokenID: tokenID,
		Denom:   denom,
	}
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(nftLink)
	store.Set(URIKey, bz)
}

// GetURIKVStoreValue get token id by uri's value from kvstore
func (k Keeper) GetURIKVStoreValue(ctx sdk.Context, uri string) (nftLink *types.NFTLinkDenom, found bool) {
	store := ctx.KVStore(k.storeKey)
	URIKey := types.GetURIKey(uri)
	bz := store.Get(URIKey)
	if len(bz) == 0 {
		return nil, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &nftLink)
	return nftLink, true
}

// RemoveURIKVStoreValue remove uri from kvstore
func (k Keeper) RemoveURIKVStoreValue(ctx sdk.Context, uri string) {
	store := ctx.KVStore(k.storeKey)
	URIKey := types.GetURIKey(uri)
	store.Delete(URIKey)
}

// IsURIExists check the uri's existing
func (k Keeper) IsURIExists(ctx sdk.Context, uri string) error {
	nftLink, found := k.GetURIKVStoreValue(ctx, uri)
	if found {
		return sdkerrors.Wrap(types.ErrWrongURI, fmt.Sprintf("current URI %v already exists and linked with %v token from denom %v", uri, nftLink.TokenID, nftLink.Denom))
	}

	return nil
}

// SetDigitalHashKVStoreValue set digitalhash's value to kvstore
func (k Keeper) SetDigitalHashKVStoreValue(ctx sdk.Context, digitalHash, tokenID, denom string) {
	store := ctx.KVStore(k.storeKey)
	DigitalHashKey := types.GetDigitalHashKey(digitalHash)
	nftLink := types.NFTLinkDenom{
		TokenID: tokenID,
		Denom:   denom,
	}

	bz := k.cdc.MustMarshalBinaryLengthPrefixed(nftLink)
	store.Set(DigitalHashKey, bz)
}

// GetDigitalHashKVStoreValue get token id by digitalhash's value from kvstore
func (k Keeper) GetDigitalHashKVStoreValue(ctx sdk.Context, digitalHash string) (nftLink *types.NFTLinkDenom, found bool) {
	store := ctx.KVStore(k.storeKey)
	DigitalHashKey := types.GetDigitalHashKey(digitalHash)
	bz := store.Get(DigitalHashKey)
	if len(bz) == 0 {
		return nil, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &nftLink)
	return nftLink, true
}

// IsURIExists check the uri's existing
func (k Keeper) IsDigitalHashExists(ctx sdk.Context, digitalHash string) error {
	nftLink, found := k.GetDigitalHashKVStoreValue(ctx, digitalHash)
	if found {
		return sdkerrors.Wrap(types.ErrWrongDigitalHash, fmt.Sprintf("current DigitalHash %v already exists and linked with %v token from denom %v", digitalHash, nftLink.TokenID, nftLink.Denom))
	}

	return nil
}

// RemoveDigitalHashKVStoreValue remove digitalHash from kvstore
func (k Keeper) RemoveDigitalHashKVStoreValue(ctx sdk.Context, digitalHash string) {
	store := ctx.KVStore(k.storeKey)
	DigitalHashKey := types.GetDigitalHashKey(digitalHash)
	store.Delete(DigitalHashKey)
}
