package main

type TGenerationMethods string
type TGeneratorType string

const (
	// GenerationAPI indicates that the list is generated by calling an API
	GenerationAPI TGenerationMethods = "API"
	// GenerationEvents indicates that the list is generated by listening to on-chain events
	GenerationEvents TGenerationMethods = "Events"
	// GenerationExternalList indicates that the list is generated by retrieving a list from an external source
	GenerationExternalList TGenerationMethods = "External"
	// GenerationLegacyList is the same as GenerationExternalList, but it is for deprecated lists
	GenerationLegacyList TGenerationMethods = "Legacy"

	// GeneratorToken indicates that the list is a token list
	GeneratorToken TGeneratorType = "Token"
	// GeneratorPool indicates that the list is a pool list
	GeneratorPool TGeneratorType = "Pool"
)

type TGenerators struct {
	Exec             func()
	Name             string
	Description      string
	GenerationMethod TGenerationMethods
	GeneratorType    TGeneratorType
	Tags             []string //
}

var GENERATORS = map[string]TGenerators{
	`1inch`: {
		Exec:             build1InchTokenList,
		Name:             `1Inch`,
		Description:      `A list of tokens available in 1Inch DeFi / DEX aggregator`,
		GenerationMethod: GenerationAPI,
		GeneratorType:    GeneratorToken,
	},
	`coingecko`: {
		Exec:             buildCoingeckoTokenList,
		Name:             `CoinGecko`,
		Description:      `A list of tokens available showing in CoinGecko data agregator.`,
		GenerationMethod: GenerationAPI,
		GeneratorType:    GeneratorToken,
	},
	`cowswap`: {
		Exec:             buildCowswapTokenList,
		Name:             `Cow Swap`,
		Description:      `A list of tokens available for trading on CoW Swap, a DEX focused on MEV protection.`,
		GenerationMethod: GenerationExternalList,
		GeneratorType:    GeneratorToken,
	},
	`curve`: {
		Exec:             buildCurveTokenList,
		Name:             `Curve`,
		Description:      `A list of tokens available for trading on Curve, the largest stableswap.`,
		GenerationMethod: GenerationAPI,
		GeneratorType:    GeneratorToken,
	},
	`defillama`: {
		Exec:             buildDefillamaTokenList,
		Name:             `DefiLlama`,
		Description:      `A list of tokens available in DefiLlama token service`,
		GenerationMethod: GenerationExternalList,
		GeneratorType:    GeneratorToken,
	},
	`ledger`: {
		Exec:             buildLedgersTokenList,
		Name:             `Ledger`,
		Description:      `A list of tokens supported in Ledger Live App`,
		GenerationMethod: GenerationExternalList,
		GeneratorType:    GeneratorToken,
	},
	`messari`: {
		Exec:             buildMessariTokenList,
		Name:             `Messari`,
		Description:      `A list of tokens registered in Messari`,
		GenerationMethod: GenerationAPI,
		GeneratorType:    GeneratorToken,
	},
	`optimism`: {
		Exec:             buildOptimismTokenList,
		Name:             `Optimism`,
		Description:      `A list of okens used as the source of truth for the Optimism Gateway.`,
		GenerationMethod: GenerationExternalList,
		GeneratorType:    GeneratorToken,
	},
	`paraswap`: {
		Exec:             buildParaswapTokenList,
		Name:             `Paraswap`,
		Description:      `A list of tokens available for trading on Paraswap DEX`,
		GenerationMethod: GenerationAPI,
		GeneratorType:    GeneratorToken,
	},
	`portals`: {
		Exec:             buildPortalsTokenList,
		Name:             `Portals`,
		Description:      `A list of tokens available for trading on Portals DEX.`,
		GenerationMethod: GenerationAPI,
		GeneratorType:    GeneratorToken,
	},
	`sushiswap-pairs`: {
		Exec:             buildSushiswapPairsTokenList,
		Name:             `SushiSwap (token pairs)`,
		Description:      `A list of token used in the SushiSwap Liquidity Pools.`,
		GenerationMethod: GenerationEvents,
		GeneratorType:    GeneratorToken,
	},
	`sushiswap-pools`: {
		Exec:             buildSushiswapPoolsTokenList,
		Name:             `SushiSwap (pools)`,
		Description:      `A list of Liquidity Pool available on SushiSwap DEX.`,
		GenerationMethod: GenerationEvents,
		GeneratorType:    GeneratorPool,
	},
	`sushiswap`: {
		Exec:             buildSushiswapTokenList,
		Name:             `SushiSwap`,
		Description:      `A list of tokens available on SushiSwap DEX.`,
		GenerationMethod: GenerationExternalList,
		GeneratorType:    GeneratorToken,
	},
	`tokenlistooor`: {
		Exec:             buildTokenListooorList,
		Name:             `Tokenlistooor`,
		Description:      `An aggregated list of tokens from Paraswap, Yearn, and Curve`,
		GenerationMethod: GenerationAPI,
		GeneratorType:    GeneratorToken,
	},
	`uniswap-pairs`: {
		Exec:             buildUniswapPairsTokenList,
		Name:             `UniSwap (pairs)`,
		Description:      `A list of token pairs (liquidity pools) available for trading on UniSwap.`,
		GenerationMethod: GenerationEvents,
		GeneratorType:    GeneratorToken,
	},
	`uniswap-pools`: {
		Exec:             buildUniswapPoolsTokenList,
		Name:             `UniSwap (pools)`,
		Description:      `A list of Liquidity Pool available on Uniswap V2 DEX.`,
		GenerationMethod: GenerationEvents,
		GeneratorType:    GeneratorPool,
	},
	`uniswap`: {
		Exec:             buildUniswapTokenList,
		Name:             `UniSwap`,
		Description:      `A list of tokens available on UniSwap DEX.`,
		GenerationMethod: GenerationExternalList,
		GeneratorType:    GeneratorToken,
	},
	`wido`: {
		Exec:             buildWidoTokenList,
		Name:             `Wido`,
		Description:      `A list of tokens supported by the Wido Router`,
		GenerationMethod: GenerationAPI,
		GeneratorType:    GeneratorToken,
	},
	`yearn-extended`: {
		Exec:             buildYearnExtendedTokenList,
		Name:             `Yearn Extended`,
		Description:      `A list of tokens available for depositing in Yearn, the tokens that represent yVaults and any related tokens.`,
		GenerationMethod: GenerationExternalList,
		GeneratorType:    GeneratorToken,
	},
	`yearn`: {
		Exec:             buildYearnTokenList,
		Name:             `Yearn`,
		Description:      `A list of Yearn's vaults and their underlying tokens.`,
		GenerationMethod: GenerationAPI,
		GeneratorType:    GeneratorToken,
	},
}
