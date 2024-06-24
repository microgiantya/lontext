package lontext

var (
	_ltxStaff = map[int]*ltxStaff{
		0: {severity: "EMERGENCY", color: "\033[38;1m"},
		1: {severity: "ALERT", color: "\033[37;1m"},
		2: {severity: "CRITICAL", color: "\033[31;1m"},
		3: {severity: "ERROR", color: "\033[31m"},
		4: {severity: "WARNING", color: "\033[33m"},
		5: {severity: "NOTICE", color: "\033[36m"},
		6: {severity: "INFORMATIONAL", color: "\033[34m"},
		7: {severity: "DEBUG", color: "\033[35m"},
		8: {severity: "", color: "\033[0m"},
	}
)

type ltxStaff struct {
	severity string
	color    string
}
