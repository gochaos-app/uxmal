package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	k8s "github.com/mental12345/uxmal/k8s/create"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "creates a simple jobs given a set of parameters.",
	Long: `uxmal create --name "job Name" --img "Image" --cmd "Command, you can use double quotes for commnds that requiere arguments"`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		image, _ := cmd.Flags().GetString("img")
		command, _ := cmd.Flags().GetString("cmd")
		script, _ := cmd.Flags().GetString("script")
		k8s.K8sJobs(name, image, command)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().String("name", "", "Name of job")
	createCmd.Flags().String("img", "", "Container image to used")
	createCmd.Flags().String("cmd", "", "Command")
	createCmd.Flags().String("script", "", "script")

}
