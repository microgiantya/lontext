package lontext

var (
	commonLontextDataChannels lontextDataChannels
)

type lontextDataChannels []chan lontextData

func newLontextDataChannels(separate bool) (ltxDataChannels lontextDataChannels) {
	switch separate {
	case true:
		for i := 0; i < 8; i++ {
			ltxDataChannels = append(ltxDataChannels, make(chan lontextData))
		}
	default:
		if commonLontextDataChannels == nil {
			for i := 0; i < 8; i++ {
				commonLontextDataChannels = append(commonLontextDataChannels, make(chan lontextData))
			}
		}
		ltxDataChannels = commonLontextDataChannels
	}
	return
}
