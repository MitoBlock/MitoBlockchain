package mitoblockchaindev

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"mitoblockchaindev/testutil/sample"
	mitoblockchaindevsimulation "mitoblockchaindev/x/mitoblockchaindev/simulation"
	"mitoblockchaindev/x/mitoblockchaindev/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = mitoblockchaindevsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateDiscountToken = "op_weight_msg_create_discount_token"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateDiscountToken int = 100

	opWeightMsgCreateDiscountTokenStatus = "op_weight_msg_create_discount_token_status"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateDiscountTokenStatus int = 100

	opWeightMsgDeleteDiscountTokenStatus = "op_weight_msg_delete_discount_token_status"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteDiscountTokenStatus int = 100

	opWeightMsgCreateMembershipToken = "op_weight_msg_create_membership_token"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateMembershipToken int = 100

	opWeightMsgCreateMembershipTokenStatus = "op_weight_msg_create_membership_token_status"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateMembershipTokenStatus int = 100

	opWeightMsgDeleteMembershipTokenStatus = "op_weight_msg_delete_membership_token_status"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteMembershipTokenStatus int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	mitoblockchaindevGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&mitoblockchaindevGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateDiscountToken int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateDiscountToken, &weightMsgCreateDiscountToken, nil,
		func(_ *rand.Rand) {
			weightMsgCreateDiscountToken = defaultWeightMsgCreateDiscountToken
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateDiscountToken,
		mitoblockchaindevsimulation.SimulateMsgCreateDiscountToken(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateDiscountTokenStatus int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateDiscountTokenStatus, &weightMsgCreateDiscountTokenStatus, nil,
		func(_ *rand.Rand) {
			weightMsgCreateDiscountTokenStatus = defaultWeightMsgCreateDiscountTokenStatus
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateDiscountTokenStatus,
		mitoblockchaindevsimulation.SimulateMsgCreateDiscountTokenStatus(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteDiscountTokenStatus int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteDiscountTokenStatus, &weightMsgDeleteDiscountTokenStatus, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteDiscountTokenStatus = defaultWeightMsgDeleteDiscountTokenStatus
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteDiscountTokenStatus,
		mitoblockchaindevsimulation.SimulateMsgDeleteDiscountTokenStatus(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateMembershipToken int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateMembershipToken, &weightMsgCreateMembershipToken, nil,
		func(_ *rand.Rand) {
			weightMsgCreateMembershipToken = defaultWeightMsgCreateMembershipToken
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateMembershipToken,
		mitoblockchaindevsimulation.SimulateMsgCreateMembershipToken(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateMembershipTokenStatus int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateMembershipTokenStatus, &weightMsgCreateMembershipTokenStatus, nil,
		func(_ *rand.Rand) {
			weightMsgCreateMembershipTokenStatus = defaultWeightMsgCreateMembershipTokenStatus
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateMembershipTokenStatus,
		mitoblockchaindevsimulation.SimulateMsgCreateMembershipTokenStatus(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteMembershipTokenStatus int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteMembershipTokenStatus, &weightMsgDeleteMembershipTokenStatus, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteMembershipTokenStatus = defaultWeightMsgDeleteMembershipTokenStatus
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteMembershipTokenStatus,
		mitoblockchaindevsimulation.SimulateMsgDeleteMembershipTokenStatus(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
