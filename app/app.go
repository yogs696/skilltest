package app

import (
	"fmt"
	"log"
	"net"
	"os"
	"sort"

	"github.com/yogs696/skilltest/config"
)

var (
	// Application name.
	// Default from config
	AppName = config.Of.App.Name

	// Description of the application
	// Default from config
	AppDesc = config.Of.App.Desc

	// Application version.
	// Will set from makefile, otherwise the default is v0.0.0
	Version = "v.0.0.0"

	// Datetime of last application has been builded
	LastBuildAt string

	// To force set service into maintenance mode hard
	// REST API service will not available during hard maintenance.
	HardMaintenance = "false"
)

// AppArgs defines given argumen of any app services.
// If have new services, please add needed argument here,
// so the "Up()" function can be called on this code
type AppArgs struct {
	// Network listener for REST API server
	NL net.Listener

	// Address REST API server address
	Address string
}

// Map of function that will be called on Up() method based on their order.
// If have new services, just create new file and their method and register here
var orderUp = map[int]func(arg *AppArgs){
	1: dbUp,
	6: apiUp,
	8: doEntV1Register,
	9: apiSaveRouteList,
}

// Map of function that will be called on Down() method based on their order.
// If have new services, just create new file and their method and register here
var orderDown = map[int]func(){
	1: apiDown,
	4: dbDown,
	// 4: queueDown,
}

// Up will turn up all services.
func Up(args *AppArgs, manual ...int) {
	log.Println("++--------------------[UP...]--------------------++")

	// Do manual start up if have manual argument given
	// Otherwise do auto
	if len(manual) > 0 {
		for _, v := range manual {
			orderUp[v](args)
		}
	} else {
		keys := make([]int, 0, len(orderUp))
		for k := range orderUp {
			keys = append(keys, k)
		}
		sort.Ints(keys)

		for _, k := range keys {
			orderUp[k](args)
		}
	}
}

// Down will turn down all services.
func Down(manual ...int) {
	log.Println("++--------------------[DOWN.]--------------------++")

	// Do manual down if have manual argument given
	// Otherwise do auto
	if len(manual) > 0 {
		for _, v := range manual {
			orderDown[v]()
		}
	} else {
		printOutDown(fmt.Sprintf("Running cleanup task on PID: %v ... \n", os.Getpid()))

		keys := make([]int, 0, len(orderDown))
		for k := range orderDown {
			keys = append(keys, k)
		}
		sort.Ints(keys)

		for _, k := range keys {
			orderDown[k]()
		}
	}
}

// Helper function to print out message when servces up
func printOutUp(s string) {
	log.Printf("[UP...] - %v\n", s)
}

// Helper function to print out message when servces down
func printOutDown(s string) {
	log.Printf("[DOWN.] - %v\n", s)
}
