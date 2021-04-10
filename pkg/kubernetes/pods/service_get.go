package pods

import (
	"context"

	k8 "pod-chef-back-end/pkg/kubernetes"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type PodService struct {
	k8rnetesClient *kubernetes.Clientset
}

func NewPodService(k8rnetesClient *kubernetes.Clientset) *PodService {

	return &PodService{k8rnetesClient: k8rnetesClient}
}

func (serviceHandler *PodService) GetPodsByNodeAndNamespaceService(node string, namespace string) ([]k8.Pod, error) {

	pods, err := serviceHandler.k8rnetesClient.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		return nil, err
	}

	var response []k8.Pod

	for _, element := range pods.Items {
		if node == element.Spec.NodeName {
			response = append(response, k8.Pod{
				State:        element.Status.Phase,
				RestartCount: element.Status.ContainerStatuses[len(element.Status.ContainerStatuses)-1].RestartCount,
				Name:         element.GetObjectMeta().GetName()})
		}
	}

	return response, err
}
