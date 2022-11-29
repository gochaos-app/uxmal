package k8s

import (
	"context"
	"log"

	k8s "github.com/gochaos-app/uxmal/k8s"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func DeleteJobs(name string, ns string) {
	clientset, _ := k8s.K8sConfig()
	if ns == "" {
		ns = "default"
	}
	err := clientset.BatchV1().Jobs(ns).Delete(context.TODO(), name, metav1.DeleteOptions{})

	if err != nil {
		log.Println(err)
		log.Fatalln("Failed to delete K8s job.")
	}

	//print job details
	log.Println("Deleted K8s job successfully")
	pod_name := k8s.JobDetails(name, ns)
	log.Println("Pod name: ", pod_name)
}

func DeleteCronjobs(name string, ns string) {
	clientset, _ := k8s.K8sConfig()
	if ns == "" {
		ns = "default"
	}
	err := clientset.BatchV1beta1().CronJobs(ns).Delete(context.TODO(), name, metav1.DeleteOptions{})

	if err != nil {
		log.Println(err)
		log.Fatalln("Failed to delete K8s cronjob.")
	}

	//print job details
	log.Println("Deleted K8s cronjob successfully")

}
