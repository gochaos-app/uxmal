package cmd

import (
	run "github.com/mental12345/uxmal/k8s/run"
	"github.com/spf13/cobra"
)

// runJobCmd represents the runJob command
var runJobCmd = &cobra.Command{
	Use:   "run",
	Short: "Creates a job and outputs the status as well as any generated log.",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		image, _ := cmd.Flags().GetString("img")
		command, _ := cmd.Flags().GetString("cmd")
		ns, _ := cmd.Flags().GetString("ns")
		run.RunJob(name, image, command, ns)
	},
}

func init() {
	jobCmd.AddCommand(runJobCmd)
	runJobCmd.Flags().String("name", "", "Name of job")
	runJobCmd.MarkFlagRequired("name")
	runJobCmd.Flags().String("img", "", "Container image to used")
	runJobCmd.MarkFlagRequired("img")
	runJobCmd.Flags().String("cmd", "", "Command")
	runJobCmd.MarkFlagRequired("cmd")
	runJobCmd.Flags().String("ns", "", "Namespace")

}
