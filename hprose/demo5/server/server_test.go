package server

import (
	"errors"
	"fmt"
	"github.com/hprose/hprose-golang/v3/rpc"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/reverse"
	"github.com/hprose/hprose-golang/v3/rpc/socket"
	cmap "github.com/orcaman/concurrent-map"
	"net"
	"testing"
	"time"
)

type SyncTaskProxy struct {
	SyncTask func(param Param) (string, error)
}
type AsyncTaskProxy struct {
	AsyncTask func(param Param) error
}
type StateProxy struct {
	RunningState func() (string, error)
}

type ProxyCache struct {
	SyncTaskProxy  *SyncTaskProxy
	AsyncTaskProxy *AsyncTaskProxy
	StateTaskProxy *StateProxy
}

type Param struct {
	Id      int    `json:"id"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

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
	RemoteService      = rpc.NewService()
	Caller             *reverse.Caller
	ReverseCaller      *reverse.Caller
	RemoteServiceCache cmap.ConcurrentMap
)

func TestServer(t *testing.T) {
	RemoteService.AddAllMethods(new(Arith), "Arith")
	server, err := net.Listen("tcp", "127.0.0.1:8412")
	if err != nil {
		t.Log(err)
	}
	err = RemoteService.Bind(server)
	Caller = reverse.NewCaller(RemoteService)
	RemoteService.Codec = rpc.NewServiceCodec(rpc.WithDebug(true))

	RemoteServiceCache = cmap.New()
	handler := RemoteService.GetHandler("socket")
	socketHandler := handler.(*socket.Handler)
	socketHandler.OnError = func(con net.Conn, err error) {
		fmt.Println("server OnError")
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	socketHandler.OnAccept = func(conn net.Conn) net.Conn {
		fmt.Printf("client accept :%s\n", conn.RemoteAddr().String())
		fmt.Println(conn.LocalAddr().String() + "->" + conn.RemoteAddr().String() + " closed on client")
		var syncTaskProxy = SyncTaskProxy{}
		var asyncTaskProxy = AsyncTaskProxy{}
		var stateProxy = StateProxy{}
		Caller.UseService(&syncTaskProxy, "syn")
		Caller.UseService(&asyncTaskProxy, "syn")
		Caller.UseService(&stateProxy, "syn")
		proxyCache := ProxyCache{}
		proxyCache.SyncTaskProxy = &syncTaskProxy
		proxyCache.AsyncTaskProxy = &asyncTaskProxy
		proxyCache.StateTaskProxy = &stateProxy
		msg, err := proxyCache.SyncTaskProxy.SyncTask(Param{Id: 1, Status: 1, Message: "dafads"})
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(msg)
		return conn
	}
	socketHandler.OnClose = func(conn net.Conn) {
		fmt.Println("server OnClose")
	}
	time.Sleep(86400 * time.Second)
}
