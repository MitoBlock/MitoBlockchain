package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"mitoblockchaindev/x/mitoblockchaindev/types"
)

var _ = strconv.Itoa(0)

func CmdDeleteMembershipTokenStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-membership-token-status [membership-token-status-id] [token-id]",
		Short: "Broadcast message delete-membership-token-status",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argMembershipTokenStatusID, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argTokenID, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteMembershipTokenStatus(
				clientCtx.GetFromAddress().String(),
				argMembershipTokenStatusID,
				argTokenID,
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
