package logging

import (
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	// log is the base logger of the framework
	log       = &zapLogger{}
	dunno     = []byte("???")
	centerDot = []byte("·")
	dot       = []byte(".")
	slash     = []byte("/")
)

// A Logger writes logs to a handler
type Logger interface {
	// Panic logs a error level message then panics
	Panic(msg string, ctx ...interface{})
	// // Error logs an error level message
	// Error(msg string, ctx ...interface{})
	// // Warn logs a warning level message
	// Warn(msg string, ctx ...interface{})
	// // Info logs an information level message
	Info(msg string, ctx ...interface{})
	// // Debug logs a debug level message. This may be very verbose
	// Debug(msg string, ctx ...interface{})
	// // New returns a child logger with the given context
	// New(ctx ...interface{}) Logger
	// // Sync the logger cache
	// Sync() error
}

// 产生log的结构，为了是链式输出，整个日志
// 现在不需要这样的结构
// 这里

// zapLogger is an implementation of logger using Uber's zap library
type zapLogger struct {
	zap    *zap.SugaredLogger
	ctx    []interface{}
	parent *zapLogger
}

// zap要和jaeger联合
// func WithTrace(ctx context.Context) *zap.SugaredLogger {
// }
// New returns a child logger with the given context
func (l *zapLogger) New(ctx ...interface{}) Logger {
	// 这个地方先临时使用
	return &zapLogger{
		zap:    log.zap,
		ctx:    ctx,
		parent: l,
	}
}

// 要提供一个gin的中间件
// GetLogger returns a context logger for the given module
func GetLogger(moduleName string) Logger {
	l := log.New("module", moduleName)
	return l
}

// Info logs an information level message
func (l *zapLogger) Info(msg string, ctx ...interface{}) {

	l.zap.Infow(msg, ctx...)
}

// Panic logs a error level message then panics
func (l *zapLogger) Panic(msg string, ctx ...interface{}) {
	panicData := msg + "\n"
	for i := 0; i < len(ctx); i += 2 {
		panicData += fmt.Sprintf("\t%v : %v\n", ctx[i], ctx[i+1])
	}
	panic(panicData)
}

// Initialize starts the base logger used by all Hexya components
func Initialize() {
	logConfig := zap.NewProductionConfig()
	if viper.GetBool("Debug") {
		logConfig = zap.NewDevelopmentConfig()
	}

	plainLog, err := logConfig.Build()
	if err != nil {
		panic(err)
	}
	log.zap = plainLog.Sugar()

	log.Info("Sake Starting...")
}
