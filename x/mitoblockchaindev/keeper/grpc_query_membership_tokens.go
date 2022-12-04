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

func (k Keeper) MembershipTokens(goCtx context.Context, req *types.QueryMembershipTokensRequest) (*types.QueryMembershipTokensResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// Define a variable that will store a list of membership tokens
	var membershipTokens []*types.MembershipToken

	// Get context with the information about the environment
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the key-value module store using the store key (in our case store key is "chain")
	store := ctx.KVStore(k.storeKey)

	// Get the part of the store that keeps membershipTokens (using membership token key, which is "MembershipToken-value-")
	postStore := prefix.NewStore(store, []byte(types.MembershipTokenKey))

	// Paginate the membership tokens store based on PageRequest
	pageRes, err := query.Paginate(postStore, req.Pagination, func(key []byte, value []byte) error {
		var membershipToken types.MembershipToken
		if err := k.cdc.Unmarshal(value, &membershipToken); err != nil {
			return err
		}

		membershipTokens = append(membershipTokens, &membershipToken)

		return nil
	})

	// Throw an error if pagination failed
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Return a struct containing a list of membership tokens and pagination info
	return &types.QueryMembershipTokensResponse{MembershipToken: membershipTokens, Pagination: pageRes}, nil
}
