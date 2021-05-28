package kubernetes

import (
	"context"
	"net/http"

	httpError "pod-chef-back-end/pkg/errors"

	"github.com/labstack/gommon/log"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//GetPodsByNodeAndNamespace method responsible for getting pod data from the namespace and node
func (repo *KubernetesRepository) GetPodsByNodeAndNamespace(node string, namespace string) (interface{}, error) {
	//struct with the needed values from the pods
	type KubernetesRepository struct {
		State        v1.PodPhase
		RestartCount int32
		Name         string
	}

	//call driven adapter responsible for getting pods from the kubernetes cluster
	pods, err := repo.Clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})

	//verify if there is an error and then what kind of error it is
	if statusError, isStatus := err.(*errors.StatusError); isStatus && statusError.Status().Reason == metav1.StatusReasonNotFound {
		//no pods found
		log.Error(err)
		return nil, &httpError.Error{Err: err, Code: http.StatusNotFound, Message: "No pods found"}
	} else if err != nil {
		log.Error(err)
		return nil, &httpError.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	var response []*KubernetesRepository

	//filter each field from the kubernetes pod struct
	for _, element := range pods.Items {
		//verify if this pod belongs in node with name(parameter)
		if node == element.Spec.NodeName {
			//adds new node to the response
			response = append(response, &KubernetesRepository{
				State:        element.Status.Phase,
				RestartCount: element.Status.ContainerStatuses[len(element.Status.ContainerStatuses)-1].RestartCount,
				Name:         element.GetObjectMeta().GetName()})
		}
	}

	return response, nil

}
