package k8s

import (
	"log"
	"time"

	create "github.com/mental12345/uxmal/k8s/create"
	logs "github.com/mental12345/uxmal/k8s/logs"
	status "github.com/mental12345/uxmal/k8s/status"
)

func RunJob(name string, image string, cmd string, ns string) {
	if ns == "" {
		ns = "default"
	}
	create.K8sJobs(name, image, cmd, ns)

	for {
		time.Sleep(2 * time.Second)
		statusJob, _ := status.GetJobsStatus(name, ns)

		if statusJob == 0 {
			log.Println("Success!! :)")
			logs.LogsJobs(name, ns)
			break
		}
		if statusJob == 2 {
			log.Println("Fail!! :(")
			logs.LogsJobs(name, ns)
			break
		}
	}
}
