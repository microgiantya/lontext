package lontext

type lontextData struct {
	data        interface{}
	fileName    string
	fileLineNum string
	uniqueID    string
	version     string
	view        view
	severity    int
}
