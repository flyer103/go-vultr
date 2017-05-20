package vultr

import (
	"log"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:           "server",
	Short:         "Manipulate your server in Vultr",
	SilenceErrors: true,
	SilenceUsage:  true,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Hi server")
	},
}

func init() {
	RootCmd.AddCommand(serverCmd)
}
