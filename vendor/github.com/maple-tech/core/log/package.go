// Package log wraps up the logging functionality for the service
package log

import (
	"errors"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/maple-tech/core/config"
)

var logger = logrus.New()
var traceLogger = logrus.New()
var rotate *Logger

// Initialize starts the logging processes needed, must take place after initial configuration is loaded
func Initialize(cfg *config.OptionsLogging) error {
	//Ensure we have a valid configuration Options provided, or go get it, or error out
	if cfg == nil {
		if config.IsLoaded() == false {
			return errors.New("core.Config package must be initialized before using the logging functions")
		}

		cfg = &(config.Get().Logging)
	}

	//In development mode, we can ditch the old logs to keep things cleaner
	if config.IsDevelopment() {
		if _, err := os.Stat(cfg.Path); err == nil {
			os.Remove(cfg.Path)
		}

		//Tells Logrus to log the method name. I chose this as dev only since it may leak source code paths
		//logger.SetReportCaller(true)
	}

	if cfg.Format == "JSON" {
		logger.Formatter = &logrus.JSONFormatter{}
	} else {
		logger.Formatter = &logrus.TextFormatter{DisableColors: !cfg.Colors}
	}

	switch cfg.Level {
	case "DEBUG":
		logger.Level = logrus.DebugLevel
	case "INFO":
		logger.Level = logrus.InfoLevel
	case "ERROR":
		logger.Level = logrus.ErrorLevel
	default:
		logger.Level = logrus.WarnLevel
	}

	//Check if we want the rotator functionality
	if cfg.Rotate {
		rotate = &Logger{
			Filename:   cfg.Path,
			MaxSize:    cfg.MaxSize,
			MaxBackups: cfg.Backups,
			MaxAge:     cfg.MaxAge,
			Compress:   cfg.Compress,
		}

		logger.Out = rotate
	} else {
		f, err := os.OpenFile(cfg.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0750)
		if err != nil {
			return fmt.Errorf("failed to open log file for writing; error = %s", err.Error())
		}

		logger.Out = f
	}

	//Setup the trace logger, which logs account usage
	traceLogger.Formatter = &logrus.JSONFormatter{}
	traceLogger.Level = logrus.TraceLevel
	traceLogger.Out = os.Stdout

	return nil
}

// Shutdown cleans up the rotator for the logging mechanism
func Shutdown() {
	if rotate != nil {
		rotate.Close()
	}
}

// Get returns the current logger pointer
func Get() *logrus.Logger {
	return logger
}

// GetTraceLogger returns the current traceLogger pointer
func GetTraceLogger() *logrus.Logger {
	return traceLogger
}
