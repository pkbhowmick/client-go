package api

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	"k8s.io/apimachinery/pkg/api/resource"

	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"k8s.io/client-go/tools/clientcmd"

	"k8s.io/client-go/util/homedir"
)

func int32Ptr(i int32) *int32 {
	return &i
}

func int64Ptr(i int64) *int64 {
	return &i
}

func strPtr(s string) *string {
	return &s
}

func CreateClientSet() kubernetes.Interface {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		// fmt.Println(home)
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	//fmt.Println(*kubeconfig)
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	var clientset kubernetes.Interface
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	return clientset
}

func CreateDeployment() {
	clientset := CreateClientSet()
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "go-api-server",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "go-rest-api",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "go-rest-api",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "go-rest-api",
							Image: "pkbhowmick/go-rest-api:2.0.1",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 8080,
								},
							},
						},
					},
				},
			},
		},
	}

	fmt.Println("Creating deployment...")
	result, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created deployment %q\n", result.GetObjectMeta().GetName())
}

func GetDeployment() {
	fmt.Println("Listing all deployment objects ...")
	clientset := CreateClientSet()
	deploymentClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	list, err := deploymentClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, item := range list.Items {
		fmt.Printf("%s (%d replicas)\n", item.Name, *item.Spec.Replicas)
	}
}

func CreateStatefulSet() {
	fmt.Println("Create statefulset cmd is ok")
	clientset := CreateClientSet()
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
									MountPath: "/db/data",
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

	fmt.Println("Creating StatefulSet...")
	result, err := stsClient.Create(context.TODO(), statefulSet, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created StatefulSet: %q\n", result.GetObjectMeta().GetName())
}

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
		panic(err)
	}
	fmt.Printf("Created ReplicaSet %q\n", result.GetObjectMeta().GetName())
}

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
		panic(err)
	}
	fmt.Printf("Created Daemonset %q\n", result.GetObjectMeta().GetName())
}

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
		panic(err)
	}
	fmt.Printf("Created Job %q\n", result.GetObjectMeta().GetName())
}
