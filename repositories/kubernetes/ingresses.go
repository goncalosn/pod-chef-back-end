package kubernetes

import (
	"context"
	"net/http"

	httpError "pod-chef-back-end/pkg/errors"

	"github.com/labstack/gommon/log"
	apiv1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (ingressHandler *KubernetesRepository) GetIngress(name string) (interface{}, error) {
	list, err := ingressHandler.Clientset.NetworkingV1().Ingresses(apiv1.NamespaceDefault).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Error(err)
		return nil, &httpError.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	var response *networkingv1.Ingress

	for _, ingress := range list.Items {
		if ingress.Name == name {
			response = &ingress
		}
	}

	return response, nil
}

func (ingressHandler *KubernetesRepository) CreateIngress(namespace string, name string, host string) (interface{}, error) {
	ingressClient := ingressHandler.Clientset.NetworkingV1().Ingresses(namespace)

	var pathType string = "Prefix"

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

	ingress, err := ingressClient.Create(context.TODO(), ingress, v1.CreateOptions{})

	if err != nil {
		log.Error(err)
		return nil, &httpError.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	return ingress, nil
}
