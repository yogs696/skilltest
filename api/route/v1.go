package route

import (
	"github.com/labstack/echo/v4"
	"github.com/yogs696/skilltest/api/middleware"
)

func v1(e *echo.Echo) *echo.Group {
	return e.Group("/v1")
}

func v1WithMiddlerware(e *echo.Echo) *echo.Group {
	return e.Group("/v1", middleware.JWTAuth()...)
}
