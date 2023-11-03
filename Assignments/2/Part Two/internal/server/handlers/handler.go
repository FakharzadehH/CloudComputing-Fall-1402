package handlers

import (
	"github.com/FakharzadehH/CloudComputing-Fall-1402/internal/domain/payloads"
	"github.com/FakharzadehH/CloudComputing-Fall-1402/internal/repository"
	"github.com/FakharzadehH/CloudComputing-Fall-1402/service"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	repo *repository.Repository
	svcs *service.Service
}

func New(repo *repository.Repository, svcs *service.Service) *Handler {
	return &Handler{
		repo: repo,
		svcs: svcs,
	}
}

func (h *Handler) CheckWeather() echo.HandlerFunc {
	type request struct {
		payloads.CheckWeatherRequest
	}
	type response struct {
		*payloads.CheckWeatherResponse
	}
	return func(c echo.Context) error {
		var req request
		if err := c.Bind(&req); err != nil {
			return err
		}

		res, err := h.svcs.CheckWeather(c.Request().Context(), req.CheckWeatherRequest)
		if err != nil {
			return err
		}

		return c.JSON(200, response{
			CheckWeatherResponse: res,
		})
	}
}
