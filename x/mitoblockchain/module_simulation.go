package mitoblockchain

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/mitoblock/mitoblockchain/testutil/sample"
	mitoblockchainsimulation "github.com/mitoblock/mitoblockchain/x/mitoblockchain/simulation"
	"github.com/mitoblock/mitoblockchain/x/mitoblockchain/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = mitoblockchainsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
opWeightMsgCreateDiscountToken = "op_weight_msg_create_discount_token"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateDiscountToken int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	mitoblockchainGenesis := types.GenesisState{
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&mitoblockchainGenesis)
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
		mitoblockchainsimulation.SimulateMsgCreateDiscountToken(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
