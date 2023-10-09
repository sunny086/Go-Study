package gva_router

import (
	"github.com/gin-gonic/gin"
)

var hr HelloRouter

func InitRouter() *gin.Engine {
	r := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	group := r.Group("")
	hr.InitHelloRouter(group)
	return r
}
