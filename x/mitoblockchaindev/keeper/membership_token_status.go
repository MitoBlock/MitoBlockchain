package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"mitoblockchaindev/x/mitoblockchaindev/types"
)

// GetMembershipTokenStatusCount get the total number of membershipTokenStatus
func (k Keeper) GetMembershipTokenStatusCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.MembershipTokenStatusCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetMembershipTokenStatusCount set the total number of membershipTokenStatus
func (k Keeper) SetMembershipTokenStatusCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.MembershipTokenStatusCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendMembershipTokenStatus appends a membershipTokenStatus in the store with a new id and update the count
func (k Keeper) AppendMembershipTokenStatus(
	ctx sdk.Context,
	membershipTokenStatus types.MembershipTokenStatus,
) uint64 {
	// Create the membershipTokenStatus
	count := k.GetMembershipTokenStatusCount(ctx)

	// Set the ID of the appended value
	membershipTokenStatus.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MembershipTokenStatusKey))
	appendedValue := k.cdc.MustMarshal(&membershipTokenStatus)
	store.Set(GetMembershipTokenStatusIDBytes(membershipTokenStatus.Id), appendedValue)

	// Update membershipTokenStatus count
	k.SetMembershipTokenStatusCount(ctx, count+1)

	return count
}

// SetMembershipTokenStatus set a specific membershipTokenStatus in the store
func (k Keeper) SetMembershipTokenStatus(ctx sdk.Context, membershipTokenStatus types.MembershipTokenStatus) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MembershipTokenStatusKey))
	b := k.cdc.MustMarshal(&membershipTokenStatus)
	store.Set(GetMembershipTokenStatusIDBytes(membershipTokenStatus.Id), b)
}

// GetMembershipTokenStatus returns a membershipTokenStatus from its id
func (k Keeper) GetMembershipTokenStatus(ctx sdk.Context, id uint64) (val types.MembershipTokenStatus, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MembershipTokenStatusKey))
	b := store.Get(GetMembershipTokenStatusIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMembershipTokenStatus removes a membershipTokenStatus from the store
func (k Keeper) RemoveMembershipTokenStatus(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MembershipTokenStatusKey))
	store.Delete(GetMembershipTokenStatusIDBytes(id))
}

// GetAllMembershipTokenStatus returns all membershipTokenStatus
func (k Keeper) GetAllMembershipTokenStatus(ctx sdk.Context) (list []types.MembershipTokenStatus) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MembershipTokenStatusKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.MembershipTokenStatus
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetMembershipTokenStatusIDBytes returns the byte representation of the ID
func GetMembershipTokenStatusIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetMembershipTokenStatusIDFromBytes returns ID in uint64 format from a byte array
func GetMembershipTokenStatusIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
