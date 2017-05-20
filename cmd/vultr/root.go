package vultr

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	longMsg = "go-vultr is a simple command line tool that can interact with Vultr. " +
		"It will make your life easier using Vultr"

	RootCmd = &cobra.Command{
		Use:          "go-vultr",
		Short:        "go-vultr is a simple command line tool that can interact with Vultr.",
		Long:         longMsg,
		SilenceUsage: true,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("Hi go-vultr")
		},
	}
)

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
}
