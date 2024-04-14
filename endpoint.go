package logger

import (
	"encoding/json"
	"fmt"
)

const (
	loggerCommonFormat         = "%s %s% 13s %s:%s %s%s"
	loggerCommonFormatUniqueID = "%s %s% 13s [%s] %s:%s %s%s"
)

var (
	showPlainLine = func(t loggerData) {
		fmt.Println(viewPlain(t))
	}
	showJSONLine = func(t loggerData) {
		fmt.Println(viewJSON(t))
	}
	dropLogLine = func(_ loggerData) {}
)

func viewPlain(v loggerData) (logLine string) {
	var s string

	switch t := v.data.(type) {
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
		s = fmt.Sprintf("%+v", v.data)
	}

	switch v.uniqueID {
	case "":
		logLine = fmt.Sprintf(loggerCommonFormat, v.version, _loggerStaff[v.severity].color, _loggerStaff[v.severity].severity, v.fileName, v.fileLineNum, s, _loggerStaff[8].color)
	default:
		logLine = fmt.Sprintf(loggerCommonFormatUniqueID, v.version, _loggerStaff[v.severity].color, _loggerStaff[v.severity].severity, v.uniqueID, v.fileName, v.fileLineNum, s, _loggerStaff[8].color)
	}
	return
}

func viewJSON(v loggerData) (logLine string) {
	var message string

	switch t := v.data.(type) {
	case error:
		if t == nil {
			return
		}
		message = t.Error()
	case string:
		if t == "" {
			return
		}
		message = t
	default:
		message = fmt.Sprintf("%+v", v.data)
	}

	var loggerViewJSONType = loggerViewJSONType{}
	switch v.uniqueID {
	case "":
	default:
		loggerViewJSONType.UniqieID = v.uniqueID

	}
	loggerViewJSONType.Version = v.version
	loggerViewJSONType.Severity = _loggerStaff[v.severity].severity
	loggerViewJSONType.UniqieID = v.uniqueID
	loggerViewJSONType.FileName = v.fileName
	loggerViewJSONType.FileLineNum = v.fileLineNum
	loggerViewJSONType.Message = message

	logLineBytes, _ := json.Marshal(loggerViewJSONType)

	logLine = string(logLineBytes)
	return
}
