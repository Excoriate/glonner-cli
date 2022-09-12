package logger

import (
	"go.uber.org/zap"
)

func makeLogger() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	return sugar
}

func NewLogger() ILogger {
	return Logger{logger: makeLogger()}
}

type ILogger interface {
	LogError(args ...interface{})
	LogWarn(args ...interface{})
	LogDebug(args ...interface{})
	LogInfo(args ...interface{})
	LogInfoF(template string, args ...interface{})
	LogErrorF(template string, args ...interface{})
	LogWarnF(template string, args ...interface{})
	LogDebugF(template string, args ...interface{})
	LogFatalf(template string, args ...interface{})
	LogFatal(args ...interface{})
}

type Logger struct {
	logger *zap.SugaredLogger
}

func (l Logger) LogInfo(args ...interface{}) {
	l.logger.Info(args...)
}

func (l Logger) LogInfoF(template string, args ...interface{}) {
	l.logger.Infof(template, args...)
}

func (l Logger) LogError(args ...interface{}) {
	l.logger.Error(args...)
}

func (l Logger) LogErrorF(template string, args ...interface{}) {
	l.logger.Errorf(template, args...)
}

func (l Logger) LogDebug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l Logger) LogDebugF(template string, args ...interface{}) {
	l.logger.Debugf(template, args...)
}

func (l Logger) LogWarn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l Logger) LogWarnF(template string, args ...interface{}) {
	l.logger.Warnf(template, args...)
}

func (l Logger) LogFatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l Logger) LogFatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args...)
}
