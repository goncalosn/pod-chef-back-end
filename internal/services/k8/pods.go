package k8

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type K8HandlerSrv struct {
	Clientset *kubernetes.Clientset
}

type Pod struct {
	State        v1.PodPhase `json:"state"`
	RestartCount int32       `json:"restartCount"`
	Name         string      `json:"name"`
}

func (serviceHandler *K8HandlerSrv) GetPodsByNodeAndNamespaceService(node string, namespace string) ([]Pod, error) {

	pods, err := serviceHandler.Clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		return nil, err
	}

	var response []Pod

	for _, element := range pods.Items {
		if node == element.Spec.NodeName {
			response = append(response, Pod{
				element.Status.Phase,
				element.Status.ContainerStatuses[len(element.Status.ContainerStatuses)-1].RestartCount,
				element.GetObjectMeta().GetName()})
		}
	}

	return response, err
}
