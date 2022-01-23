package cmd

import (
	status "github.com/mental12345/uxmal/k8s/status"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "uxmal status checks the status of a given job name in k8s.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		jobName, _ := cmd.Flags().GetString("name")
		status.GetJobsStatus(jobName)

	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
	statusCmd.Flags().String("name", "", "Name of job")
	statusCmd.MarkFlagRequired("name")

}
