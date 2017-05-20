package vultr

import (
	"github.com/spf13/cobra"

	"github.com/flyer103/go-vultr/pkg/apikeyctrl"
)

var apikeyDeleteCmd = &cobra.Command{
	Use:           "delete",
	Short:         "Delete your api key",
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctrl, err := apikeyctrl.New()
		if err != nil {
			return err
		}

		return ctrl.Delete()
	},
}

func init() {
	apikeyCmd.AddCommand(apikeyDeleteCmd)
}
