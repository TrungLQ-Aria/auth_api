package response

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func R200(c echo.Context, data interface{}) error {
	return response(c, http.StatusOK, data, "")
}

func R400(c echo.Context, data interface{}, error string) error {
	if error == "" {
		error = ErrBadRequest
	}

	return response(c, http.StatusBadRequest, data, error)
}

func R401(c echo.Context, data interface{}, error string) error {
	if error == "" {
		error = ErrUnauthorized
	}

	return response(c, http.StatusBadRequest, data, error)
}

func R404(c echo.Context, data interface{}, error string) error {
	if error == "" {
		error = ErrNotFound
	}

	return response(c, http.StatusBadRequest, data, error)
}

func response(c echo.Context, statusCode int, data interface{}, err string) error {
	return c.JSON(statusCode, echo.Map{"data": data, "message": err})
}
