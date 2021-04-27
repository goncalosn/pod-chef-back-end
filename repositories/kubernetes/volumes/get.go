package volumes

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//Get all services
func (serviceHandler *KubernetesClient) GetVolumes() (interface{}, error) {
	response, err := serviceHandler.Clientset.StorageV1().VolumeAttachments().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return response.Items, nil
}
