package log

import (
	"context"
	"sync"
	"go.uber.org/zap"
)

var (
	l    *zap.SugaredLogger
	lock sync.Mutex
)

const (
	Contextkey = "decex:loggerInCtx"
)

func Logger() *zap.SugaredLogger {
	if l != nil {
		return l
	}
	lock.Lock()
	defer lock.Unlock()
	if l != nil {
		return l
	}
	logger, _ := zap.NewProduction()
	l = logger.Sugar()
	return l
}

func WithContext(c context.Context) *zap.SugaredLogger {
	if c == nil {
		return Logger()
	}
	if v, ok := c.Value(Contextkey).(*zap.SugaredLogger); ok {
		return v
	}
	return Logger()
}
