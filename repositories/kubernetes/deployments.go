package kubernetes

import (
	"context"
	"net/http"

	pkg "pod-chef-back-end/pkg"

	"github.com/labstack/gommon/log"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//GetDeploymentByNameAndNamespace method responsible for getting a deployment by it's namespace from the namespace and name
func (repo *KubernetesRepository) GetDeploymentByNameAndNamespace(name string, namespace string) (interface{}, error) {
	//data structure which will be returned
	type Deploy struct {
		Name      string
		Namespace string
		Images    []apiv1.Container
		Status    v1.DeploymentStatus
	}

	//call driven adapter responsible for getting a deployment from the kubernetes cluster
	response, err := repo.Clientset.AppsV1().Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		//print the error stack
		log.Error(err)

		//return a custom error
		return nil, &pkg.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	return response, nil
}

//CreateDeployment method responsible for creating a deployment from namespace, name, replicas and image
func (repo *KubernetesRepository) CreateDeployment(namespace string, name string, replicas *int32, image string, containerPort int32) (interface{}, error) {
	//call driven adapter responsible for getting a deployment from the kubernetes cluster
	deploymentsClient := repo.Clientset.AppsV1().Deployments(namespace)

	//structure of the deployment
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
									ContainerPort: containerPort,
								},
							},
						},
					},
				},
			},
		},
	}

	// create deployment
	res, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		//print the error stack
		log.Error(err)

		//return a custom error
		return nil, &pkg.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	return res.Name, nil
}
