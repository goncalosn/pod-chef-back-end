package handlers

import (
	"net/http"

	. "pod-chef-back-end/pkg/api/k8"

	"k8s.io/client-go/kubernetes"

	"github.com/labstack/echo/v4"
)

func GetHandlers(e *echo.Echo, clientset *kubernetes.Clientset) {

	h := PodsHandler{Client: clientset}
	dh := DeploymentHandler{Client: clientset}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/pods", h.GetPods)
	e.GET("/deploy/list", dh.ListDeployments)
}

func PostHandlers(e *echo.Echo, clientset *kubernetes.Clientset) {
	dh := DeploymentHandler{Client: clientset}

	e.POST("/deploy/default-create", dh.CreateDefaultDeployment)
	e.POST("/deploy/file-create", dh.CreateFileDeployment)
}

func DeleteHandlers(e *echo.Echo, clientset *kubernetes.Clientset) {
	dh := DeploymentHandler{Client: clientset}

	e.DELETE("/deploy/delete/:depname", dh.DeleteDeployment)
}
