package services

import (
	"context"
	"net/http"

	httpError "pod-chef-back-end/pkg/errors"

	"github.com/labstack/gommon/log"
	v1 "k8s.io/api/core/v1"
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

//TODO add namespace
func (serviceHandler *KubernetesClient) CreateService(service *v1.Service) (interface{}, error) {
	type Service struct {
		Name string
	}
	servicesClient := serviceHandler.Clientset.CoreV1().Services(v1.NamespaceDefault)

	res, err := servicesClient.Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		log.Error(err)
		return nil, &httpError.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	return Service{res.Name}, nil
}
