package http

import (
	"net/http"
	httpError "pod-chef-back-end/pkg/errors"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func (h *HTTPHandler) GetDeploymentsByNamespace(c echo.Context) error {
	response, err := h.KubernetesServices.GetDeployments()

	if err != nil {
		kubernetesError := err.(*httpError.Error)
		return c.JSON(kubernetesError.Code, kubernetesError)
	}

	return c.JSONPretty(http.StatusOK, response, " ")
}

func (h *HTTPHandler) CreateDeployment(c echo.Context) error {
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

	response, err := h.KubernetesServices.CreateDeployment(c.Request().Header.Get("Token"), &replicasI32, image)

	if err != nil {
		// kubernetesError := err.(*httpError.Error)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSONPretty(http.StatusCreated, response, " ")
}

func (h *HTTPHandler) DeleteDeployment(c echo.Context) error {
	name := c.QueryParam("name")

	if name == "" {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	response, err := h.KubernetesServices.DeleteDeployment(name)

	if err != nil {
		return err
	}

	return c.JSONPretty(http.StatusOK, response, " ")
}
