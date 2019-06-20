package buried

import (
	"github.com/sirupsen/logrus"
	"time"
)

type Logger struct {
	logrus.Logger
}

func NewLog(opts ...Option) *Logger {

	l := &Logger{}
	l.Formatter = &logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
	}
	l.Level = logrus.InfoLevel
	l.Hooks = make(logrus.LevelHooks)

	b := &buried{w: new(writer)}
	for _, opt := range opts {
		opt(b)
	}
	l.AddHook(b)

	return l
}
