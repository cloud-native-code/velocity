package logger

import "go.uber.org/zap"

func Build(b LoggerBuilder) *zap.Logger {
	return b.buildLogger()
}
