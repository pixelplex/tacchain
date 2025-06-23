package app

import (
	erc20types "github.com/cosmos/evm/x/erc20/types"
	"github.com/ethereum/go-ethereum/common"
)

// NOTE: This is the WToken contract on EVM Testnet (Saint Peterburg)
const WTACContract = "0xCf61405b7525F09f4E7501fc831fE7cbCc823d4c"

const gTACDenom = "ugtac"

var GTACTokenPair = erc20types.NewTokenPair(
	common.HexToAddress(WTACContract),
	gTACDenom,
	erc20types.OWNER_MODULE,
)

var TacTokenPairs = []erc20types.TokenPair{GTACTokenPair}
