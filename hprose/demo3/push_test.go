package demo3

import (
	"context"
	"fmt"
	"github.com/hprose/hprose-golang/v3/rpc"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/log"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/push"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
	"time"
)

func TestPush(t *testing.T) {
	service := push.NewBroker(rpc.NewService())
	server, err := net.Listen("tcp", "127.0.0.1:8412")
	assert.NoError(t, err)
	err = service.Bind(server)
	assert.NoError(t, err)
	service.AddFunction(func(ctx context.Context, name string) string {
		serviceContext := rpc.GetServiceContext(ctx)
		producer := serviceContext.Items().GetInterface("producer").(push.Producer)
		producer.Push("ooxx", "test")
		return "hello " + name
	}, "hello")
	time.Sleep(time.Millisecond * 5)

	client1 := rpc.NewClient("tcp://127.0.0.1/")
	client1.Use(log.Plugin.IOHandler)
	prosumer1 := push.NewProsumer(client1, "1")
	prosumer1.OnError = func(e error) {
		fmt.Println(e.Error())
	}
	prosumer1.OnSubscribe = func(topic string) {
		fmt.Println(topic, "is subscribed.")
	}
	prosumer1.OnUnsubscribe = func(topic string) {
		fmt.Println(topic, "is unsubscribed.")
	}
	client2 := rpc.NewClient("tcp://127.0.0.1/")
	//client2.Use(log.Plugin.IOHandler)
	prosumer2 := push.NewProsumer(client2, "2")
	prosumer1.Subscribe("test", func(data int, from string) {
		fmt.Printf("%v from %v\n", data, from)
	})
	prosumer1.Subscribe("test2", func(message push.Message) {
		fmt.Println(message)
	})
	time.Sleep(time.Millisecond * 100)
	client1.Invoke("hello", []interface{}{"world"})
	prosumer2.Push(1, "test", "1")
	// var wg sync.WaitGroup
	// n := 1000
	// wg.Add(n)
	// for i := 0; i < n; i++ {
	// 	go func(i int) {
	// 		prosumer2.Push(i, "test", "1")
	// 		wg.Done()
	// 	}(i)
	// }
	// wg.Wait()
	time.Sleep(time.Millisecond * 100)

	server.Close()

	time.Sleep(time.Millisecond * 100)

	server, _ = net.Listen("tcp", "127.0.0.1:8412")
	_ = service.Bind(server)

	time.Sleep(time.Millisecond * 1000)

	prosumer2.Push(2, "test", "1")

	// wg.Add(n)
	// for i := 0; i < n; i++ {
	// 	go func(i int) {
	// 		prosumer2.Push(i, "test", "1")
	// 		wg.Done()
	// 	}(i)
	// }
	// wg.Wait()
	time.Sleep(time.Millisecond * 1000)

	prosumer1.Unsubscribe("test")
	prosumer1.Unsubscribe("test2")

	assert.NoError(t, err)
	server.Close()
}
