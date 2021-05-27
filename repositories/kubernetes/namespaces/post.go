package namespaces

import (
	"context"
	"net/http"

	httpError "pod-chef-back-end/pkg/errors"

	"github.com/labstack/gommon/log"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//Get all namespaces
func (serviceHandler *KubernetesClient) CreateNamespace(name string) (interface{}, error) {
	namespace := &apiv1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}

	namespace, err := serviceHandler.Clientset.CoreV1().Namespaces().Create(context.TODO(), namespace, metav1.CreateOptions{})

	if err != nil {
		log.Error(err)
		return nil, &httpError.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	return namespace, nil

}
