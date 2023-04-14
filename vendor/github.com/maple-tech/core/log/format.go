package log

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

//Debug logs a debug message
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

//Debugf logs a debug message using formatting like fmt.Printf
func Debugf(str string, args ...interface{}) {
	logger.Debug(fmt.Sprintf(str, args...))
}

//Info logs an info message
func Info(args ...interface{}) {
	logger.Info(args...)
}

//Infof logs an info message using formatting like fmt.Printf
func Infof(str string, args ...interface{}) {
	logger.Info(fmt.Sprintf(str, args...))
}

//Error logs an error message
func Error(args ...interface{}) {
	logger.Error(args...)
}

//Errorf logs an error message using formatting like fmt.Printf
func Errorf(str string, args ...interface{}) {
	logger.Error(fmt.Sprintf(str, args...))
}

//Err logs an error message from an error object. Uses fmt.Printf
//for the bulk work, uses the last argument as the error object.
//Example:
//	err := someMethod()
//	log.Err("failed on id '%s'", 123, err)
func Err(msg string, args ...interface{}) {
	logger.WithFields(logrus.Fields{
		"err": args[len(args)-1],
	}).Error(fmt.Sprintf(msg, args[:len(args)-1]...))
}

//ForHTTPPanic wraps logrus.WithFields and fills in the property fields from an http.Request object
func ForHTTPPanic(req *http.Request, err interface{}) *logrus.Entry {
	return logger.WithFields(logrus.Fields{
		"mod":         "web",
		"err":         err,
		"method":      req.Method,
		"host":        req.Host,
		"url":         req.URL.Path,
		"query":       req.URL.RawQuery,
		"remote-addr": req.RemoteAddr,
		"user-agent":  req.UserAgent(),
	})
}

//ForHTTPResponse wraps logrus.WithFields and fills in the property fields from an http.Request object
//with the status and code, intended for error types (unless development is enabled)
func ForHTTPResponse(req *http.Request, status int, code string) *logrus.Entry {
	return logger.WithFields(logrus.Fields{
		"mod":         "web-req",
		"status":      status,
		"code":        code,
		"method":      req.Method,
		"host":        req.Host,
		"url":         req.URL.Path,
		"query":       req.URL.RawQuery,
		"remote-addr": req.RemoteAddr,
		"user-agent":  req.UserAgent(),
	})
}

//TraceEndpoint performs a log operation for tracing account usage based on the endpoint
func TraceEndpoint(userID, companyID *string, r *http.Request, status int) {
	usr := ""
	if userID != nil {
		usr = *userID
	}

	cmp := ""
	if companyID != nil {
		cmp = *companyID
	}

	traceLogger.WithFields(logrus.Fields{
		"type":        "endpoint",
		"user":        usr,
		"company":     cmp,
		"status":      status,
		"method":      r.Method,
		"url":         r.URL.Path,
		"query":       r.URL.RawQuery,
		"remote-addr": r.RemoteAddr,
		"user-agent":  r.UserAgent(),
	}).Info()
}
