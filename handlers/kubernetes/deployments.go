package kubernetes

import (
	"net/http"
	pkg "pod-chef-back-end/pkg"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

//createDeployment create a deployment with the given image and replicas
func (h *HTTPHandler) createDeployment(c echo.Context) error {
	//geting form data
	replicas := c.FormValue("replicas")
	image := c.FormValue("image")

	//checking data for empty values
	if replicas == "" || image == "" {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	//get the token's claims
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)
	role := claims["role"].(string)

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
	response, err := h.kubernetesServices.CreateDeployment(email, role, &replicasI32, image)
	if err != nil {
		//type assertion of custom Error to default error
		kubernetesError := err.(*pkg.Error)

		//return the error sent by the service
		return c.JSON(kubernetesError.Code, kubernetesError)
	}

	return c.JSONPretty(http.StatusCreated, response, " ")
}

//getDeploymentsByUser get all deployments inside a namespace
func (h *HTTPHandler) getDeploymentsByUser(c echo.Context) error {
	//get the token's claims
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	//call driver adapter responsible for getting the deployments from the mongo database
	response, err := h.kubernetesServices.GetDeploymentsByUser(email)

	if err != nil {
		//type assertion of custom Error to default error
		mongoError := err.(*pkg.Error)

		//return the error sent by the service
		return c.JSON(mongoError.Code, mongoError)
	}

	return c.JSONPretty(http.StatusOK, response, " ")
}

//getDeploymentsByUserAndName getting a deployment
func (h *HTTPHandler) getDeploymentByUserAndName(c echo.Context) error {
	//geting query data
	name := c.QueryParam("name")

	//checking data for empty values
	if name == "" {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	//get the token's claims
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	//call driver adapter responsible for getting a deployment from the mongo database
	response, err := h.kubernetesServices.GetDeploymentByUserAndName(email, name)

	if err != nil {
		//type assertion of custom Error to default error
		mongoError := err.(*pkg.Error)

		//return the error sent by the service
		return c.JSON(mongoError.Code, mongoError)
	}

	return c.JSONPretty(http.StatusOK, response, " ")
}

//deleteDeploymentByUserAndName delete a deployment by name
func (h *HTTPHandler) deleteDeploymentByUserAndName(c echo.Context) error {
	//geting query data
	name := c.QueryParam("name")

	//checking data for empty values
	if name == "" {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	//get the token's claims
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	//call driver adapter responsible for getting a deployment from the mongo database
	response, err := h.kubernetesServices.DeleteDeploymentByUserAndUUID(email, name)

	if err != nil {
		//type assertion of custom Error to default error
		mongoError := err.(*pkg.Error)

		//return the error sent by the service
		return c.JSON(mongoError.Code, mongoError)
	}

	return c.JSONPretty(http.StatusOK, response, " ")
}
