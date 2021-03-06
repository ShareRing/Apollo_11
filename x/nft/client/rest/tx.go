package rest

import (
	"github.com/ShareRing/nft-module/x/nft/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/gorilla/mux"
	"net/http"
)

func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router,
	cdc *codec.Codec, queryRoute string) {
	// Transfer an NFT to an address
	r.HandleFunc(
		"/nfts/transfer",
		transferNFTHandler(cdc, cliCtx),
	).Methods("POST")

	// Update an NFT metadata
	r.HandleFunc(
		"/nfts/collection/{denom}/nft/{id}/metadata",
		editNFTMetadataHandler(cdc, cliCtx),
	).Methods("PUT")

	// Update an NFT DigitalHash
	r.HandleFunc(
		"/nfts/collection/{denom}/nft/{id}/digitalhash",
		editNFTDigitalHashHandler(cdc, cliCtx),
	).Methods("PUT")

	// Mint an NFT
	r.HandleFunc(
		"/nfts/mint",
		mintNFTHandler(cdc, cliCtx),
	).Methods("POST")

	// Burn an NFT
	r.HandleFunc(
		"/nfts/collection/{denom}/nft/{id}/burn",
		burnNFTHandler(cdc, cliCtx),
	).Methods("PUT")
}

type transferNFTReq struct {
	BaseReq   rest.BaseReq `json:"base_req"`
	Denom     string       `json:"denom"`
	ID        string       `json:"id"`
	Recipient string       `json:"recipient"`
}

func transferNFTHandler(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req transferNFTReq
		if !rest.ReadRESTReq(w, r, cdc, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		fromAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		recipient, err := sdk.AccAddressFromBech32(req.Recipient)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// create the message
		msg := types.NewMsgTransferNFT(fromAddr, recipient, req.Denom, req.ID)

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type editNFTMetadataReq struct {
	BaseReq  rest.BaseReq `json:"base_req"`
	Denom    string       `json:"denom"`
	ID       string       `json:"id"`
	TokenURI string       `json:"tokenURI"`
}

func editNFTMetadataHandler(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req editNFTMetadataReq
		if !rest.ReadRESTReq(w, r, cdc, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		fromAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// create the message
		msg := types.NewMsgEditNFTMetadata(fromAddr, req.ID, req.Denom, req.TokenURI)

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type editNFTDigitalHashReq struct {
	BaseReq     rest.BaseReq `json:"base_req"`
	Denom       string       `json:"denom"`
	ID          string       `json:"id"`
	DigitalHash string       `json:"digital_hash"`
}

func editNFTDigitalHashHandler(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req editNFTDigitalHashReq
		if !rest.ReadRESTReq(w, r, cdc, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		fromAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// create the message
		msg := types.NewMsgEditNFTDigitalHash(fromAddr, req.ID, req.Denom, req.DigitalHash)

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type mintNFTReq struct {
	BaseReq     rest.BaseReq   `json:"base_req"`
	Recipient   sdk.AccAddress `json:"recipient"`
	Denom       string         `json:"denom"`
	ID          string         `json:"id"`
	TokenURI    string         `json:"tokenURI"`
	DigitalHash string         `json:"digital_hash"`
}

func mintNFTHandler(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req mintNFTReq
		if !rest.ReadRESTReq(w, r, cdc, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		fromAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// create the message
		msg := types.NewMsgMintNFT(fromAddr, req.Recipient, req.ID, req.Denom, req.TokenURI, req.DigitalHash)

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type burnNFTReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Denom   string       `json:"denom"`
	ID      string       `json:"id"`
}

func burnNFTHandler(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req burnNFTReq
		if !rest.ReadRESTReq(w, r, cdc, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		fromAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// create the message
		msg := types.NewMsgBurnNFT(fromAddr, req.ID, req.Denom)
		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}
