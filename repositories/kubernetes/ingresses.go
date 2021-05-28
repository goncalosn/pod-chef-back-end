package kubernetes

import (
	"context"
	"net/http"

	httpError "pod-chef-back-end/pkg/errors"

	"github.com/labstack/gommon/log"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//GetIngressByNameAndNamespace method responsible for getting an ingress by it's name and namespace
func (ingressHandler *KubernetesRepository) GetIngressByNameAndNamespace(name string, namespace string) (interface{}, error) {
	//call driven adapter responsible for getting a ingress from the kubernetes cluster
	response, err := ingressHandler.Clientset.NetworkingV1().Ingresses(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		//print the error stack
		log.Error(err)

		//return a custom error
		return nil, &httpError.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	return response, nil
}

//CreateIngress method responsible for creating an ingress by it's name and namespace and a host
func (ingressHandler *KubernetesRepository) CreateIngress(namespace string, name string, host string) (interface{}, error) {
	//call driven adapter responsible for dealing wtih ingresses
	ingressClient := ingressHandler.Clientset.NetworkingV1().Ingresses(namespace)

	//path type in which the ingress will be redirecting the traffic to the user
	var pathType string = "Prefix"

	//data structure that will be used to create the ingress
	ingress := &networkingv1.Ingress{
		ObjectMeta: v1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: networkingv1.IngressSpec{
			Rules: []networkingv1.IngressRule{
				{
					Host: host,
					IngressRuleValue: networkingv1.IngressRuleValue{
						HTTP: &networkingv1.HTTPIngressRuleValue{
							Paths: []networkingv1.HTTPIngressPath{
								{
									Path:     "/",
									PathType: (*networkingv1.PathType)(&pathType),
									Backend: networkingv1.IngressBackend{
										Service: &networkingv1.IngressServiceBackend{
											Name: "gateway",
											Port: networkingv1.ServiceBackendPort{
												Number: 8080,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	//call driven adapter responsible for creating a ingress inside the kubernetes cluster
	ingress, err := ingressClient.Create(context.TODO(), ingress, v1.CreateOptions{})

	if err != nil {
		//print the error stack
		log.Error(err)

		//return a custom error
		return nil, &httpError.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	return ingress, nil
}
