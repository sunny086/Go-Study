package demo6

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/hprose/hprose-golang/v3/rpc"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/log"
	"sync"
	"testing"
	"time"
)

type AgentInfo struct {
	Uuid       string  `json:"uuid"`       // 设备uuid
	Ip         string  `json:"ip"`         // 设备ip
	CpuNum     int     `json:"cpuNum"`     // cpu核数
	MemSize    int     `json:"memSize"`    // 内存大小
	DiskSize   int     `json:"diskSize"`   // 磁盘大小
	CpuUsage   float64 `json:"cpuUsage"`   // cpu使用率
	MemUsage   float64 `json:"memUsage"`   // 内存使用率
	DiskUsage  float64 `json:"diskUsage"`  // 磁盘使用率
	InputFlow  float64 `json:"inputFlow"`  // 输入流量
	OutputFlow float64 `json:"outputFlow"` // 输出流量
}

type AgentInfoRpcFunc struct {
	SendAgentInfo func(info AgentInfo) error
}

func TestAgentInfoClient1(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	//每间隔五秒发送一次
	// 创建一个定时器，每隔五秒触发一次
	ticker := time.NewTicker(5 * time.Second)

	// 启动一个无限循环，等待定时器事件
	for {
		select {
		case <-ticker.C:
			SendAgentInfo()
		}
	}
	wg.Wait()
}

func SendAgentInfo() error {
	client := rpc.NewClient("tcp://127.0.0.1/")
	client.Use(log.Plugin)
	var proxy struct {
		SendAgentInfo func(info AgentInfo) error
	}
	client.UseService(&proxy, "AgentInfoRpcFunc")
	var info = AgentInfo{
		Uuid:       "client1",
		Ip:         "127.0.0.1",
		CpuNum:     1,
		MemSize:    1024,
		DiskSize:   1024,
		CpuUsage:   0.1,
		MemUsage:   0.1,
		DiskUsage:  0.1,
		InputFlow:  0.1,
		OutputFlow: 0.1,
	}
	err := proxy.SendAgentInfo(info)
	return err
}

func TestReadStream(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "10.25.10.216:6379",
		Password: "Netvine123#@!",
		DB:       4,
	})
	result, err := client.XInfoStream(context.Background(), "agent:device:client1").Result()
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
	streams, err := client.XReadStreams(context.Background(), "agent:device:client1", "0").Result()
	if err != nil {
		t.Error(err)
	}
	t.Log(streams)
	var agentList []AgentInfo
	for _, message := range streams {
		for _, xMessage := range message.Messages {
			fmt.Println(xMessage.ID)
			fmt.Println(xMessage.Values)
		}
	}
	t.Log(agentList)

}
