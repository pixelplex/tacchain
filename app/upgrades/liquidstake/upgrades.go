package liquidstake_upgrade

import (
	"context"

	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"

	"github.com/Asphere-xyz/tacchain/app/upgrades"
	"github.com/Asphere-xyz/tacchain/utils"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	evmerc20types "github.com/cosmos/evm/x/erc20/types"
)

// UpgradeName defines the on-chain upgrade name
const UpgradeName = "liquidstake"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: storetypes.StoreUpgrades{
		Added:   []string{"utacliquidstake", "epochs"},
		Deleted: []string{},
	},
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
		lsmBondCommonAddress, err := utils.GenerateAddressFromDenom(lsmBondDenom)
		if err != nil {
			return newVM, err
		}
		ak.BankKeeper.SetDenomMetaData(ctx, GTACMetadata)

		ak.Erc20Keeper.SetNativePrecompile(sdkCtx, lsmBondCommonAddress)

		lsmTokenPair := evmerc20types.NewTokenPair(lsmBondCommonAddress, lsmBondDenom, evmerc20types.OWNER_MODULE)

		ak.Erc20Keeper.SetToken(sdkCtx, lsmTokenPair)

		return newVM, nil
	}
}
