package cli

import (
    "strconv"
	
	"github.com/spf13/cobra"
    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/mitoblock/mitoblockchain/x/mitoblockchain/types"
)

var _ = strconv.Itoa(0)

func CmdCreateMembershipToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-membership-token [timestamp] [activity-name] [score] [message] [membership-duration] [eligible-companies] [expiry-date]",
		Short: "Broadcast message createMembershipToken",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
      		 argTimestamp := args[0]
             argActivityName := args[1]
             argScore := args[2]
             argMessage := args[3]
             argMembershipDuration := args[4]
             argEligibleCompanies := args[5]
             argExpiryDate := args[6]
            
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateMembershipToken(
				clientCtx.GetFromAddress().String(),
				argTimestamp,
				argActivityName,
				argScore,
				argMessage,
				argMembershipDuration,
				argEligibleCompanies,
				argExpiryDate,
				
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

    return cmd
}