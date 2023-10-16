package demo2

import (
	"errors"
	"fmt"
	"github.com/hprose/hprose-golang/v3/rpc"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/log"
	"github.com/hprose/hprose-golang/v3/rpc/socket"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
	"time"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func TestAddNetRPCMethodsServer(t *testing.T) {
	service := rpc.NewService()
	service.Codec = rpc.NewServiceCodec(rpc.WithDebug(true))
	service.AddNetRPCMethods(new(Arith), "Arith")
	server, err := net.Listen("tcp", "127.0.0.1:8412")
	assert.NoError(t, err)
	err = service.Bind(server)
	assert.NoError(t, err)

	handler := service.GetHandler("socket")
	socketHandler := handler.(*socket.Handler)
	socketHandler.OnError = func(con net.Conn, err error) {
		fmt.Println("server OnError")
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	socketHandler.OnAccept = func(conn net.Conn) net.Conn {
		fmt.Printf("client accept :%s\n", conn.RemoteAddr().String())
		return conn
	}
	socketHandler.OnClose = func(conn net.Conn) {
		fmt.Println("server OnClose")
	}

	time.Sleep(time.Second * 86400)
	server.Close()
}

func TestAddNetRPCMethodsClient(t *testing.T) {
	client := rpc.NewClient("tcp://127.0.0.1/")
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
}
