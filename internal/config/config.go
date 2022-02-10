// @description:
// @file: config.go
// @date: 2021/11/16

// Package config 读取配置文件。
package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"redis-viewer/internal/util"

	"github.com/spf13/viper"
)

// Config represents the main config for the application.
type Config struct {
	Mode string `mapstructure:"mode"`

	// client
	Addr string `mapstructure:"addr"`

	// sentinel
	MasterName    string   `mapstructure:"master_name"`
	SentinelAddrs []string `mapstructure:"sentinel_addrs"`

	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`

	Count int `mapstructure:"count"`
}

// LoadConfig loads a users config and creates the config if it does not exist.
func LoadConfig() {
	if runtime.GOOS != "windows" {
		homeDir, err := util.GetHomeDirectory()
		if err != nil {
			log.Fatal(err)
		}

		err = util.CreateDirectory(filepath.Join(homeDir, ".config", "redis-viewer"))
		if err != nil {
			log.Fatal(err)
		}

		viper.AddConfigPath("$HOME/.config/redis-viewer")
	} else {
		viper.AddConfigPath("$HOME")
	}

	viper.SetConfigName("redis-viewer")
	viper.SetConfigType("yml")

	viper.SetDefault("mode", "client")
	viper.SetDefault("count", 20)

	if err := viper.SafeWriteConfig(); err != nil {
		if os.IsNotExist(err) {
			err = viper.WriteConfig()
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal(err)
		}
	}
}

// GetConfig returns the users config.
func GetConfig() (config Config) {
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Error parsing config", err)
	}

	return
}
