package bpm

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

type request struct {
	ArrivalTIme string `json:"arrival_time"`
	SendTime    int64  `json:"send_time"`
}

type test1 struct {
	ArrivalTIme int64 `json:"arrival_time"`
	SendTime    int64 `json:"send_time"`
}

func TestMQ(t *testing.T) {
	r := gin.Default() //Default返回一个默认路由引擎
	r.GET("/hello", func(c *gin.Context) {
		var req request
		err := c.ShouldBindJSON(&req)
		if err != nil {
			fmt.Println(err.Error())
		}
		//转json string
		bytes, err := json.Marshal(req)
		if err != nil {
			fmt.Println(err.Error())
		}
		var t test1
		err = json.Unmarshal(bytes, &t)
		if err != nil {
			fmt.Println(err.Error())
		}
	})
	r.Run(":7777")
}
