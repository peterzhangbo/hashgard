package params

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type IssueParams struct {
	Name            string  `json:"name"`
	Symbol          string  `json:"symbol"`
	TotalSupply     sdk.Int `json:"total_supply"`
	MintingFinished bool    `json:"minting_finished"`
}