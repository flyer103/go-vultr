package vultr

import (
	"fmt"

	"github.com/spf13/cobra"
)

var regionsCmd = &cobra.Command{
	Use:           "regions",
	Short:         "List all regions.",
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := NewVultrClient()
		if err != nil {
			return err
		}

		res, err := client.RegionsList()
		if err != nil {
			return err
		}

		data, err := PrettyJsonString(res)
		if err != nil {
			return err
		}
		fmt.Println(data)

		return nil
	},
}

func init() {
	RootCmd.AddCommand(regionsCmd)
}
