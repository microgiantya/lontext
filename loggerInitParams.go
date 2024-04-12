package logger

type LoggerInitParams struct {
	Severity       float64
	UniqueIDPrefix string
	Version        string
	separate       bool
	View           loggerView
}

func (t *LoggerInitParams) fixSeverity() {
	if t.Severity < 0 || t.Severity > 7 {
		t.Severity = 0
	}
}

func (t *LoggerInitParams) fillSeverity(loggers loggerTransmissions) {
	t.fixSeverity()
	_severity := int(t.Severity)

	for i := 0; i <= _severity; i++ {
		switch t.View {
		case "json":
			loggers[i].setLoggerFunc(showJSONLine)
		default:
			loggers[i].setLoggerFunc(showPlainLine)
		}
	}
	for i := _severity + 1; i < 8; i++ {
		loggers[i].setLoggerFunc(dropLogLine)
	}
}
