package gateways

import (
	"go.uber.org/zap"
	"goHexBoilerplate/src/shared/contracts/gateways"
)

type ZapLogger struct {
	logger *zap.SugaredLogger
}

func NewZapLogger() (gateways.Logger, error) {
	log, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	sugar := log.Sugar()
	defer log.Sync()

	return &ZapLogger{logger: sugar}, nil
}

func (l *ZapLogger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *ZapLogger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

func (l *ZapLogger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l *ZapLogger) Fatalln(args ...interface{}) {
	l.logger.Fatal(args)
}
