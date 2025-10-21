package middleware

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/yogs696/skilltest/config"
	"github.com/yogs696/skilltest/internal/entity/std"
)

func JWTAuth() []echo.MiddlewareFunc {
	return []echo.MiddlewareFunc{
		JWTValidateToken(),
		// JWTVerifySecretKey,
	}
}

func JWTValidateToken() func(next echo.HandlerFunc) echo.HandlerFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningMethod: "RS256",
		SigningKey:    config.Of.App.GetPublicKey(),
		ErrorHandler: func(c echo.Context, err error) error {
			apiResp := std.APIResponseError(std.StatusForbidden, errors.New("unauthorized"))
			return c.JSON(int(apiResp.StatusCode), apiResp.Body)
		},
	})
}

func JWTVerifySecretKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		usr := c.Get("user").(*jwt.Token)
		claims := usr.Claims.(jwt.MapClaims)
		seck := claims["sec"].(string)

		if config.Of.App.GetSecretKey() != seck {
			err := errors.New("your token has been expired")
			apiResp := std.APIResponseError(std.StatusForbidden, err)
			return c.JSON(int(apiResp.StatusCode), apiResp.Body)
		}

		return next(c)
	}
}
