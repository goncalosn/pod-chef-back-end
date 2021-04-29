package namespaces

import (
	"context"
	"net/http"

	httpError "pod-chef-back-end/pkg/errors"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//Get all namespaces
func (serviceHandler *KubernetesClient) GetNamespaces() ([]string, error) {
	namespaces, err := serviceHandler.Clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		return nil, &httpError.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	var response []string

	//filter each field from the kubernetes namespace struct
	for _, element := range namespaces.Items {
		//namespace system list
		systemNamespaces := []string{"kube-node-lease", "kube-public", "kube-system", "local-path-storage"}

		if !contains(systemNamespaces, element.Name) {
			//adds namespace name to the response
			response = append(response, element.Name)
		}

	}

	return response, nil

}

func contains(array []string, str string) bool {
	for _, element := range array {
		if element == str {
			return true
		}
	}
	return false
}
