package logger

import (
	"context"
	"fmt"
	"path"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

type Logger struct {
	ctx            context.Context
	uniqueIDPrefix string
	uniqueID       atomic.Int64
	wg             sync.WaitGroup
	loggers        loggerTransmissions
	channels       loggerDataChannels
	version        string
	view           loggerView
}

func fixUniqueIDPrefix(uniqueIDPrefix string) (fixedUniqueIDPrefix string) {
	fixedUniqueIDPrefix = uniqueIDPrefix
	if fixedUniqueIDPrefix == "" {
		fixedUniqueIDPrefix = "unknown"
	}
	return
}

func fixVersion(version string) (fixedVersion string) {
	fixedVersion = version
	if fixedVersion == "" {
		fixedVersion = "v0.0.0"
	}
	return
}

func fixView(view loggerView) (fixedView loggerView) {
	fixedView = view
	if fixedView != LoggerViewPlain && fixedView != LoggerViewJSON {
		fixedView = LoggerViewPlain
	}
	return
}

func newLogger(ctx context.Context, params *LoggerInitParams) (logger *Logger) {
	params.View = fixView(params.View)
	params.Version = fixVersion(params.Version)
	params.UniqueIDPrefix = fixUniqueIDPrefix(params.UniqueIDPrefix)

	loggers, needStart := newLoggerTransmissions(params.separate, params.View)
	params.fillSeverity(loggers)

	logger = &Logger{
		ctx:            ctx,
		uniqueIDPrefix: params.UniqueIDPrefix,
		loggers:        loggers,
		channels:       newLoggerDataChannels(params.separate),
		version:        params.Version,
		view:           params.View,
	}

	logger.uniqueID.Store(getLoggerUniqueIDFromCache(logger.uniqueIDPrefix))

	if needStart {
		logger.wg.Add(1)
		go logger.receive()
		logger.wg.Wait()
	}
	return
}

func NewLoggerWithCancel(params *LoggerInitParams) (logger *Logger, cancel context.CancelFunc) {
	params.separate = true
	ctx, cancel := context.WithCancel(context.Background())
	logger = newLogger(ctx, params)
	return
}

func NewLoggerContext(ctx context.Context, params *LoggerInitParams) *Logger {
	params.separate = true
	return newLogger(ctx, params)
}

func NewLoggerContextWithCancel(ctx context.Context, params *LoggerInitParams) (logger *Logger, cancel context.CancelFunc) {
	params.separate = true
	_ctx, cancel := context.WithCancel(ctx)
	logger = newLogger(_ctx, params)
	return
}

func NewCommonLoggerWithCancel(params *LoggerInitParams) (*Logger, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	return newLogger(ctx, params), cancel
}

func NewCommonLoggerContext(ctx context.Context, params *LoggerInitParams) *Logger {
	return newLogger(ctx, params)
}

func (t *Logger) UpdateUniqueID() {
	t.uniqueID.Store(getLoggerUniqueIDFromCache(t.uniqueIDPrefix))
}

func (t *Logger) log(severity int, data interface{}) {
	var fileName string
	_, file, line, ok := runtime.Caller(2)
	if ok {
		fileName = path.Base(file)
	}

	t.channels[severity] <- loggerData{
		fileName:    fileName,
		fileLineNum: fmt.Sprint(line),
		severity:    severity,
		uniqueID:    fmt.Sprintf("%s-%08X", t.uniqueIDPrefix, uint32(t.uniqueID.Load())),
		data:        data,
		version:     t.version,
		view:        t.view,
	}
}

func (t *Logger) LogEmergency(data interface{}) {
	t.log(0, data)
}

func (t *Logger) LogAlert(data interface{}) {
	t.log(1, data)
}

func (t *Logger) LogCritical(data interface{}) {
	t.log(2, data)
}

func (t *Logger) LogError(data interface{}) {
	t.log(3, data)
}

func (t *Logger) LogWarning(data interface{}) {
	t.log(4, data)
}

func (t *Logger) LogNotice(data interface{}) {
	t.log(5, data)
}

func (t *Logger) LogInformational(data interface{}) {
	t.log(6, data)
}

func (t *Logger) LogDebug(data interface{}) {
	t.log(7, data)
}

func (t *Logger) Deadline() (deadline time.Time, ok bool) {
	return t.ctx.Deadline()
}

func (t *Logger) Done() <-chan struct{} {
	return t.ctx.Done()
}

func (t *Logger) Err() error {
	return t.ctx.Err()
}

func (t *Logger) Value(v any) any {
	return t.ctx.Value(v)
}

func (t *Logger) receive() {
	t.wg.Done()
	for {
		select {
		case v := <-t.channels[0]:
			t.loggers[0].doTransmission(v)
		case v := <-t.channels[1]:
			t.loggers[1].doTransmission(v)
		case v := <-t.channels[2]:
			t.loggers[2].doTransmission(v)
		case v := <-t.channels[3]:
			t.loggers[3].doTransmission(v)
		case v := <-t.channels[4]:
			t.loggers[4].doTransmission(v)
		case v := <-t.channels[5]:
			t.loggers[5].doTransmission(v)
		case v := <-t.channels[6]:
			t.loggers[6].doTransmission(v)
		case v := <-t.channels[7]:
			t.loggers[7].doTransmission(v)
		case <-t.ctx.Done():
			return
		}
	}
}
