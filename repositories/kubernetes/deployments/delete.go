package deployments

import (
	"context"
	"net/http"
	httpError "pod-chef-back-end/pkg/errors"

	"github.com/labstack/gommon/log"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (serviceHandler *KubernetesClient) DeleteDeployment(name string) (interface{}, error) {
	deletePolicy := metav1.DeletePropagationForeground
	deploymentsClient := serviceHandler.Clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

	if err := deploymentsClient.Delete(context.TODO(), name, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		log.Error(err)
		return nil, &httpError.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	return nil, nil
}
