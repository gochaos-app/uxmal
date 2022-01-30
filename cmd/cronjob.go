package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cronjobCmd represents the cronjob command
var cronjobCmd = &cobra.Command{
	Use:   "cron",
	Short: "Operations with cronjobs in kubernetes cluster.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use create, delete logs or status sub-commands")
	},
}

func init() {
	rootCmd.AddCommand(cronjobCmd)
}
