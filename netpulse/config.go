package netpulse

import (
	"net/http"

	"github.com/ursiform/logger"
)

type Config struct {
	group     string
	Handler   http.Handler `json:"-"`
	Interface string       `json:"interface,omitempty"`
	LogLevel  string       `json:"loglevel,omitempty"`
	Port      int          `json:"port,omitempty"`
	Service   string       `json:"service,omitempty"`
	Version   string       `json:"version,omitempty"`
	logLevel  int
}

func initConfig(config *Config) *Config {
	if config == nil {
		config = new(Config)
	}
	if config.group == "" {
		config.group = group
	}
	if config.LogLevel == "" {
		config.LogLevel = "silent"
	}
	if level, ok := logger.LogLevel[config.LogLevel]; !ok {
		format := "LogLevel=\"%s\" is invalid; using \"%s\" [%d]"
		logger.MustError(format, config.LogLevel, "debug", errLogLevel)
		config.LogLevel = "debug"
		config.logLevel = logger.Debug
	} else {
		config.logLevel = level
	}
	return config
}
