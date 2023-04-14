package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		startTime := time.Now()               // starting time request
		ctx.Next()                            // process request
		endTime := time.Now()                 // end time reqeust
		latencyTime := endTime.Sub(startTime) // calculate execution time
		requestMethod := ctx.Request.Method   // request method
		requestURI := ctx.Request.RequestURI  // request route
		statusCode := ctx.Writer.Status()     // status code
		clientIP := ctx.ClientIP()            // client ID

		log.WithFields(log.Fields{
			"latency_time":   latencyTime,
			"request_method": requestMethod,
			"req_uri":        requestURI,
			"status_code":    statusCode,
			"client_ip":      clientIP,
		}).Info("http request")

		ctx.Next()

	}
}
