package logger

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const (
	p = "logger"
)

var (
	Version = "v0.0.0"

	_trace   = _level{"T", "\033[37m"}
	_debug   = _level{"D", "\033[30m"}
	_verbose = _level{"V", "\033[36m"}
	_info    = _level{"I", "\033[34m"}
	_notice  = _level{"N", "\033[35m"}
	_warning = _level{"W", "\033[33m"}
	_error   = _level{"E", "\033[31m"}
	_fatal   = _level{"F", "\033[1;0;41m"}
	_default = "\033[0m"

	pass = func(t logType) { _log(t) }
	skip = func(_ logType) {}

	ll = []logger{
		{
			f: pass,
		},
		{
			f: pass,
		},
		{
			f: pass,
		},
		{
			f: pass,
		},
		{
			f: pass,
		},
		{
			f: pass,
		},
		{
			f: pass,
		},
		{
			f: pass,
		},
	}

	tCh = make(chan logType)
	dCh = make(chan logType)
	vCh = make(chan logType)
	iCh = make(chan logType)
	nCh = make(chan logType)
	wCh = make(chan logType)
	eCh = make(chan logType)
	fCh = make(chan logType)

	wg sync.WaitGroup

	ctx    context.Context
	cancel context.CancelFunc
)

func listener() {
	wg.Add(1)
	defer wg.Done()

	for {
		select {
		case v := <-tCh:
			ll[7].log(v)
		case v := <-dCh:
			ll[6].log(v)
		case v := <-vCh:
			ll[5].log(v)
		case v := <-iCh:
			ll[4].log(v)
		case v := <-nCh:
			ll[3].log(v)
		case v := <-wCh:
			ll[2].log(v)
		case v := <-eCh:
			ll[1].log(v)
		case v := <-fCh:
			ll[0].log(v)
		case <-ctx.Done():
			return
		}
	}
}

func _log(v logType) {
	s := ""
	switch t := v.Data.(type) {
	case error:
		if t == nil {
			return
		}
		s = t.Error()
	case string:
		if t == "" {
			return
		}
		s = t
	default:
		s = fmt.Sprintf("%+v", v.Data)
	}

	if v.P == "" {
		fmt.Printf("[%s] %s%s%s\n", Version, v.Level.C, s, _default)
		return
	}

	fmt.Printf("[%s] %s%s: %s: %s%s\n", Version, v.Level.C, v.Level.R, v.P, s, _default)
}

func T(p string, data interface{}) {
	tCh <- logType{_trace, p, data}
}

func LogTrace(p string, data interface{}) {
	T(p, data)
}

func D(p string, data interface{}) {
	dCh <- logType{_debug, p, data}
}

func LogDebug(p string, data interface{}) {
	D(p, data)
}

func V(p string, data interface{}) {
	vCh <- logType{_verbose, p, data}
}

func LogVerbose(p string, data interface{}) {
	V(p, data)
}

func I(p string, data interface{}) {
	iCh <- logType{_info, p, data}
}

func LogInfo(p string, data interface{}) {
	I(p, data)
}

func I_(p string, data interface{}) {
	_log(logType{_info, p, data})
}

func LogInfoBypassLogLevel(p string, data interface{}) {
	_log(logType{_info, p, data})
}

func N(p string, data interface{}) {
	nCh <- logType{_notice, p, data}
}

func LogNotice(p string, data interface{}) {
	N(p, data)
}

func W(p string, data interface{}) {
	wCh <- logType{_warning, p, data}
}

func LogWarning(p string, data interface{}) {
	W(p, data)
}

func E(p string, data interface{}) {
	eCh <- logType{_error, p, data}
}

func LogError(p string, data interface{}) {
	E(p, data)
}

func E_(p string, data interface{}) {
	_log(logType{_error, p, data})
}

func LogErrorBypassLogLevel(p string, data interface{}) {
	E_(p, data)
}

