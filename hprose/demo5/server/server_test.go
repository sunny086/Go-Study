package server

import (
	"errors"
	"fmt"
	"github.com/hprose/hprose-golang/v3/rpc"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/reverse"
	"github.com/hprose/hprose-golang/v3/rpc/socket"
	cmap "github.com/orcaman/concurrent-map"
	"github.com/stretchr/testify/assert"
	"net"
	"sync"
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

//func (t *Arith) Multiply(args *Args, reply *int) error {
//	*reply = args.A * args.B
//	return nil
//}
//
//func (t *Arith) Divide(args *Args, quo *Quotient) error {
//	if args.B == 0 {
//		return errors.New("divide by zero")
//	}
//	quo.Quo = args.A / args.B
//	quo.Rem = args.A % args.B
//	return nil
//}

func (t *Arith) Multiply(args *Args) (int, error) {
	return args.A * args.B, nil
}

func (t *Arith) Divide(args *Args) (Quotient, error) {
	if args.B == 0 {
		return Quotient{}, errors.New("divide by zero")
	}
	return Quotient{Quo: args.A / args.B, Rem: args.A % args.B}, nil
}

var (
	Caller             *reverse.Caller
	ReverseCaller      *reverse.Caller
	RemoteServiceCache cmap.ConcurrentMap
	wg                 sync.WaitGroup
)

func TestServer(t *testing.T) {
	wg.Add(1)
	RemoteService := rpc.NewService()
	RemoteService.Codec = rpc.NewServiceCodec(rpc.WithDebug(true))
	RemoteService.AddAllMethods(new(Arith), "Arith")
	server, err := net.Listen("tcp", "127.0.0.1:8412")
	assert.NoError(t, err)
	err = RemoteService.Bind(server)
	Caller = reverse.NewCaller(RemoteService)

	RemoteServiceCache = cmap.New()
	handler := RemoteService.GetHandler("socket")
	socketHandler := handler.(*socket.Handler)
	socketHandler.OnError = func(con net.Conn, err error) {
		fmt.Println("server OnError")
		assert.NoError(t, err)
	}
	socketHandler.OnAccept = func(conn net.Conn) net.Conn {
		fmt.Printf("client accept :%s\n", conn.RemoteAddr().String())
		fmt.Println(conn.LocalAddr().String() + "->" + conn.RemoteAddr().String() + " closed on client")
		go ProxyCacheInit()
		return conn
	}
	socketHandler.OnClose = func(conn net.Conn) {
		fmt.Println("server OnClose")
	}
	wg.Wait()
}

func ProxyCacheInit() {
	var syncTaskProxy = SyncTaskProxy{}
	var asyncTaskProxy = AsyncTaskProxy{}
	var stateProxy = StateProxy{}
	Caller.UseService(&syncTaskProxy, "1")
	Caller.UseService(&asyncTaskProxy, "1")
	Caller.UseService(&stateProxy, "1")
	proxyCache := ProxyCache{}
	proxyCache.SyncTaskProxy = &syncTaskProxy
	proxyCache.AsyncTaskProxy = &asyncTaskProxy
	proxyCache.StateTaskProxy = &stateProxy
	RemoteServiceCache.Set("proxyCache", proxyCache)
	ExecuteTest()
}

func ExecuteTest() {
	//for循环测试长连接是否只构建一次
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("start execute test")
		proxy, _ := RemoteServiceCache.Get("proxyCache")
		proxyCache := proxy.(ProxyCache)
		msg, err := proxyCache.SyncTaskProxy.SyncTask(Param{Status: 1, Message: "hello client ~~~"})
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("client返回值：" + msg)
	}
}
