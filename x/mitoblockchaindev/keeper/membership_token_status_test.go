package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "mitoblockchaindev/testutil/keeper"
	"mitoblockchaindev/testutil/nullify"
	"mitoblockchaindev/x/mitoblockchaindev/keeper"
	"mitoblockchaindev/x/mitoblockchaindev/types"
)

func createNMembershipTokenStatus(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.MembershipTokenStatus {
	items := make([]types.MembershipTokenStatus, n)
	for i := range items {
		items[i].Id = keeper.AppendMembershipTokenStatus(ctx, items[i])
	}
	return items
}

func TestMembershipTokenStatusGet(t *testing.T) {
	keeper, ctx := keepertest.MitoblockchaindevKeeper(t)
	items := createNMembershipTokenStatus(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetMembershipTokenStatus(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestMembershipTokenStatusRemove(t *testing.T) {
	keeper, ctx := keepertest.MitoblockchaindevKeeper(t)
	items := createNMembershipTokenStatus(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveMembershipTokenStatus(ctx, item.Id)
		_, found := keeper.GetMembershipTokenStatus(ctx, item.Id)
		require.False(t, found)
	}
}

func TestMembershipTokenStatusGetAll(t *testing.T) {
	keeper, ctx := keepertest.MitoblockchaindevKeeper(t)
	items := createNMembershipTokenStatus(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllMembershipTokenStatus(ctx)),
	)
}

func TestMembershipTokenStatusCount(t *testing.T) {
	keeper, ctx := keepertest.MitoblockchaindevKeeper(t)
	items := createNMembershipTokenStatus(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetMembershipTokenStatusCount(ctx))
}
