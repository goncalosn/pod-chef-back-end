package handlers

import (
	"net/http"

	. "pod-chef-back-end/pkg/api/k8"

	"k8s.io/client-go/kubernetes"

	"github.com/labstack/echo/v4"
)

func GetHandlers(e *echo.Echo, clientset *kubernetes.Clientset) {

	h := &K8Handler{Clientset: clientset}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/pods", h.GetPodsByNodeAndNamespace)
	// e.GET("/foods", GetServices)
}
