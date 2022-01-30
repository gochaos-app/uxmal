package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// jobCmd represents the job command
var jobCmd = &cobra.Command{
	Use:   "job",
	Short: "Operations with jobs in kubernetes cluster.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use create, delete logs or status sub-commands")
	},
}

func init() {
	rootCmd.AddCommand(jobCmd)
}
