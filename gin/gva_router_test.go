package gin

import (
	"Go-Study/gin/gva_router"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
	"time"
)

//type server interface {
//	ListenAndServe() error
//}

func initServer(address string, router *gin.Engine) http.Server {
	return http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    180 * time.Second,
		WriteTimeout:   180 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func TestGvaRouter(t *testing.T) {
	engine := gva_router.InitRouter()
	server := initServer(":7071", engine)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Print(err.Error())
	}
	//engine.Run(":7071")
}
