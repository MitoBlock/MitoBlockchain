package keeper

import (
	"context"

	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"mitoblockchaindev/x/mitoblockchaindev/types"
)

func (k msgServer) DeleteMembershipTokenStatus(goCtx context.Context, msg *types.MsgDeleteMembershipTokenStatus) (*types.MsgDeleteMembershipTokenStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	membershipTokenStatus, exist := k.GetMembershipTokenStatus(ctx, msg.MembershipTokenStatusID)
	if !exist {
		return nil, sdkerrors.Wrapf(types.ErrID, "membership token status doesnt exist")
	}

	if msg.TokenID != membershipTokenStatus.TokenID {
		return nil, sdkerrors.Wrapf(types.ErrID, "Membership token Id does not exist for which status with token Id %d was made", msg.TokenID)
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MembershipTokenStatusKey))
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, membershipTokenStatus.Id)
	store.Delete(bz)

	return &types.MsgDeleteMembershipTokenStatusResponse{}, nil
}
