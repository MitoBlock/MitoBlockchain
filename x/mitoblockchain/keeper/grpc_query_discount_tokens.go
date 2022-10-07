package keeper

import (
    "context"

    "github.com/cosmos/cosmos-sdk/store/prefix"
    sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/cosmos/cosmos-sdk/types/query"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"

    "github.com/mitoblock/mitoblockchain/x/mitoblockchain/types"
)

func (k Keeper) DiscountTokens(c context.Context, req *types.QueryDiscountTokensRequest) (*types.QueryDiscountTokensResponse, error) {
    // Throw an error if request is nil
    if req == nil {
        return nil, status.Error(codes.InvalidArgument, "invalid request")
    }

    // Define a variable that will store a list of posts
    var discountTokens []*types.DiscountToken

    // Get context with the information about the environment
    ctx := sdk.UnwrapSDKContext(c)

    // Get the key-value module store using the store key (in our case store key is "chain")
    store := ctx.KVStore(k.storeKey)

    // Get the part of the store that keeps discountTokens (using discount token key, which is "DiscountToken-value-")
    postStore := prefix.NewStore(store, []byte(types.DiscountTokenKey))

    // Paginate the discount tokens store based on PageRequest
    pageRes, err := query.Paginate(postStore, req.Pagination, func(key []byte, value []byte) error {
        var discountToken types.DiscountToken
        if err := k.cdc.Unmarshal(value, &discountToken); err != nil {
            return err
        }

        discountTokens = append(discountTokens, &discountToken)

        return nil
    })

    // Throw an error if pagination failed
    if err != nil {
        return nil, status.Error(codes.Internal, err.Error())
    }

    // Return a struct containing a list of discount tokens and pagination info
    return &types.QueryDiscountTokensResponse{DiscountToken: discountTokens, Pagination: pageRes}, nil
}