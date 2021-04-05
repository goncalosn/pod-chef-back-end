package k8

import (
	"net/http"

	. "pod-chef-back-end/internal/services/k8"

	"github.com/labstack/echo/v4"
	"k8s.io/client-go/kubernetes"
)

type K8Handler struct {
	Clientset *kubernetes.Clientset
}

//GetPodsByNodeAndNamespace - GET - returns all the pods from the namespace
func (h *K8Handler) GetPodsByNodeAndNamespace(c echo.Context) error {
	c.Logger().Info("GetPodsByNodeAndNamespace request")

	namespace := c.FormValue("namespace")
	node := c.FormValue("node")

	if namespace == "" || node == "" {
		return c.JSON(http.StatusBadRequest, "invalid form")
	}

	serviceHandler := &K8HandlerSrv{Clientset: h.Clientset}

	response, err := serviceHandler.GetPodsByNodeAndNamespaceService(node, namespace)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSONPretty(http.StatusOK, response, " ")
}
