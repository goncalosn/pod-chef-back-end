package deployments

import (
	"net/http"
	httpError "pod-chef-back-end/pkg/errors"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func (h *HTTPHandler) CreateDefaultDeployment(c echo.Context) error {
	name := c.FormValue("name")
	replicas := c.FormValue("replicas")
	image := c.FormValue("image")

	if name == "" || replicas == "" || image == "" {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	replicasI64, err := strconv.ParseInt(replicas, 10, 32)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, "Internal error")
	}
	replicasI32 := int32(replicasI64)

	response, err := h.DeploymentServices.CreateDefaultDeployment(name, &replicasI32, image)

	if err != nil {
		kubernetesError := err.(*httpError.Error)
		return c.JSON(kubernetesError.Code, kubernetesError)
	}

	return c.JSONPretty(http.StatusCreated, response, " ")
}

func (h *HTTPHandler) CreateFileDeployment(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		kubernetesError := err.(*httpError.Error)
		return c.JSON(kubernetesError.Code, kubernetesError)
	}

	response, err := h.DeploymentServices.CreateFileDeployment(file)

	if err != nil {
		kubernetesError := err.(*httpError.Error)
		return c.JSON(kubernetesError.Code, kubernetesError)
	}

	return c.JSONPretty(http.StatusCreated, response, " ")
}
