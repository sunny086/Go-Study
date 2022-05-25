package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//r := gin.Default()
	//r.LoadHTMLGlob("tem/*")
	//r.GET("/index", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "template.html", gin.H{"title": "我是测试", "ce": "123456"})
	//})
	//
	//r.Run()

	r := gin.Default()
	r.LoadHTMLGlob("tem/**/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user/index.html", gin.H{"title": "我是测试", "address": "www.5lmh.com"})
	})
	r.Run()
}
