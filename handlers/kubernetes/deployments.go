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
	//body structure
	type body struct {
		Replicas string `json:"replicas"`
		Image    string `json:"image"`
	}

	data := new(body)

	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	//checking data for empty values
	if data.Replicas == "" || data.Image == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	//get the token's claims
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)
	role := claims["role"].(string)

	//parsing replicas to int64
	replicasI64, err := strconv.ParseInt(data.Replicas, 10, 32)
	if err != nil {
		log.Error(err)

		//error parsing data, as it contained data other than numbers
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})

	}

	//parse replicas to int32
	replicasI32 := int32(replicasI64)

	//call driver adapter responsible for creating the deployment in the kubernetes cluster
	response, err := h.kubernetesServices.CreateDeployment(email, role, &replicasI32, data.Image)
	if err != nil {
		//type assertion of custom Error to default error
		kubernetesError := err.(*pkg.Error)

		//return the error sent by the service
		return c.JSON(kubernetesError.Code, kubernetesError)
	}

	return c.JSONPretty(http.StatusCreated, response, " ")
}

//getDeploymentsByUser get all deployments inside a namespace
func (h *HTTPHandler) getMyDeployments(c echo.Context) error {
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

//getDeploymentsByUser get all deployments inside a namespace
func (h *HTTPHandler) getDeploymentsByUser(c echo.Context) error {
	//body structure
	type body struct {
		Email string `json:"email"`
	}

	data := new(body)

	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	//checking data for empty values
	if data.Email == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	//call driver adapter responsible for getting the deployments from the mongo database
	response, err := h.kubernetesServices.GetDeploymentsByUser(data.Email)

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
	//body structure
	type body struct {
		Name string `json:"name"`
	}

	data := new(body)

	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	//checking data for empty values
	if data.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	//get the token's claims
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	//call driver adapter responsible for getting a deployment from the mongo database
	response, err := h.kubernetesServices.GetDeploymentByUserAndName(email, data.Name)

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
	//body structure
	type body struct {
		Name string `json:"name"`
	}

	data := new(body)

	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	//checking data for empty values
	if data.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	//get the token's claims
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	//call driver adapter responsible for getting a deployment from the mongo database
	response, err := h.kubernetesServices.DeleteDeploymentByUserAndUUID(email, data.Name)

	if err != nil {
		//type assertion of custom Error to default error
		mongoError := err.(*pkg.Error)

		//return the error sent by the service
		return c.JSON(mongoError.Code, mongoError)
	}

	return c.JSONPretty(http.StatusOK, response, " ")
}
