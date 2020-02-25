package log

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

const (
	reqIDValue = "reqid"
)

func ContextWithReqID(ctx context.Context, reqid string) context.Context {
	return context.WithValue(ctx, reqIDValue, reqid)
}

func GetReqID(ctx context.Context) string {
	reqid := ctx.Value(reqIDValue)
	if reqid != nil {
		return reqid.(string)
	}
	return ""
}

type Logger struct {
	logger   log.Logger
	location string
	reqid    string
}

func NewLogger(logger log.Logger) *Logger {
	return &Logger{
		logger:   logger,
		location: "root",
	}
}

func (l Logger) collectKeyVal(keyval ...interface{}) []interface{} {
	if len(l.reqid) != 0 {
		keyval = append([]interface{}{"reqid", l.reqid}, keyval...)
	}

	if len(l.location) != 0 {
		keyval = append([]interface{}{"location", l.location}, keyval...)
	}

	return keyval
}

func (l Logger) Info(keyval ...interface{}) {
	_ = level.Info(l.KitLogger()).Log(keyval...)
}

func (l Logger) Error(keyval ...interface{}) {
	_ = level.Error(l.KitLogger()).Log(keyval...)
}

func (l Logger) Infom(msg string, keyval ...interface{}) {
	keyval = append([]interface{}{"msg", msg}, keyval...)
	_ = level.Info(l.KitLogger()).Log(keyval...)
}

func (l Logger) Errorm(msg string, keyval ...interface{}) {
	keyval = append([]interface{}{"msg", msg}, keyval...)
	_ = level.Error(l.KitLogger()).Log(keyval...)
}

func (l Logger) KitLogger() log.Logger {
	return log.With(l.logger, l.collectKeyVal()...)
}

func (l Logger) WithLocation(location string) Logger {
	newLogger := l
	newLogger.location = location
	return newLogger
}

func (l Logger) WithReqID(reqid string) Logger {
	newLogger := l
	newLogger.reqid = reqid
	return newLogger
}
