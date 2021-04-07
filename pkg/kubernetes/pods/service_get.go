package pods

import (
	"context"

	. "pod-chef-back-end/pkg/kubernetes"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type PodService struct {
	KubernetesClient *kubernetes.Clientset
}

func NewPodService(KubernetesClient *kubernetes.Clientset) *PodService {

	return &PodService{KubernetesClient: KubernetesClient}
}

func (serviceHandler *PodService) GetPodsByNodeAndNamespaceService(node string, namespace string) ([]Pod, error) {

	pods, err := serviceHandler.KubernetesClient.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		return nil, err
	}

	var response []Pod

	for _, element := range pods.Items {
		if node == element.Spec.NodeName {
			response = append(response, Pod{
				State:        element.Status.Phase,
				RestartCount: element.Status.ContainerStatuses[len(element.Status.ContainerStatuses)-1].RestartCount,
				Name:         element.GetObjectMeta().GetName()})
		}
	}

	return response, err
}
