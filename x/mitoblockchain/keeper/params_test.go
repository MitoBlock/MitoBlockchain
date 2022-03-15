package keeper_test

import (
	"testing"

	testkeeper "github.com/mitoblock/mitoblockchain/testutil/keeper"
	"github.com/mitoblock/mitoblockchain/x/mitoblockchain/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MitoblockchainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
