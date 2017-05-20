package vultr

import (
	"fmt"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:           "server",
	Short:         "Manipulate your server in Vultr. The default action is to list your servers.",
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := NewVultrClient()
		if err != nil {
			return err
		}

		info, err := client.ServerList()
		if err != nil {
			return err
		}

		data, err := PrettyJsonString(info)
		if err != nil {
			return err
		}
		fmt.Println(data)

		return nil
	},
}

func init() {
	RootCmd.AddCommand(serverCmd)
}
