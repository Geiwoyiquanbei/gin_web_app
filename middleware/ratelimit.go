package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

func RareLimitMiddleware(duration time.Duration, capacity int64) func(con *gin.Context) {
	bucket := ratelimit.NewBucket(duration, capacity)
	return func(con *gin.Context) {
		if bucket.TakeAvailable(2) == 2 {
			con.Next()
		}
		con.String(http.StatusOK, "limit rate...")
		con.Abort()
		return
	}
}
