package client

import (
	"errors"
	"fmt"
	"github.com/hprose/hprose-golang/v3/rpc"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/log"
	"net"
	"strconv"
	"testing"
	"time"
)

type Param struct {
	Id      int    `json:"id"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type SyncTaskProxy struct {
}

func (t *SyncTaskProxy) SyncTask(param Param) (result string, err error) {
	if param.Status == -1 {
		return "", errors.New("STATUS ERROR")
	}
	result = param.Message + strconv.Itoa(param.Id) + strconv.Itoa(param.Status)
	return
}

func TestClient(t *testing.T) {
	client := rpc.NewClient("tcp://127.0.0.1/")
	// 创建Socket传输
	socketTransport := rpc.SocketTransport(client)
	socketTransport.OnConnect = func(c net.Conn) net.Conn {
		fmt.Println(c.LocalAddr().String() + "->" + c.RemoteAddr().String() + " connected")
		return c
	}
	socketTransport.OnClose = func(c net.Conn) {
		fmt.Println(c.LocalAddr().String() + "->" + c.RemoteAddr().String() + " closed on client")
	}
	client.Use(log.Plugin)

	client.UseService(new(SyncTaskProxy), "syn")

	var proxy struct {
		GetTask func(param Param) (result string, err error)
	}
	client.UseService(&proxy, "task")
	result, err := proxy.GetTask(Param{Id: 1, Status: 1, Message: "hello"})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result)
	time.Sleep(86400 * time.Second)
}
