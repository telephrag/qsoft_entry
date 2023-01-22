package when

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Endpoint(ctx *gin.Context) {
	yearStr := ctx.Param("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		err := fmt.Errorf("%w: %s", ErrYearNotAnInteger, yearStr)

		ctx.Error(&gin.Error{
			Err:  err,
			Type: gin.ErrorTypePublic,
			Meta: "when.Endpoint:YearNotAnInteger",
		})
		ctx.String(http.StatusBadRequest, err.Error()+"\n")
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
	}
	b := days >> 63 // fast abs
	days = (days ^ b) - b
	ctx.String(http.StatusOK, respFmt, days)
}
