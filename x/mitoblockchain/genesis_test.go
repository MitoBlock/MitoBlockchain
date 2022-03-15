package mitoblockchain_test

import (
	"testing"

	keepertest "github.com/mitoblock/mitoblockchain/testutil/keeper"
	"github.com/mitoblock/mitoblockchain/testutil/nullify"
	"github.com/mitoblock/mitoblockchain/x/mitoblockchain"
	"github.com/mitoblock/mitoblockchain/x/mitoblockchain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MitoblockchainKeeper(t)
	mitoblockchain.InitGenesis(ctx, *k, genesisState)
	got := mitoblockchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
