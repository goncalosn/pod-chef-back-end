package services

import (
	"net/http"
	httpError "pod-chef-back-end/pkg/errors"

	"github.com/labstack/echo/v4"
)

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

func (h *HTTPHandler) GetServiceByNameAndNamespace(c echo.Context) error {
	name := c.QueryParam("name")
	namespace := c.QueryParam("namespace")

	if namespace == "" || name == "" {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	response, err := h.ServiceServices.GetServiceByNameAndNamespace(name, namespace)

	if err != nil {
		kubernetesError := err.(*httpError.Error)
		return c.JSON(kubernetesError.Code, kubernetesError)
	}

	return c.JSONPretty(http.StatusOK, response, " ")
}
