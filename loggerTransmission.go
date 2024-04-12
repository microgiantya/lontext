package logger

var (
	commonLoggerTransmissions loggerTransmissions
)

type loggerTransmission struct {
	f loggerFunc
}

func (t *loggerTransmission) setLoggerFunc(f loggerFunc) {
	t.f = f
}

func (t *loggerTransmission) doTransmission(v loggerData) {
	t.f(v)
}

type loggerTransmissions []*loggerTransmission

func newLoggerTransmissions(separate bool, veiw loggerView) (loggerTransmissions loggerTransmissions, needStart bool) {
	switch separate {
	case true:
		for i := 0; i < 8; i++ {
			loggerTransmission := &loggerTransmission{}
			switch veiw {
			case loggerViewJSON:
				loggerTransmission.setLoggerFunc(showJSONLine)
			default:
				loggerTransmission.setLoggerFunc(showPlainLine)

			}
			loggerTransmission.setLoggerFunc(showPlainLine)
			loggerTransmissions = append(loggerTransmissions, loggerTransmission)
		}
		needStart = true
	default:
		if commonLoggerTransmissions == nil {
			for i := 0; i < 8; i++ {
				loggerTransmission := &loggerTransmission{}
				loggerTransmission.setLoggerFunc(showPlainLine)
				commonLoggerTransmissions = append(commonLoggerTransmissions, loggerTransmission)
			}
			needStart = true
		}
		loggerTransmissions = commonLoggerTransmissions
	}
	return
}
