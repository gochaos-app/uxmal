package cmd

import (
	del "github.com/gochaos-app/uxmal/k8s/delete"

	"github.com/spf13/cobra"
)

// deleteCronCmd represents the deleteCron command
var deleteCronCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a cronjob by name.",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		ns, _ := cmd.Flags().GetString("ns")
		del.DeleteCronjobs(name, ns)
	},
}

func init() {
	cronjobCmd.AddCommand(deleteCronCmd)
	deleteCronCmd.Flags().String("name", "", "Name of job")
	deleteCronCmd.MarkFlagRequired("name")
	deleteCronCmd.Flags().String("ns", "", "Namespace")
}
