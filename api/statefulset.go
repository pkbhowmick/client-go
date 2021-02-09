package api

import (
	"context"
	"fmt"

	"k8s.io/client-go/util/retry"

	appsv1 "k8s.io/api/apps/v1"

	apiv1 "k8s.io/api/core/v1"

	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var stsName, image string

func SetStsName(sts string) {
	stsName = sts
}

func SetImage(img string) {
	image = img
}

func CreateStatefulSet() {
	fmt.Println("Creating Service ...")
	clientset := CreateClientSet()
	svcClient := clientset.CoreV1().Services(apiv1.NamespaceDefault)
	service := &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "mongo-service",
			Labels: map[string]string{
				"app": "mongodb",
			},
		},
		Spec: apiv1.ServiceSpec{
			Ports: []apiv1.ServicePort{
				{
					Port: 80,
					Name: "web",
				},
			},
			ClusterIP: apiv1.ClusterIPNone,
			Selector: map[string]string{
				"app": "mongodb",
			},
		},
	}
	result, err := svcClient.Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Service %q created\n", result.GetObjectMeta().GetName())
	fmt.Println("Creating StatefulSet...")
	stsClient := clientset.AppsV1().StatefulSets(apiv1.NamespaceDefault)
	statefulSet := &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name: "mongo-sts",
		},
		Spec: appsv1.StatefulSetSpec{
			Replicas: int32Ptr(3),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "mongodb",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "mongodb",
					},
				},
				Spec: apiv1.PodSpec{
					TerminationGracePeriodSeconds: int64Ptr(10),
					Containers: []apiv1.Container{
						{
							Name:  "mongo",
							Image: "mongo",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "db-port",
									ContainerPort: 27017,
								},
							},
							VolumeMounts: []apiv1.VolumeMount{
								{
									Name:      "mongo-vol",
									MountPath: "/data/db",
								},
							},
						},
					},
				},
			},
			VolumeClaimTemplates: []apiv1.PersistentVolumeClaim{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "mongo-vol",
					},
					Spec: apiv1.PersistentVolumeClaimSpec{
						AccessModes:      []apiv1.PersistentVolumeAccessMode{apiv1.ReadWriteOnce},
						StorageClassName: strPtr("standard"),
						Resources: apiv1.ResourceRequirements{
							Requests: apiv1.ResourceList{
								apiv1.ResourceStorage: resource.MustParse("1Gi"),
							},
						},
					},
				},
			},
			ServiceName: "mongodb-service",
		},
	}

	resultSts, errSts := stsClient.Create(context.TODO(), statefulSet, metav1.CreateOptions{})
	if errSts != nil {
		fmt.Println(errSts)
		return
	}
	fmt.Printf("Created StatefulSet: %q\n", resultSts.GetObjectMeta().GetName())
}

func ListStatefulSet() {
	fmt.Println("***** Listing all StatefulSets *****")
	clientset := CreateClientSet()
	stsClient := clientset.AppsV1().StatefulSets(apiv1.NamespaceDefault)
	list, err := stsClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, item := range list.Items {
		fmt.Printf("---> %s (%d replicas)\n", item.Name, *item.Spec.Replicas)
	}
}

func DeleteStatefulSet() {
	fmt.Printf("Deleteing StatefulSet: %q\n", stsName)
	clientset := CreateClientSet()
	stsClient := clientset.AppsV1().StatefulSets(apiv1.NamespaceDefault)
	err := stsClient.Delete(context.TODO(), stsName, metav1.DeleteOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%q successfully deleted\n", stsName)
}

func UpdateStatefulSet() {
	fmt.Printf("Updating StatefulSet %q replicas to %d\n", stsName, replicas)
	clientset := CreateClientSet()
	stsClient := clientset.AppsV1().StatefulSets(apiv1.NamespaceDefault)
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		result, getErr := stsClient.Get(context.TODO(), stsName, metav1.GetOptions{})
		if getErr != nil {
			panic(fmt.Errorf("Failed to get latest version of StatefulSet: %v", getErr))
		}
		result.Spec.Replicas = int32Ptr(replicas)
		result.Spec.Template.Spec.Containers[0].Image = image
		_, updateErr := stsClient.Update(context.TODO(), result, metav1.UpdateOptions{})
		return updateErr
	})
	if retryErr != nil {
		panic(fmt.Errorf("Update failed: %v", retryErr))
	}
	fmt.Printf("Statefulset %q Successfully updated\n", stsName)
}
