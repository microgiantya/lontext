package lontext

import (
	"fmt"
)

const (
	fileNameLenLimit        = 14
	lineNumLenLimit         = 4
	ltxCommonFormat         = "%s %s% 13s %14s:%-4s %s%s"
	ltxCommonFormatUniqueID = "%s %s% 13s [%s] %14s:%-4s %s%s"
)

var (
	showPlainLine = func(t lontextData) {
		for _, line := range viewPlain(t) {
			fmt.Println(line)
		}
	}
	showJSONLine = func(t lontextData) {
		for _, line := range viewJSON(t) {
			fmt.Println(line)
		}
	}
	dropLogLine = func(_ lontextData) {}
)

func cutFileName(fileName string) (cuttedFileName string) {
	if len(fileName) > fileNameLenLimit {
		cuttedFileName = "~" + fileName[len(fileName)-fileNameLenLimit+1:]
		return
	}
	cuttedFileName = fileName
	return
}
