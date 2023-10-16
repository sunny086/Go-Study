package client

import (
	"errors"
	"fmt"
	"github.com/hprose/hprose-golang/v3/rpc"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/log"
	"github.com/stretchr/testify/assert"
	"net"
	"strconv"
	"testing"
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

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func TestClient(t *testing.T) {
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
	{
		result, err := proxy.Divide(Args{3, 2})
		assert.Equal(t, Quotient{1, 1}, result)
		assert.NoError(t, err)
	}
	{
		_, err := proxy.Divide(Args{3, 0})
		assert.EqualError(t, err, "divide by zero")
	}
	//time.Sleep(86400 * time.Second)
}
