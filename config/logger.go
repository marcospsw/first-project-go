package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NewLogger() *Logger {
	writer := io.Writer(os.Stdout)
	flags := log.Ldate | log.Ltime
	return &Logger{
		debug:   log.New(writer, "DEBUG: ", flags),
		info:    log.New(writer, "INFO: ", flags),
		warning: log.New(writer, "WARNING: ", flags),
		err:     log.New(writer, "ERROR: ", flags),
		writer:  writer,
	}
}

// Create Non-Formatted Logs
func (l *Logger) Debug(values ...interface{}) {
	l.debug.Println(values...)
}
func (l *Logger) Info(values ...interface{}) {
	l.info.Println(values...)
}
func (l *Logger) Warning(values ...interface{}) {
	l.warning.Println(values...)
}
func (l *Logger) Err(values ...interface{}) {
	l.err.Println(values...)
}

// Create Formatted Logs
func (l *Logger) Debugf(format string, values ...interface{}) {
	l.debug.Printf(format, values...)
}
func (l *Logger) Infof(format string, values ...interface{}) {
	l.info.Printf(format, values...)
}
func (l *Logger) Warningf(format string, values ...interface{}) {
	l.warning.Printf(format, values...)
}
func (l *Logger) Errf(format string, values ...interface{}) {
	l.err.Printf(format, values...)
}
