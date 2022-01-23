package cmd

import (
	del "github.com/mental12345/uxmal/k8s/delete"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a job given by name",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		del.DeleteJobs(name)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().String("name", "", "Name of job")
	deleteCmd.MarkFlagRequired("name")
}
