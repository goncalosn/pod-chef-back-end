package deployments

import (
	"context"
	"net/http"

	httpError "pod-chef-back-end/pkg/errors"

	"github.com/labstack/gommon/log"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Deploy struct {
	Name string
}

func (serviceHandler *KubernetesClient) CreateDefaultDeployment(name string, replicas *int32, image string) (interface{}, error) {
	deploymentsClient := serviceHandler.Clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

	deploymentsList, err := deploymentsClient.List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		//service error
		log.Error(err)
		return nil, &httpError.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	for _, dep := range deploymentsList.Items {
		if dep.Name == name {
			// TODO: Throw new error, "deploy allready exists"
			return &httpError.Error{Code: http.StatusInternalServerError, Message: "that deploy allready exists, change name?", Err: nil}, nil
		}
	}

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "demo",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "demo",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "web",
							Image: image,
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	// result is the full deployment created
	_, err = deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		//service error
		log.Error(err)
		return nil, &httpError.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	return nil, nil
}

func (serviceHandler *KubernetesClient) CreateFileDeployment(dep *appsv1.Deployment) (interface{}, error) {
	deploymentsClient := serviceHandler.Clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

	_, err := deploymentsClient.Create(context.TODO(), dep, metav1.CreateOptions{})
	if err != nil {
		//service error
		log.Error(err)
		return nil, &httpError.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	return nil, nil
}

// True means there is a deploy with the same name
func (serviceHandler *KubernetesClient) CheckRepeatedDeployName(name string, namespace string) (bool, error) {
	deploymentsClient := serviceHandler.Clientset.AppsV1().Deployments(namespace)
	deploymentsList, err := deploymentsClient.List(context.TODO(), metav1.ListOptions{})

	// deploymentsList, err := serviceHandler.Clientset.AppsV1().Deployments().List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		//service error
		log.Error(err)
		return true, &httpError.Error{Err: err, Code: http.StatusInternalServerError, Message: "internal error"}
	}

	for _, dep := range deploymentsList.Items {
		if dep.Name == name {
			// TODO: Throw new error, "deploy allready exists"
			return true, nil
		}
	}

	return false, nil
}
