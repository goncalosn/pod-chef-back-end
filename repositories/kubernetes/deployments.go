package kubernetes

import (
	"context"
	"net/http"

	httpError "pod-chef-back-end/pkg/errors"

	"github.com/labstack/gommon/log"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (repo *KubernetesRepository) GetDeployments() (interface{}, error) {
	type Deploy struct {
		Name      string
		Namespace string
		Images    []apiv1.Container
		Status    v1.DeploymentStatus
	}

	list, err := repo.Clientset.AppsV1().Deployments(apiv1.NamespaceDefault).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Error(err)
		return nil, &httpError.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	var response []Deploy

	for _, dep := range list.Items {
		response = append(response, Deploy{
			Name:      dep.Name,
			Namespace: dep.Namespace,
			Images:    dep.Spec.Template.Spec.Containers,
			Status:    dep.Status,
		})
	}

	return response, nil
}

func (repo *KubernetesRepository) CreateDeployment(namespace string, name string, replicas *int32, image string) (interface{}, error) {
	deploymentsClient := repo.Clientset.AppsV1().Deployments(namespace)

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			Labels: map[string]string{
				"run": "app",
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"run": "app",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"run": "app",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "web",
							Image: image,
							Ports: []apiv1.ContainerPort{
								{
									ContainerPort: 8080,
								},
							},
						},
					},
				},
			},
		},
	}

	// result is the full deployment created
	res, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		log.Error(err)
		return nil, &httpError.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	return res.Name, nil
}

func (repo *KubernetesRepository) DeleteDeployment(name string) (interface{}, error) {
	deletePolicy := metav1.DeletePropagationForeground
	deploymentsClient := repo.Clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

	if err := deploymentsClient.Delete(context.TODO(), name, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		log.Error(err)
		return nil, &httpError.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	return nil, nil
}
