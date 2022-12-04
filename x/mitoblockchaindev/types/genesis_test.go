package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"mitoblockchaindev/x/mitoblockchaindev/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				DiscountTokenStatusList: []types.DiscountTokenStatus{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				DiscountTokenStatusCount: 2,
				MembershipTokenStatusList: []types.MembershipTokenStatus{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				MembershipTokenStatusCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated discountTokenStatus",
			genState: &types.GenesisState{
				DiscountTokenStatusList: []types.DiscountTokenStatus{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid discountTokenStatus count",
			genState: &types.GenesisState{
				DiscountTokenStatusList: []types.DiscountTokenStatus{
					{
						Id: 1,
					},
				},
				DiscountTokenStatusCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated membershipTokenStatus",
			genState: &types.GenesisState{
				MembershipTokenStatusList: []types.MembershipTokenStatus{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid membershipTokenStatus count",
			genState: &types.GenesisState{
				MembershipTokenStatusList: []types.MembershipTokenStatus{
					{
						Id: 1,
					},
				},
				MembershipTokenStatusCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
