package pods

import (
	"net/http"
	ports "pod-chef-back-end/internal/core/ports"
	httpError "pod-chef-back-end/pkg/errors"

	"github.com/labstack/echo/v4"
)

type HTTPHandler struct {
	PodServices ports.PodServices
}

func NewHTTPHandler(podService ports.PodServices) *HTTPHandler {
	return &HTTPHandler{
		PodServices: podService,
	}
}

func (h *HTTPHandler) GetPodsByNodeAndNamespace(c echo.Context) error {
	namespace := c.QueryParam("namespace")
	node := c.QueryParam("node")

	if namespace == "" || node == "" {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	response, err := h.PodServices.GetPodsByNodeAndNamespace(node, namespace)

	if err != nil {
		kubernetesError := err.(*httpError.Error)
		return c.JSON(kubernetesError.Code, kubernetesError)
	}

	return c.JSONPretty(http.StatusOK, response, " ")
}
