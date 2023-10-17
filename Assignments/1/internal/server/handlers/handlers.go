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

func (h *Handler) UserSignUp() echo.HandlerFunc {
	type request struct {
		payloads.SignUpRequest
	}
	type response struct {
		payloads.GenericMessageResponse
	}

	return func(c echo.Context) error {
		var req request
		if err := c.Bind(&req); err != nil {
			return err
		}
		//res,err :=
		return nil
	}
}

func (h *Handler) CheckStatus() echo.HandlerFunc {
	type request struct {
		payloads.CheckStatusRequest
	}
	type response struct {
		payloads.GenericMessageResponse
	}

	return func(c echo.Context) error {
		var req request
		if err := c.Bind(&req); err != nil {
			return err
		}
		//res,err :=
		return nil
		//res, err :=
	}
}
