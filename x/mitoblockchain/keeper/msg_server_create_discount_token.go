package keeper

import (
	"context"

    "github.com/mitoblock/mitoblockchain/x/mitoblockchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)


func (k msgServer) CreateDiscountToken(goCtx context.Context,  msg *types.MsgCreateDiscountToken) (*types.MsgCreateDiscountTokenResponse, error) {
	  // Get the context
	ctx := sdk.UnwrapSDKContext(goCtx)

     // Create variable of type DiscountToken
	 var discountToken = types.DiscountToken{
        Creator: msg.Creator,
		Timestamp: msg.Timestamp,
		ActivityName: msg.ActivityName,
		Score: msg.Score,
		Message: msg.Message,
		DiscountValue: msg.DiscountValue,
		EligibleCompanies: msg.EligibleCompanies,
		ItemType: msg.ItemType,
		ExpiryDate: msg.ExpiryDate,
    }

	// Add a discount token to the store and get back the ID
	id := k.AppendDiscountToken(ctx, discountToken)

	return &types.MsgCreateDiscountTokenResponse{Id: id}, nil
}
