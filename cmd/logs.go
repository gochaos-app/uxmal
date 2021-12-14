package cmd

import (
	"github.com/spf13/cobra"
	logs "github.com/mental12345/uxmal/k8s/logs"

)

// logsCmd represents the logs command
var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Get logs of a given job.",
	Run: func(cmd *cobra.Command, args []string) {
		jobName := args[0]
		logs.LogsJobs(jobName)
	},
}

func init() {
	rootCmd.AddCommand(logsCmd)
}
