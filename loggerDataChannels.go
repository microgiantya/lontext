package logger

var (
	commonLoggerDataChannels loggerDataChannels
)

type loggerDataChannels []chan loggerData

func newLoggerDataChannels(separate bool) (loggerDataChannels loggerDataChannels) {
	switch separate {
	case true:
		for i := 0; i < 8; i++ {
			loggerDataChannels = append(loggerDataChannels, make(chan loggerData))
		}
	default:
		if commonLoggerDataChannels == nil {
			for i := 0; i < 8; i++ {
				commonLoggerDataChannels = append(commonLoggerDataChannels, make(chan loggerData))
			}
		}
		loggerDataChannels = commonLoggerDataChannels
	}
	return
}
