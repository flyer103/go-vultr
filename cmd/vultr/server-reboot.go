package vultr

import (
	"github.com/spf13/cobra"

	pvultr "github.com/flyer103/go-vultr/pkg/vultr"
)

var rebootAll bool

var serverRebootCmd = &cobra.Command{
	Use:           "reboot",
	Short:         "Reboot your server(s)",
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return CmdDo(pvultr.APINameServerReboot, rebootAll, args)
	},
}

func init() {
	serverRebootCmd.Flags().BoolVarP(&rebootAll, "all", "a", false, "reboot all servers?")

	serverCmd.AddCommand(serverRebootCmd)
}
