package log

import "go.uber.org/zap"

var Logger, _ = zap.NewProduction()

func SetDefaultLogger(logger *zap.Logger) {
	Logger = logger
}
