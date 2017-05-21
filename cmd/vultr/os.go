package vultr

import (
	"fmt"

	"github.com/spf13/cobra"
)

var osCmd = &cobra.Command{
	Use:           "os",
	Short:         "List all OS",
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := NewVultrClient()
		if err != nil {
			return err
		}

		res, err := client.OSList()
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
	RootCmd.AddCommand(osCmd)
}
