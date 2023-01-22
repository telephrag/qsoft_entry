package when

import "github.com/gin-gonic/gin"

func PingPong(ctx *gin.Context) {
	if ping := ctx.GetHeader("X-PING"); ping == "ping" {
		ctx.Header("X-PONG", "pong")
	}

	ctx.Next()
}
