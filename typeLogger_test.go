package logger

/*
func TestLogCommon(t *testing.T) {
	ff := []float64{-1, 0, 1, 3, 4, 5, 6, 7, 8}
	for _, f := range ff {
		for _, s := range []_severity{emergencySeverity, alertSeverity, criticalSeverity, errorSeverity, warningSeverity, noticeSeverity, informationalSeverity, debugSeverity} {
			ctx, cancel := context.WithCancel(context.Background())
			Initialization(ctx, &loggerInitializationParams{f})
			{
				expected := fmt.Sprintf("v0.0.0 %s%s logger_test.go(18): message%s", logStaff[s].color, logStaff[s].severity, resetColor)
				actual := logCommon(logData{
					fileName: "logger_test.go",
					line:     "18",
					severity: s,
					uniqueID: "",
					data:     "message",
				})
				assert.Equal(t, expected, actual)
			}
			{
				expected := fmt.Sprintf("v0.0.0 %s%s [test] logger_test.go(27): message%s", logStaff[s].color, logStaff[s].severity, resetColor)
				actual := logCommon(logData{
					fileName: "logger_test.go",
					line:     "27",
					severity: s,
					uniqueID: "test",
					data:     "message",
				})
				assert.Equal(t, expected, actual)
			}
			cancel()
			time.Sleep(time.Microsecond * 200)
		}
	}
}
*/
