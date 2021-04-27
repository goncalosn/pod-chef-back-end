package services

import (
	"net/http"
	ports "pod-chef-back-end/internal/core/ports"
	httpError "pod-chef-back-end/pkg/errors"

	"github.com/labstack/echo/v4"
)

type HTTPHandler struct {
	ServiceServices ports.ServiceServices
}

func NewHTTPHandler(serviceService ports.ServiceServices) *HTTPHandler {
	return &HTTPHandler{
		ServiceServices: serviceService,
	}
}

func (h *HTTPHandler) GetServicesByNamespace(c echo.Context) error {
	namespace := c.QueryParam("namespace")

	if namespace == "" {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	response, err := h.ServiceServices.GetServicesByNamespace(namespace)

	if err != nil {
		kubernetesError := err.(*httpError.Error)
		return c.JSON(kubernetesError.Code, kubernetesError)
	}

	return c.JSONPretty(http.StatusOK, response, " ")
}
