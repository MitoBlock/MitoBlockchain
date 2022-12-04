package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"mitoblockchaindev/x/mitoblockchaindev/types"
)

var _ = strconv.Itoa(0)

func CmdCreateDiscountToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-discount-token [timestamp] [activity-name] [score] [message] [discount-value] [eligible-companies] [item-type] [expiry-date]",
		Short: "Broadcast message createDiscountToken",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTimestamp := args[0]
			argActivityName := args[1]
			argScore := args[2]
			argMessage := args[3]
			argDiscountValue := args[4]
			argEligibleCompanies := args[5]
			argItemType := args[6]
			argExpiryDate := args[7]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateDiscountToken(
				clientCtx.GetFromAddress().String(),
				argTimestamp,
				argActivityName,
				argScore,
				argMessage,
				argDiscountValue,
				argEligibleCompanies,
				argItemType,
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
