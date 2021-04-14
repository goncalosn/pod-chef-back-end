package deployments

import (
	"net/http"
	httpError "pod-chef-back-end/pkg/errors"

	"github.com/labstack/echo/v4"
)

func (h *HTTPHandler) GetDeployments(c echo.Context) error {

	response, err := h.DeploymentServices.GetDeployments()

	if err != nil {
		kubernetesError := err.(*httpError.Error)
		return c.JSON(kubernetesError.Code, kubernetesError)
	}

	return c.JSONPretty(http.StatusOK, response, " ")
}
