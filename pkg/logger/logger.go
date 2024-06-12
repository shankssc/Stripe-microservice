package logger

import (
	"go.uber.org/zap"
)

type Logger struct {
	*zap.SugaredLogger
}

func NewLogger(level string) *Logger {
	var logger *zap.Logger
	var err error

	switch level {
	case "debug":
		logger, err = zap.NewDevelopment()
	default:
		logger, err = zap.NewProduction()
	}

	if err != nil {
		panic(err)
	}

	return &Logger{logger.Sugar()}
}

func (l *Logger) Info(args ...interface{}) {
	l.SugaredLogger.Info(args...)
}
