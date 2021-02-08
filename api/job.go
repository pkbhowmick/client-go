package api

import (
	"context"
	"fmt"

	batchv1 "k8s.io/api/batch/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateJob() {
	fmt.Println("Creating Job . . .")
	clientset := CreateClientSet()
	jobClient := clientset.BatchV1().Jobs(apiv1.NamespaceDefault)
	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: "job-test",
		},
		Spec: batchv1.JobSpec{
			Template: apiv1.PodTemplateSpec{
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:    "go-version",
							Image:   "golang:alpine",
							Command: []string{"go", "version"},
						},
					},
					RestartPolicy: apiv1.RestartPolicyNever,
				},
			},
			BackoffLimit: int32Ptr(2),
		},
	}
	result, err := jobClient.Create(context.TODO(), job, metav1.CreateOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Created Job %q\n", result.GetObjectMeta().GetName())
}
