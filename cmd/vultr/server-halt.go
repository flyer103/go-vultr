package vultr

import (
	"github.com/spf13/cobra"

	pvultr "github.com/flyer103/go-vultr/pkg/vultr"
)

var haltAll bool

var serverHaltCmd = &cobra.Command{
	Use:           "halt",
	Short:         "Halt your server(s)",
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return CmdDo(pvultr.APINameServerHalt, haltAll, args)
	},
}

func init() {
	serverHaltCmd.Flags().BoolVarP(&haltAll, "all", "a", false, "halt all servers?")

	serverCmd.AddCommand(serverHaltCmd)
}
