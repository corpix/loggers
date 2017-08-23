package main

import (
	"github.com/sirupsen/logrus"

	logrusLogger "github.com/corpix/logger/logger/logrus"
	"github.com/corpix/logger/logger/prefixwrapper"
)

func main() {
	l := prefixwrapper.New(
		"prefix > ",
		logrusLogger.New(logrus.New()),
	)

	_, err := l.Write([]byte("hello"))
	if err != nil {
		panic(err)
	}

	l.Debug("Hidden")
	l.Print("Info")
	l.Error("Error")
	l.Fatal("Fatal")
}
