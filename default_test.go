package logger

import (
	"fmt"
	"io"
	"os"
	"testing"
	"time"
)

var (
	defaultStdOut *os.File
	pipeOut       *os.File
	pipeIn        *os.File
	set           bool
	lines         []string
)

func switchStdOut() (err error) {
	switch set {
	case true:
		os.Stdout = defaultStdOut
		if pipeIn != nil && pipeOut != nil {
			pipeIn.Close()
			var bb []byte
			bb, err = io.ReadAll(pipeOut)
			if err != nil {
				fmt.Println("readAll err:", err)
				return
			}
			lines = append(lines, fmt.Sprintf("%q", bb))
			pipeOut.Close()
		}
		set = false
	default:
		lines = []string{}
		defaultStdOut = os.Stdout
		pipeOut, pipeIn, err = os.Pipe()
		if err != nil {
			return
		}
		os.Stdout = pipeIn
		set = true
	}
	return
}

func Test1(t *testing.T) {
	_ = switchStdOut()

	ctx, cancel := NewLoggerWithCancel(&LoggerInitParams{
		Severity: 7,
	})
	ctx.LogError("log message")
	time.Sleep(time.Second)
	ctx.Close()
	time.Sleep(time.Second)
	cancel()
	_ = switchStdOut()

	fmt.Println("lines:", lines)
}
func Test2(t *testing.T) {
	_ = switchStdOut()

	ctx, cancel := NewLoggerWithCancel(&LoggerInitParams{
		Severity:       7,
		UniqueIDPrefix: "api",
	})
	ctx.LogError("log message")
	time.Sleep(time.Second)
	cancel()
	_ = switchStdOut()

	fmt.Println("lines:", lines)
}

func Test3(t *testing.T) {
	ctx := NewLogger(&LoggerInitParams{
		Severity:       7,
		UniqueIDPrefix: "apiV2",
		View:           LoggerViewJSON,
	})
	ctx.LogEmergency("message")
	ctx.LogAlert("message")
	ctx.LogCritical("message")
	ctx.LogError("message")

	ctx.IncrementUniqueID()

	ctx.LogWarning("message")
	ctx.LogNotice("message")
	ctx.LogInformational("message")
	ctx.LogDebug("message")
}
