package cmd

import (
	del "github.com/gochaos-app/uxmal/k8s/delete"
	"github.com/spf13/cobra"
)

// deleteJobCmd represents the deleteJob command
var deleteJobCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a job by name.",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		ns, _ := cmd.Flags().GetString("ns")
		del.DeleteJobs(name, ns)
	},
}

func init() {
	jobCmd.AddCommand(deleteJobCmd)
	deleteJobCmd.Flags().String("name", "", "Name of job")
	deleteJobCmd.MarkFlagRequired("name")
	deleteJobCmd.Flags().String("ns", "", "Namespace")

}
