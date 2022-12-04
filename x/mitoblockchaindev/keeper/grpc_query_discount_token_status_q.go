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

func (k Keeper) DiscountTokenStatusQ(c context.Context, req *types.QueryDiscountTokenStatusQRequest) (*types.QueryDiscountTokenStatusQResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// Define a variable that will store a list of discount token status
	var discountTokenStatuses []*types.DiscountTokenStatus

	// Get context with the information about the environment
	ctx := sdk.UnwrapSDKContext(c)

	// Get the key-value module store using the store key (in this case store key is "chain")
	store := ctx.KVStore(k.storeKey)

	// Get the part of the store that keeps discount tokens (using discount token key, which is "DiscountToken-value-")
	discountTokenStatusStore := prefix.NewStore(store, []byte(types.DiscountTokenStatusKey))

	// Get the discount token by ID
	discountToken, _ := k.GetDiscountToken(ctx, req.Id)

	// Get the discount token ID
	discountTokenID := discountToken.Id

	// Paginate the discount tokens store based on PageRequest
	pageRes, err := query.Paginate(discountTokenStatusStore, req.Pagination, func(key []byte, value []byte) error {
		var discountTokenStatus types.DiscountTokenStatus
		if err := k.cdc.Unmarshal(value, &discountTokenStatus); err != nil {
			return err
		}

		if discountTokenStatus.TokenID == discountTokenID {
			discountTokenStatuses = append(discountTokenStatuses, &discountTokenStatus)
		}

		return nil
	})

	// Throw an error if pagination failed
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryDiscountTokenStatusQResponse{DiscountToken: &discountToken, DiscountTokenStatus: discountTokenStatuses, Pagination: pageRes}, nil
}
