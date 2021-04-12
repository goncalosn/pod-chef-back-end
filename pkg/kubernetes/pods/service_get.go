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

type PodService struct {
	k8rnetesClient *kubernetes.Clientset
}

func NewPodService(k8rnetesClient *kubernetes.Clientset) *PodService {

	return &PodService{k8rnetesClient: k8rnetesClient}
}

//Get all pods in node and namespace
func (serviceHandler *PodService) GetPodsByNodeAndNamespaceService(node string, namespace string) (interface{}, error) {
	//struct with the needed values from the pods
	type Pod struct {
		State        v1.PodPhase
		RestartCount int32
		Name         string
	}

	//list all pods
	pods, err := serviceHandler.k8rnetesClient.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})

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

	var response []*Pod

	//filter each field from the kubernetes pod struct
	for _, element := range pods.Items {
		//verify if this pod belongs in node with name(parameter)
		if node == element.Spec.NodeName {
			//adds new node to the response
			response = append(response, &Pod{
				State:        element.Status.Phase,
				RestartCount: element.Status.ContainerStatuses[len(element.Status.ContainerStatuses)-1].RestartCount,
				Name:         element.GetObjectMeta().GetName()})
		}
	}

	return response, httpError.NewHTTPError(nil, http.StatusFound, "Pods found")

}
