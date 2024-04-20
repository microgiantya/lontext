package logger

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	fileNameLenLimit           = 14
	lineNumLenLimit            = 4
	loggerCommonFormat         = "%s %s% 13s %14s:%-4s %s%s"
	loggerCommonFormatUniqueID = "%s %s% 13s [%s] %14s:%-4s %s%s"
)

var (
	showPlainLine = func(t loggerData) {
		for _, line := range viewPlain(t) {
			fmt.Println(line)
		}
	}
	showJSONLine = func(t loggerData) {
		for _, line := range viewJSON(t) {
			fmt.Println(line)
		}
	}
	dropLogLine = func(_ loggerData) {}
)

func cutFileName(fileName string) (cuttedFileName string) {
	if len(fileName) > fileNameLenLimit {
		cuttedFileName = fmt.Sprintf("~%s", fileName[len(fileName)-fileNameLenLimit+1:])
		return
	}
	cuttedFileName = fileName
	return
}

func viewPlain(v loggerData) (logLines []string) {
	var messageRaw string

	switch t := v.data.(type) {
	case error:
		if t == nil {
			return
		}
		messageRaw = t.Error()
	case string:
		if t == "" {
			return
		}
		messageRaw = strings.Trim(t, "\n\t")
	default:
		messageRaw = fmt.Sprintf("%+v", v.data)
	}

	cuttedFileName := cutFileName(v.fileName)
	switch v.uniqueID {
	case "":
		for _, logLine := range strings.Split(messageRaw, "\n") {
			logLines = append(logLines, fmt.Sprintf(loggerCommonFormat, v.version, _loggerStaff[v.severity].color, _loggerStaff[v.severity].severity, cuttedFileName, v.fileLineNum, logLine, _loggerStaff[8].color))
		}
	default:
		for _, logLine := range strings.Split(messageRaw, "\n") {
			logLines = append(logLines, fmt.Sprintf(loggerCommonFormatUniqueID, v.version, _loggerStaff[v.severity].color, _loggerStaff[v.severity].severity, v.uniqueID, cuttedFileName, v.fileLineNum, logLine, _loggerStaff[8].color))
		}
	}
	return
}

func viewJSON(v loggerData) (logLines []string) {
	var messageRaw string

	switch t := v.data.(type) {
	case error:
		if t == nil {
			return
		}
		messageRaw = t.Error()
	case string:
		if t == "" {
			return
		}
		messageRaw = strings.Trim(t, "\n\t")
	default:
		messageRaw = fmt.Sprintf("%+v", v.data)
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
	loggerViewJSONType.Message = messageRaw

	logLineBytes, _ := json.Marshal(loggerViewJSONType)

	logLines = strings.Split(string(logLineBytes), "\n")
	return
}
