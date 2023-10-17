package server

import "github.com/labstack/echo/v4"

func ErrorHandler() echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		if err != nil {
			if httpErr, ok := err.(*echo.HTTPError); ok {
				c.JSON(httpErr.Code, httpErr.Message)
				return
			}
		}
	}
}
