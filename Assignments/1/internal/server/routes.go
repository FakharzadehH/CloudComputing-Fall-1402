package server

import (
	"github.com/FakharzadehH/CloudComputing-Fall-1402/internal/server/handlers"
	"github.com/labstack/echo/v4"
)

func routes(e *echo.Echo) {
	api := e.Group("/api")
	api.POST("/submit-request", handlers.UserSignUp())
	api.GET("/check-request", handlers.CheckStatus())
}
