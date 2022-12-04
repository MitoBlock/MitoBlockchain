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

func CmdCreateMembershipTokenStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-membership-token-status [token-id] [timestamp] [status]",
		Short: "Broadcast message create-membership-token-status",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTokenID, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argTimestamp := args[1]
			argStatus := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateMembershipTokenStatus(
				clientCtx.GetFromAddress().String(),
				argTokenID,
				argTimestamp,
				argStatus,
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
