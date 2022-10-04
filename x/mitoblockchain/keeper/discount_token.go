package keeper

import (
    "encoding/binary"

    "github.com/cosmos/cosmos-sdk/store/prefix"
    sdk "github.com/cosmos/cosmos-sdk/types"

    "mitoblockchain/x/mitoblockchain/types"
)

func (k Keeper) GetDiscountTokenCount(ctx sdk.Context) uint64 {
    // Get the store using storeKey (which is "discount token") and DiscountTokenCountKey (which is "DiscountToken/count/")
    store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.DiscountTokenCountKey))

    // Convert the DiscountTokenCountKey to bytes
    byteKey := []byte(types.DiscountTokenCountKey)

    // Get the value of the count
    bz := store.Get(byteKey)

    // Return zero if the count value is not found (for example, it's the first discount token)
    if bz == nil {
        return 0
    }

    // Convert the count into a uint64
    return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetDiscountTokenCount(ctx sdk.Context, count uint64) {
    // Get the store using storeKey (which is "discount token") and DiscountTokenCountKey (which is "DiscountToken/count/")
    store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.DiscountTokenCountKey))

    // Convert the DiscountTokenCountKey to bytes
    byteKey := []byte(types.DiscountTokenCountKey)

    // Convert count from uint64 to string and get bytes
    bz := make([]byte, 8)
    binary.BigEndian.PutUint64(bz, count)

    // Set the value of Post/count/ to count
    store.Set(byteKey, bz)
}