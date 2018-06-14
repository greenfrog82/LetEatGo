package main

import (
	"github.com/gin-gonic/gin"
	"routers"
	"middlewares"
)

func main() {
	r := gin.Default()
	
	middlewares.InitMiddleware(r)
	routers.InitRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080
}