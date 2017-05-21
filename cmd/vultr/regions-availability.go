package vultr

import (
	"fmt"

	"github.com/spf13/cobra"

	pvultr "github.com/flyer103/go-vultr/pkg/vultr"
)

var neeAPIKey bool

var regionsAvailabilityCmd = &cobra.Command{
	Use:           "availability",
	Short:         "show availability plans of the targeted DCID",
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 || args[0] == "" {
			return pvultr.ErrNoDCID
		}

		client, err := NewVultrClient()
		if err != nil {
			return err
		}

		res, err := client.RegionsAvailability(args[0], neeAPIKey)
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
	regionsAvailabilityCmd.Flags().BoolVar(&neeAPIKey, "needAPIKey", false, "need API Key?")

	regionsCmd.AddCommand(regionsAvailabilityCmd)
}
