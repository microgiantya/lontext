package lontext

var (
	transmissionsCommon transmissions
)

type transmission struct {
	f ltxFunc
}

func (t *transmission) setLontextFunc(f ltxFunc) {
	t.f = f
}

func (t *transmission) doTransmission(v lontextData) {
	t.f(v)
}

type transmissions []*transmission

func newLontextTransmissions(separate bool, veiw view) (ltxTransmissions transmissions, needStart bool) {
	switch separate {
	case true:
		for i := 0; i < 8; i++ {
			ltxTransmission := &transmission{}
			switch veiw {
			case lontextViewJSON:
				ltxTransmission.setLontextFunc(showJSONLine)
			default:
				ltxTransmission.setLontextFunc(showPlainLine)
			}
			ltxTransmission.setLontextFunc(showPlainLine)
			ltxTransmissions = append(ltxTransmissions, ltxTransmission)
		}
		needStart = true
	default:
		if transmissionsCommon == nil {
			for i := 0; i < 8; i++ {
				ltxTransmission := &transmission{}
				ltxTransmission.setLontextFunc(showPlainLine)
				transmissionsCommon = append(transmissionsCommon, ltxTransmission)
			}
			needStart = true
		}
		ltxTransmissions = transmissionsCommon
	}
	return
}
