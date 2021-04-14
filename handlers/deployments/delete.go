package deployments

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *HTTPHandler) DeleteDeployment(c echo.Context) error {
	name := c.QueryParam("name")

	if name == "" {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	response, err := h.DeploymentServices.DeleteDeployment(name)

	if err != nil {
		return err
	}

	return c.JSONPretty(http.StatusOK, response, " ")
}
