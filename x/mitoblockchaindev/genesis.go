package mitoblockchaindev

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"mitoblockchaindev/x/mitoblockchaindev/keeper"
	"mitoblockchaindev/x/mitoblockchaindev/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the discountTokenStatus
	for _, elem := range genState.DiscountTokenStatusList {
		k.SetDiscountTokenStatus(ctx, elem)
	}

	// Set discountTokenStatus count
	k.SetDiscountTokenStatusCount(ctx, genState.DiscountTokenStatusCount)
	// Set all the membershipTokenStatus
	for _, elem := range genState.MembershipTokenStatusList {
		k.SetMembershipTokenStatus(ctx, elem)
	}

	// Set membershipTokenStatus count
	k.SetMembershipTokenStatusCount(ctx, genState.MembershipTokenStatusCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.DiscountTokenStatusList = k.GetAllDiscountTokenStatus(ctx)
	genesis.DiscountTokenStatusCount = k.GetDiscountTokenStatusCount(ctx)
	genesis.MembershipTokenStatusList = k.GetAllMembershipTokenStatus(ctx)
	genesis.MembershipTokenStatusCount = k.GetMembershipTokenStatusCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
