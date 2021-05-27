package deployments

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func (h *HTTPHandler) CreateDeployment(c echo.Context) error {
	token := c.FormValue("token")
	replicas := c.FormValue("replicas")
	image := c.FormValue("image")

	if replicas == "" || image == "" {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	replicasI64, err := strconv.ParseInt(replicas, 10, 32)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, "Internal error")

	}
	replicasI32 := int32(replicasI64)

	response, err := h.DeploymentServices.CreateDeployment(token, &replicasI32, image)

	if err != nil {
		// kubernetesError := err.(*httpError.Error)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSONPretty(http.StatusCreated, response, " ")
}
