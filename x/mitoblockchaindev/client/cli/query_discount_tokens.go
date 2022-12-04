package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"mitoblockchaindev/x/mitoblockchaindev/types"
)

var _ = strconv.Itoa(0)

func CmdDiscountTokens() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "discount-tokens",
		Short: "Query discountTokens",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryDiscountTokensRequest{}

			res, err := queryClient.DiscountTokens(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
