// Package conf unmarshal config from viper.
package conf

import (
	"log"

	"github.com/spf13/viper"
)

// Config represents the main config for the application.
type Config struct {
	Addrs    []string `mapstructure:"addrs"`
	DB       int      `mapstructure:"db"`
	Username string   `mapstructure:"username"`
	Password string   `mapstructure:"password"`

	// sentinel
	MasterName string `mapstructure:"master_name"`

	Limit int64 `mapstructure:"limit"` // default 20
}

// Get returns the users config.
func Get() (config Config) {
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Error parsing config", err)
	}

	return
}
