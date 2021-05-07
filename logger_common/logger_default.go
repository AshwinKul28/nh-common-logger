package logger_common

// import (
// 	"errors"
// 	"io"
// 	"log"
// 	"os"
// 	"strings"
// )

// type Logger interface {
// 	Info(value ...interface{})
// 	Debug(value ...interface{})
// 	Error(value ...interface{})
// 	Fatal(value ...interface{})
// 	Warning(value ...interface{})
// 	Critical(value ...interface{})
// }

// const (
// 	LEVEL_DEBUG = iota
// 	LEVEL_INFO
// 	LEVEL_WARNING
// 	LEVEL_ERROR
// 	LEVEL_CRITICAL
// )

// var Reset = "\033[0m"
// var Red = "\033[31m"
// var Green = "\033[32m"
// var Yellow = "\033[33m"
// var Blue = "\033[34m"
// var Purple = "\033[35m"
// var Cyan = "\033[36m"
// var Gray = "\033[37m"
// var White = "\033[97m"

// var (
// 	ErrInvalidLogLevel = errors.New("invalid log level")
// 	defaultLogger      = logger{
// 		Level:  LEVEL_DEBUG,
// 		Logger: log.New(os.Stderr, "", log.Ltime|log.Lshortfile),
// 	}
// 	logLevels = map[string]int{
// 		"DEBUG":    LEVEL_DEBUG,
// 		"INFO":     LEVEL_INFO,
// 		"ERROR":    LEVEL_ERROR,
// 		"WARNING":  LEVEL_WARNING,
// 		"CRITICAL": LEVEL_CRITICAL,
// 	}
// )

// type logger struct {
// 	Level  int
// 	Prefix string
// 	Logger *log.Logger
// }

// func DefaultLogger() Logger {
// 	return defaultLogger
// }

// func NewLogger(level string, out io.Writer, prefix string) (Logger, error) {
// 	l, ok := logLevels[strings.ToUpper(level)]
// 	if !ok {
// 		return defaultLogger, ErrInvalidLogLevel
// 	}
// 	return logger{
// 		Level:  l,
// 		Prefix: prefix,
// 		Logger: log.New(out, "", log.LstdFlags),
// 	}, nil
// }

// func (l logger) Debug(value ...interface{}) {
// 	if l.Level > LEVEL_DEBUG {
// 		return
// 	}
// 	l.prependLog(Green, "DEBUG:", value)
// }

// func (l logger) Info(value ...interface{}) {
// 	if l.Level > LEVEL_INFO {
// 		return
// 	}
// 	l.prependLog(Blue, "INFO:", value)
// }

// func (l logger) Warning(value ...interface{}) {
// 	if l.Level > LEVEL_WARNING {
// 		return
// 	}
// 	l.prependLog(Yellow, "WARNING:", value)
// }

// func (l logger) Error(value ...interface{}) {
// 	if l.Level > LEVEL_ERROR {
// 		return
// 	}
// 	l.prependLog(Red, "ERROR:", value)
// }

// func (l logger) Critical(value ...interface{}) {
// 	l.prependLog(Red, "CRITICAL:", value)
// }

// func (l logger) Fatal(value ...interface{}) {
// 	l.prependLog(Cyan, "FATAL:", value)
// 	os.Exit(1)
// }

// func (l logger) prependLog(color string, level string, value ...interface{}) {
// 	l.Logger.Println(color, l.Prefix, level, value, Reset)
// }
