package liquidstake_upgrade

import (
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

const (
	LsmBondDenom = "stk/utac"
	DisplayDenom = "gTAC"
)

var GTACMetadata = banktypes.Metadata{
	Description: "Liquid Staked TAC token",
	DenomUnits: []*banktypes.DenomUnit{
		{
			Denom:    LsmBondDenom,
			Exponent: 0,
		},
		{
			Denom:    DisplayDenom,
			Exponent: 18,
		},
	},
	Base:    LsmBondDenom,
	Display: DisplayDenom,
	Name:    "Gravity TAC",
	Symbol:  "gTAC",
	URI:     "",
	URIHash: "",
}
