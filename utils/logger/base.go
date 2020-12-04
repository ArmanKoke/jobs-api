package logger

import (
	"jobs-api/utils/env"

	"go.uber.org/zap"
)

// LogEnv interface
type LogEnv interface {
	ZapLogger() *zap.Logger
}

// NewLogEnv instance
func NewLogEnv() LogEnv {
	log := new(Log)
	log.Env = env.AppDebugMode()
	return log
}
