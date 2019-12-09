package logger

import (
	"fmt"
	"os"
)

type Logger interface {
	Debugf(format string, a ...interface{}) (n int, err error)
}

type logger struct {
}

func NewLogger() Logger {
	return &logger{}
}

func (l logger) Debugf(format string, a ...interface{}) (int, error) {
	return fmt.Fprintf(os.Stdout, format, a...)
}
