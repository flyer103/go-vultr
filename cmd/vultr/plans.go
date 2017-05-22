package vultr

import (
	"fmt"

	"github.com/spf13/cobra"
)

var planType string

var plansCmd = &cobra.Command{
	Use:           "plans",
	Short:         "List plans",
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := NewVultrClient()
		if err != nil {
			return err
		}

		res, err := client.PlansList(planType)
		if err != nil {
			return err
		}

		data, err := PrettyJsonString(res)
		if err != nil {
			return err
		}
		fmt.Println(data)

		return nil
	},
}

func init() {
	plansCmd.Flags().StringVarP(&planType, "type", "t", "all",
		"the type of plans to return. Possible values: 'all', 'vc2', 'ssd', 'vdc2', 'dedicated'")

	RootCmd.AddCommand(plansCmd)
}
