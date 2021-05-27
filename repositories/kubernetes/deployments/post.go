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

func (serviceHandler *KubernetesClient) CreateDeployment(namespace string, name string, replicas *int32, image string) (interface{}, error) {
	deploymentsClient := serviceHandler.Clientset.AppsV1().Deployments(namespace)

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

	return Deploy{res.Name}, nil
}
