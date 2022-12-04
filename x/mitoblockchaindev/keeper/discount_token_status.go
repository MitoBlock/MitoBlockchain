package keeper

import (
	"encoding/binary"

	"mitoblockchaindev/x/mitoblockchaindev/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetDiscountTokenStatusCount get the total number of discountTokenStatus
func (k Keeper) GetDiscountTokenStatusCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.DiscountTokenStatusCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetDiscountTokenStatusCount set the total number of discountTokenStatus
func (k Keeper) SetDiscountTokenStatusCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.DiscountTokenStatusCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendDiscountTokenStatus appends a discountTokenStatus in the store with a new id and update the count
func (k Keeper) AppendDiscountTokenStatus(
	ctx sdk.Context,
	discountTokenStatus types.DiscountTokenStatus,
) uint64 {
	// Create the discountTokenStatus
	count := k.GetDiscountTokenStatusCount(ctx)

	// Set the ID of the appended value
	discountTokenStatus.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DiscountTokenStatusKey))
	appendedValue := k.cdc.MustMarshal(&discountTokenStatus)
	store.Set(GetDiscountTokenStatusIDBytes(discountTokenStatus.Id), appendedValue)

	// Update discountTokenStatus count
	k.SetDiscountTokenStatusCount(ctx, count+1)

	return count
}

// SetDiscountTokenStatus set a specific discountTokenStatus in the store
func (k Keeper) SetDiscountTokenStatus(ctx sdk.Context, discountTokenStatus types.DiscountTokenStatus) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DiscountTokenStatusKey))
	b := k.cdc.MustMarshal(&discountTokenStatus)
	store.Set(GetDiscountTokenStatusIDBytes(discountTokenStatus.Id), b)
}

// GetDiscountTokenStatus returns a discountTokenStatus from its id
func (k Keeper) GetDiscountTokenStatus(ctx sdk.Context, id uint64) (val types.DiscountTokenStatus, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DiscountTokenStatusKey))
	b := store.Get(GetDiscountTokenStatusIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDiscountTokenStatus removes a discountTokenStatus from the store
func (k Keeper) RemoveDiscountTokenStatus(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DiscountTokenStatusKey))
	store.Delete(GetDiscountTokenStatusIDBytes(id))
}

// GetAllDiscountTokenStatus returns all discountTokenStatus
func (k Keeper) GetAllDiscountTokenStatus(ctx sdk.Context) (list []types.DiscountTokenStatus) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DiscountTokenStatusKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.DiscountTokenStatus
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetDiscountTokenStatusIDBytes returns the byte representation of the ID
func GetDiscountTokenStatusIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetDiscountTokenStatusIDFromBytes returns ID in uint64 format from a byte array
func GetDiscountTokenStatusIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
