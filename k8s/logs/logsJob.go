package k8s

import (
	"context"
	"fmt"
	"log"

	k8s "github.com/mental12345/uxmal/k8s"
	status "github.com/mental12345/uxmal/k8s/status"
	v1 "k8s.io/api/core/v1"
)

func LogsJobs(name string, ns string) error {
	clientset, _ := k8s.K8sConfig()
	if ns == "" {
		ns = "default"
	}
	podName := k8s.JobDetails(name, ns)

	podLogOpts := v1.PodLogOptions{}

	podLogRequest := clientset.CoreV1().Pods(ns).GetLogs(podName, &podLogOpts)
	stream, err := podLogRequest.Stream(context.TODO())
	if err != nil {
		return err
	}
	defer stream.Close()
	x, err := status.GetJobsStatus(name, ns)
	if x == 0 || x == 2 {
		buffer := make([]byte, 4000)
		numBytes, err := stream.Read(buffer)
		if err != nil {
			log.Println(err)
			return err
		}
		message := string(buffer[:numBytes])
		fmt.Print(message)
	} else {
		log.Println(err)
		return err
	}

	log.Println("finished")
	return nil
}
