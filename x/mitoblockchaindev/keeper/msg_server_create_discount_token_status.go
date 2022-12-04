package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"mitoblockchaindev/x/mitoblockchaindev/types"
)

func (k msgServer) CreateDiscountTokenStatus(goCtx context.Context, msg *types.MsgCreateDiscountTokenStatus) (*types.MsgCreateDiscountTokenStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the discount token exists for which a status is being created
	discountToken, found := k.GetDiscountToken(ctx, msg.TokenID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Create variable of type discount token status
	var discountTokenStatus = types.DiscountTokenStatus{
		Creator:   msg.Creator,
		Id:        msg.Id,
		Timestamp: msg.Timestamp,
		Status:    msg.Status,
		TokenID:   msg.TokenID,
		CreatedAt: ctx.BlockHeight(),
	}

	// Check if the status is older than the token. If more than 100 blocks, then return error.
	if discountTokenStatus.CreatedAt > discountToken.CreatedAt+1000000 {
		return nil, sdkerrors.Wrapf(types.ErrTokenStatusOld, "status created at %d is older than token created at %d", discountTokenStatus.CreatedAt, discountToken.CreatedAt)
	}

	id := k.AppendDiscountTokenStatus(ctx, discountTokenStatus)

	return &types.MsgCreateDiscountTokenStatusResponse{Id: id}, nil
}
