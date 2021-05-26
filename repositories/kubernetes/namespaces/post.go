package namespaces

import (
	"context"
	"net/http"

	httpError "pod-chef-back-end/pkg/errors"

	"github.com/labstack/gommon/log"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// //Get all namespaces
// func (serviceHandler *KubernetesClient) AddNamespaces(name string) (interface{}, error) {
// 	namespace, err := serviceHandler.Clientset.CoreV1().Namespaces().Create(context.TODO(), metav1.ObjectMeta{Name: name}, metav1.CreateOptions{})

// 	if err != nil {
// 		return nil, &httpError.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
// 	}

// 	return namespace, nil

// }

// True means there is a namespace with the same name
func (serviceHandler *KubernetesClient) CheckRepeatedNamespace(name string) (bool, error) {
	namespaces, err := serviceHandler.Clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		//service error
		log.Error(err)
		return true, &httpError.Error{Err: err, Code: http.StatusInternalServerError, Message: "internal error"}
	}

	for _, namepace := range namespaces.Items {
		if namepace.Name == name {
			//returns true if there is already a namespace with the same name
			return true, nil
		}
	}

	return false, nil
}
