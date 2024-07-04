[![go report card](https://goreportcard.com/badge/github.com/microgiantya/lontext "go report card")](https://goreportcard.com/report/github.com/microgiantya/lontext)

# Logger
## _Another go logger package_
Logger package was inspired by asterisk PBX logger.

##### Features
- Logger implements context.Context interface, which allow easy use it as logger and/or as Context
- Common/separate settings for each Logger instance
- Auto/manual increment for "transaction" like based on prefix
- Supported output formats plain and json
- Colored output for plain logs (journalctl -a)

##### Examples
- Separate Logger instance with cancel (cancel() freeing Logger recources):
```go
import "github.com/microgiantya/logger"
...
ctx, cancel := logger.NewLoggerWithCancel(&LoggerInitParams{
	Severity: 7,
})
defer cancel()
```

- Separate Logger instance (Close() freeing Logger recources):
```go
import "github.com/microgiantya/logger"
...
ctx := logger.NewLogger(&LoggerInitParams{
	Severity: 7,
})
defer ctx.Close()
```

- Plain output:
```go
...
ctx := NewLogger(&LoggerInitParams{
		Severity:       7,
		UniqueIDPrefix: "apiV1",
})
ctx.LogEmergency("message")
ctx.LogAlert("message")
ctx.LogCritical("message")
ctx.LogError("message")

ctx.IncrementUniqueID()

ctx.LogWarning("message")
ctx.LogNotice("message")
ctx.LogInformational("message")
ctx.LogDebug("message")
```

```
v0.0.0     EMERGENCY [apiV1-00000000] default_test.go:83 message
v0.0.0         ALERT [apiV1-00000000] default_test.go:84 message
v0.0.0      CRITICAL [apiV1-00000000] default_test.go:85 message
v0.0.0         ERROR [apiV1-00000000] default_test.go:86 message
v0.0.0       WARNING [apiV1-00000001] default_test.go:90 message
v0.0.0        NOTICE [apiV1-00000001] default_test.go:91 message
v0.0.0 INFORMATIONAL [apiV1-00000001] default_test.go:92 message
v0.0.0         DEBUG [apiV1-00000001] default_test.go:93 message
```

- JSON output:
```go
...
ctx := NewLogger(&LoggerInitParams{
	Severity:       7,
	UniqueIDPrefix: "apiV2",
	View:           LoggerViewJSON,
})
ctx.LogEmergency("message")
ctx.LogAlert("message")
ctx.LogCritical("message")
ctx.LogError("message")

ctx.IncrementUniqueID()

ctx.LogWarning("message")
ctx.LogNotice("message")
ctx.LogInformational("message")
ctx.LogDebug("message")
```

```
{"version":"v0.0.0","severity":"EMERGENCY","unique_id":"apiV2-00000000","file_name":"default_test.go","file_line_num":"84","message":"message"}
{"version":"v0.0.0","severity":"ALERT","unique_id":"apiV2-00000000","file_name":"default_test.go","file_line_num":"85","message":"message"}
{"version":"v0.0.0","severity":"CRITICAL","unique_id":"apiV2-00000000","file_name":"default_test.go","file_line_num":"86","message":"message"}
{"version":"v0.0.0","severity":"ERROR","unique_id":"apiV2-00000000","file_name":"default_test.go","file_line_num":"87","message":"message"}
{"version":"v0.0.0","severity":"WARNING","unique_id":"apiV2-00000001","file_name":"default_test.go","file_line_num":"91","message":"message"}
{"version":"v0.0.0","severity":"NOTICE","unique_id":"apiV2-00000001","file_name":"default_test.go","file_line_num":"92","message":"message"}
{"version":"v0.0.0","severity":"INFORMATIONAL","unique_id":"apiV2-00000001","file_name":"default_test.go","file_line_num":"93","message":"message"}
{"version":"v0.0.0","severity":"DEBUG","unique_id":"apiV2-00000001","file_name":"default_test.go","file_line_num":"94","message":"message"}
```

##### License

MIT

