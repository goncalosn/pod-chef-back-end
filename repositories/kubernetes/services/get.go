package services

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//Get all services
func (serviceHandler *KubernetesClient) GetServicesByNamespace(namespace string) (interface{}, error) {
	//struct with the needed values from the services
	type KubernetesService struct {
		Name      string
		Kind      v1.ServiceType
		CreatedAt metav1.Time
	}

	services, err := serviceHandler.Clientset.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
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
func (serviceHandler *KubernetesClient) GetServiceByNameAndNamespace(name string, namespace string) (interface{}, error) {
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

	services, err := serviceHandler.Clientset.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
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
