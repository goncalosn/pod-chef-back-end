package kubernetes

import (
	"context"
	"net/http"

	pkg "pod-chef-back-end/pkg"

	"github.com/labstack/gommon/log"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//GetNodeByName method responsible for getting a node from the cluster
func (repo *KubernetesRepository) GetNodeByName(name string) (interface{}, error) {
	//data structure which will be returned
	type Node struct {
		Name       string
		Allocable  map[string]interface{}
		Capacity   map[string]interface{}
		Conditions []interface{}
		CreatedAt  metav1.Time
	}

	//list the node from the pool with the given name
	node, err := repo.Clientset.CoreV1().Nodes().Get(context.TODO(), name, metav1.GetOptions{})

	//verify if there is an error and then what kind of error it is
	if statusError, isStatus := err.(*errors.StatusError); isStatus && statusError.Status().Reason == metav1.StatusReasonNotFound {
		//node not found
		return nil, &pkg.Error{Err: err, Code: http.StatusNotFound, Message: "Node not found"}
	} else if err != nil {
		log.Error(err)
		return nil, &pkg.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	var conditions []interface{}

	//mapping through the node conditions
	for _, element := range node.Status.Conditions {
		type Condition struct {
			Type               v1.NodeConditionType
			Status             v1.ConditionStatus
			LastTransitionTime metav1.Time
		}

		//filter fields of conditions
		condition := &Condition{
			Type:               element.Type,
			Status:             element.Status,
			LastTransitionTime: element.LastTransitionTime,
		}

		//adds filtered condition to array of conditions
		conditions = append(conditions, condition)
	}

	//adds the node to the response
	response := &Node{
		Name: node.Name,
		Allocable: map[string]interface{}{
			"Cpu":     node.Status.Allocatable.Cpu(),
			"Memory":  node.Status.Allocatable.Memory().AsApproximateFloat64(),
			"Storage": node.Status.Allocatable.StorageEphemeral().AsApproximateFloat64(),
			"Pods":    node.Status.Allocatable.Pods(),
		},
		Capacity: map[string]interface{}{
			"Cpu":     node.Status.Capacity.Cpu(),
			"Memory":  node.Status.Capacity.Memory().AsApproximateFloat64(),
			"Storage": node.Status.Capacity.StorageEphemeral().AsApproximateFloat64(),
			"Pods":    node.Status.Capacity.Pods(),
		},
		Conditions: conditions,
		CreatedAt:  node.GetCreationTimestamp(),
	}

	return response, nil
}

//GetNodes method responsible for getting nodes from the cluster
func (repo *KubernetesRepository) GetNodes() (interface{}, error) {
	//struct with the needed values from the nodes
	type Node struct {
		Name      string
		Roles     []string
		CreatedAt metav1.Time
	}

	//list all nodes from the cluster
	nodes, err := repo.Clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})

	//verify if there is an error and then what kind of error it is
	if statusError, isStatus := err.(*errors.StatusError); isStatus && statusError.Status().Reason == metav1.StatusReasonNotFound {
		//no nodes found
		log.Error(err)
		return nil, &pkg.Error{Err: err, Code: http.StatusNotFound, Message: "No nodes found"}
	} else if err != nil {
		log.Error(err)
		return nil, &pkg.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
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
