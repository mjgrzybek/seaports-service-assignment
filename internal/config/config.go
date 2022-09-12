package config

import "github.com/spf13/viper"

const (
	GracefulShutdownTimeout = "gracefulShutdownTimeout"
	LogLevel                = "logLevel"
	PortsFile               = "portsFile"
)

func init() {
	viper.SetDefault(GracefulShutdownTimeout, "1s")
	viper.SetDefault(LogLevel, "info")
}
