package vultr

import (
	"github.com/spf13/cobra"

	pvultr "github.com/flyer103/go-vultr/pkg/vultr"
)

var destroyAll bool

var serverDestroyCmd = &cobra.Command{
	Use:           "destroy",
	Short:         "Destroy your server(s)",
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return CmdDo(pvultr.APINameServerDestroy, destroyAll, args)
	},
}

func init() {
	serverDestroyCmd.Flags().BoolVarP(&destroyAll, "all", "a", false, "destroy all servers?")

	serverCmd.AddCommand(serverDestroyCmd)
}
