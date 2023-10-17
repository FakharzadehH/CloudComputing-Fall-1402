package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Start() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.HTTPErrorHandler = ErrorHandler()
	//TODO : Add routes and if middlewares
	e.Logger.Fatal(e.Start(":1323"))
}
