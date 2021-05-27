package namespaces

import (
	"context"
	"net/http"

	httpError "pod-chef-back-end/pkg/errors"

	"github.com/labstack/gommon/log"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//Get all namespaces
func (serviceHandler *KubernetesClient) DeleteNamespace(name string) (interface{}, error) {
	err := serviceHandler.Clientset.CoreV1().Namespaces().Delete(context.TODO(), name, metav1.DeleteOptions{})

	if err != nil {
		log.Error(err)
		return nil, &httpError.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	return nil, nil

}
