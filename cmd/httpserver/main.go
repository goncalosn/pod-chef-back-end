package main

import (
	"net/http"

	"github.com/labstack/gommon/log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"

	emailHandlers "pod-chef-back-end/handlers/email"
	kubernetesHandlers "pod-chef-back-end/handlers/kubernetes"
	mongoHandlers "pod-chef-back-end/handlers/mongo"
	emailServices "pod-chef-back-end/internal/core/services/email"
	kubernetesServices "pod-chef-back-end/internal/core/services/kubernetes"
	mongoServices "pod-chef-back-end/internal/core/services/mongo"
	emailRepo "pod-chef-back-end/repositories/email"
	kubernetesRepo "pod-chef-back-end/repositories/kubernetes"
	mongoRepo "pod-chef-back-end/repositories/mongo"
)

func main() {
	//setup env file
	viper.SetConfigFile("../../.env")
	//read env file
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	var isLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(viper.Get("TOKEN_SECRET").(string)),
	})

	//initalize kubernetes access configurations
	kubernetesRepository := kubernetesRepo.NewKubernetesRepository()
	kubernetesService := kubernetesServices.NewKubernetesService(kubernetesRepository)

	//initalize mongo access configurations
	mongoRepository := mongoRepo.NewMongoRepository(viper.GetViper())
	mongoService := mongoServices.NewMongoService(mongoRepository)

	//initalize email access configurations
	emailRepository := emailRepo.NewEmailRepository(viper.GetViper())
	emailService := emailServices.NewEmailService(emailRepository, mongoRepository)

	//initialize the kubernetes http handlers
	kubernetesHandler := kubernetesHandlers.NewHTTPHandler(kubernetesService)
	kubernetesHandlers.Handlers(e, kubernetesHandler, isLoggedIn)

	//initialize the mongo http handlers
	mongoHandler := mongoHandlers.NewHTTPHandler(mongoService, viper.GetViper())
	mongoHandlers.Handlers(e, mongoHandler, isLoggedIn)

	//initialize the email http handlers
	emailHandler := emailHandlers.NewHTTPHandler(emailService, viper.GetViper())
	emailHandlers.Handlers(e, emailHandler, isLoggedIn)

	e.GET("/api", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	e.Logger.Fatal(e.Start(":1323"))

}
