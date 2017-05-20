package vultr

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/flyer103/go-vultr/pkg/apikeyctrl"
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

func GetAPIKey() (string, error) {
	ctrl, err := apikeyctrl.New()
	if err != nil {
		return "", err
	}

	apikey, err := ctrl.Get()
	if err != nil {
		return "", err
	}

	return apikey, nil
}
