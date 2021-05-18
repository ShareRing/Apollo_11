package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ShareRing/nft-module/x/nft/types"
)

func TestIsExistsExtraInfo(t *testing.T) {
	app, ctx := createTestApp(false)

	// create a new nft with id = "id" and owner = "address"
	// MintNFT shouldn't fail when collection does not exist
	nft := types.NewBaseNFT(id, address, tokenURI1, digitalHash1)
	err := app.NFTKeeper.MintNFT(ctx, denom, &nft)
	require.NoError(t, err)

	//send one more nft token with same uri
	nftSameURI := types.NewBaseNFT(id2, address, tokenURI1, digitalHash2)
	err = app.NFTKeeper.MintNFT(ctx, denom, &nftSameURI)
	require.Error(t, err)

	//send one more nft token with same digital hash
	nftSameHash := types.NewBaseNFT(id3, address, tokenURI2, digitalHash1)
	err = app.NFTKeeper.MintNFT(ctx, denom, &nftSameHash)
	require.Error(t, err)

	//update old token and add new
	nft.TokenURI = tokenURI2
	err = app.NFTKeeper.UpdateNFT(ctx,denom, &nft)

	err = app.NFTKeeper.MintNFT(ctx, denom, &nftSameURI)
	require.NoError(t, err)

	//drop old token and add new
	err = app.NFTKeeper.DeleteNFT(ctx, denom, nft.GetID())
	require.NoError(t, err)

	err = app.NFTKeeper.MintNFT(ctx, denom, &nftSameHash)
	require.NoError(t, err)

}

