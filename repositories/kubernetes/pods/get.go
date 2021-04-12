package pods

import (
	"context"
	"net/http"

	httpError "pod-chef-back-end/pkg/errors"

	"github.com/labstack/gommon/log"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

//this service's dependencies
type KubernetesClient struct {
	Clientset *kubernetes.Clientset
}

//service in charge of dealing with GET requests and nodes
func New(clientset *kubernetes.Clientset) *KubernetesClient {
	return &KubernetesClient{
		Clientset: clientset,
	}
}

//Get all pods in node and namespace
func (serviceHandler *KubernetesClient) GetPodsByNodeAndNamespace(node string, namespace string) (interface{}, error) {
	//struct with the needed values from the pods
	type KubernetesClient struct {
		State        v1.PodPhase
		RestartCount int32
		Name         string
	}

	//list all pods
	pods, err := serviceHandler.Clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})

	//verify if there is an error and then what kind of error it is
	if statusError, isStatus := err.(*errors.StatusError); isStatus && statusError.Status().Reason == metav1.StatusReasonNotFound {
		//pods not found
		log.Error(err)
		return nil, httpError.NewHTTPError(err, http.StatusNotFound, "No pods found")
	} else if err != nil {
		//service error
		log.Error(err)
		return nil, httpError.NewHTTPError(err, http.StatusInternalServerError, "Internal error")
	}

	var response []*KubernetesClient

	//filter each field from the kubernetes pod struct
	for _, element := range pods.Items {
		//verify if this pod belongs in node with name(parameter)
		if node == element.Spec.NodeName {
			//adds new node to the response
			response = append(response, &KubernetesClient{
				State:        element.Status.Phase,
				RestartCount: element.Status.ContainerStatuses[len(element.Status.ContainerStatuses)-1].RestartCount,
				Name:         element.GetObjectMeta().GetName()})
		}
	}

	return response, httpError.NewHTTPError(nil, http.StatusFound, "KubernetesClients found")

}
