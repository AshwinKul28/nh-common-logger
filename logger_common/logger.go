package logger_common

import (
	"os"

	"github.com/rs/zerolog"
)

var logger zerolog.Logger

func ConfigureLogger() {
	// if value := config.Get("APP_ENV"); value == "development" {
	// 	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	// 	// zerolog.SetGlobalLevel(zerolog.DebugLevel)
	// 	logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()
	// } else {
	// 	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	// 	logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()
	// }
	logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()
}

func With() zerolog.Context {
	return logger.With()
}

func Level(level zerolog.Level) zerolog.Logger {
	return logger.Level(level)
}

func Err(err error) *zerolog.Event {
	return logger.Err(err)
}

func Trace() *zerolog.Event {
	return logger.Trace()
}

func Debug() *zerolog.Event {
	return logger.Debug()
}

func Info() *zerolog.Event {
	return logger.Info()
}

func Warn() *zerolog.Event {
	return logger.Warn()
}

func Error() *zerolog.Event {
	return logger.Error()
}

func Fatal() *zerolog.Event {
	return logger.Fatal()
}

func Panic() *zerolog.Event {
	return logger.Panic()
}

func WithLevel(level zerolog.Level) *zerolog.Event {
	return logger.WithLevel(level)
}

func Log() *zerolog.Event {
	return logger.Log()
}

func Print(v ...interface{}) {
	logger.Print(v...)
}

func Printf(format string, v ...interface{}) {
	logger.Printf(format, v...)
}
