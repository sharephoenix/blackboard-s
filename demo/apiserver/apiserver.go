package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//关闭gin debug
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	//提交生产api
	router.POST("/test", productionOrderHandler)
	//测试api是否正常访问
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.Run(":8080")
}

func productionOrderHandler(c *gin.Context) {
	message := "未知错误"
	c.String(200, message)
	return
}