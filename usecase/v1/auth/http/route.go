package uv1authroute

import (
	"github.com/labstack/echo/v4"
	"github.com/yogs696/skilltest/pkg/kemu"
	"github.com/yogs696/skilltest/usecase/v1/auth"
)

type domainService struct {
	s    auth.Service
	kemu *kemu.Mutex
}

func RegisterRoute(v1 *echo.Group, s auth.Service, k *kemu.Mutex) {
	// Setup domain service
	ds := &domainService{
		s:    s,
		kemu: k,
	}

	// Create root user group
	ug := v1.Group("/auth") // <- Route group (and also prefix) "user"

	ug.POST("/register", ds.Register) // <- user register
	ug.POST("/login", ds.Login)       // <- user login

}
