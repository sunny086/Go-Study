package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main() {

	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		request, _ := http.NewRequest("GET", "http://api.uomg.com/api/rand.img3", nil)
		client := http.Client{}
		response, _ := client.Do(request)
		defer response.Body.Close()
		fmt.Println(response)
		all, _ := ioutil.ReadAll(response.Body)
		//强制浏览器下载
		//context.Header("Content-Disposition", "attachment; filename="+url.QueryEscape("operate_log.xlsx"))
		context.Data(http.StatusOK, "application/octet-stream", all)
	})
	//r.POST("/xxxpost", getting)
	//r.PUT("/xxxput")
	//监听端口默认为8080
	r.Run(":9000")

}
