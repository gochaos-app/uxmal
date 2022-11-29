package k8s

import (
	"context"
	"log"
	"strings"

	k8s "github.com/gochaos-app/uxmal/k8s"
	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/api/batch/v1beta1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func K8sJobs(name string, image string, cmd string, ns string) {
	clientset, _ := k8s.K8sConfig()
	if ns == "" {
		ns = "default"
	}
	jobs := clientset.BatchV1().Jobs(ns)
	var jobbackoff int32 = 0

	jobSpec := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: batchv1.JobSpec{
			Template: v1.PodTemplateSpec{
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:    name,
							Image:   image,
							Command: strings.Split(cmd, " "),
						},
					},
					RestartPolicy: v1.RestartPolicyNever,
				},
			},
			BackoffLimit: &jobbackoff,
		},
	}

	_, err := jobs.Create(context.TODO(), jobSpec, metav1.CreateOptions{})
	if err != nil {
		log.Println(err)
		log.Fatalln("Failed to create K8s job.")
	}

	//print job details
	log.Println("Created K8s job successfully")
	pod_name := k8s.JobDetails(name, ns)
	log.Println("Pod name: ", pod_name)
}

func K8sCronJobs(name string, image string, cmd string, schedule string, ns string) {
	clientset, _ := k8s.K8sConfig()

	if ns == "" {
		ns = "default"
	}
	var jobbackoff int32 = 0
	cronjobs := clientset.BatchV1beta1().CronJobs(ns)

	cronjobSpec := &v1beta1.CronJob{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: v1beta1.CronJobSpec{
			Schedule:          schedule,
			ConcurrencyPolicy: v1beta1.ForbidConcurrent,
			JobTemplate: v1beta1.JobTemplateSpec{
				Spec: batchv1.JobSpec{
					Template: v1.PodTemplateSpec{
						Spec: v1.PodSpec{
							Containers: []v1.Container{
								{
									Name:    name,
									Image:   image,
									Command: strings.Split(cmd, " "),
								},
							},
							RestartPolicy: v1.RestartPolicyNever,
						},
					},
					BackoffLimit: &jobbackoff,
				},
			},
		},
	}
	_, err := cronjobs.Create(context.TODO(), cronjobSpec, metav1.CreateOptions{})
	if err != nil {
		log.Println(err)
		log.Println("Failed to create a K8s cronjob")
	}

}
