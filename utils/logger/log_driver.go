package logger

import (
	"jobs-api/utils/env"

	"go.elastic.co/apm/module/apmzap"
	"go.uber.org/zap"
)

// Log struct
type Log struct {
	Env string
}

// ZapLogger instance
func (l *Log) ZapLogger() *zap.Logger {
	var log *zap.Logger
	var err error
	var wrappedCore = zap.WrapCore((&apmzap.Core{
		FatalFlushTimeout: 10000,
	}).WrapCore)

	if l.Env == env.DefaultDebugMode {
		log, err = zap.NewProduction(wrappedCore)
	} else {
		log, err = zap.NewDevelopment()
	}

	if err != nil {
		log = zap.NewExample()
		log.Warn("Unable to set up the logger. Replaced with example one which shouldn't fail", zap.Error(err))
	}
	zap.ReplaceGlobals(log)
	return log
}
