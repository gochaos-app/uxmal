package k8s

import (
	"log"
	"time"

	create "github.com/mental12345/uxmal/k8s/create"
	logs "github.com/mental12345/uxmal/k8s/logs"
	status "github.com/mental12345/uxmal/k8s/status"
)

func RunJob(name string, image string, cmd string) {
	create.K8sJobs(name, image, cmd)

	for {
		time.Sleep(2 * time.Second)
		statusJob, _ := status.GetJobsStatus(name)

		if statusJob == 0 {
			log.Println("Success!! :)")
			logs.LogsJobs(name)
			break
		}
		if statusJob == 2 {
			log.Println("Fail!! :(")
			logs.LogsJobs(name)
			break
		}
	}
}
