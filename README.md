###### Warning - antipattern content!!!

# Ltx - _another go logger package_
Ltx package was inspired by [Asterisk PBX](https://github.com/asterisk/asterisk) logger.

##### Features
- Ltx extends [context.Context](https://pkg.go.dev/context#Context) implementation, which allow easy use it as logger and/or as [context.Context](https://pkg.go.dev/context#Context)
- Common/separate settings for each Ltx instance
- Auto/manual increment for "transaction" like based on prefix
- Supported output formats plain and json
- Colored output for plain logs (journalctl -a)

##### Examples // TODO move to examples folder
- Separate Ltx instance with cancel (cancel() freeing Ltx recources):
```go
import "github.com/microgiantya/ltx"
...
ltx, cancel := ltx.NewLtxWithCancel(&LtxInitParams{
	Severity: 7,
})
defer cancel()
```

- Separate Ltx instance (Close() freeing Ltx resources):
```go
import "github.com/microgiantya/ltx"
...
ltx := ltx.NewLtx(&LtxInitParams{
	Severity: 7,
})
defer ltx.Close()
```

- Plain output:
```go
...
ltx := NewLtx(&LtxInitParams{
		Severity:       7,
		UniqueIDPrefix: "apiV1",
})
ltx.LogEmergency("message")
ltx.LogAlert("message")
ltx.LogCritical("message")
ltx.LogError("message")

ltx.IncrementUniqueID()

ltx.LogWarning("message")
ltx.LogNotice("message")
ltx.LogInformational("message")
ltx.LogDebug("message")
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
ltx := NewLtx(&LtxInitParams{
	Severity:       7,
	UniqueIDPrefix: "apiV2",
	View:           LoggerViewJSON,
})
ltx.LogEmergency("message")
ltx.LogAlert("message")
ltx.LogCritical("message")
ltx.LogError("message")

ltx.IncrementUniqueID()

ltx.LogWarning("message")
ltx.LogNotice("message")
ltx.LogInformational("message")
ltx.LogDebug("message")
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

