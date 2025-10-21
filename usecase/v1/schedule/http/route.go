package uv1schedulehttp

import (
	"github.com/labstack/echo/v4"
	"github.com/yogs696/skilltest/pkg/kemu"
	"github.com/yogs696/skilltest/usecase/v1/schedule"
)

type domainService struct {
	s    schedule.Service
	kemu *kemu.Mutex
}

func RegisterRoute(v1 *echo.Group, s schedule.Service, k *kemu.Mutex) {
	// Setup domain service
	ds := &domainService{
		s:    s,
		kemu: k,
	}

	// Create route schedule group
	pg := v1.Group("/schedule")         // <- Route group (and also prefix) "schedule"
	pg.GET("/list", ds.List)            // <- create schedule
	pg.POST("/create", ds.Create)       // <- create schedule
	pg.PUT("/update/:id", ds.Update)    // <- create schedule
	pg.DELETE("/delete/:id", ds.Delete) // <- delete schedule

	pg.POST("/test", ds.Testing) // <- testing

}
