package config

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// The configuration that underlying all config
type config struct {
	// Base App configuration
	App app `json:"app"`

	// Database configuration
	Database database `json:"database"`

	// External API/Microservices configuration
	External external `json:"external"`
}

// Of is the config context that will be called by another package
var Of config

// Run viper setup and marshaling the config once at the runtime
func init() {
	viper.SetConfigName(".config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../")
	viper.AddConfigPath("/etc/skill-test/")
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Config file has been changed, re-load that")
		load()
	})

	load()
}

// Load and marshaling the config file
func load() {
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("[ConfigLoad][Read] - Err: %s", err.Error())
		panic(err)
	}

	if err := viper.Unmarshal(&Of); err != nil {
		log.Printf("[ConfigLoad][Unmarshal] - Err: %s", err.Error())
		panic(err)
	}
}
