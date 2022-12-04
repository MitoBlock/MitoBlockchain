package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "mitoblockchaindev/testutil/keeper"
	"mitoblockchaindev/x/mitoblockchaindev/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MitoblockchaindevKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
