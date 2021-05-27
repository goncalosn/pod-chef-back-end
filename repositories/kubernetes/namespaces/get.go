package namespaces

import (
	"context"
	"net/http"

	httpError "pod-chef-back-end/pkg/errors"

	"github.com/labstack/gommon/log"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//Get all namespaces
func (serviceHandler *KubernetesClient) GetNamespaces() ([]string, error) {
	namespaces, err := serviceHandler.Clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		log.Error(err)
		return nil, &httpError.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	var response []string

	//filter each field from the kubernetes namespace struct
	for _, element := range namespaces.Items {
		response = append(response, element.Name)
	}

	return response, nil

}
