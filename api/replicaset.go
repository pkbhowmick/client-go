package api

import (
	"context"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateReplicaSet() {
	fmt.Println("Creating replicaset...")
	clientset := CreateClientSet()
	replicaSetClient := clientset.AppsV1().ReplicaSets(apiv1.NamespaceDefault)
	replicaSet := &appsv1.ReplicaSet{
		ObjectMeta: metav1.ObjectMeta{
			Name: "replicaset-example",
		},
		Spec: appsv1.ReplicaSetSpec{
			Replicas: int32Ptr(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"tier": "frontend",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"tier": "frontend",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "php-redis",
							Image: "gcr.io/google_samples/gb-frontend:v3",
						},
					},
				},
			},
		},
	}
	result, err := replicaSetClient.Create(context.TODO(), replicaSet, metav1.CreateOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Created ReplicaSet %q\n", result.GetObjectMeta().GetName())
}
