package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"

	kubernetesHandlers "pod-chef-back-end/handlers/kubernetes"
	mongoHandlers "pod-chef-back-end/handlers/mongo"
	kubernetesServices "pod-chef-back-end/internal/core/services/kubernetes"
	mongoServices "pod-chef-back-end/internal/core/services/mongo"
	emailRepo "pod-chef-back-end/repositories/email"
	kubernetesRepo "pod-chef-back-end/repositories/kubernetes"
	mongoRepo "pod-chef-back-end/repositories/mongo"
)

func main() {
	//setup env file
	viper.SetEnvPrefix("api")
	viper.AutomaticEnv()

	e := echo.New()

	// Middleware
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		// AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAccessControlRequestHeaders},
		AllowHeaders: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	var isLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(viper.GetString("TOKEN_SECRET")),
	})

	//initalize email access configurations
	emailRepository := emailRepo.NewEmailRepository(viper.GetViper())
	//initalize mongo access configurations
	mongoRepository := mongoRepo.NewMongoRepository(viper.GetViper())
	//initalize kubernetes access configurations
	kubernetesRepository := kubernetesRepo.NewKubernetesProdClient()

	mongoService := mongoServices.NewMongoService(mongoRepository, emailRepository, kubernetesRepository)
	kubernetesService := kubernetesServices.NewKubernetesService(kubernetesRepository, mongoRepository)

	//initialize the kubernetes http handlers
	kubernetesHandler := kubernetesHandlers.NewHTTPHandler(kubernetesService)
	kubernetesHandlers.Handlers(e, kubernetesHandler, isLoggedIn)

	//initialize the mongo http handlers
	mongoHandler := mongoHandlers.NewHTTPHandler(mongoService, viper.GetViper())
	mongoHandlers.Handlers(e, mongoHandler, isLoggedIn)

	e.GET("/api", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	e.Logger.Fatal(e.Start(":1323"))

}
