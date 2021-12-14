package cmd

import (
	"github.com/spf13/cobra"
	status "github.com/mental12345/uxmal/k8s/status"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "uxmal status checks the status of a given job name in k8s.",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		jobName := args[0]        
		status.GetJobsStatus(jobName)

	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
