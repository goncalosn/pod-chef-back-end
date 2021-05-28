package http

import (
	"net/http"
	httpError "pod-chef-back-end/pkg/errors"

	"github.com/labstack/echo/v4"
)

func (h *HTTPHandler) GetNode(c echo.Context) error {
	node := c.QueryParam("node")

	if node == "" {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	response, err := h.KubernetesServices.GetNode(node)

	if err != nil {
		kubernetesError := err.(*httpError.Error)
		return c.JSON(kubernetesError.Code, kubernetesError)
	}

	return c.JSONPretty(http.StatusOK, response, " ")
}

func (h *HTTPHandler) GetNodes(c echo.Context) error {
	response, err := h.KubernetesServices.GetNodes()

	if err != nil {
		kubernetesError := err.(*httpError.Error)
		return c.JSON(kubernetesError.Code, kubernetesError)
	}

	return c.JSONPretty(http.StatusOK, response, " ")
}
