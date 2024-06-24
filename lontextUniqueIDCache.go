package lontext

import (
	"sync"
)

var (
	lontextUniqueIDCache sync.Map
)

func getLontextUniqueIDFromCache(prefix string) (uniqueID int64) {
	if rawUniqueID, ok := lontextUniqueIDCache.Load(prefix); ok {
		uniqueID = rawUniqueID.(int64) + 1
	}
	lontextUniqueIDCache.Store(prefix, uniqueID)
	return
}
