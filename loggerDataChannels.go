package logger

type loggerDataChannels []chan loggerData

func newLoggerDataChannels() (logDataChannels loggerDataChannels) {
	for i := 0; i < 8; i++ {
		logDataChannels = append(logDataChannels, make(chan loggerData))
	}
	return
}
