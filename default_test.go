package logger

import (
	"context"
	"testing"
	"time"
)

// func TestGetUniqueID(t *testing.T) {
// 	data := []string{"00000001", "00000002", "00000003"}
// 	for _, expected := range data {
// 		actual := getUniqueID()
// 		assert.Equal(t, expected, actual)
// 	}
// }

func TestNewDefaultLoggerWithCancel(t *testing.T) {

	var (
		ctx1    *Logger
		ctx2    *Logger
		ctx3    *Logger
		ctx4    *Logger
		cancel1 context.CancelFunc
		cancel2 context.CancelFunc
		// cancel3 context.CancelFunc
		// cancel4 context.CancelFunc
	)
	{
		ctx1, cancel1 = NewLoggerWithCancel(&loggerInitParams{
			UniqueIDPrefix: "apiV1",
		})
		ctx1.LogEmergency("message")
		ctx1.LogAlert("message")
		ctx1.LogCritical("message")
		ctx1.LogError("message")
		ctx1.LogWarning("message")
		ctx1.LogNotice("message")
		ctx1.LogInformational("message")
		ctx1.LogDebug("message")
		time.Sleep(time.Millisecond * 100)

		ctx1.UpdateUniqueID()

		ctx1.LogEmergency("message")
		ctx1.LogAlert("message")
		ctx1.LogCritical("message")
		ctx1.LogError("message")
		ctx1.LogWarning("message")
		ctx1.LogNotice("message")
		ctx1.LogInformational("message")
		ctx1.LogDebug("message")
		time.Sleep(time.Millisecond * 100)
	}
	{
		ctx2, cancel2 = NewLoggerWithCancel(&loggerInitParams{
			Severity:       7,
			UniqueIDPrefix: "apiV2",
		})
		ctx2.LogEmergency("message")
		ctx2.LogAlert("message")
		ctx2.LogCritical("message")
		ctx2.LogError("message")
		ctx2.LogWarning("message")
		ctx2.LogNotice("message")
		ctx2.LogInformational("message")
		ctx2.LogDebug("message")
		time.Sleep(time.Millisecond * 100)

		ctx2.UpdateUniqueID()

		ctx2.LogEmergency("message")
		ctx2.LogAlert("message")
		ctx2.LogCritical("message")
		ctx2.LogError("message")
		ctx2.LogWarning("message")
		ctx2.LogNotice("message")
		ctx2.LogInformational("message")
		ctx2.LogDebug("message")
		time.Sleep(time.Millisecond * 100)
	}

	time.Sleep(time.Second)
	cancel1()
	cancel2()
	time.Sleep(time.Second)

	{
		ctx1, cancel1 = NewLoggerWithCancel(&loggerInitParams{
			Severity:       5,
			Version:        "v1.10.8",
			UniqueIDPrefix: "apiV1",
		})
		ctx1.LogEmergency("message")
		ctx1.LogAlert("message")
		ctx1.LogCritical("message")
		ctx1.LogError("message")
		ctx1.LogWarning("message")
		ctx1.LogNotice("message")
		ctx1.LogInformational("message")
		ctx1.LogDebug("message")
		time.Sleep(time.Millisecond * 100)

		ctx1.UpdateUniqueID()

		ctx1.LogEmergency("message")
		ctx1.LogAlert("message")
		ctx1.LogCritical("message")
		ctx1.LogError("message")
		ctx1.LogWarning("message")
		ctx1.LogNotice("message")
		ctx1.LogInformational("message")
		ctx1.LogDebug("message")
		time.Sleep(time.Millisecond * 100)
	}
	{
		ctx2, cancel2 = NewLoggerWithCancel(&loggerInitParams{
			Severity:       7,
			Version:        "v2.5.94",
			UniqueIDPrefix: "apiV2",
			View:           loggerViewJSON,
		})
		ctx2.LogEmergency("message")
		ctx2.LogAlert("message")
		ctx2.LogCritical("message")
		ctx2.LogError("message")
		ctx2.LogWarning("message")
		ctx2.LogNotice("message")
		ctx2.LogInformational("message")
		ctx2.LogDebug("message")
		time.Sleep(time.Millisecond * 100)

		ctx2.UpdateUniqueID()

		ctx2.LogEmergency("message")
		ctx2.LogAlert("message")
		ctx2.LogCritical("message")
		ctx2.LogError("message")
		ctx2.LogWarning("message")
		ctx2.LogNotice("message")
		ctx2.LogInformational("message")
		ctx2.LogDebug("message")
		time.Sleep(time.Millisecond * 100)
	}

	time.Sleep(time.Second)
	cancel1()
	cancel2()

	ctx3 = NewCommonLoggerContext(context.Background(), &loggerInitParams{
		UniqueIDPrefix: "default",
	})
	ctx3.LogEmergency("message")
	ctx3.LogAlert("message")
	ctx3.LogCritical("message")
	ctx3.LogError("message")
	ctx3.LogWarning("message")
	ctx3.LogNotice("message")
	ctx3.LogInformational("message")
	ctx3.LogDebug("message")

	ctx3.UpdateUniqueID()

	ctx3.LogEmergency("message")
	ctx3.LogAlert("message")
	ctx3.LogCritical("message")
	ctx3.LogError("message")
	ctx3.LogWarning("message")
	ctx3.LogNotice("message")
	ctx3.LogInformational("message")
	ctx3.LogDebug("message")

	ctx4 = NewCommonLoggerContext(context.Background(), &loggerInitParams{
		UniqueIDPrefix: "test",
	})

	ctx4.LogEmergency("message")
	ctx4.LogAlert("message")
	ctx4.LogCritical("message")
	ctx4.LogError("message")
	ctx4.LogWarning("message")
	ctx4.LogNotice("message")
	ctx4.LogInformational("message")
	ctx4.LogDebug("message")
}

func TestNewLoggerContext(t *testing.T) {}
