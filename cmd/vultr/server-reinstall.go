package vultr

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	pvultr "github.com/flyer103/go-vultr/pkg/vultr"
)

var reinstallAll bool
var ErrNoSUBID = errors.New("No SUBID")

var serverReinstallCmd = &cobra.Command{
	Use:           "reinstall",
	Short:         "Reinstall server",
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

		if reinstallAll {
			res, err := client.ServerReinstallAll()
			if err != nil {
				return err
			}

			dataBytes, err := json.MarshalIndent(res, "", "    ")
			if err != nil {
				return err
			}

			fmt.Println(string(dataBytes))
			return nil
		}

		if len(args) == 0 || args[0] == "" {
			return ErrNoSUBID
		}

		return client.ServerReinstall(args[0])
	},
}

func init() {
	serverCmd.AddCommand(serverReinstallCmd)

	serverReinstallCmd.Flags().BoolVarP(&reinstallAll, "all", "a", false, "Reinstall all server?")
}
