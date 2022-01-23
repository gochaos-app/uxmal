package cmd

import (
	k8s "github.com/mental12345/uxmal/k8s/create"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "creates a simple jobs given a set of parameters.",
	Long:  `uxmal create --name "job Name" --img "Image" --cmd "Command" --script script.txt`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		image, _ := cmd.Flags().GetString("img")
		command, _ := cmd.Flags().GetString("cmd")

		k8s.K8sJobs(name, image, command)

	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().String("name", "", "Name of job")
	createCmd.MarkFlagRequired("name")
	createCmd.Flags().String("img", "", "Container image to used")
	createCmd.MarkFlagRequired("img")
	createCmd.Flags().String("cmd", "", "Command")
	createCmd.MarkFlagRequired("cmd")

}
