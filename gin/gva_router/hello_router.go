package gva_router

import "github.com/gin-gonic/gin"

type HelloRouter struct {
}

func (s *HelloRouter) InitHelloRouter(Router *gin.RouterGroup) {
	r := Router.Group("")

	r.GET("hello", Hello())

}

func Hello() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "hello world",
		})
	}
}