func F(p string, data interface{}) {
	fCh <- logType{_fatal, p, data}
}

func LogFatal(p string, data interface{}) {
	F(p, data)
}

func GetFromContext(ctx context.Context, key string) (s string) {
	switch v := ctx.Value(key).(type) {
	case *Helper:
		s = v.Get()
		return
	}
	return
}

func GetValueFromContext(ctx context.Context, key string) (s string) {
	switch v := ctx.Value(key).(type) {
	case *Helper:
		s = v.Get()
		return
	}
	return
}

func SetToContext(ctx context.Context, key string, value string) {
	switch t := ctx.Value(key).(type) {
	case *Helper:
		t.Set(value)
	}
}

func SaveKeyToContext(ctx context.Context, key string, value string) {
	switch t := ctx.Value(key).(type) {
	case *Helper:
		t.Set(value)
	}
}

func LogInit(fl float64) {
	var fNil []int
	var fFill []int

	switch true {
	case fl >= 7:
		fNil = []int{}
		fFill = []int{0, 1, 2, 3, 4, 5, 6, 7}
	case fl >= 6:
		fNil = []int{7}
		fFill = []int{0, 1, 2, 3, 4, 5, 6}
	case fl >= 5:
		fNil = []int{6, 7}
		fFill = []int{0, 1, 2, 3, 4, 5}
	case fl >= 4:
		fNil = []int{5, 6, 7}
		fFill = []int{0, 1, 2, 3, 4}
	case fl >= 3:
		fNil = []int{4, 5, 6, 7}
		fFill = []int{0, 1, 2, 3}
	case fl >= 2:
		fNil = []int{3, 4, 5, 6, 7}
		fFill = []int{0, 1, 2}
	case fl >= 1:
		fNil = []int{2, 3, 4, 5, 6, 7}
		fFill = []int{0, 1}
	default:
		fNil = []int{1, 2, 3, 4, 5, 6, 7}
		fFill = []int{0}
	}

	cancel()
	wg.Wait()

	for _, i := range fNil {
		ll[i].set(skip)
	}

	for _, i := range fFill {
		ll[i].set(pass)
	}

	ctx, cancel = context.WithCancel(context.Background())

	go listener()

	T(p, "enabled: T")
	D(p, "enabled: D")
	V(p, "enabled: V")
	I(p, "enabled: I")
	N(p, "enabled: N")
	W(p, "enabled: W")
	E(p, "enabled: E")
	F(p, "enabled: F")
}

func Prep(ctx context.Context, p string) (_s string) {
	_s = fmt.Sprintf("%s %s", ctx.Value("tx"), p)

	return
}

type logFunc func(t logType)

type logger struct {
	f logFunc
}

func (t *logger) set(f logFunc) {
	t.f = f
}

func (t *logger) log(v logType) {
	t.f(v)
}

type logType struct {
	Level _level
	P     string
	Data  interface{}
}

type _level struct {
	R string
	C string
}

type Helper struct {
	sync.Mutex
	T time.Time
	A string
}

// LogCurrentAction
func (t *Helper) Get() (s string) {
	t.Lock()

	s = fmt.Sprintf(`action: %s starting at: %s (%v);`, t.A, t.T.Format("2006-01-02 15:04:05"), time.Since(t.T))

	t.Unlock()

	return
}

func (t *Helper) GetActionInfo() (s string) {
	t.Lock()

	s = fmt.Sprintf(`action: %s starting at: %s (%v);`, t.A, t.T.Format("2006-01-02 15:04:05"), time.Since(t.T))

	t.Unlock()

	return
}

// SetAction
func (t *Helper) Set(s string) {
	t.Lock()

	t.T = time.Now()
	t.A = s

	t.Unlock()
}

func (t *Helper) SetAction(s string) {
	t.Lock()

	t.T = time.Now()
	t.A = s

	t.Unlock()
}

func init() {
	ctx, cancel = context.WithCancel(context.Background())

	go listener()
}
