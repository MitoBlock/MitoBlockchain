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

func createNDiscountTokenStatus(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.DiscountTokenStatus {
	items := make([]types.DiscountTokenStatus, n)
	for i := range items {
		items[i].Id = keeper.AppendDiscountTokenStatus(ctx, items[i])
	}
	return items
}

func TestDiscountTokenStatusGet(t *testing.T) {
	keeper, ctx := keepertest.MitoblockchaindevKeeper(t)
	items := createNDiscountTokenStatus(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetDiscountTokenStatus(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestDiscountTokenStatusRemove(t *testing.T) {
	keeper, ctx := keepertest.MitoblockchaindevKeeper(t)
	items := createNDiscountTokenStatus(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveDiscountTokenStatus(ctx, item.Id)
		_, found := keeper.GetDiscountTokenStatus(ctx, item.Id)
		require.False(t, found)
	}
}

func TestDiscountTokenStatusGetAll(t *testing.T) {
	keeper, ctx := keepertest.MitoblockchaindevKeeper(t)
	items := createNDiscountTokenStatus(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllDiscountTokenStatus(ctx)),
	)
}

func TestDiscountTokenStatusCount(t *testing.T) {
	keeper, ctx := keepertest.MitoblockchaindevKeeper(t)
	items := createNDiscountTokenStatus(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetDiscountTokenStatusCount(ctx))
}
