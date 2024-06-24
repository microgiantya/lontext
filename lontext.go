// Copyright 2024 Evgeny Artemev. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Warning: go anti pattern.
//
// Package lontext is a union of logger entity and [context] package.
// Provides easy logging messages with unique label for related events like transactions, requests - responses, etc.
//
// Inspired by Asterisk PBX (https://github.com/asterisk/asterisk) logger outputs.
package lontext

import (
	"context"
	"fmt"
	"io"
	"path"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// Lontext struct contains [context.Context] which make this package go anti pattern.
type Lontext struct {
	ctx       context.Context //nolint // antipattern
	chanClose chan struct{}
	prefix    string
	version   string
	view      view
	ltxs      transmissions
	channels  lontextDataChannels
	wg        sync.WaitGroup
	uniqueID  atomic.Int64
	writer    io.Writer
}

func newLontextContext(ctx context.Context, params *LontextInitParams) (ltx *Lontext) {
	params.View = fixView(params.View)
	params.Version = fixVersion(params.Version)
	params.Prefix = fixPrefix(params.Prefix)

	ltxs, needStart := newLontextTransmissions(params.separate, params.View)
	params.fillSeverity(ltxs)

	ltx = &Lontext{
		ctx:      ctx,
		prefix:   params.Prefix,
		ltxs:     ltxs,
		channels: newLontextDataChannels(params.separate),
		version:  params.Version,
		view:     params.View,
	}

	if params.Writer != nil {
		ltx.writer = params.Writer
	}

	ltx.uniqueID.Store(getLontextUniqueIDFromCache(ltx.prefix))

	if params.needChanClose {
		ltx.chanClose = make(chan struct{})
	}

	if needStart {
		ltx.wg.Add(1)
		go ltx.receive()
		ltx.wg.Wait()
	}
	return
}

// NewLontextCommon return new instance of [Lontext] pointer using [context.Background].
// Close method must be called to close [Lontext] receiver.
func NewLontextCommon(params *LontextInitParams) (ltx *Lontext) {
	params.needChanClose = true
	ltx = newLontextContext(context.Background(), params)
	return
}

// NewLontextCommonWithCancel return new instance of [Lontext] pointer using [context.WithCancel].
// [Lontext] receiver will closed when cancel is called.
func NewLontextCommonWithCancel(params *LontextInitParams) (ltx *Lontext, cancel context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	ltx = newLontextContext(ctx, params)
	return
}

// NewLontextCommon return new instance of [Lontext] pointer using provided [context.Context].
// Close method must be called to close [Lontext] receiver, otherwise receiver will
// close when provided [context.Context] is Done.
func NewLontextCommonContext(ctx context.Context, params *LontextInitParams) (ltx *Lontext) {
	params.needChanClose = true
	ltx = newLontextContext(ctx, params)
	return
}

// NewLontextCommonWithCancelContext same as [NewLontextCommonWithCancel] but use provided [context.Context].
// Returned cancel allow cancel child tasks.
func NewLontextCommonWithCancelContext(ctx context.Context, params *LontextInitParams) (ltx *Lontext, cancel context.CancelFunc) {
	_ctx, cancel := context.WithCancel(ctx)
	ltx = newLontextContext(_ctx, params)
	return
}

// Close() method used to close Lontext receiver.
func NewLontext(params *LontextInitParams) (ltx *Lontext) {
	params.separate = true
	ltx = NewLontextCommon(params)
	return
}

// Done() method used to close Lontext receiver.
func NewLontextWithCancel(params *LontextInitParams) (ltx *Lontext, cancel context.CancelFunc) {
	params.separate = true
	ltx, cancel = NewLontextCommonWithCancel(params)
	return
}

// Close() method used to close [Lontext] receiver.
func NewLontextContext(ctx context.Context, params *LontextInitParams) (ltx *Lontext) {
	params.separate = true
	ltx = NewLontextCommonContext(ctx, params)
	return
}

// Done() method used to close Lontext receiver.
func NewLontextContextWithCancel(ctx context.Context, params *LontextInitParams) (ltx *Lontext, cancel context.CancelFunc) {
	params.separate = true
	ltx, cancel = NewLontextCommonWithCancelContext(ctx, params)
	return
}

func (t *Lontext) write(severity int, data interface{}) {
	var fileName string
	_, file, line, ok := runtime.Caller(2)
	if ok {
		fileName = path.Base(file)
	}

	t.channels[severity] <- lontextData{
		fileName:    fileName,
		fileLineNum: fmt.Sprint(line),
		severity:    severity,
		uniqueID:    fmt.Sprintf("%s-%08X", t.prefix, uint32(t.uniqueID.Load())),
		data:        data,
		version:     t.version,
		view:        t.view,
	}
}

func (t *Lontext) receive() {
	t.wg.Done()
	for {
		select {
		case v := <-t.channels[0]:
			t.ltxs[0].doTransmission(v)
		case v := <-t.channels[1]:
			t.ltxs[1].doTransmission(v)
		case v := <-t.channels[2]:
			t.ltxs[2].doTransmission(v)
		case v := <-t.channels[3]:
			t.ltxs[3].doTransmission(v)
		case v := <-t.channels[4]:
			t.ltxs[4].doTransmission(v)
		case v := <-t.channels[5]:
			t.ltxs[5].doTransmission(v)
		case v := <-t.channels[6]:
			t.ltxs[6].doTransmission(v)
		case v := <-t.channels[7]:
			t.ltxs[7].doTransmission(v)
		case <-t.ctx.Done():
			if t.chanClose == nil {
				return
			}
		case <-t.chanClose:
			return
		}
	}
}

func (t *Lontext) IncrementUniqueID() {
	t.uniqueID.Store(getLontextUniqueIDFromCache(t.prefix))
}

// Emergency writes message using emergency severity. The [os.Exit](1) after is called.
func (t *Lontext) Emergency(data interface{}) {
	t.write(0, data)
}

// Alert writes message using alert severity.
func (t *Lontext) Alert(data interface{}) {
	t.write(1, data)
}

// Critical writes message using critical severity.
func (t *Lontext) Critical(data interface{}) {
	t.write(2, data)
}

// Error writes message using error severity.
func (t *Lontext) Error(data interface{}) {
	t.write(3, data)
}

// Warning writes message using warning severity.
func (t *Lontext) Warning(data interface{}) {
	t.write(4, data)
}

// Notice writes message using notice severity.
func (t *Lontext) Notice(data interface{}) {
	t.write(5, data)
}

// Informational writes message using informational severity.
func (t *Lontext) Informational(data interface{}) {
	t.write(6, data)
}

// Debug writes message using debug severity.
func (t *Lontext) Debug(data interface{}) {
	t.write(7, data)
}

// Deadline purely wrap [context.Context] Deadline method.
func (t *Lontext) Deadline() (deadline time.Time, ok bool) {
	return t.ctx.Deadline()
}

// Done purely wrap [context.Context] Done method.
func (t *Lontext) Done() <-chan struct{} {
	return t.ctx.Done()
}

// Err purely wrap [context.Context] Err method.
func (t *Lontext) Err() error {
	return t.ctx.Err()
}

// Value purely wrap [context.Context] Value method.
func (t *Lontext) Value(v any) any {
	return t.ctx.Value(v)
}

// Close release [Lontext] resources.
func (t *Lontext) Close() {
	if t.chanClose != nil {
		t.chanClose <- struct{}{}
		return
	}
}

// TODO // write flush method
