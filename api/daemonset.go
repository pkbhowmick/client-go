package api

import (
	"context"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"

	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateDaemonSet() {
	fmt.Println("Creating DaemonSet ...")
	clientset := CreateClientSet()
	daemonSetClient := clientset.AppsV1().DaemonSets(apiv1.NamespaceDefault)
	daemonSet := &appsv1.DaemonSet{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-daemonset",
		},
		Spec: appsv1.DaemonSetSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "elasticsearch",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "elasticsearch",
					},
				},
				Spec: apiv1.PodSpec{
					Tolerations: []apiv1.Toleration{
						{
							Key:    "node-role.kubernetes.io/master",
							Effect: "NoSchedule",
						},
					},
					Containers: []apiv1.Container{
						{
							Name:  "elasticsearch",
							Image: "elasticsearch:7.10.1",
							Resources: apiv1.ResourceRequirements{
								Limits: apiv1.ResourceList{
									apiv1.ResourceMemory: resource.MustParse("200Mi"),
								},
								Requests: apiv1.ResourceList{
									apiv1.ResourceMemory: resource.MustParse("200Mi"),
									apiv1.ResourceCPU:    resource.MustParse("100m"),
								},
							},
							VolumeMounts: []apiv1.VolumeMount{
								{
									Name:      "varlog",
									MountPath: "/var/log",
								},
								{
									Name:      "varlibdockercontainers",
									MountPath: "/var/lib/docker/containers",
								},
							},
						},
					},
					TerminationGracePeriodSeconds: int64Ptr(30),
					Volumes: []apiv1.Volume{
						{
							Name: "varlog",
							VolumeSource: apiv1.VolumeSource{
								HostPath: &apiv1.HostPathVolumeSource{
									Path: "/var/log",
								},
							},
						},
						{
							Name: "varlibdockercontainers",
							VolumeSource: apiv1.VolumeSource{
								HostPath: &apiv1.HostPathVolumeSource{
									Path: "/var/lib/docker/containers",
								},
							},
						},
					},
				},
			},
		},
	}
	result, err := daemonSetClient.Create(context.TODO(), daemonSet, metav1.CreateOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Created Daemonset %q\n", result.GetObjectMeta().GetName())
}
