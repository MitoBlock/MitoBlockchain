package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "mitoblockchaindev/testutil/keeper"
	"mitoblockchaindev/x/mitoblockchaindev/keeper"
	"mitoblockchaindev/x/mitoblockchaindev/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.MitoblockchaindevKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
