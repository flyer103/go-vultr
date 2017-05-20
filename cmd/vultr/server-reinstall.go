package vultr

import (
	"github.com/spf13/cobra"

	pvultr "github.com/flyer103/go-vultr/pkg/vultr"
)

var reinstallAll bool

var serverReinstallCmd = &cobra.Command{
	Use:           "reinstall",
	Short:         "Reinstall server(s)",
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return CmdDo(pvultr.APINameServerReinstall, reinstallAll, args)
	},
}

func init() {
	serverReinstallCmd.Flags().BoolVarP(&reinstallAll, "all", "a", false, "reinstall all server?")

	serverCmd.AddCommand(serverReinstallCmd)
}
