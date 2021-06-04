## NFT-MODULE
### DESCRIPTION
NFT module represents ability to use non-fungible tokens in your blockchains.  
### MAIN STRUCTS
#### NFT STRUCT
NFT token follow next structure:  
```
{
    id: string  
    owner: string  
    token_uri: string  
    digital_hash: string  
}
```

- id - unique token identifier, the current collection cannot have tokens with the same identifier
- owner - address of current owner of nft token
- token_uri - link to the file to which the token grants right
  - Attention: it must be unique in all blockchain
- digital_hash - unique sequence generated from the digital version of the item to which the rights are granted
  - Attention: it must be unique in all blockchain

#### COLLECTION STRUCT
NFT tokens are collected in a special collection with a unique name, which means that the uniqueness of the token is calculated by the collection name and token ID.
```
{
    demon: string  
    nfts:  []nft-tokens  
}
```

- denom - unique collection name
  - Attention: collection with the same name cannot exist
- ntfs - array of nft tokens, which minted in current collection
### REST ENDPOINTS
#### QUERY
##### GetSupply
purpose: get total supply of a collection of NFTs  
type: get  
path: "/nft/supply/{denom}"   
**{demon}** - name of collection, from your want get total supply   

##### GetOwner
purpose: get the NFTs owned by an account address  
type: get  
path: "/nft/owner/{delegatorAddr}"   
**{delegatorAddr}** - name of collection, from your want get total supply
##### GetOwnerByDenom
purpose: get the NFTs owned by an account address and selected denom  
type: get  
path: "/nft/owner/{delegatorAddr}/collection/{denom}"   
**{delegatorAddr}** - name of collection, from your want get total supply  
**{demon}** - name of collection, from your want get nft

##### GetCollection
purpose:  get all the NFT from a given collection   
type: get  
path: "/nft/collection/{denom}"   
**{demon}** - name of collection, from your want get nft

##### GetDenoms
purpose: get all denoms  
type: get  
path: "/nft/denoms"

##### GetNFT
purpose:   
type: get  
path: "/nft/collection/{denom}/nft/{id}"   
**{demon}** - name of collection, from your want get nft  
**{id}** - id of needed nft  

##### GetNFTbyDigitalHash
purpose: get NFT from all collection by unique digital-hash  
type: get  
path: "/nft/collection/nftByDigitalHash/{digital-hash}"   
**{digital-hash}** - unique digital hash

#### TX
##### TransferNFT
purpose: transfer an NFT to an address   
type: post  
path: "/nfts/transfer"  
struct in:  
```
{
	"base_req": rest.BaseReq
	"denom": string
	"id": string
	"recipient": string
}
```

##### EditNFTMetadata
purpose: update an NFT metadata   
type: put  
path: "/nfts/collection/{denom}/nft/{id}/metadata"  
**{demon}** - name of collection, where nft need to be edit    
**{id}** - id of needed nft  
struct in:
```
{
	"base_req": rest.BaseReq
	"denom": string
	"id": string
	"tokenURI": string
}
```

##### EditNFTDigitalHash
purpose: update an NFT DigitalHash   
type: put  
path: "/nfts/collection/{denom}/nft/{id}/digitalhash"  
**{demon}** - name of collection, where nft need to be edit    
**{id}** - id of needed nft  
struct in:
```
{
	"base_req": rest.BaseReq
	"denom": string
	"id": string
	"digital_hash": string
}
```
##### MintNFT
purpose: mint an NFT   
type: post  
path: "/nfts/mint"
struct in:
```
{
	base_req: rest.BaseReq
	recipient: string
	denom: string
	id: string
	tokenURI: string
	digital_hash: string
}
```

##### BurnNFT
purpose: burn an NFT   
type: put  
path: "/nfts/collection/{denom}/nft/{id}/burn"  
**{demon}** - name of collection, where nft need to be burn    
**{id}** - id of needed nft  
struct in:
```
{
	base_req: rest.BaseReq
	denom: string
	id: string
}
```

### CLI COMMANDS
#### QUERY
##### QuerySupply
purpose: get total supply of a collection of NFTs    
cmd example: *cli query nft supply \[denom]    
\[denom] - name of collection, from your want get total supply  
struct out:
```
num: string
```

##### QueryOwner
purpose: get the NFTs owned by an account address    
cmd example: *cli query nft owner \[accountAddress] \[denom]  
\[accountAddress]- address of account  
\[denom] - name of collection, from your want get nfts, optional parameter  
struct out:
```
{
  address: string
  idCollections: types.IDCollections
}
```

##### QueryCollection
purpose: get all the NFTs from a given collection"    
cmd example: *cli query nft collection \[denom]   
\[denom] - name of collection, from your want nfts  
struct out:  
```
{
 denom: string
 nfts:
    [
      {
        id: string
        owner: string
        token_uri: string
        digital_hash: string
      }
    ]
}
```

##### QueryDenoms
purpose: queries all denominations of all collections of NFTs      
cmd example: *cli query denoms  
struct out:  
``` 
[nameDemon1: string, ..., nameDemonN: string]
```

##### QueryNFT
purpose: query a single NFT from a collection      
cmd example: *cli query nft token \[denom] \[ID] 
\[denom] - name of collection, from your want get total supply  
\[ID] - id of nft that your need  
struct out:
```
{
   id: string
   owner: string
   token_uri: string
   digital_hash: string
}
```

##### QueryNFTByDigitalHash
purpose:    
cmd example: *cli query nft token-by-digital-hash \[digital-hash]    
\[digital-hash] - unique digital hash from token  
struct out:
```
{
denom: string
basenft:
        {
           id: string
           owner: string
           token_uri: string
           digital_hash: string
        }
}        
```

#### TX
##### TransferNFT
purpose: transfer a NFT to a recipient    
cmd example: *cli tx nft transfer \[sender] \[recipient] \[denom] \[tokenID]  
\[sender] -  address of sender account  
\[recipient] - address of recipient account  
\[denom] - name of collection contains nft  
\[tokenID] - id of target nft  

##### EditNFTMetadata
purpose: edit the metadata of an NFT    
cmd example: *cli tx nft edit-metadata \[denom] \[tokenID] \[--tokenURI]  
\[denom] -  name of collection    
\[tokenID] - token id  
\[--tokenURI] - new value for uri  

##### EditNFTDigitalHash
purpose: edit the digital hash of an NFT    
cmd example: *cli tx nft edit-digital-hash \[denom] \[tokenID] \[--digital-hash]  
\[denom] -  name of collection    
\[tokenID] - token id  
\[--digital-hash] - new value for digital-hash

##### MintNFT
purpose: mint an NFT and set the owner to the recipient    
cmd example: *cli tx nft mint \[denom] \[tokenID] \[recipient] \[--from] \[--tokenURI] \[--digital-hash]  
\[denom] -  name of collection    
\[tokenID] - token id  
\[recipient] - address - repicient  
\[--from] - alias or addres who mint token  
\[--tokenURI] - new value for uri  
\[--digital-hash] - new value for digital-hash  

##### BurnNFT
purpose: burn an NFT      
cmd example: *cli tx nft burn \[denom] \[tokenID]  
\[denom] -  name of collection    
\[tokenID] - token id  
