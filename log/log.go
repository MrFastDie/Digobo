package log

import (
	"Digobo/config"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

var Debug, Info, Warning, Error *log.Logger

func Init() {
	var lvl = config.Config.Log.Lvl
	var toStdout = config.Config.Log.ToStdout

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	t := time.Now().Format("2006-01-02")
	file, err := os.OpenFile(dir + "/log/" + t + ".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	var debugHandle, infoHandle, warningHandle, errorHandle io.Writer

	if toStdout {
		debugHandle = io.MultiWriter(file, os.Stdout)
		infoHandle = io.MultiWriter(file, os.Stdout)
		warningHandle = io.MultiWriter(file, os.Stdout)
		errorHandle = io.MultiWriter(file, os.Stderr)
	} else {
		debugHandle = io.Writer(file)
		infoHandle = io.Writer(file)
		warningHandle = io.Writer(file)
		errorHandle = io.Writer(file)
	}

	switch strings.ToLower(lvl) {
	case "debug":
	case "info":
		debugHandle = ioutil.Discard
	case "warning":
		debugHandle = ioutil.Discard
		infoHandle = ioutil.Discard
	case "error":
		debugHandle = ioutil.Discard
		infoHandle = ioutil.Discard
		warningHandle = ioutil.Discard
	default:
		debugHandle = ioutil.Discard
		infoHandle = ioutil.Discard
	}

	Debug = log.New(debugHandle, "Debug: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(infoHandle, "Info: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(warningHandle, "Warning: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(errorHandle, "Warning: ", log.Ldate|log.Ltime|log.Lshortfile)
}