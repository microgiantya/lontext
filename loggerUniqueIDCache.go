package logger

import (
	"sync"
)

var (
	loggerUniqueIDPrefixCache sync.Map
)

func getLoggerUniqueIDFromCache(uniqueIDPrefix string) (uniqueID int64) {
	rawUniqueID, ok := loggerUniqueIDPrefixCache.Load(uniqueIDPrefix)
	if ok {
		uniqueID = rawUniqueID.(int64) + 1
	}
	loggerUniqueIDPrefixCache.Store(uniqueIDPrefix, uniqueID)
	return
}
