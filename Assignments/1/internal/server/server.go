package server

import (
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
	e.IPExtractor = echo.ExtractIPDirect()
	db, err := config.NewGORMConnection(config.GetConfig())
	if err != nil {
		return err
	}
	repo := repository.NewRepository(db)
	svcs := service.New(repo)
	handler := handlers.New(repo, svcs)
	routes(e, handler)
	return e.Start(":1323")
}
