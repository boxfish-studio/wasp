package chain

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/iotaledger/wasp/tools/wasp-cli/cli/cliclients"
	"github.com/iotaledger/wasp/tools/wasp-cli/cli/config"
	"github.com/iotaledger/wasp/tools/wasp-cli/log"
)

func initListContractsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list-contracts",
		Short: "List deployed contracts in chain",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			client := cliclients.WaspClient()
			contracts, _, err := client.ChainsApi.
				GetContracts(context.Background(), config.GetCurrentChainID().String()).
				Execute()

			log.Check(err)

			log.Printf("Total %d contracts in chain %s\n", len(contracts), config.GetCurrentChainID())

			header := []string{
				"hname",
				"name",
				"description",
				"proghash",
				"owner fee",
				"validator fee",
			}
			rows := make([][]string, len(contracts))
			i := 0
			for _, contract := range contracts {
				rows[i] = []string{
					contract.HName,
					contract.Name,
					contract.Description,
					contract.ProgramHash,
				}
				i++
			}
			log.PrintTable(header, rows)
		},
	}
}
