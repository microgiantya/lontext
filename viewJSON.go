package lontext

import (
	"encoding/json"
	"fmt"
	"strings"
)

type lontextViewJSONType struct {
	Version     string `json:"version"`
	Severity    string `json:"severity"`
	UniqieID    string `json:"unique_id,omitempty"`
	FileName    string `json:"file_name"`
	FileLineNum string `json:"file_line_num"`
	Message     string `json:"message"`
}

func viewJSON(v lontextData) (logLines []string) {
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

	var ltxViewJSONType = lontextViewJSONType{}
	switch v.uniqueID {
	case "":
	default:
		ltxViewJSONType.UniqieID = v.uniqueID
	}
	ltxViewJSONType.Version = v.version
	ltxViewJSONType.Severity = _ltxStaff[v.severity].severity
	ltxViewJSONType.UniqieID = v.uniqueID
	ltxViewJSONType.FileName = v.fileName
	ltxViewJSONType.FileLineNum = v.fileLineNum
	ltxViewJSONType.Message = messageRaw

	logLineBytes, _ := json.Marshal(ltxViewJSONType)

	logLines = strings.Split(string(logLineBytes), "\n")
	return
}
