package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug 	*log.Logger
	info 		*log.Logger
	warning *log.Logger
	err 		*log.Logger
	writer 	io.Writer
}

func NewLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, prefix, log.Ldate|log.Ltime)

	return &Logger{
		debug: 		log.New(writer, "DEBUG: ", logger.Flags()),
		info: 		log.New(writer, "INFO: ", logger.Flags()),
		warning: 	log.New(writer, "WARNING: ", logger.Flags()),
		err: 			log.New(writer, "ERROR: ", logger.Flags()),
		writer: 	writer,
	}
}

// Logs não formatados
func (logger *Logger) Debug(values ...interface{}) {
	logger.debug.Println(values...)
}

func (logger *Logger) Info(values ...interface{}) {
	logger.info.Println(values...)
}

func (logger *Logger) Warning(values ...interface{}) {
	logger.warning.Println(values...)
}

func (logger *Logger) Error(values ...interface{}) {
	logger.err.Println(values...)
}

// Logs formatados
func (logger *Logger) Debugf(format string, values ...interface{}) {
	logger.debug.Printf(format, values...)
}

func (logger *Logger) Infof(format string, values ...interface{}) {
	logger.info.Printf(format, values...)
}

func (logger *Logger) Warningf(format string, values ...interface{}) {
	logger.warning.Printf(format, values...)
}

func (logger *Logger) Errorf(format string, values ...interface{}) {
	logger.err.Printf(format, values...)
}