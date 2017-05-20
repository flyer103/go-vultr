package vultr

import (
	"fmt"

	"github.com/spf13/cobra"
)

var apikeyGetCmd = &cobra.Command{
	Use:           "get",
	Short:         "Get the stored api key",
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		apikey, err := GetAPIKey()
		if err != nil {
			return err
		}

		fmt.Println("API KEY:", apikey)
		return nil
	},
}

func init() {
	apikeyCmd.AddCommand(apikeyGetCmd)
}
