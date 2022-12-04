package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"mitoblockchaindev/x/mitoblockchaindev/types"
)

func (k msgServer) CreateMembershipTokenStatus(goCtx context.Context, msg *types.MsgCreateMembershipTokenStatus) (*types.MsgCreateMembershipTokenStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the membership token exists for which a status is being created
	membershipToken, found := k.GetMembershipToken(ctx, msg.TokenID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Create variable of type membership token status
	var membershipTokenStatus = types.MembershipTokenStatus{
		Creator:   msg.Creator,
		Id:        msg.Id,
		Timestamp: msg.Timestamp,
		Status:    msg.Status,
		TokenID:   msg.TokenID,
		CreatedAt: ctx.BlockHeight(),
	}

	// Check if the status is older than the token. If more than 1000000 blocks, then return error.
	if membershipTokenStatus.CreatedAt > membershipToken.CreatedAt+1000000 {
		return nil, sdkerrors.Wrapf(types.ErrTokenStatusOld, "status created at %d is older than token created at %d", membershipTokenStatus.CreatedAt, membershipToken.CreatedAt)
	}

	id := k.AppendMembershipTokenStatus(ctx, membershipTokenStatus)

	return &types.MsgCreateMembershipTokenStatusResponse{Id: id}, nil
}
