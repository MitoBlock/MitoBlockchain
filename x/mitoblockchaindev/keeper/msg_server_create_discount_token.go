package keeper

import (
	"context"

	"mitoblockchaindev/x/mitoblockchaindev/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateDiscountToken(goCtx context.Context, msg *types.MsgCreateDiscountToken) (*types.MsgCreateDiscountTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Create variable of type DiscountToken
	var discountToken = types.DiscountToken{
		Creator:           msg.Creator,
		Timestamp:         msg.Timestamp,
		ActivityName:      msg.ActivityName,
		Score:             msg.Score,
		Message:           msg.Message,
		DiscountValue:     msg.DiscountValue,
		EligibleCompanies: msg.EligibleCompanies,
		ItemType:          msg.ItemType,
		ExpiryDate:        msg.ExpiryDate,
		CreatedAt:         ctx.BlockHeight(),
	}

	// Add a discount token to the store and get back the ID
	id := k.AppendDiscountToken(ctx, discountToken)

	return &types.MsgCreateDiscountTokenResponse{Id: id}, nil
}
