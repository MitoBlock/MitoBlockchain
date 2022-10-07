package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mitoblock/mitoblockchain/x/mitoblockchain/types"
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

	// Set the value of DiscountToken/count/ to count
	store.Set(byteKey, bz)
}

func (k Keeper) AppendDiscountToken(ctx sdk.Context, discountToken types.DiscountToken) uint64 {
	// Get the current number of discount tokens in the store
	count := k.GetDiscountTokenCount(ctx)

	// Assign an ID to the discount token based on the number of discount tokens in the store
	discountToken.Id = count

	// Get the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.DiscountTokenKey))

	// Convert the discount token ID into bytes
	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, discountToken.Id)

	// Marshal the discount token into bytes
	appendedValue := k.cdc.MustMarshal(&discountToken)

	// Insert the discount token bytes using discount token ID as a key
	store.Set(byteKey, appendedValue)

	// Update the discount token count
	k.SetDiscountTokenCount(ctx, count+1)
	return count
}
