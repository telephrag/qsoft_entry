package when

import "github.com/gin-gonic/gin"

// Runs service on router `r` at `addr`
func Run(r *gin.Engine, addr string) error {
	r.Use(Logger())
	r.Use(PingPong)
	r.GET("/when/:year", Endpoint)
	return r.Run(addr)
}
