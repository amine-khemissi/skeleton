package logger

import (
	"context"
	"fmt"
	"io"
	"os"
	"sync"
)

var gLogger Logger

var lock = sync.Mutex{}

type Logger interface {
	Debug(ctx context.Context, args ...interface{})
	Info(ctx context.Context, args ...interface{})
	Warn(ctx context.Context, args ...interface{})
	Error(ctx context.Context, args ...interface{})
}
type Level int

const (
	Debug Level = iota
	Info
	Warn
	Error
)

type logger struct {
	min Level
	dst io.Writer
}

func (l *logger) Debug(ctx context.Context, args ...interface{}) {
	l.log(ctx, Debug, args)
}
func (l *logger) Info(ctx context.Context, args ...interface{}) {
	l.log(ctx, Info, args)
}
func (l *logger) Warn(ctx context.Context, args ...interface{}) {
	l.log(ctx, Warn, args)
}
func (l *logger) Error(ctx context.Context, args ...interface{}) {
	l.log(ctx, Error, args)
}

func (l *logger) log(ctx context.Context, level Level, args []interface{}) {
	lock.Lock()
	defer lock.Unlock()
	if l.min > level {
		return
	}
	l.dst.Write([]byte(fmt.Sprintln(args...)))
}

type Option func(l *logger)

func SetLevel(level Level) Option {
	return func(l *logger) {
		l.min = level
	}
}
func SetDst(dst string) Option {
	//todo : dst should be able to handle tcp/htt[ stream in addition to file system output
	return func(l *logger) {
		f, err := os.Open(dst)
		if err != nil {
			panic(err)
		}
		l.dst = f
	}
}

func Instance() Logger {
	lock.Lock()
	defer lock.Unlock()
	if gLogger == nil {
		gLogger = New()
	}
	return gLogger
}

func New(opts ...Option) Logger {

	// todo: New should not be exported, it should be configured at starting phase
	// todo : add timestamp and additional ctx values
	gLogger := &logger{
		min: Debug,
		dst: nil,
	}
	for _, opt := range opts {
		opt(gLogger)
	}
	if gLogger.dst == nil {
		gLogger.dst = os.Stdout
	}
	return gLogger
}
