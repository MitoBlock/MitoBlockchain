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
		Timestamp: msg.timestamp,
		ActivityName: msg.activityName,
		Score: msg.score,
		Message: msg.message,
		DiscountValue: msg.discountValue,
		eligibleCompanies: msg.eligibleCompanies
		ItemType: msg.itemType,
		ExpiryDate: msg.expiryDate,
    }

	// Add a discount token to the store and get back the ID
	id := k.AppendDiscountToken(ctx, discountToken)

	return &types.MsgCreateDiscountTokenResponse{Id: id}, nil
}
