package middlewares

import (
	"github.com/gin-gonic/gin"
	"time"
	"log"
)

func InitMiddleware(router *gin.Engine) {
	router.Use(Logger())	
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}