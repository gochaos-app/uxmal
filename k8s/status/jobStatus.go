package k8s

import (
	"context"
	"log"

	k8s "github.com/gochaos-app/uxmal/k8s"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetJobsStatus(name string, ns string) (int, error) {
	clientset, _ := k8s.K8sConfig()
	if ns == "" {
		ns = "default"
	}
	jobs := clientset.BatchV1().Jobs(ns)
	job, err := jobs.Get(context.TODO(), name, metav1.GetOptions{})

	if err != nil {
		log.Println("Error: ", err)
		return 2, err
	}

	if job.Status.Active > 0 {
		log.Println("Job is still running")
		return 1, nil
	} else if job.Status.Succeeded > 0 {
		log.Println("Job was succesful")
		return 0, nil
	} else {
		log.Println("Job failed")
		return 2, nil

	}
}
