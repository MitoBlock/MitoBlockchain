package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mitoblock/mitoblockchain/x/mitoblockchain/types"
)

func (k Keeper) GetMembershipTokenCount(ctx sdk.Context) uint64 {
	// Get the store using storeKey (which is "mitoblockchain") and MembershipTokenCountKey (which is "MembershipToken/count/")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.MembershipTokenCountKey))

	// Convert the MembershipTokenCountKey to bytes
	byteKey := []byte(types.MembershipTokenCountKey)

	// Get the value of the count
	bz := store.Get(byteKey)

	// Return zero if the count value is not found (for example, it's the first Membership token)
	if bz == nil {
		return 0
	}

	// Convert the count into a uint64
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetMembershipTokenCount(ctx sdk.Context, count uint64) {
	// Get the store using storeKey (which is "mitoblockchain") and MembershipTokenCountKey (which is "MembershipToken/count/")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.MembershipTokenCountKey))

	// Convert the MembershipTokenCountKey to bytes
	byteKey := []byte(types.MembershipTokenCountKey)

	// Convert count from uint64 to string and get bytes
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)

	// Set the value of MembershipToken/count/ to count
	store.Set(byteKey, bz)
}

func (k Keeper) AppendMembershipToken(ctx sdk.Context, membershipToken types.MembershipToken) uint64 {
	// Get the current number of membershipTokens in the store
	count := k.GetMembershipTokenCount(ctx)

	// Assign an ID to the membership token based on the number of membership tokens in the store
	membershipToken.Id = count

	// Get the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.MembershipTokenKey))

	// Convert the membership token ID into bytes
	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, membershipToken.Id)

	// Marshal the membership token into bytes
	appendedValue := k.cdc.MustMarshal(&membershipToken)

	// Insert the membership token bytes using membership token ID as a key
	store.Set(byteKey, appendedValue)

	// Update the membership token count
	k.SetMembershipTokenCount(ctx, count+1)
	return count
}
