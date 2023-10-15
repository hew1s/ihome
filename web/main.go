package main

import (
	"github.com/gin-gonic/gin"
	"web/controller"
)

// 添加gin框架
func main() {
	r := gin.Default()
	r.Static("/home", "view")
	r.GET("/api/v1.0/session", controller.GetSession)

	err := r.Run("127.0.0.1:9002")
	if err != nil {
		return
	}
}
