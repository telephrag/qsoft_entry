package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const DAY_UNIX_SEC = 60 * 60 * 24

var (
	ErrYearNotAnInteger = errors.New("integer expected as a request parameter")
)

func main() {
	router := gin.New()

	router.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		logFmt := "%-30s | %-8s | %4s %3d %s %s"
		if params.ErrorMessage == "" { // ErrorMessage already contains endline symbol
			logFmt += "\n"
		}
		return fmt.Sprintf(logFmt,
			params.TimeStamp.UTC().Format("2006-01-02 15:04:05.999999999"),
			params.Latency,
			params.Method,
			params.StatusCode,
			params.Request.URL,
			params.ErrorMessage,
		)
	}))

	router.GET("/when/:year", func(ctx *gin.Context) {
		yearStr := ctx.Param("year")
		year, err := strconv.Atoi(yearStr)
		if err != nil {
			err := fmt.Errorf("%w: %s", ErrYearNotAnInteger, yearStr)

			ctx.Error(&gin.Error{
				Err:  err,
				Type: gin.ErrorTypePublic,
			})
			ctx.String(http.StatusBadRequest, err.Error())
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		when := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()
		now := time.Now().Unix()
		days := (now - when) / DAY_UNIX_SEC

		var respFmt string
		if days >= 0 {
			respFmt = fmt.Sprint("Days gone: %d\n")
		} else {
			respFmt = fmt.Sprint("Days left: %d\n")
			days *= -1 // TODO: Use a better way of taking a module of int64
		}
		ctx.String(http.StatusOK, respFmt, days)
	})

	router.Run("localhost:3000")
}
