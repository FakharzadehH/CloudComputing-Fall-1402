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

		img1, err := c.FormFile("image1")
		if err != nil {
			return err
		}
		img2, err := c.FormFile("image2")
		if err != nil {
			return err
		}
		email := c.FormValue("email")
		lastName := c.FormValue("last_name")
		nationalID := c.FormValue("national_id")
		if err != nil {
			return err
		}

		var req request
		req.Email = email
		req.LastName = lastName
		req.NationalID = nationalID
		req.Image1 = img1
		req.Image2 = img2

		res, err := h.svcs.SubmitRequest(c.Request().Context(), req.SignUpRequest, c.RealIP())
		if err != nil {
			return err
		}

		return c.JSON(200, response{
			GenericMessageResponse: *res,
		})
	}
}

func (h *Handler) CheckStatus() echo.HandlerFunc {
	type request struct {
		payloads.CheckStatusRequest
	}
	type response struct {
		*payloads.GenericMessageResponse
	}

	return func(c echo.Context) error {
		var req request
		if err := c.Bind(&req); err != nil {
			return err
		}
		res, err := h.svcs.CheckStatus(c.Request().Context(), req.CheckStatusRequest, c.RealIP())
		if err != nil {
			return err
		}
		return c.JSON(200, response{
			GenericMessageResponse: res,
		})
	}
}
