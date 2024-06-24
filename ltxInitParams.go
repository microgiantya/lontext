package lontext

import "io"

type LontextInitParams struct {
	Prefix        string
	Version       string
	View          view
	Severity      float64
	Writer        io.Writer
	separate      bool
	needChanClose bool
}

func (t *LontextInitParams) fixSeverity() {
	if t.Severity < 0 || t.Severity > 7 {
		t.Severity = 0
	}
}

func (t *LontextInitParams) fillSeverity(ltxs transmissions) {
	t.fixSeverity()
	_severity := int(t.Severity)

	for i := 0; i <= _severity; i++ {
		switch t.View {
		case "json":
			ltxs[i].setLontextFunc(showJSONLine)
		default:
			ltxs[i].setLontextFunc(showPlainLine)
		}
	}
	for i := _severity + 1; i < 8; i++ {
		ltxs[i].setLontextFunc(dropLogLine)
	}
}
