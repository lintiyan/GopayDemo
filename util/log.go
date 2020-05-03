package util

import (
	"os"

	"github.com/sirupsen/logrus"
)

var ErrorLog *logrus.Logger
var AccessLog *logrus.Logger
var errorLogFile = "./tmp/error.log"
var accessLogFile = "./tmp/access.log"

func init() {
	initErrorLog()
	initAccessLog()
}

func initErrorLog() {
	ErrorLog = logrus.New()
	ErrorLog.SetFormatter(&logrus.JSONFormatter{})
	file, err := os.OpenFile(errorLogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(err)
	}
	ErrorLog.SetOutput(file)
}

func initAccessLog() {
	AccessLog = logrus.New()
	AccessLog.SetFormatter(&logrus.JSONFormatter{})
	file, err := os.OpenFile(accessLogFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	AccessLog.SetOutput(file)
}
