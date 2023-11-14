package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

type test struct {
	A string
	B any
	C interface{}
}

func TestBaseUse(t *testing.T) {
	var test test
	test.A = "123"
	test.B = "123"
	test.C = "123"

	r := gin.Default() //Default返回一个默认路由引擎
	r.GET("/hello", func(c *gin.Context) {
		username := c.Query("username")
		fmt.Println("username:", username)
		c.JSON(200, gin.H{
			"msg": "hello world",
			"a":   test.A,
			"b":   test.B,
			"c":   test.C,
		})
	})
	//		http://127.0.0.1:8080/?username=123
	r.Run()
}
