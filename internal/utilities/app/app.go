package app

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Config *viper.Viper
	Logger *logrus.Logger
}

func New(
	c *viper.Viper,
	l *logrus.Logger,
) *AppConfig {
	return &AppConfig{
		Config: c,
		Logger: l,
	}
}
