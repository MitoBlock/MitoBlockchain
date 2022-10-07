package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/mitoblock/mitoblockchain/x/mitoblockchain/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdDiscountTokens() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "discount-tokens [timestamp] [activity-name] [score] [message] [discount-value] [eligible-companies] [item-type] [expiry-date]",
		Short: "Query discountTokens",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqTimestamp := args[0]
			reqActivityName := args[1]
			reqScore := args[2]
			reqMessage := args[3]
			reqDiscountValue := args[4]
			reqEligibleCompanies := args[5]
			reqItemType := args[6]
			reqExpiryDate := args[7]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryDiscountTokensRequest{

				Timestamp:         reqTimestamp,
				ActivityName:      reqActivityName,
				Score:             reqScore,
				Message:           reqMessage,
				DiscountValue:     reqDiscountValue,
				EligibleCompanies: reqEligibleCompanies,
				ItemType:          reqItemType,
				ExpiryDate:        reqExpiryDate,
			}

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
