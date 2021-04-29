package deployments

import (
	"context"
	"net/http"

	httpError "pod-chef-back-end/pkg/errors"

	"github.com/labstack/gommon/log"
	v1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (serviceHandler *KubernetesClient) GetDeployments() (interface{}, error) {
	type Deploy struct {
		Name      string
		Namespace string
		Images    []apiv1.Container
		Status    v1.DeploymentStatus
	}

	list, err := serviceHandler.Clientset.AppsV1().Deployments(apiv1.NamespaceDefault).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		//service error
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
