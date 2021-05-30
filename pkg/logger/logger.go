package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

var (
	l *log.Logger
)

const (
	green   = "\033[97;42m"
	white   = "\033[90;47m"
	yellow  = "\033[90;43m"
	red     = "\033[97;41m"
	blue    = "\033[97;44m"
	magenta = "\033[97;45m"
	cyan    = "\033[97;46m"
	reset   = "\033[0m"
)

type LogLevel int8

const (
	debug LogLevel = 0
	info  LogLevel = 1
	warn  LogLevel = 2
	err   LogLevel = 3
	fatal LogLevel = 4
)

func init() {
	//f, _ := os.OpenFile("./test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	writer := io.MultiWriter(os.Stdout)
	l = log.New(writer, "", log.Lshortfile|log.Ldate|log.Ltime)

}

func logLevelString(level LogLevel) (levelStr string, levelColor string) {
	switch level {
	case debug:
		return "DEBUG", green
	case info:
		return "INFO ", blue
	case warn:
		return "WARN ", yellow
	case err:
		return "ERROR", red
	case fatal:
		return "FATAL", magenta
	default:
		return "", white
	}
}

func printf(level LogLevel, format string, args ...interface{}) {
	levelStr, color := logLevelString(level)
	prefix := fmt.Sprintf("%s %s %s |", color, levelStr, reset)
	_ = l.Output(3, fmt.Sprintf("%s %s ", prefix, fmt.Sprintf(format, args...)))
	if level == fatal {
		os.Exit(1)
	}
}

func Debugf(format string, args ...interface{}) {
	printf(debug, format, args...)
}

func Infof(format string, args ...interface{}) {
	printf(info, format, args...)
}

func Warnf(format string, args ...interface{}) {
	printf(warn, format, args...)
}

func Errorf(format string, args ...interface{}) {
	printf(err, format, args...)
}

func Fatalf(format string, args ...interface{}) {
	printf(fatal, format, args...)
}
