package cmd

import (
	logs "github.com/mental12345/uxmal/k8s/logs"
	"github.com/spf13/cobra"
)

// logsCmd represents the logs command
var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Get logs of a given job.",
	Run: func(cmd *cobra.Command, args []string) {
		jobName, _ := cmd.Flags().GetString("name")
		logs.LogsJobs(jobName)
	},
}

func init() {
	rootCmd.AddCommand(logsCmd)
	logsCmd.Flags().String("name", "", "Name of job")
	logsCmd.MarkFlagRequired("name")
}
