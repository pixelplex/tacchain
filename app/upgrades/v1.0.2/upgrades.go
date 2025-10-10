package v102

// Upgrade for implementing liquid stake module

import (
	"context"

	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"

	"github.com/Asphere-xyz/tacchain/app/upgrades"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	evmerc20types "github.com/cosmos/evm/x/erc20/types"

	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/crypto/sha3"

)

// UpgradeName defines the on-chain upgrade name
const UpgradeName = "v1.0.2"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: storetypes.StoreUpgrades{
		Added:   []string{"utacliquidstake", "epochs"},
		Deleted: []string{},
	},
}

func generateAddressFromDenom(denom string) (common.Address, error) {
	hash := sha3.NewLegacyKeccak256()
	if _, err := hash.Write([]byte(denom)); err != nil {
		return common.Address{}, err
	}
	return common.BytesToAddress(hash.Sum(nil)), nil
}

func CreateUpgradeHandler(
	mm upgrades.ModuleManager,
	configurator module.Configurator,
	ak *upgrades.AppKeepers,
) upgradetypes.UpgradeHandler {
	return func(ctx context.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		newVM, err := mm.RunMigrations(ctx, configurator, fromVM)
		if err != nil {
			return newVM, err
		}

		sdkCtx := sdk.UnwrapSDKContext(ctx)

		// Register gTAC token pair
		lsmBondDenom := ak.LiquidStakeKeeper.LiquidBondDenom(sdkCtx)
		lsmBondCommonAddress, err := generateAddressFromDenom(lsmBondDenom)
		if err != nil {
			return newVM, err
		}
		ak.BankKeeper.SetDenomMetaData(ctx, GTACMetadata)

		erc20Params := ak.Erc20Keeper.GetParams(sdkCtx)
		erc20Params.NativePrecompiles = append(erc20Params.NativePrecompiles, lsmBondCommonAddress.String())
		if err := ak.Erc20Keeper.SetParams(sdkCtx, erc20Params); err != nil {
			return newVM, err
		}

		lsmTokenPair := evmerc20types.NewTokenPair(lsmBondCommonAddress, lsmBondDenom, evmerc20types.OWNER_MODULE)

		ak.Erc20Keeper.SetToken(sdkCtx, lsmTokenPair)

		return newVM, nil
	}
}
