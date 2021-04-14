package nodes

import (
	"net/http"
	ports "pod-chef-back-end/internal/core/ports"
	httpError "pod-chef-back-end/pkg/errors"

	"github.com/labstack/echo/v4"
)

type HTTPHandler struct {
	NodeServices ports.NodeServices
}

func NewHTTPHandler(nodeService ports.NodeServices) *HTTPHandler {
	return &HTTPHandler{
		NodeServices: nodeService,
	}
}

func (h *HTTPHandler) GetNode(c echo.Context) error {
	node := c.QueryParam("node")

	if node == "" {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	response, err := h.NodeServices.GetNode(node)

	if err != nil {
		kubernetesError := err.(*httpError.Error)
		return c.JSON(kubernetesError.Code, kubernetesError)
	}

	return c.JSONPretty(http.StatusOK, response, " ")
}

func (h *HTTPHandler) GetNodes(c echo.Context) error {
	response, err := h.NodeServices.GetNodes()

	if err != nil {
		kubernetesError := err.(*httpError.Error)
		return c.JSON(kubernetesError.Code, kubernetesError)
	}

	return c.JSONPretty(http.StatusOK, response, " ")
}
