package main

import (
	"github.com/sirupsen/logrus"

	logrusLogger "github.com/corpix/logger/logrus"
)

func main() {
	l := logrusLogger.New(logrus.New())

	l.Debug("Hidden")
	l.Print("Info")
	l.Error("Error")
	l.Fatal("Fatal")
}
