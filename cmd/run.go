package cmd

import (
	run "github.com/mental12345/uxmal/k8s/run"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Creates a job and outputs the status as well as any generated log.",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		image, _ := cmd.Flags().GetString("img")
		command, _ := cmd.Flags().GetString("cmd")
		run.RunJob(name, image, command)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().String("name", "", "Name of job")
	runCmd.MarkFlagRequired("name")
	runCmd.Flags().String("img", "", "Container image to used")
	runCmd.MarkFlagRequired("img")
	runCmd.Flags().String("cmd", "", "Command")
	runCmd.MarkFlagRequired("cmd")
}
