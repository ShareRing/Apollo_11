package keeper_test

import (
	simapp2 "github.com/ShareRing/nft-module/simapp"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/ShareRing/nft-module/x/nft/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// nolint: deadcode unused
var (
	denom        = "test-denom"
	denom2       = "test-denom2"
	denom3       = "test-denom3"
	id           = "1"
	id2          = "2"
	id3          = "3"
	address      = types.CreateTestAddrs(1)[0]
	address2     = types.CreateTestAddrs(2)[1]
	address3     = types.CreateTestAddrs(3)[2]
	tokenURI1    = "https://google.com/token-1.json"
	tokenURI2    = "https://google.com/token-2.json"
	tokenURI3    = "https://google.com/token-3.json"
	tokenURI4    = "https://google.com/token-4.json"
	tokenURI5    = "https://google.com/token-5.json"
	tokenURI6    = "https://google.com/token-6.json"
	digitalHash1 = "https://google.com/digitalHash-1.json"
	digitalHash2 = "https://google.com/digitalHash-2.json"
	digitalHash3 = "https://google.com/digitalHash-3.json"
	digitalHash4 = "https://google.com/digitalHash-4.json"
	digitalHash5 = "https://google.com/digitalHash-5.json"
	digitalHash6 = "https://google.com/digitalHash-6.json"
)

func createTestApp(isCheckTx bool) (*simapp2.SimApp, sdk.Context) {
	app := simapp2.Setup(isCheckTx)

	ctx := app.BaseApp.NewContext(isCheckTx, abci.Header{})

	return app, ctx
}
