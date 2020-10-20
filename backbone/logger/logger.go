package logger

import (
	"context"
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/amine-khemissi/skeleton/backbone/config"
)

func init() {
	go reloadLogger(&lock)
}

var gLogger Logger

var lock = sync.Mutex{}

type Logger interface {
	Debug(ctx context.Context, args ...interface{})
	Info(ctx context.Context, args ...interface{})
	Warn(ctx context.Context, args ...interface{})
	Error(ctx context.Context, args ...interface{})
	init()
}
type Level int

const (
	Debug Level = iota
	Info
	Warn
	Error
)

var str2Level = map[string]Level{
	"debug":   Debug,
	"info":    Info,
	"warning": Warn,
	"error":   Error,
}

func (l Level) ToString() string {
	for k, v := range str2Level {
		if v == l {
			return k
		}
	}
	return ""
}

const (
	moduleName     = "logger"
	levelKey       = "level"
	destinationKey = "destination"
)

type logger struct {
	min     Level
	dstName string
	dst     io.Writer
	release func()
	l       *sync.Mutex
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
	l.dst.Write([]byte(fmt.Sprint("[", time.Now().UTC(), "][", level.ToString(), "]") + fmt.Sprintln(args...)))
}

func Instance() Logger {
	lock.Lock()
	defer lock.Unlock()
	if gLogger == nil {
		gLogger = loadLogger(&lock)
		gLogger.init()
	}
	return gLogger
}
func loadLogger(l *sync.Mutex) *logger {

	logLevelStr, isString := config.Instance().Get(moduleName, levelKey).(string)
	if !isString {
		panic(fmt.Sprintln("expected", moduleName, "::", levelKey, "to be a string"))
	}
	logLevel, isLevel := str2Level[logLevelStr]
	if !isLevel {
		panic(fmt.Sprintln("unknown", moduleName, "::", levelKey, "value:", logLevelStr))
	}
	logDst, isString := config.Instance().Get(moduleName, destinationKey).(string)
	if !isString {
		panic(fmt.Sprintln("expected", moduleName, "::", destinationKey, "to be a string"))
	}
	tmpLogger := &logger{
		min:     logLevel,
		dstName: logDst,
		l:       l,
	}

	return tmpLogger
}

func (l *logger) init() {
	f, err := os.OpenFile(l.dstName,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644)
	if err != nil {
		panic(err)
	}
	l.dst = f
	l.release = func() {
		f.Close()
	}
}
func reloadLogger(l *sync.Mutex) {

	for {
		time.Sleep(3 * time.Second)

		tmpLogger := loadLogger(l)
		if gLogger != nil &&
			gLogger.(*logger).dstName == tmpLogger.dstName &&
			gLogger.(*logger).min == tmpLogger.min {
			continue
		}
		tmpLogger.init()
		lock.Lock()
		gLogger.(*logger).release()
		gLogger = tmpLogger
		lock.Unlock()
	}
}
