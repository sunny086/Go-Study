package demo4

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hprose/hprose-golang/v3/rpc"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/log"
	"github.com/stretchr/testify/assert"
	"net"
	"reflect"
	"testing"
	"time"
)

func TestTcp(t *testing.T) {
	service := rpc.NewService()
	service.AddMissingMethod(func(ctx context.Context, name string, args []interface{}) (result []interface{}, err error) {
		serviceContext := rpc.GetServiceContext(ctx)
		data, err := json.Marshal(args)
		if err != nil {
			return nil, err
		}
		return []interface{}{name + string(data) + serviceContext.RemoteAddr.String()}, nil
	})
	method := service.Get("*")
	assert.Equal(t, reflect.Func, method.Func().Kind())
	assert.Equal(t, []reflect.Type{reflect.TypeOf(""), reflect.TypeOf([]interface{}{})}, method.Parameters())
	assert.True(t, method.ReturnError())
	assert.Nil(t, method.Options())
	socketHandler := rpc.SocketHandler(service)
	socketHandler.OnAccept = func(c net.Conn) net.Conn {
		fmt.Println(c.RemoteAddr().String() + "->" + c.LocalAddr().String() + " accepted")
		return c
	}
	socketHandler.OnClose = func(c net.Conn) {
		fmt.Println(c.RemoteAddr().String() + "->" + c.LocalAddr().String() + " closed on server")
	}
	socketHandler.OnError = func(c net.Conn, e error) {
		if c != nil {
			fmt.Println(c.RemoteAddr().String()+"->"+c.LocalAddr().String(), e)
		} else {
			fmt.Println(e)
		}
	}
	server, err := net.Listen("tcp", "127.0.0.1:8412")
	assert.NoError(t, err)
	err = service.Bind(server)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 5)

	client := rpc.NewClient("tcp://127.0.0.1/")
	socketTransport := rpc.SocketTransport(client)
	socketTransport.OnConnect = func(c net.Conn) net.Conn {
		fmt.Println(c.LocalAddr().String() + "->" + c.RemoteAddr().String() + " connected")
		return c
	}
	socketTransport.OnClose = func(c net.Conn) {
		fmt.Println(c.LocalAddr().String() + "->" + c.RemoteAddr().String() + " closed on client")
	}
	client.Use(log.Plugin)
	var proxy struct {
		Hello func(name string) string
	}
	client.UseService(&proxy)
	proxy.Hello("world")
	server.Close()
}

type HelloService struct{}

func (h *HelloService) SayHello(name string) string {
	fmt.Printf("Received a request from client: %s\n", name)
	return "Hello, " + name
}

func TestClient(t *testing.T) {

}

func TestServer(t *testing.T) {

}
