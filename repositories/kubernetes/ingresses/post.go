package ingresses

import (
	"context"
	"net/http"

	httpError "pod-chef-back-end/pkg/errors"

	uuid "github.com/satori/go.uuid"

	"github.com/labstack/gommon/log"
	networkingv1 "k8s.io/api/networking/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (ingressHandler *KubernetesClient) CreateIngress(namespace string, host string) (interface{}, error) {
	ingressClient := ingressHandler.Clientset.NetworkingV1().Ingresses(namespace)

	uuid := uuid.NewV4()

	var pathType string = "Prefix"

	ingress := &networkingv1.Ingress{
		ObjectMeta: v1.ObjectMeta{Name: uuid.String(), Namespace: namespace},
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
