package handlers

import (
	"net/http"

	. "pod-chef-back-end/pkg/api/k8"

	"k8s.io/client-go/kubernetes"

	"github.com/labstack/echo/v4"
)

func GetHandlers(e *echo.Echo, clientset *kubernetes.Clientset) {

	h := PodsHandler{Client: clientset}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/pods", h.GetPods)
	// e.GET("/foods", GetServices)
}

func PostHandlers(e *echo.Echo, clientset *kubernetes.Clientset) {

	dh := DeploymentHandler{Client: clientset}

	e.GET("/deploy/list", dh.ListDeployments)
	e.POST("/deploy/create", dh.CreateDeployment)
	e.POST("/deploy/delete/:depname", dh.DeleteDeployment)
}
