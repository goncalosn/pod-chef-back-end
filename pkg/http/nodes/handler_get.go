package nodes

import (
	"net/http"
	. "pod-chef-back-end/pkg/domain/nodes"
	httpError "pod-chef-back-end/pkg/errors"

	"github.com/labstack/echo/v4"
)

type NodeHandler struct {
	NodeInteractor NodeInteractor
}

func (h *NodeHandler) GetNode(c echo.Context) error {
	node := c.FormValue("node")

	if node == "" {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	response, err := h.NodeInteractor.GetNodeInteractor(node)

	if err != nil {
		kubernetesError := err.(httpError.KubernetesError)
		return c.JSON(kubernetesError.GetStatus(), kubernetesError)
	}

	return c.JSONPretty(http.StatusOK, response, " ")
}

func (h *NodeHandler) GetNodes(c echo.Context) error {
	response, err := h.NodeInteractor.GetNodesInteractor()

	if err != nil {
		kubernetesError := err.(httpError.KubernetesError)
		return c.JSON(kubernetesError.GetStatus(), kubernetesError)
	}

	return c.JSONPretty(http.StatusOK, response, " ")
}
