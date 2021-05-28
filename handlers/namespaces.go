package http

import (
	"net/http"
	httpError "pod-chef-back-end/pkg/errors"

	"github.com/labstack/echo/v4"
)

//GetNamespaces get all namespaces from the kubernetes cluster
func (h *HTTPHandler) GetNamespaces(c echo.Context) error {

	//call driver adapter responsible for getting the deployments from the kubernetes cluster
	response, err := h.kubernetesServices.GetNamespaces()

	if err != nil {
		//type assertion of custom Error to default error
		kubernetesError := err.(*httpError.Error)

		//return the error sent by the service
		return c.JSON(kubernetesError.Code, kubernetesError)
	}

	return c.JSONPretty(http.StatusOK, response, " ")
}
