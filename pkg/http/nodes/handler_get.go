package nodes

import (
	"net/http"
	. "pod-chef-back-end/pkg/domain/nodes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type NodeHandler struct {
	NodeInteractor NodeInteractor
}

//GetNodeStatsService - GET - returns all the pods from the namespace
func (h *NodeHandler) GetNodeStatsService(c echo.Context) error {
	log.Info("GetNodeStatsService request")

	node := c.FormValue("node")

	if node == "" {
		return c.JSON(http.StatusBadRequest, "invalid form")
	}

	response, err := h.NodeInteractor.GetNodeStatsServiceInteractor(node)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSONPretty(http.StatusOK, response, " ")
}
