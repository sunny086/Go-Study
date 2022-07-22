package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

var siganChannel = make(chan os.Signal, 1)

func main() {
	ContextWithCancel()
	//ContextWithTimeout()
	//ContextWithDeadline()
	//TimeoutContext()
}

func ContextWithCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println(ctx.Value("msg"))
				fmt.Println("kill1")
				return
			default:

			}
		}
	}()
	fmt.Println("开始")
	time.AfterFunc(5*time.Second, func() {
		ctx = context.WithValue(ctx, "msg", "10秒后调用cancel()")
		cancel()
		fmt.Println("结束")
	})
	Exit()
}

func Exit() {
	signal.Notify(siganChannel, os.Kill, os.Interrupt)
	<-siganChannel
}

func TimeoutContext() {
	// 创建一个子节点的context,3秒后自动超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	go watch(ctx, "监控1")
	go watch(ctx, "监控2")
	time.Sleep(80 * time.Second)
	cancel()
	Exit()
}

// 单独的监控协程
func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "收到信号，监控退出,time=", time.Now().Unix())
			return
		default:
			fmt.Println(name, "goroutine监控中,time=", time.Now().Unix())
			time.Sleep(1 * time.Second)
		}
	}
}
