package app

import (
	"fmt"

	upgradetypes "cosmossdk.io/x/upgrade/types"

	"github.com/Asphere-xyz/tacchain/app/upgrades"
	liquidstake_upgrade "github.com/Asphere-xyz/tacchain/app/upgrades/liquidstake"
	v0010 "github.com/Asphere-xyz/tacchain/app/upgrades/v0.0.10"
	v0011 "github.com/Asphere-xyz/tacchain/app/upgrades/v0.0.11"
	v009 "github.com/Asphere-xyz/tacchain/app/upgrades/v0.0.9"
	v050tov053 "github.com/Asphere-xyz/tacchain/app/upgrades/v050-to-v053"
)

// Upgrades list of chain upgrades
var Upgrades = []upgrades.Upgrade{
	v009.Upgrade,
	v0010.Upgrade,
	v0011.Upgrade,
	liquidstake_upgrade.Upgrade,
	v050tov053.Upgrade,
}

// RegisterUpgradeHandlers registers the chain upgrade handlers
func (app *TacChainApp) RegisterUpgradeHandlers() {
	keepers := upgrades.AppKeepers{
		AccountKeeper:         &app.AccountKeeper,
		ParamsKeeper:          &app.ParamsKeeper,
		ConsensusParamsKeeper: &app.ConsensusParamsKeeper,
		Codec:                 app.appCodec,
		GetStoreKey:           app.GetKey,
		LiquidStakeKeeper:     &app.LiquidStakeKeeper,
		BankKeeper:            app.BankKeeper,
		Erc20Keeper:           &app.Erc20Keeper,
	}
	app.GetStoreKeys()
	// register all upgrade handlers
	for _, upgrade := range Upgrades {
		app.UpgradeKeeper.SetUpgradeHandler(
			upgrade.UpgradeName,
			upgrade.CreateUpgradeHandler(
				app.ModuleManager,
				app.configurator,
				&keepers,
			),
		)
	}

	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(fmt.Sprintf("failed to read upgrade info from disk %s", err))
	}

	if app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		return
	}

	// register store loader for current upgrade
	for _, upgrade := range Upgrades {
		if upgradeInfo.Name == upgrade.UpgradeName {
			app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &upgrade.StoreUpgrades))
			break
		}
	}
}
