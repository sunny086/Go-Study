package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/disk"
	"os"
	"runtime"
	"strconv"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func main() {
	println(`系统类型：`, runtime.GOOS)

	println(`系统架构：`, runtime.GOARCH)

	println(`CPU 核数：`, runtime.GOMAXPROCS(0))

	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	println(`电脑名称：`, name)

	rate := runtime.MemProfileRate

	fmt.Println(rate)

	d, _ := disk.Usage("/")
	var diskTotal, diskUsed, diskUsedPercent float64
	diskTotal = float64(d.Total / GB)
	diskUsed = float64(d.Used / GB)
	diskUsedPercent, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", d.UsedPercent), 64)
	fmt.Println(diskTotal)
	fmt.Println(diskUsed)
	fmt.Println(diskUsedPercent)
}
