package keeper

import (
	"fmt"
	"strconv"

	"github.com/ShareRing/nft-module/x/nft/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	abci "github.com/tendermint/tendermint/abci/types"
)

// query endpoints supported by the NFT Querier
const (
	QuerySupply           = "supply"
	QueryOwner            = "owner"
	QueryOwnerByDenom     = "ownerByDenom"
	QueryCollection       = "collection"
	QueryDenoms           = "denoms"
	QueryNFT              = "nft"
	QueryNFTByDigitalHash = "nftByDigitalHash"
)

// NewQuerier is the module level router for state queries
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case QuerySupply:
			return querySupply(ctx, path[1:], req, k)
		case QueryOwner:
			return queryOwner(ctx, path[1:], req, k)
		case QueryOwnerByDenom:
			return queryOwnerByDenom(ctx, path[1:], req, k)
		case QueryCollection:
			return queryCollection(ctx, path[1:], req, k)
		case QueryDenoms:
			return queryDenoms(ctx, path[1:], req, k)
		case QueryNFT:
			return queryNFT(ctx, path[1:], req, k)
		case QueryNFTByDigitalHash:
			return queryNFTByDigitalHash(ctx, path[1:], req, k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown nft query endpoint "+path[0])
		}
	}
}

func querySupply(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var params types.QueryCollectionParams

	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("incorrectly formatted request data %v", err.Error()))
	}

	collection, found := k.GetCollection(ctx, params.Denom)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrUnknownCollection, fmt.Sprintf("unknown denom %s", params.Denom))
	}

	bz := []byte(strconv.Itoa(collection.Supply()))
	return bz, nil
}

func queryOwner(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var params types.QueryBalanceParams

	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, err.Error())
	}

	owner := k.GetOwner(ctx, params.Owner)
	bz, err := types.ModuleCdc.MarshalJSON(owner)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryOwnerByDenom(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var params types.QueryBalanceParams

	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, err.Error())
	}

	var owner types.Owner

	idCollection, _ := k.GetOwnerByDenom(ctx, params.Owner, params.Denom)
	owner.Address = params.Owner
	owner.IDCollections = append(owner.IDCollections, idCollection).Sort()

	bz, err := types.ModuleCdc.MarshalJSON(owner)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryCollection(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var params types.QueryCollectionParams

	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, err.Error())
	}

	collection, found := k.GetCollection(ctx, params.Denom)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrUnknownCollection, fmt.Sprintf("unknown denom %s", params.Denom))
	}

	// use Collections custom JSON to make the denom the key of the object
	collections := types.NewCollections(collection)
	//bz, err := collections.MarshalJSON()
	bz, err := types.ModuleCdc.MarshalJSON(collections)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryDenoms(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) ([]byte, error) {
	denoms := k.GetDenoms(ctx)

	bz, err := types.ModuleCdc.MarshalJSON(denoms)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryNFT(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var params types.QueryNFTParams

	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, err.Error())
	}

	nft, err := k.GetNFT(ctx, params.Denom, params.TokenID)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrUnknownNFT, fmt.Sprintf("invalid NFT #%s from collection %s", params.TokenID, params.Denom))
	}

	bz, err := types.ModuleCdc.MarshalJSON(nft)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryNFTByDigitalHash(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var params types.QueryNFTsByDigitalHashParams

	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, err.Error())
	}

	if len(params.Denom) == 0 {
		params.Denom = k.GetDenoms(ctx)
	}

	nft, err := k.GetNFTByDigitalHash(ctx, params.Denom, params.DigitalHash)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrWrongDigitalHash, fmt.Sprintf("invalid DigitalHash #%s", params.DigitalHash))
	}

	bz, err := types.ModuleCdc.MarshalJSON(nft)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}