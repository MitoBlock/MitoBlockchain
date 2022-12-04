package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"mitoblockchaindev/x/mitoblockchaindev/types"
)

func (k Keeper) MembershipTokenStatusQ(goCtx context.Context, req *types.QueryMembershipTokenStatusQRequest) (*types.QueryMembershipTokenStatusQResponse, error) {
	// Throw an error if request is nil
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// Define a variable that will store a list of membership token status
	var membershipTokenStatuses []*types.MembershipTokenStatus

	// Get context with the information about the environment
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the key-value module store using the store key (in this case store key is "chain")
	store := ctx.KVStore(k.storeKey)

	// Get the part of the store that keeps membership tokens (using membership token key, which is "MembershipToken-value-")
	membershipTokenStatusStore := prefix.NewStore(store, []byte(types.MembershipTokenStatusKey))

	// Get the membership token by ID
	membershipToken, _ := k.GetMembershipToken(ctx, req.Id)

	// Get the membership token ID
	membershipTokenID := membershipToken.Id

	// Paginate the membership tokens store based on PageRequest
	pageRes, err := query.Paginate(membershipTokenStatusStore, req.Pagination, func(key []byte, value []byte) error {
		var membershipTokenStatus types.MembershipTokenStatus
		if err := k.cdc.Unmarshal(value, &membershipTokenStatus); err != nil {
			return err
		}

		if membershipTokenStatus.TokenID == membershipTokenID {
			membershipTokenStatuses = append(membershipTokenStatuses, &membershipTokenStatus)
		}

		return nil
	})

	// Throw an error if pagination failed
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Return a struct containing a list of membership tokens and pagination info
	return &types.QueryMembershipTokenStatusQResponse{MembershipToken: &membershipToken, MembershipTokenStatus: membershipTokenStatuses, Pagination: pageRes}, nil
}
