package cmd

import (
	del "github.com/mental12345/uxmal/k8s/delete"
	"github.com/spf13/cobra"
)

// deleteJobCmd represents the deleteJob command
var deleteJobCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a job by name.",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		del.DeleteJobs(name)
	},
}

func init() {
	jobCmd.AddCommand(deleteJobCmd)
	deleteJobCmd.Flags().String("name", "", "Name of job")
	deleteJobCmd.MarkFlagRequired("name")
}
