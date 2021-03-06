package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger gin.HandlerFunc
func Logger() gin.HandlerFunc {

	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] %s %s %d %s",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC822),
			param.TimeStamp.Format(time.RFC822),
			// param.Method,
			param.Path,
			param.StatusCode,
			param.Latency)
	})

}
