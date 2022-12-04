package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mitoblockchaindev/x/mitoblockchaindev/types"
)

func (k Keeper) MembershipTokenStatusAll(c context.Context, req *types.QueryAllMembershipTokenStatusRequest) (*types.QueryAllMembershipTokenStatusResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var membershipTokenStatuss []types.MembershipTokenStatus
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	membershipTokenStatusStore := prefix.NewStore(store, types.KeyPrefix(types.MembershipTokenStatusKey))

	pageRes, err := query.Paginate(membershipTokenStatusStore, req.Pagination, func(key []byte, value []byte) error {
		var membershipTokenStatus types.MembershipTokenStatus
		if err := k.cdc.Unmarshal(value, &membershipTokenStatus); err != nil {
			return err
		}

		membershipTokenStatuss = append(membershipTokenStatuss, membershipTokenStatus)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMembershipTokenStatusResponse{MembershipTokenStatus: membershipTokenStatuss, Pagination: pageRes}, nil
}

func (k Keeper) MembershipTokenStatus(c context.Context, req *types.QueryGetMembershipTokenStatusRequest) (*types.QueryGetMembershipTokenStatusResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	membershipTokenStatus, found := k.GetMembershipTokenStatus(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetMembershipTokenStatusResponse{MembershipTokenStatus: membershipTokenStatus}, nil
}
