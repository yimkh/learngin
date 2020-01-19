package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{ // 返回一个JSON，状态码是200，gin.H是map[string]interface{}的简写
			"message": "pong",
		})
	})
	router.Run() // 启动服务，并默认监听8080端口

	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("myfile") //获取文件
		filename := file.Filename
		size := file.Size
		header := file.Header
		c.JSON(200, gin.H{
			"filename": filename,
			"size":     size,
			"header":   header,
		})
	})
	router.Run()
}
