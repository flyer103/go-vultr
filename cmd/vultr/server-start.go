package vultr

import (
	"github.com/spf13/cobra"

	pvultr "github.com/flyer103/go-vultr/pkg/vultr"
)

var startAll bool

var serverStartCmd = &cobra.Command{
	Use:           "start",
	Short:         "Start your server(s)",
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return CmdDo(pvultr.APINameServerStart, startAll, args)
	},
}

func init() {
	serverStartCmd.Flags().BoolVarP(&startAll, "all", "a", false, "start all servers?")

	serverCmd.AddCommand(serverStartCmd)
}
