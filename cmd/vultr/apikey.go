package vultr

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	apikey string
)

var apikeyCmd = &cobra.Command{
	Use:           "apikey",
	Short:         "Manage your api key.",
	SilenceErrors: true,
	SilenceUsage:  true,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Hi apikey")
	},
}

func init() {
	RootCmd.AddCommand(apikeyCmd)

	apikeyCmd.Flags().StringVar(&apikey, "apikey", "", "Your Vultr API Key")
}
