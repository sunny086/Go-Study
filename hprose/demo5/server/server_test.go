package server

import (
	"errors"
	"fmt"
	"github.com/hprose/hprose-golang/v3/rpc"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/log"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/reverse"
	"github.com/hprose/hprose-golang/v3/rpc/socket"
	cmap "github.com/orcaman/concurrent-map"
	"net"
	"strconv"
	"testing"
	"time"
)

type SyncTaskProxy struct {
	SyncTask func(task Task) (Task, error)
}
type AsyncTaskProxy struct {
	AsyncTask func(task Task) error
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

type Task struct {
}

func (t *Task) GetTask(param Param) (result string, err error) {
	if param.Status == -1 {
		return "", errors.New("STATUS ERROR")
	}
	result = param.Message + strconv.Itoa(param.Id) + strconv.Itoa(param.Status)
	return
}

var (
	protocol           = "tcp"
	port               = 8412
	RemoteService      = rpc.NewService()
	Caller             *reverse.Caller
	ReverseCaller      *reverse.Caller
	RemoteServiceCache cmap.ConcurrentMap
)

func TestServer(t *testing.T) {
	Caller = reverse.NewCaller(RemoteService)
	RemoteService.Codec = rpc.NewServiceCodec(rpc.WithDebug(true))
	RemoteService.AddAllMethods(new(Task), "task")
	var server net.Listener
	var err error
	server, err = net.Listen(protocol, fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		t.Log(err)
	}
	err = RemoteService.Bind(server)

	RemoteServiceCache = cmap.New()
	handler := RemoteService.GetHandler("socket")
	socketHandler := handler.(*socket.Handler)
	socketHandler.OnError = func(con net.Conn, err error) {
		fmt.Println("server OnError")
	}
	socketHandler.OnAccept = func(conn net.Conn) net.Conn {
		fmt.Printf("client accept :%s\n", conn.RemoteAddr().String())
		return conn
	}
	socketHandler.OnClose = func(conn net.Conn) {
		fmt.Println("server OnClose")
	}

	var syncTaskProxy = SyncTaskProxy{}
	var asyncTaskProxy = AsyncTaskProxy{}
	var stateProxy = StateProxy{}
	Caller.UseService(&syncTaskProxy, "uuid")
	Caller.UseService(&asyncTaskProxy, "uuid")
	Caller.UseService(&stateProxy, "uuid")

	time.Sleep(86400 * time.Second)
}

func TestClient(t *testing.T) {
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
		GetTask func(param Param) (result string, err error)
	}
	client.UseService(&proxy)
	proxy.GetTask(Param{Id: 1, Status: 1, Message: "hello"})
	time.Sleep(86400 * time.Second)
}
