package api

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net"
	"sync"
	"time"

	echopprof "github.com/hiko1129/echo-pprof"
	"github.com/labstack/echo/v4"
	"github.com/yogs696/skilltest/api/helper"
	"github.com/yogs696/skilltest/api/route"
	"github.com/yogs696/skilltest/config"
)

// Instance defines web framework API instance
type Instance struct {
	Echo                     *echo.Echo
	RouteGroup               map[string]*echo.Group
	RouteGroupWithMiddleware map[string]*echo.Group
	Version                  string

	mu sync.Mutex
}

// New creates new API server instance and start the server immediately
func New() *Instance {
	// Create echo instance
	e := echo.New()

	// Custom any config
	e.HideBanner = true
	e.HidePort = true
	e.HTTPErrorHandler = helper.DefaultJsonErrorHandlerConfig.JsonErrorHandler
	e.Validator = helper.NewValidator()

	// Register route group
	rg := route.RegisterGroup(e)
	rgWithMiddleware := route.RegisterGroupWithMiddleware(e)

	// Register the middleware
	echopprof.Wrap(e)

	return &Instance{
		Echo:                     e,
		RouteGroup:               rg,
		RouteGroupWithMiddleware: rgWithMiddleware,
		Version:                  echo.Version,
	}
}

// Start start up the REST API Server
func (i *Instance) Start(nl net.Listener, address string) error {
	if i != nil {
		// Lock the mutex
		i.mu.Lock()
		defer i.mu.Unlock()

		// Set network listener
		i.Echo.Listener = nl

		return i.Echo.Start(address)
	}

	return errors.New("api server intance already available")
}

// Stale to shutdown the REST API Server
func (i *Instance) Stale() {
	if i != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := i.Echo.Shutdown(ctx); err != nil {
			log.Printf("[API] - Failed to shutting down the API server: %v \n", err.Error())
		} else {
			log.Println("[API] - Server has been shut down")
		}
	}
}

// SaveRouteList will save all registedred route list into json file in the user home directory
func (i *Instance) SaveRouteList() error {
	fname := config.Of.App.ResolveFilePathInWorkDir("route-list.json")

	// Get route list
	data, err := json.Marshal(i.Echo.Routes())
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(fname, data, 0644); err != nil {
		return err
	}

	return nil
}
