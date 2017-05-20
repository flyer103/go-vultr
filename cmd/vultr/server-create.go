package vultr

import (
	"github.com/spf13/cobra"
)

var (
	dcID      string
	vpsPlanID string
	osID      string
)

var serverCreateCmd = &cobra.Command{
	Use:           "create",
	Short:         "Create a server",
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := NewVultrClient()
		if err != nil {
			return err
		}

		return client.ServerCreate(dcID, vpsPlanID, osID)
	},
}

func init() {
	serverCreateCmd.Flags().StringVar(&dcID, "dcid", "", "Vultr DCID")
	serverCreateCmd.Flags().StringVar(&vpsPlanID, "vps-plan-id", "", "Vultr VPSPlanID")
	serverCreateCmd.Flags().StringVar(&osID, "osid", "", "Vultr OSID")

	serverCmd.AddCommand(serverCreateCmd)
}
