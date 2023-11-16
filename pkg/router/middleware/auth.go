package routermiddleware

import (
	"ema_sound_clone_api/config"
	"ema_sound_clone_api/internal/utils/auth"
	"ema_sound_clone_api/internal/utils/response"
	"github.com/labstack/echo/v4"
	"strconv"
	"strings"
)

func DevAPIKeyAuthentication(env config.Env) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			apikey := c.QueryParam("apikey")
			if env.DevApiKey != apikey {
				return response.R401(c, nil, "invalid api key")
			}

			return next(c)
		}
	}
}

func AdminAuthentication(env config.Env) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if bearerToken := c.Request().Header.Get("Authorization"); bearerToken != "" {
				if !strings.HasPrefix(strings.ToLower(bearerToken), "bearer") {
					return response.R401(c, nil, "")
				}

				token := strings.Split(bearerToken, " ")[1]

				claims, err := auth.UnSign(env.AdminJWTKey, token)
				if err != nil || claims.Valid() != nil {
					return response.R401(c, nil, "")
				}

				accountID, err := strconv.ParseInt(claims.Subject, 10, 32)
				if err != nil {
					return response.R401(c, nil, "")
				}

				c.Set("userId", accountID)

				return next(c)
			}

			return response.R401(c, nil, "")
		}
	}
}
