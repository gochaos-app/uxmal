package cmd

import (
	status "github.com/mental12345/uxmal/k8s/status"
	"github.com/spf13/cobra"
)

// statusJobCmd represents the statusJob command
var statusJobCmd = &cobra.Command{
	Use:   "status",
	Short: "uxmal status checks the status of a given job name in k8s.",
	Run: func(cmd *cobra.Command, args []string) {
		jobName, _ := cmd.Flags().GetString("name")
		status.GetJobsStatus(jobName)
	},
}

func init() {
	jobCmd.AddCommand(statusJobCmd)
	statusJobCmd.Flags().String("name", "", "Name of job")
	statusJobCmd.MarkFlagRequired("name")

}
