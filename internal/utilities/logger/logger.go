package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func New(c *viper.Viper) *logrus.Logger {
	l := logrus.New()

	l.SetFormatter(&logrus.JSONFormatter{})
	l.SetLevel(logrus.Level(c.GetInt32("log.level")))

	return l
}
