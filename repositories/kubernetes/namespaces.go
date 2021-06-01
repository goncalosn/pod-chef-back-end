package kubernetes

import (
	"context"
	"net/http"

	pkg "pod-chef-back-end/pkg"

	"github.com/labstack/gommon/log"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//CreateNamespace method responsible for creating a namespace
func (repo *KubernetesRepository) CreateNamespace(name string) (interface{}, error) {
	//data structure used to create the namespace
	namespace := &apiv1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}

	//call driven adapter responsible for getting a deployment from the kubernetes cluster
	namespace, err := repo.Clientset.CoreV1().Namespaces().Create(context.TODO(), namespace, metav1.CreateOptions{})

	if err != nil {
		//print the error stack
		log.Error(err)

		//return a custom error
		return nil, &pkg.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	return namespace, nil

}

//DeleteNamespace method responsible for deleting a namespace from the kubernetes cluster
func (repo *KubernetesRepository) DeleteNamespace(name string) (interface{}, error) {
	//call driven adapter responsible for getting a deployment from the kubernetes cluster
	err := repo.Clientset.CoreV1().Namespaces().Delete(context.TODO(), name, metav1.DeleteOptions{})

	if err != nil {
		//print the error stack
		log.Error(err)

		//return a custom error
		return nil, &pkg.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	return nil, nil

}
