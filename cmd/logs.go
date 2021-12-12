package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// logsCmd represents the logs command
var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Get logs of a given job.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("logs called")
		jobName := args[0]
		fmt.Println(jobName)
	},
}

func init() {
	rootCmd.AddCommand(logsCmd)
}
