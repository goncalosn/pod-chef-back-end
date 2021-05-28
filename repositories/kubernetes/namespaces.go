package kubernetes

import (
	"context"
	"net/http"

	httpError "pod-chef-back-end/pkg/errors"

	"github.com/labstack/gommon/log"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (repo *KubernetesRepository) GetNamespaces() ([]string, error) {
	namespaces, err := repo.Clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})

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

func (repo *KubernetesRepository) CreateNamespace(name string) (interface{}, error) {
	namespace := &apiv1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}

	namespace, err := repo.Clientset.CoreV1().Namespaces().Create(context.TODO(), namespace, metav1.CreateOptions{})

	if err != nil {
		log.Error(err)
		return nil, &httpError.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	return namespace, nil

}

func (repo *KubernetesRepository) DeleteNamespace(name string) (interface{}, error) {
	err := repo.Clientset.CoreV1().Namespaces().Delete(context.TODO(), name, metav1.DeleteOptions{})

	if err != nil {
		log.Error(err)
		return nil, &httpError.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	return nil, nil

}
