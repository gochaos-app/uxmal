package cmd

import (
	k8s "github.com/gochaos-app/uxmal/k8s/create"
	"github.com/spf13/cobra"
)

// createCronCmd represents the createCron command
var createCronCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a simple cronjob given a set of parameters.",
	Long:  `uxmal cron create --name "myCronjob" --img "Image" --cmd "ls -ltr" --sch "* * * * *" `,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		image, _ := cmd.Flags().GetString("img")
		command, _ := cmd.Flags().GetString("cmd")
		schedule, _ := cmd.Flags().GetString("sch")
		ns, _ := cmd.Flags().GetString("ns")
		k8s.K8sCronJobs(name, image, command, schedule, ns)
	},
}

func init() {
	cronjobCmd.AddCommand(createCronCmd)
	createCronCmd.Flags().String("name", "", "Name of job")
	createCronCmd.MarkFlagRequired("name")
	createCronCmd.Flags().String("img", "", "Container image to used")
	createCronCmd.MarkFlagRequired("img")
	createCronCmd.Flags().String("cmd", "", "Command")
	createCronCmd.MarkFlagRequired("cmd")
	createCronCmd.Flags().String("sch", "", "Schedule")
	createCronCmd.MarkFlagRequired("sch")
	createCronCmd.Flags().String("ns", "", "Namespace")

}
