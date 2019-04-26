package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

func envBool(name string, defaultValue bool) bool {
	var defaultString string

	if defaultValue {
		defaultString = "true"
	} else {
		defaultString = "false"
	}
	return lowerEnvDefault(name, defaultString) == "true"
}

func envDefault(name, defaultValue string) string {
	val := os.Getenv(name)
	if val == "" {
		return defaultValue
	}
	return val
}

func lowerEnvDefault(name, defaultValue string) string {
	return strings.ToLower(envDefault(name, defaultValue))
}

// DBConnectionString forms & returns the DBConnectionString
func DBConnectionString() string {
	return fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=%s",
		DBHost(),
		DBName(),
		DBUser(),
		DBPassword(),
		SSLMode(),
	)
}

// DBHost returns the POSTGRES_HOST if exists, otherwise localhost
func DBHost() string {
	return envDefault("POSTGRES_HOST", "localhost")
}

// DBName returns the POSTGRES_DB if exists, otherwise postgres
func DBName() string {
	return envDefault("POSTGRES_DB", "postgres")
}

// DBPassword returns the POSTGRES_PASSWORD if exists, otherwise ""
func DBPassword() string {
	return envDefault("POSTGRES_PASSWORD", "\"\"")
}

// DBUser returns the POSTGRES_USER if exists, otherwise postgres
func DBUser() string {
	return envDefault("POSTGRES_USER", "postgres")
}

// LogFormat returns a logrus.Formatter that can be used to configure the
// logger. Defaults to the logrus.TextFormatter.
func LogFormat() logrus.Formatter {
	format := lowerEnvDefault("LOG_FORMAT", "")
	switch format {
	case "json":
		return &logrus.JSONFormatter{}
	}
	return &logrus.TextFormatter{}
}

// LogLevel Returns the logrus log level to use. Defaults to Info
func LogLevel() logrus.Level {
	level := lowerEnvDefault("LOG_LEVEL", "")
	switch level {
	case "debug":
		return logrus.DebugLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	default:
	}
	return logrus.InfoLevel
}

// PostAnalyze returns whether or not to run an analyze after doing a full
// vacuum
func PostAnalyze() bool {
	return envBool("POST_ANALYZE", true)
}

// PreAnalyze returns whether or not to run an analzye to update stats before
// Running a full vacuum
func PreAnalyze() bool {
	return envBool("PRE_ANALYZE", true)
}

// SSLMode returns the SSL_MODE if exists, otherwise the default of require.
// See https://godoc.org/github.com/lib/pq#hdr-Connection_String_Parameters
func SSLMode() string {
	return envDefault("SSL_MODE", "require")
}

// Verbose returns whether or not we should run queries in verbose mode.
// Default is false
func Verbose() bool {
	return envBool("VERBOSE", false)
}