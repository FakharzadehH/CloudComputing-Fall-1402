package server

import (
	"github.com/FakharzadehH/CloudComputing-Fall-1402/internal/config"
	"github.com/FakharzadehH/CloudComputing-Fall-1402/internal/logger"
	"github.com/FakharzadehH/CloudComputing-Fall-1402/internal/repository"
	"github.com/FakharzadehH/CloudComputing-Fall-1402/internal/server/handlers"
	"github.com/FakharzadehH/CloudComputing-Fall-1402/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Start() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.HTTPErrorHandler = ErrorHandler()
	db, err := config.NewGORMConnection(config.GetConfig())
	if err != nil {
		logger.Logger().Fatal("err when connecting to db", "err", err)
	}
	repo := repository.NewRepository(db)
	svcs := service.New(repo)
	handler := handlers.New(repo, svcs)
	routes(e, handler)
	//TODO : Add routes and if middlewares
	e.Logger.Fatal(e.Start(":1323"))
}
