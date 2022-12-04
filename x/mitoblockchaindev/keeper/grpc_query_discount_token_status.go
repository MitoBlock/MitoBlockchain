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

func (k Keeper) DiscountTokenStatusAll(c context.Context, req *types.QueryAllDiscountTokenStatusRequest) (*types.QueryAllDiscountTokenStatusResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var discountTokenStatuss []types.DiscountTokenStatus
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	discountTokenStatusStore := prefix.NewStore(store, types.KeyPrefix(types.DiscountTokenStatusKey))

	pageRes, err := query.Paginate(discountTokenStatusStore, req.Pagination, func(key []byte, value []byte) error {
		var discountTokenStatus types.DiscountTokenStatus
		if err := k.cdc.Unmarshal(value, &discountTokenStatus); err != nil {
			return err
		}

		discountTokenStatuss = append(discountTokenStatuss, discountTokenStatus)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllDiscountTokenStatusResponse{DiscountTokenStatus: discountTokenStatuss, Pagination: pageRes}, nil
}

func (k Keeper) DiscountTokenStatus(c context.Context, req *types.QueryGetDiscountTokenStatusRequest) (*types.QueryGetDiscountTokenStatusResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	discountTokenStatus, found := k.GetDiscountTokenStatus(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetDiscountTokenStatusResponse{DiscountTokenStatus: discountTokenStatus}, nil
}
