package keeper

import (
	"context"

	"mitoblockchaindev/x/mitoblockchaindev/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateMembershipToken(goCtx context.Context, msg *types.MsgCreateMembershipToken) (*types.MsgCreateMembershipTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Create variable of type Membership token
	var membershipToken = types.MembershipToken{
		Creator:            msg.Creator,
		Timestamp:          msg.Timestamp,
		ActivityName:       msg.ActivityName,
		Score:              msg.Score,
		Message:            msg.Message,
		MembershipDuration: msg.MembershipDuration,
		EligibleCompanies:  msg.EligibleCompanies,
		ExpiryDate:         msg.ExpiryDate,
		CreatedAt:          ctx.BlockHeight(),
	}

	// Add a Membership token to the store and get back the ID
	id := k.AppendMembershipToken(ctx, membershipToken)

	return &types.MsgCreateMembershipTokenResponse{Id: id}, nil
}
