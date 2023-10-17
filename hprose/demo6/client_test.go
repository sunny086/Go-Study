package demo6

import (
	"github.com/hprose/hprose-golang/v3/rpc"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/log"
	"sync"
	"testing"
)

func TestClient1(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	client := rpc.NewClient("tcp://127.0.0.1/")
	client.Use(log.Plugin)
	var proxy struct {
		Hello func(name string) string
	}
	client.UseService(&proxy)
	t.Log(proxy.Hello("client1"))
	wg.Wait()
}

func TestClient2(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	client := rpc.NewClient("tcp://127.0.0.1/")
	client.Use(log.Plugin)
	var proxy struct {
		Hello func(name string) string
	}
	client.UseService(&proxy)
	t.Log(proxy.Hello("client2"))
	wg.Wait()
}

func TestClient3(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	client := rpc.NewClient("tcp://127.0.0.1/")
	client.Use(log.Plugin)
	var proxy struct {
		Hello func(name string) string
	}
	client.UseService(&proxy)
	t.Log(proxy.Hello("client3"))
	wg.Wait()
}

func TestClient4(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	client := rpc.NewClient("tcp://127.0.0.1/")
	client.Use(log.Plugin)
	var proxy struct {
		Hello func(name string) string
	}
	client.UseService(&proxy)
	t.Log(proxy.Hello("client4"))
	wg.Wait()
}

func TestClient5(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	client := rpc.NewClient("tcp://127.0.0.1/")
	client.Use(log.Plugin)
	var proxy struct {
		Hello func(name string) string
	}
	client.UseService(&proxy)
	t.Log(proxy.Hello("client5"))
	wg.Wait()
}

func TestClient6(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	client := rpc.NewClient("tcp://127.0.0.1/")
	client.Use(log.Plugin)
	var proxy struct {
		Hello func(name string) string
	}
	client.UseService(&proxy)
	t.Log(proxy.Hello("client6"))
	wg.Wait()
}

func TestClient7(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	client := rpc.NewClient("tcp://127.0.0.1/")
	client.Use(log.Plugin)
	var proxy struct {
		Hello func(name string) string
	}
	client.UseService(&proxy)
	t.Log(proxy.Hello("client7"))
	wg.Wait()
}

func TestClient8(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	client := rpc.NewClient("tcp://127.0.0.1/")
	client.Use(log.Plugin)
	var proxy struct {
		Hello func(name string) string
	}
	client.UseService(&proxy)
	t.Log(proxy.Hello("client8"))
	wg.Wait()
}

func TestClient9(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	client := rpc.NewClient("tcp://127.0.0.1/")
	client.Use(log.Plugin)
	var proxy struct {
		Hello func(name string) string
	}
	client.UseService(&proxy)
	t.Log(proxy.Hello("client9"))
	wg.Wait()
}
