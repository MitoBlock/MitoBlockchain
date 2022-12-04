package types

const (
	// ModuleName defines the module name
	ModuleName = "mitoblockchaindev"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_mitoblockchaindev"

	// Keep track of the index of discount tokens
	DiscountTokenKey      = "DiscountToken/value/"
	DiscountTokenCountKey = "DiscountToken/count/"

	// Keep track of the index of membership tokens
	MembershipTokenKey      = "MembershipToken/value/"
	MembershipTokenCountKey = "MembershipToken/count/"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	DiscountTokenStatusKey      = "DiscountTokenStatus/value/"
	DiscountTokenStatusCountKey = "DiscountTokenStatus/count/"
)

const (
	MembershipTokenStatusKey      = "MembershipTokenStatus/value/"
	MembershipTokenStatusCountKey = "MembershipTokenStatus/count/"
)
