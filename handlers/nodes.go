package http

import (
	"net/http"
	httpError "pod-chef-back-end/pkg/errors"

	"github.com/labstack/echo/v4"
)

//GetNodeByName get node by name from the kubenretes cluster
func (h *HTTPHandler) GetNodeByName(c echo.Context) error {
	//geting query data
	node := c.QueryParam("node")

	//checking data for empty values
	if node == "" {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	//call driver adapter responsible for getting the node from the kubernetes cluster
	response, err := h.kubernetesServices.GetNodeByName(node)

	if err != nil {
		//type assertion of custom Error to default error
		kubernetesError := err.(*httpError.Error)

		//return the error sent by the service
		return c.JSON(kubernetesError.Code, kubernetesError)
	}

	return c.JSONPretty(http.StatusOK, response, " ")
}

//GetNodes get nodes from  kubernetes cluster
func (h *HTTPHandler) GetNodes(c echo.Context) error {
	//call driver adapter responsible for getting the nodes from the kubernetes cluster
	response, err := h.kubernetesServices.GetNodes()

	if err != nil {
		//type assertion of custom Error to default error
		kubernetesError := err.(*httpError.Error)

		//return the error sent by the service
		return c.JSON(kubernetesError.Code, kubernetesError)
	}

	return c.JSONPretty(http.StatusOK, response, " ")
}
