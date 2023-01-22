package when

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Logs activity of `when` service in pretty format and w/ errors
func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		logFmt := "%-30s | %-8s | %4s %3d %s %s"
		if params.ErrorMessage == "" { // ErrorMessage already contains endline symbol
			logFmt += "\n"
		}
		return fmt.Sprintf(logFmt,
			params.TimeStamp.UTC().Format("2006-01-02 15:04:05.999999999"), // timestamps are assumed to be in UTC across entire app
			params.Latency,
			params.Method,
			params.StatusCode,
			params.Request.URL,
			params.ErrorMessage,
		)
	})
}
