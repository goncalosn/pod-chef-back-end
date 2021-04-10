package nodes

import (
	"context"
	"fmt"

	k8 "pod-chef-back-end/pkg/kubernetes"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type NodeService struct {
	KubernetesClient *kubernetes.Clientset
}

func NewNodeService(kubernetesClient *kubernetes.Clientset) *NodeService {

	return &NodeService{KubernetesClient: kubernetesClient}
}

func (serviceHandler *NodeService) GetNodeStatsService(name string) (k8.Node, error) {

	node, err := serviceHandler.KubernetesClient.CoreV1().Nodes().Get(context.TODO(), name, metav1.GetOptions{})

	fmt.Println(node)

	if err != nil {
		return k8.Node{}, err
	}

	var response k8.Node

	response = k8.Node{
		MemoryPressure: node.Status.Conditions[0].Type,
		DiskPressure:   node.Status.Conditions[1].Type,
		PIDPressure:    node.Status.Conditions[2].Type,
		Ready:          node.Status.Conditions[3].Type,
	}

	return response, err
}
