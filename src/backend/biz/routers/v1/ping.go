package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {
		"message": "v1/pong",
	})
}

func Ping2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {
		"message": "v1/ping-pong",
	})
}

func Ping3(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {
		"message": "v1/home/pong",
	})
}

