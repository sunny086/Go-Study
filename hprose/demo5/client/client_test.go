package client

import (
	"errors"
	"fmt"
	"github.com/hprose/hprose-golang/v3/rpc"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/log"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/reverse"
	"github.com/stretchr/testify/assert"
	"net"
	"strconv"
	"sync"
	"testing"
)

type Param struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type SyncTaskProxy struct {
}

func (t *SyncTaskProxy) SyncTask(param Param) (result string, err error) {
	if param.Status == -1 {
		return "", errors.New("STATUS ERROR")
	} else if param.Status == 1 {
		result = "hello server"
	} else {
		result = param.Message + strconv.Itoa(param.Status)
	}
	return
}

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

var wg sync.WaitGroup

func TestClient(t *testing.T) {
	wg.Add(1)
	// 创建RPC客户端
	client := rpc.NewClient("tcp://127.0.0.1/")
	// 创建Socket传输
	socketTransport := rpc.SocketTransport(client)
	socketTransport.OnConnect = func(conn net.Conn) net.Conn {
		fmt.Println(conn.LocalAddr().String() + "->" + conn.RemoteAddr().String() + " connected")
		return conn
	}
	socketTransport.OnClose = func(conn net.Conn) {
		fmt.Println(conn.LocalAddr().String() + "->" + conn.RemoteAddr().String() + " closed on client")
	}
	client.Use(log.Plugin)
	var proxy struct {
		Multiply func(args Args) (int, error)
		Divide   func(args Args) (Quotient, error)
	}
	client.UseService(&proxy, "Arith")
	{
		result, err := proxy.Multiply(Args{3, 2})
		assert.Equal(t, 6, result)
		assert.NoError(t, err)
	}
	//{
	//	result, err := proxy.Divide(Args{3, 2})
	//	assert.Equal(t, Quotient{1, 1}, result)
	//	assert.NoError(t, err)
	//}
	//{
	//	_, err := proxy.Divide(Args{3, 0})
	//	assert.EqualError(t, err, "divide by zero")
	//}
	provider := reverse.NewProvider(client, "1")
	provider.Debug = true
	provider.AddAllMethods(new(SyncTaskProxy))
	provider.OnError = func(err error) {
		fmt.Println("provider OnError")
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	go provider.Listen()
	wg.Wait()
}
