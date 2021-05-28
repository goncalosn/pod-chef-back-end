package http

import (
	"net/http"
	httpError "pod-chef-back-end/pkg/errors"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

//GetDeploymentsByUser get all deployments inside a namespace
func (h *HTTPHandler) GetDeploymentsByUser(c echo.Context) error {
	//call driver adapter responsible for getting the deployments from the kubernetes cluster
	response, err := h.kubernetesServices.GetDeploymentsByUser(c.Request().Header.Get("Token"))

	if err != nil {
		//type assertion of custom Error to default error
		kubernetesError := err.(*httpError.Error)

		//return the error sent by the service
		return c.JSON(kubernetesError.Code, kubernetesError)
	}

	return c.JSONPretty(http.StatusOK, response, " ")
}

//GetDeploymentsByUserAndName get all deployments inside a namespace
func (h *HTTPHandler) GetDeploymentsByUserAndName(c echo.Context) error {
	//geting query data
	name := c.QueryParam("name")

	//checking data for empty values
	if name == "" {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	//call driver adapter responsible for getting a deployment from the kubernetes cluster
	response, err := h.kubernetesServices.GetDeploymentByUserAndName(c.Request().Header.Get("Token"), name)

	if err != nil {
		//type assertion of custom Error to default error
		kubernetesError := err.(*httpError.Error)

		//return the error sent by the service
		return c.JSON(kubernetesError.Code, kubernetesError)
	}

	return c.JSONPretty(http.StatusOK, response, " ")
}

//CreateDeployment create a deployment with the given image and replicas
func (h *HTTPHandler) CreateDeployment(c echo.Context) error {
	//geting form data
	replicas := c.FormValue("replicas")
	image := c.FormValue("image")

	//checking data for empty values
	if replicas == "" || image == "" {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	//parsing replicas to int64
	replicasI64, err := strconv.ParseInt(replicas, 10, 32)
	if err != nil {
		log.Error(err)

		//error parsing data, as it contained data other than numbers
		return c.JSON(http.StatusBadRequest, "Invalid request")

	}

	//parse replicas to int32
	replicasI32 := int32(replicasI64)

	//call driver adapter responsible for creating the deployment in the kubernetes cluster
	response, err := h.kubernetesServices.CreateDeployment(c.Request().Header.Get("Token"), &replicasI32, image)
	if err != nil {
		//type assertion of custom Error to default error
		kubernetesError := err.(*httpError.Error)

		//return the error sent by the service
		return c.JSON(kubernetesError.Code, kubernetesError)
	}

	return c.JSONPretty(http.StatusCreated, response, " ")
}
