package ingresses

import (
	"context"
	"net/http"

	httpError "pod-chef-back-end/pkg/errors"

	"github.com/labstack/gommon/log"
	apiv1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (ingressHandler *KubernetesClient) GetIngress(name string) (interface{}, error) {
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
