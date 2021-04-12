package nodes

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

//Get node data by name
func (serviceHandler *KubernetesClient) GetNode(name string) (interface{}, error) {
	//struct with the needed values from the node
	type Node struct {
		MemoryPressure v1.ConditionStatus
		DiskPressure   v1.ConditionStatus
		PIDPressure    v1.ConditionStatus
		Ready          v1.ConditionStatus
	}

	//list the node from the pool with the given name
	node, err := serviceHandler.Clientset.CoreV1().Nodes().Get(context.TODO(), name, metav1.GetOptions{})

	//verify if there is an error and then what kind of error it is
	if statusError, isStatus := err.(*errors.StatusError); isStatus && statusError.Status().Reason == metav1.StatusReasonNotFound {
		//node not found
		log.Error(err)
		return nil, httpError.NewHTTPError(err, http.StatusNotFound, "Node not found")
	} else if err != nil {
		//service error
		log.Error(err)
		return nil, httpError.NewHTTPError(err, http.StatusInternalServerError, "Internal error")
	}

	//adds the node to the response
	response := &Node{
		MemoryPressure: node.Status.Conditions[0].Status,
		DiskPressure:   node.Status.Conditions[1].Status,
		PIDPressure:    node.Status.Conditions[2].Status,
		Ready:          node.Status.Conditions[3].Status,
	}

	return response, nil
}

//Get all nodes in cluster
func (serviceHandler *KubernetesClient) GetNodes() (interface{}, error) {
	//struct with the needed values from the nodes
	type Node struct {
		Name      string
		Roles     []string
		CreatedAt metav1.Time
	}

	//list all nodes from the cluster
	nodes, err := serviceHandler.Clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})

	//verify if there is an error and then what kind of error it is
	if statusError, isStatus := err.(*errors.StatusError); isStatus && statusError.Status().Reason == metav1.StatusReasonNotFound {
		//no nodes found
		log.Error(err)
		return nil, httpError.NewHTTPError(err, http.StatusNotFound, "No nodes found")
	} else if err != nil {
		//service error
		log.Error(err)
		return nil, httpError.NewHTTPError(err, http.StatusInternalServerError, "Internal error")
	}

	var response []*Node

	//filter each field from the kubernetes node struct
	for _, element := range nodes.Items {
		var roles = []string{}

		//filter node roles
		if _, exists := element.GetLabels()["node-role.kubernetes.io/control-plane"]; exists {
			roles = append(roles, "control-plane")
			if _, exists := element.GetLabels()["node-role.kubernetes.io/master"]; exists {
				roles = append(roles, "master")
			}
		} else /* if _, exists := element.GetLabels()["node-role.kubernetes.io/compute"]; exists*/ {
			roles = append(roles, "slave")
		}

		//adds new node to the response
		response = append(response, &Node{
			Name:      element.Name,
			Roles:     roles,
			CreatedAt: element.GetCreationTimestamp(),
		})

	}

	return response, nil
}
