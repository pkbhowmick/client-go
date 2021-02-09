package api

import (
	"context"
	"fmt"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateSecret() {
	clientset := CreateClientSet()
	secretClient := clientset.CoreV1().Secrets(apiv1.NamespaceDefault)
	secret := &apiv1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name: "mongo-secret",
		},
		Type: apiv1.SecretTypeOpaque,
		Data: map[string][]byte{
			"username": []byte(`admin`),
			"password": []byte(`admin`),
		},
	}
	result, err := secretClient.Create(context.TODO(), secret, metav1.CreateOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Secret %q created\n", result.GetObjectMeta().GetName())
}

func DeleteSecret(args []string) {
	if len(args) == 0 {
		return
	}
	clientset := CreateClientSet()
	secretClient := clientset.CoreV1().Secrets(apiv1.NamespaceDefault)
	for _, secretName := range args {
		err := secretClient.Delete(context.TODO(), secretName, metav1.DeleteOptions{})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Secret %q deleted\n", secretName)
	}

}
