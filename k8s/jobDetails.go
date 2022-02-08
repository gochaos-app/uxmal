package k8s

import (
	"context"
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func JobDetails(name string, ns string) string {
	if ns == "" {
		ns = "default"
	}
	fmt.Println(ns)
	clientset, _ := K8sConfig()
	labels := "job-name=" + name
	podList, err := clientset.CoreV1().Pods(ns).List(context.TODO(), metav1.ListOptions{
		LabelSelector: labels,
	})
	if err != nil {
		log.Fatalln("Failed to get pods: ", err)
		return "error"
	}
	var jobNamePod string
	for _, pod := range podList.Items {
		jobNamePod = pod.GetName()
	}
	return jobNamePod
}
