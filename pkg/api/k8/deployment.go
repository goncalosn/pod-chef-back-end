package k8

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type DeploymentHandler struct {
	Client *kubernetes.Clientset
	//TODO: add here deploymentsClient, needs to be initialized at first
}

func (h *DeploymentHandler) CreateDeployment(c echo.Context) error {
	c.Logger().Info("post deployment")

	deploymentsClient := h.Client.AppsV1().Deployments(apiv1.NamespaceDefault)

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "demo-deployment",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(2),
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
							Image: "nginx",
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
	result, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		return c.String(http.StatusInternalServerError, "error creating deployment")
	}

	return c.JSONPretty(http.StatusOK, result, " ")
}

func (h *DeploymentHandler) ListDeployments(c echo.Context) error {
	deps, err := h.Client.AppsV1().Deployments(apiv1.NamespaceDefault).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return c.String(http.StatusInternalServerError, "error showing deployments")
	}
	return c.JSONPretty(http.StatusOK, deps, " ")
}

func (h *DeploymentHandler) DeleteDeployment(c echo.Context) error {
	deletePolicy := metav1.DeletePropagationForeground
	deploymentsClient := h.Client.AppsV1().Deployments(apiv1.NamespaceDefault)

	//! Maybe some problem here?
	depName := c.Param("depname")

	if err := deploymentsClient.Delete(context.TODO(), depName, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		return c.String(http.StatusInternalServerError, "error deleting deployment")
	}

	return c.String(http.StatusOK, "deleted deployment")
}

func int32Ptr(i int32) *int32 { return &i }
