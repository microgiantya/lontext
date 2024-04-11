package logger

type loggerViewJSONType struct {
	Version     string `json:"version"`
	Severity    string `json:"severity"`
	UniqieID    string `json:"unique_id,omitempty"`
	FileName    string `json:"file_name"`
	FileLineNum string `json:"file_line_num"`
	Message     string `json:"message"`
}
