package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func New() *viper.Viper {
	c := viper.New()

	c.SetConfigName("config")
	c.SetConfigType("json")
	c.AddConfigPath("env")
	err := c.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	return c
}
