package kubernetes

import (
	"context"
	"net/http"

	httpError "pod-chef-back-end/pkg/errors"

	"github.com/labstack/gommon/log"
	apiv1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

//Get all services
func (repo *KubernetesRepository) GetServicesByNamespace(namespace string) (interface{}, error) {
	//struct with the needed values from the services
	type KubernetesService struct {
		Name      string
		Kind      v1.ServiceType
		CreatedAt metav1.Time
	}

	services, err := repo.Clientset.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Error(err)
		return nil, &httpError.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	var response []*KubernetesService

	//filter each field from the kubernetes service struct
	for _, element := range services.Items {

		if namespace == element.GetNamespace() {
			//adds new service to the response
			response = append(response, &KubernetesService{
				Name:      element.GetName(),
				Kind:      element.Spec.Type,
				CreatedAt: element.GetCreationTimestamp(),
			})
		}
	}

	return response, nil
}

//Get service by namespace
func (repo *KubernetesRepository) GetServiceByNameAndNamespace(name string, namespace string) (interface{}, error) {
	//struct with the needed values from the services
	type KubernetesService struct {
		Name           string
		Namespace      string
		Kind           v1.ServiceType
		CreatedAt      metav1.Time
		ClusterIP      string
		LoadBalancerIP string
		Selectors      map[string]string
		Ports          []v1.ServicePort
	}

	services, err := repo.Clientset.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Error(err)
		return nil, &httpError.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	var response *KubernetesService

	//filter each field from the kubernetes service struct
	for _, element := range services.Items {

		if name == element.GetName() {
			//adds service to the response
			response = &KubernetesService{
				Name:           element.GetName(),
				Namespace:      element.GetNamespace(),
				Kind:           element.Spec.Type,
				CreatedAt:      element.GetCreationTimestamp(),
				ClusterIP:      element.Spec.ClusterIP,
				LoadBalancerIP: element.Spec.LoadBalancerIP,
				Selectors:      element.Spec.Selector,
				Ports:          element.Spec.Ports,
			}
		}
	}

	return response, nil
}

func (repo *KubernetesRepository) CreateClusterIPService(namespace string, name string) (interface{}, error) {

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

	serviceClient, err := repo.Clientset.CoreV1().Services(namespace).Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		log.Error(err)
		return nil, &httpError.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	return serviceClient, nil
}
