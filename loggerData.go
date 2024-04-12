package logger

type loggerData struct {
	fileName    string
	fileLineNum string
	severity    int
	uniqueID    string
	data        interface{}
	version     string
	view        loggerView
}
