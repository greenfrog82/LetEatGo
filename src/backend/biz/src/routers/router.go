package routers

import (
	"github.com/gin-gonic/gin"
)

import v1 "routers/v1"
import v2 "routers/v2"

func InitRouter(router *gin.Engine) {
	rv1 := router.Group("/v1") 
	{
		rv1.GET("/ping", v1.Ping)	
		rv1.GET("/ping-pong", v1.Ping2)
	}

	home := rv1.Group("/home") 
	{ 
		home.GET("/ping", v1.Ping3) 
	}
	
	rv2 := router.Group("/v2") 
	{
		rv2.GET("/ping", v2.Ping)
	}
}