package buried

import (
	"github.com/sirupsen/logrus"
)

const LogBasePath = "logs/"

type buried struct {
	w *writer
}

// 自定义日志level
func (h *buried) Levels() []logrus.Level {
	return logrus.AllLevels
}

// 自定义钩子执行（默认协程安全）
func (h *buried) Fire(e *logrus.Entry) error {

	e.Logger.Out = h.w
	return nil
}
