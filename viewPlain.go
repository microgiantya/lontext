package lontext

import (
	"fmt"
	"strings"
)

func viewPlain(v lontextData) (logLines []string) {
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
			logLines = append(logLines, fmt.Sprintf(
				ltxCommonFormat,
				v.version,
				_ltxStaff[v.severity].color,
				_ltxStaff[v.severity].severity,
				cuttedFileName,
				v.fileLineNum,
				logLine,
				_ltxStaff[8].color,
			))
		}
	default:
		for _, logLine := range strings.Split(messageRaw, "\n") {
			logLines = append(logLines, fmt.Sprintf(
				ltxCommonFormatUniqueID,
				v.version,
				_ltxStaff[v.severity].color,
				_ltxStaff[v.severity].severity,
				v.uniqueID,
				cuttedFileName,
				v.fileLineNum,
				logLine,
				_ltxStaff[8].color,
			))
		}
	}
	return
}
