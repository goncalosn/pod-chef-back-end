package kubernetes

import (
	"net/http"
	pkg "pod-chef-back-end/pkg"

	"github.com/labstack/echo/v4"
)

//getNode get node by name from the kubenretes cluster
func (h *HTTPHandler) getNode(c echo.Context) error {
	//body structure
	type body struct {
		Node string `json:"node"`
	}

	data := new(body)

	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	//checking data for empty values
	if data.Node == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	//call driver adapter responsible for getting the node from the kubernetes cluster
	response, err := h.kubernetesServices.GetNodeByName(data.Node)

	if err != nil {
		//type assertion of custom Error to default error
		kubernetesError := err.(*pkg.Error)

		//return the error sent by the service
		return c.JSON(kubernetesError.Code, kubernetesError)
	}

	return c.JSONPretty(http.StatusOK, response, " ")
}

//getNodes get nodes from  kubernetes cluster
func (h *HTTPHandler) getNodes(c echo.Context) error {
	//call driver adapter responsible for getting the nodes from the kubernetes cluster
	response, err := h.kubernetesServices.GetNodes()

	if err != nil {
		//type assertion of custom Error to default error
		kubernetesError := err.(*pkg.Error)

		//return the error sent by the service
		return c.JSON(kubernetesError.Code, kubernetesError)
	}

	return c.JSONPretty(http.StatusOK, response, " ")
}
