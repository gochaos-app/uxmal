package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "uxmal status checks the status of a given job name in k8s.",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("status called")
		jobName := args[0]
		fmt.Println(jobName)
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
