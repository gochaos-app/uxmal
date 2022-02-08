package cmd

import (
	logs "github.com/mental12345/uxmal/k8s/logs"
	"github.com/spf13/cobra"
)

// logsJobCmd represents the logsJob command
var logsJobCmd = &cobra.Command{
	Use:   "logs",
	Short: "Get logs of a given job.",
	Run: func(cmd *cobra.Command, args []string) {
		jobName, _ := cmd.Flags().GetString("name")
		ns, _ := cmd.Flags().GetString("ns")
		logs.LogsJobs(jobName, ns)
	},
}

func init() {
	jobCmd.AddCommand(logsJobCmd)
	logsJobCmd.Flags().String("name", "", "Name of job")
	logsJobCmd.MarkFlagRequired("name")
	logsJobCmd.Flags().String("ns", "", "Namespace")

}
