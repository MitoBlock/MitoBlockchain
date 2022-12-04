package mitoblockchaindev_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "mitoblockchaindev/testutil/keeper"
	"mitoblockchaindev/testutil/nullify"
	"mitoblockchaindev/x/mitoblockchaindev"
	"mitoblockchaindev/x/mitoblockchaindev/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		DiscountTokenStatusList: []types.DiscountTokenStatus{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		DiscountTokenStatusCount: 2,
		MembershipTokenStatusList: []types.MembershipTokenStatus{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		MembershipTokenStatusCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MitoblockchaindevKeeper(t)
	mitoblockchaindev.InitGenesis(ctx, *k, genesisState)
	got := mitoblockchaindev.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.DiscountTokenStatusList, got.DiscountTokenStatusList)
	require.Equal(t, genesisState.DiscountTokenStatusCount, got.DiscountTokenStatusCount)
	require.ElementsMatch(t, genesisState.MembershipTokenStatusList, got.MembershipTokenStatusList)
	require.Equal(t, genesisState.MembershipTokenStatusCount, got.MembershipTokenStatusCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
