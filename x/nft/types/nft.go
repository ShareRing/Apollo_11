package types

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/ShareRing/nft-module/x/nft/exported"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ exported.NFT = (*BaseNFT)(nil)

// BaseNFT non fungible token definition
type BaseNFT struct {
	ID          string         `json:"id,omitempty" yaml:"id"`           // id of the token; not exported to clients
	Owner       sdk.AccAddress `json:"owner" yaml:"owner"`               // account address that owns the NFT
	TokenURI    string         `json:"token_uri" yaml:"token_uri"`       // optional extra properties available for querying
	DigitalHash string         `json:"digital_hash" yaml:"digital_hash"` // optional extra properties for saving digitalHash
}

// NewBaseNFT creates a new NFT instance
func NewBaseNFT(id string, owner sdk.AccAddress, tokenURI, digitalHash string) BaseNFT {
	return BaseNFT{
		ID:          id,
		Owner:       owner,
		TokenURI:    strings.TrimSpace(tokenURI),
		DigitalHash: strings.TrimSpace(digitalHash),
	}
}

// GetID returns the ID of the token
func (bnft BaseNFT) GetID() string { return bnft.ID }

// GetOwner returns the account address that owns the NFT
func (bnft BaseNFT) GetOwner() sdk.AccAddress { return bnft.Owner }

// SetOwner updates the owner address of the NFT
func (bnft *BaseNFT) SetOwner(address sdk.AccAddress) {
	bnft.Owner = address
}

// GetTokenURI returns the path to optional extra properties
func (bnft BaseNFT) GetTokenURI() string { return bnft.TokenURI }

// GetDigitalHash returns the digitalHash from extra properties
func (bnft BaseNFT) GetDigitalHash() string { return bnft.DigitalHash }

// EditMetadata edits metadata of an nft
func (bnft *BaseNFT) EditMetadata(tokenURI string) {
	bnft.TokenURI = tokenURI
}

// EditDigitalHash edits digital hash of an nft
func (bnft *BaseNFT) EditDigitalHash(digitalHash string) {
	bnft.DigitalHash = digitalHash
}

func (bnft BaseNFT) String() string {
	return fmt.Sprintf(`ID:				%s
Owner:			%s
TokenURI:		%s
DigitalHash:	%s`,
		bnft.ID,
		bnft.Owner,
		bnft.TokenURI,
		bnft.DigitalHash,
	)
}

// ExtraBaseNFT for return extra info about nft
type ExtraBaseNFT struct {
	Denom string `json:"denom" yaml:"denom"`
	BaseNFT
}

// NewExtraBaseNFT creates a new NFT instance
func NewExtraBaseNFT(id string, owner sdk.AccAddress, tokenURI, digitalHash, denom string) *ExtraBaseNFT {
	return &ExtraBaseNFT{
		Denom: denom,
		BaseNFT: BaseNFT{
			ID:          id,
			Owner:       owner,
			TokenURI:    strings.TrimSpace(tokenURI),
			DigitalHash: strings.TrimSpace(digitalHash),
		},
	}
}

func (enft ExtraBaseNFT) String() string {
	return fmt.Sprintf(`ID:				%s
Owner:			%s
Denom:			%s
TokenURI:		%s
DigitalHash:	%s`,
		enft.ID,
		enft.Owner,
		enft.Denom,
		enft.TokenURI,
		enft.DigitalHash,
	)
}

// NFTLinkDenom define a link between nft id and his denom
type NFTLinkDenom struct {
	TokenID string `json:"token_id" yaml:"token_id"`
	Denom   string `json:"denom" yaml:"denom"`
}

// NFTs define a list of NFT
type NFTs []exported.NFT

// NewNFTs creates a new set of NFTs
func NewNFTs(nfts ...exported.NFT) NFTs {
	if len(nfts) == 0 {
		return NFTs{}
	}
	return NFTs(nfts).Sort()
}

// Append appends two sets of NFTs
func (nfts NFTs) Append(nftsB ...exported.NFT) NFTs {
	return append(nfts, nftsB...).Sort()
}

// Find returns the searched collection from the set
func (nfts NFTs) Find(id string) (nft exported.NFT, found bool) {
	index := nfts.find(id)
	if index == -1 {
		return nft, false
	}
	return nfts[index], true
}

// Update removes and replaces an NFT from the set
func (nfts NFTs) Update(id string, nft exported.NFT) (NFTs, bool) {
	index := nfts.find(id)
	if index == -1 {
		return nfts, false
	}

	return append(append(nfts[:index], nft), nfts[index+1:]...), true
}

// Remove removes an NFT from the set of NFTs
func (nfts NFTs) Remove(id string) (NFTs, bool) {
	index := nfts.find(id)
	if index == -1 {
		return nfts, false
	}

	return append(nfts[:index], nfts[index+1:]...), true
}

// String follows stringer interface
func (nfts NFTs) String() string {
	if len(nfts) == 0 {
		return ""
	}

	out := ""
	for _, nft := range nfts {
		out += fmt.Sprintf("%v\n", nft.String())
	}
	return out[:len(out)-1]
}

// Empty returns true if there are no NFTs and false otherwise.
func (nfts NFTs) Empty() bool {
	return len(nfts) == 0
}

func (nfts NFTs) find(id string) int {
	return FindUtil(nfts, id)
}

// ----------------------------------------------------------------------------
// Encoding

// NFTJSON is the exported NFT format for clients
type NFTJSON map[string]BaseNFT

// MarshalJSON for NFTs
func (nfts NFTs) MarshalJSON() ([]byte, error) {
	nftJSON := make(NFTJSON)
	for _, nft := range nfts {
		id := nft.GetID()
		bnft := NewBaseNFT(id, nft.GetOwner(), nft.GetTokenURI(), nft.GetDigitalHash())
		nftJSON[id] = bnft
	}
	return json.Marshal(nftJSON)
}

// UnmarshalJSON for NFTs
func (nfts *NFTs) UnmarshalJSON(b []byte) error {
	nftJSON := make(NFTJSON)
	if err := json.Unmarshal(b, &nftJSON); err != nil {
		return err
	}

	for id, nft := range nftJSON {
		bnft := NewBaseNFT(id, nft.GetOwner(), nft.GetTokenURI(), nft.GetDigitalHash())
		*nfts = append(*nfts, &bnft)
	}
	return nil
}

// Findable and Sort interfaces
func (nfts NFTs) ElAtIndex(index int) string { return nfts[index].GetID() }
func (nfts NFTs) Len() int                   { return len(nfts) }
func (nfts NFTs) Less(i, j int) bool         { return strings.Compare(nfts[i].GetID(), nfts[j].GetID()) == -1 }
func (nfts NFTs) Swap(i, j int)              { nfts[i], nfts[j] = nfts[j], nfts[i] }

var _ sort.Interface = NFTs{}

// Sort is a helper function to sort the set of coins in place
func (nfts NFTs) Sort() NFTs {
	sort.Sort(nfts)
	return nfts
}
