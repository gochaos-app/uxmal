package cmd

import (
	k8s "github.com/mental12345/uxmal/k8s/create"
	"github.com/spf13/cobra"
)

// createJobCmd represents the createJob command
var createJobCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a simple job given a set of parameters.",
	Long:  `uxmal job create --name "job Name" --img "Image" --cmd "Command"`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		image, _ := cmd.Flags().GetString("img")
		command, _ := cmd.Flags().GetString("cmd")

		k8s.K8sJobs(name, image, command)
	},
}

func init() {
	jobCmd.AddCommand(createJobCmd)
	createJobCmd.Flags().String("name", "", "Name of job")
	createJobCmd.MarkFlagRequired("name")
	createJobCmd.Flags().String("img", "", "Container image to used")
	createJobCmd.MarkFlagRequired("img")
	createJobCmd.Flags().String("cmd", "", "Command")
	createJobCmd.MarkFlagRequired("cmd")

}
