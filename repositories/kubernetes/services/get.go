package services

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//Get all services
func (serviceHandler *KubernetesClient) GetServicesByNamespace(namespace string) (interface{}, error) {
	response, err := serviceHandler.Clientset.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return response.Items, nil
}
