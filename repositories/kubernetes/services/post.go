package services

import (
	"context"
	"net/http"
	httpError "pod-chef-back-end/pkg/errors"

	"github.com/labstack/gommon/log"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

//Get all services
func (serviceHandler *KubernetesClient) CreateClusterIPService(namespace string, name string) (interface{}, error) {

	service := &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			Labels: map[string]string{
				"run": "app",
			},
			Annotations: map[string]string{
				"haproxy.org/check":         "enabled",
				"haproxy.org/forwarded-for": "enabled",
				"haproxy.org/load-balance":  "roundrobin",
			},
		},
		Spec: apiv1.ServiceSpec{
			Selector: map[string]string{
				"run": "app",
			},
			Ports: []apiv1.ServicePort{
				{
					Name:       "port-1",
					Port:       80,
					Protocol:   "TCP",
					TargetPort: intstr.FromInt(8080),
				},
			},
		},
	}

	serviceClient, err := serviceHandler.Clientset.CoreV1().Services(namespace).Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		log.Error(err)
		return nil, &httpError.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	return serviceClient, nil
}
