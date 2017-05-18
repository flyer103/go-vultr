package vultr

import (
	"github.com/spf13/cobra"

	"github.com/flyer103/go-vultr/pkg/apikeyctrl"
)

var apikeyStoreCmd = &cobra.Command{
	Use:   "store",
	Short: "Store your api key",
	Long:  "Store your api key",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctrl, err := apikeyctrl.New()
		if err != nil {
			return err
		}

		if len(args) == 0 {
			return apikeyctrl.ErrEmptyAPIKey
		}

		return ctrl.Store(args[0])
	},
}

func init() {
	apikeyCmd.AddCommand(apikeyStoreCmd)
}
