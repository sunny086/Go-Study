package demo5

import (
	"errors"
	"github.com/hprose/hprose-golang/v3/rpc"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/log"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/reverse"
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

var (
	service  = rpc.NewService()
	caller   *reverse.Caller
	provider *reverse.Provider
)

func TestReverseInvokeServer(t *testing.T) {
	caller = reverse.NewCaller(service)
	server, err := net.Listen("tcp", "127.0.0.1:8412")
	assert.NoError(t, err)
	err = service.Bind(server)
	assert.NoError(t, err)
	time.Sleep(86400 * time.Second)
}

func TestReverseInvokeClient(t *testing.T) {
	client := rpc.NewClient("tcp://127.0.0.1/")
	client.Use(log.Plugin)
	provider = reverse.NewProvider(client, "1")
	provider.Debug = true
	provider.AddFunction(func(name string) string {
		return "hello " + name
	}, "hello")
	go provider.Listen()
	var proxy struct {
		Hello func(name string) (string, error)
	}
	caller.UseService(&proxy, "1")
	result, err := proxy.Hello("world")
	assert.Equal(t, "hello world", result)
	assert.NoError(t, err)

	time.Sleep(86400 * time.Second)
}
