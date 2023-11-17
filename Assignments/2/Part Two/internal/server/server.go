package server

import (
	"os"

	"github.com/FakharzadehH/CloudComputing-Fall-1402/internal/config"
	"github.com/FakharzadehH/CloudComputing-Fall-1402/internal/repository"
	"github.com/FakharzadehH/CloudComputing-Fall-1402/internal/server/handlers"
	"github.com/FakharzadehH/CloudComputing-Fall-1402/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Start() error {
	e := echo.New()
	e.Use(middleware.Logger())
	e.HTTPErrorHandler = ErrorHandler()
	rdb, err := config.NewRedisPool(config.GetConfig().Redis)
	if err != nil {
		return err
	}
	repo := repository.New(rdb)
	svcs := service.New(repo)
	h := handlers.New(repo, svcs)

	api := e.Group("/api")
	api.GET("/check-weather", h.CheckWeather())
	return e.Start(":" + os.Getenv("SERVER_PORT"))
}
