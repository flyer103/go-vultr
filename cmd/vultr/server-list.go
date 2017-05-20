package vultr

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	pvultr "github.com/flyer103/go-vultr/pkg/vultr"
)

var serverListCmd = &cobra.Command{
	Use:           "list",
	Short:         "List server info",
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		apikey, err := GetAPIKey()
		if err != nil {
			return err
		}

		cfg := &pvultr.Config{
			APIKey: apikey,
		}
		client, err := pvultr.New(cfg)
		if err != nil {
			return err
		}

		info, err := client.ServerList()
		if err != nil {
			return err
		}

		dataBytes, err := json.MarshalIndent(info, "", "    ")
		if err != nil {
			return err
		}
		fmt.Println(string(dataBytes))

		return nil
	},
}

func init() {
	serverCmd.AddCommand(serverListCmd)
}
