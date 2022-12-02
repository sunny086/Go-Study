package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestBaseUse(t *testing.T) {
	r := gin.Default() //Default返回一个默认路由引擎
	r.GET("/", func(c *gin.Context) {
		username := c.Query("username")
		fmt.Println("username:", username)
		c.JSON(200, gin.H{
			"msg": "hello world",
		})
	})
	//http://127.0.0.1:8080/?username=123
	r.Run()
}
