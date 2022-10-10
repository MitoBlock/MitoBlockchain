package simulation

import (
	"math/rand"

	"github.com/mitoblock/mitoblockchain/x/mitoblockchain/keeper"
	"github.com/mitoblock/mitoblockchain/x/mitoblockchain/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgCreateMembershipToken(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCreateMembershipToken{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the CreateMembershipToken simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CreateMembershipToken simulation not implemented"), nil, nil
	}
}
