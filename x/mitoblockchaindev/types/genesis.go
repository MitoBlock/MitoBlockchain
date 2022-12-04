package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		DiscountTokenStatusList:   []DiscountTokenStatus{},
		MembershipTokenStatusList: []MembershipTokenStatus{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in discountTokenStatus
	discountTokenStatusIdMap := make(map[uint64]bool)
	discountTokenStatusCount := gs.GetDiscountTokenStatusCount()
	for _, elem := range gs.DiscountTokenStatusList {
		if _, ok := discountTokenStatusIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for discountTokenStatus")
		}
		if elem.Id >= discountTokenStatusCount {
			return fmt.Errorf("discountTokenStatus id should be lower or equal than the last id")
		}
		discountTokenStatusIdMap[elem.Id] = true
	}
	// Check for duplicated ID in membershipTokenStatus
	membershipTokenStatusIdMap := make(map[uint64]bool)
	membershipTokenStatusCount := gs.GetMembershipTokenStatusCount()
	for _, elem := range gs.MembershipTokenStatusList {
		if _, ok := membershipTokenStatusIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for membershipTokenStatus")
		}
		if elem.Id >= membershipTokenStatusCount {
			return fmt.Errorf("membershipTokenStatus id should be lower or equal than the last id")
		}
		membershipTokenStatusIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
