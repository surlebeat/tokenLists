package main

import (
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/helpers"
)

type TPortalTokenData struct {
	Key       string            `json:"key"`
	Name      string            `json:"name"`
	Decimals  int               `json:"decimals"`
	Symbol    string            `json:"symbol"`
	Price     float64           `json:"price"`
	Address   string            `json:"address"`
	Addresses map[string]string `json:"addresses"`
	Platform  string            `json:"platform"`
	Network   string            `json:"network"`
	Images    []string          `json:"images"`
	UpdatedAt string            `json:"updatedAt"`
	CreatedAt string            `json:"createdAt"`
	Tokens    []string          `json:"tokens"`
	Liquidity float64           `json:"liquidity"`
}
type TPortalList struct {
	PageItems  int
	TotalItems int
	More       bool
	Page       int
	Tokens     []TPortalTokenData
}

func portalsMapNetworkNameToChainID(network string) uint64 {
	switch network {
	case `ethereum`:
		return 1
	case `optimism`:
		return 10
	case `fantom`:
		return 250
	case `arbitrum`:
		return 42161
	case `polygon`:
		return 137
	case `avalanche`:
		return 43114
	case `bsc`:
		return 56
	}
	return 0
}

func fetchPortalsTokenList() []TokenListToken {
	limit := 250
	page := 0
	tokens := []TokenListToken{}

	for {
		uri := `https://api.portals.fi/v2/tokens?limit=` + strconv.FormatInt(int64(limit), 10) + `&page=` + strconv.FormatInt(int64(page), 10)
		list := helpers.FetchJSON[TPortalList](uri)

		for _, token := range list.Tokens {
			logoURI := ``
			if len(token.Images) > 0 {
				logoURI = token.Images[0]
			}
			if newToken, err := SetToken(
				common.HexToAddress(token.Address),
				token.Name,
				token.Symbol,
				logoURI,
				portalsMapNetworkNameToChainID(token.Network),
				token.Decimals,
			); err == nil {
				tokens = append(tokens, newToken)
			}
		}
		if !list.More {
			break
		}
		page++
	}
	return fetchTokenList(tokens)
}

func buildPortalsTokenList() {
	tokenList := loadTokenListFromJsonFile(`portals.json`)
	tokenList.Name = "Portals Token List"
	tokenList.LogoURI = "https://portals-assets-bucket.s3.amazonaws.com/logo.png"

	tokens := fetchPortalsTokenList()
	saveTokenListInJsonFile(tokenList, tokens, `portals.json`, Standard)
}
