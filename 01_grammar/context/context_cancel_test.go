package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestCancelContext(t *testing.T) {
	stopCtx, cancelFunc := context.WithCancel(context.Background())

	// 启动一个长时间运行的协程
	go longRunningTask(stopCtx)

	// 等待 5 秒后发送取消信号
	time.Sleep(5 * time.Second)
	cancelFunc()

	// 等待一段时间以确保协程有机会响应取消信号并退出
	time.Sleep(1 * time.Second)
	fmt.Println("Main function finished")
}

func longRunningTask(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			// 当 ctx 被取消时，ctx.Done() 会返回一个关闭的通道
			// 我们可以监听这个通道来检测取消信号
			fmt.Println("Task stopped by cancel signal")
			return
		default:
			// 否则，继续执行任务的逻辑
			fmt.Println("Task is running...")
			time.Sleep(1 * time.Second)
		}
	}
}
