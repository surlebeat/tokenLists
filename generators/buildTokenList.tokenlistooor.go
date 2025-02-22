package main

import (
	"math"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/helpers"
)

type TSmolAssetsList struct {
	Version struct {
		Major int `json:"major"`
		Minor int `json:"minor"`
		Patch int `json:"patch"`
	}
	Tokens []string `json:"tokens"`
}

func buildTokenListooorList() {
	tokenList := loadTokenListFromJsonFile(`tokenlistooor.json`)
	tokenList.Name = `Tokenlistooor Token List`
	tokenList.LogoURI = `https://raw.githubusercontent.com/Migratooor/tokenLists/main/.github/tokenlistooor.svg`
	tokenList.Description = `A curated list of tokens from all the token lists on tokenlistooor.`

	/**************************************************************************
	** Create a map of all tokens from all lists and only add the missing ones
	** in it. Map are WAY faster than arrays fir our use case
	**************************************************************************/
	allTokens := make(map[uint64]map[string]TokenListToken)
	allTokensPlain := []TokenListToken{}
	listsPerChain := make(map[uint64][]string)

	allTokensPlain = append(allTokensPlain, ETHER)
	allTokensPlain = append(allTokensPlain, FTM)
	allTokensPlain = append(allTokensPlain, BSC)
	allTokensPlain = append(allTokensPlain, MATIC)
	allTokensPlain = append(allTokensPlain, MATIC_ZKEVM)
	allTokensPlain = append(allTokensPlain, XDAI)
	allTokensPlain = append(allTokensPlain, addEtherLikeToken(10))
	allTokensPlain = append(allTokensPlain, addEtherLikeToken(8453))
	allTokensPlain = append(allTokensPlain, addEtherLikeToken(324))

	/**************************************************************************
	** We want to know which tokens to add to the aggregated tokenlistooor list
	** and to do that we need to know in how many lists they are present.
	** This is chain sensitive: we need a token to be available in at least
	** 50% of the lists for a given chain to be added to the aggregated list.
	**************************************************************************/
	for name, generatorData := range GENERATORS {
		if name == `tokenlistooor` {
			continue
		}
		if generatorData.GeneratorType == GeneratorPool {
			continue
		}

		initialCount := 1
		shouldByPassCount := name == `yearn` || name == `smolAssets`
		if shouldByPassCount {
			initialCount = math.MaxInt32
		}
		tokenList := loadTokenListFromJsonFile(name + `.json`)
		for _, token := range tokenList.Tokens {
			if !helpers.IsChainIDSupported(token.ChainID) {
				continue
			}
			if _, ok := listsPerChain[token.ChainID]; !ok {
				listsPerChain[token.ChainID] = []string{}
			}

			if !helpers.Includes(listsPerChain[token.ChainID], name) {
				listsPerChain[token.ChainID] = append(listsPerChain[token.ChainID], name)
			}

			if _, ok := allTokens[token.ChainID]; !ok {
				allTokens[token.ChainID] = make(map[string]TokenListToken)
			}

			if existingToken, ok := allTokens[token.ChainID][helpers.ToAddress(token.Address)]; ok {
				allTokens[token.ChainID][helpers.ToAddress(token.Address)] = TokenListToken{
					Address:    existingToken.Address,
					Name:       helpers.SafeString(existingToken.Name, token.Name),
					Symbol:     helpers.SafeString(existingToken.Symbol, token.Symbol),
					LogoURI:    helpers.SafeString(existingToken.LogoURI, token.LogoURI),
					Decimals:   helpers.SafeInt(existingToken.Decimals, token.Decimals),
					ChainID:    token.ChainID,
					Occurrence: existingToken.Occurrence + 1,
				}
			} else {
				tokenInitialOccurence := initialCount
				if common.HexToAddress(token.Address) == common.HexToAddress(`0x9a96ec9B57Fb64FbC60B423d1f4da7691Bd35079`) { //Ajna
					tokenInitialOccurence = math.MaxInt32
				}
				allTokens[token.ChainID][helpers.ToAddress(token.Address)] = TokenListToken{
					Address:    helpers.ToAddress(token.Address),
					Name:       helpers.SafeString(token.Name, ``),
					Symbol:     helpers.SafeString(token.Symbol, ``),
					LogoURI:    helpers.SafeString(token.LogoURI, ``),
					Decimals:   helpers.SafeInt(token.Decimals, 18),
					ChainID:    token.ChainID,
					Occurrence: tokenInitialOccurence,
				}
			}
		}
	}

	for chainID, tokens := range allTokens {
		for _, token := range tokens {
			if _, ok := listsPerChain[chainID]; !ok {
				continue
			}
			chainCount := len(listsPerChain[uint64(chainID)])
			if token.Occurrence >= int(math.Ceil(float64(chainCount)*0.5)) {
				allTokensPlain = append(allTokensPlain, token)
			}
		}
	}

	tokens := fetchTokenList(allTokensPlain)
	saveTokenListInJsonFile(tokenList, tokens, `tokenlistooor.json`, Standard)
}
