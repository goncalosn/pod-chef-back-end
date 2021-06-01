package kubernetes

import (
	"net/http"
	pkg "pod-chef-back-end/pkg"

	"github.com/labstack/echo/v4"
)

//getNamespaces get all namespaces from the kubernetes cluster
func (h *HTTPHandler) getNamespaces(c echo.Context) error {

	//call driver adapter responsible for getting the deployments from the kubernetes cluster
	response, err := h.kubernetesServices.GetNamespaces()

	if err != nil {
		//type assertion of custom Error to default error
		kubernetesError := err.(*pkg.Error)

		//return the error sent by the service
		return c.JSON(kubernetesError.Code, kubernetesError)
	}

	return c.JSONPretty(http.StatusOK, response, " ")
}
