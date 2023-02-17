package net

import (
	"fmt"
	"github.com/shirou/gopsutil/net"
	"testing"
	"time"
)

// 这个不对 有问题
func TestNetSpeed1(t *testing.T) {
	lastStats := make(map[string]net.IOCountersStat)

	for {
		time.Sleep(1 * time.Second)

		netStats, err := net.IOCounters(true)
		if err != nil {
			fmt.Println("Error getting network stats:", err)
			continue
		}

		for _, stat := range netStats {
			lastStat, ok := lastStats[stat.Name]
			if !ok {
				lastStats[stat.Name] = stat
				continue
			}

			speed := (float64(stat.BytesSent-lastStat.BytesSent) / 1024) / float64(time.Second)
			fmt.Printf("Interface: %s, Speed: %.2fKB/s\n", stat.Name, speed)

			lastStats[stat.Name] = stat
		}
	}
}

// 这个有数据 但是不是我想要的 暂存一下 有时间调试
func TestNetSpeed2(t *testing.T) {
	lastCounters, _ := net.IOCounters(false)
	lastTime := time.Now()

	for {
		time.Sleep(time.Second)
		currentCounters, _ := net.IOCounters(false)
		currentTime := time.Now()

		duration := currentTime.Sub(lastTime).Seconds()
		for _, counter := range currentCounters {
			lastCounter := findCounter(lastCounters, counter.Name)
			if lastCounter == nil {
				continue
			}
			rxBytes := float64(counter.BytesRecv - lastCounter.BytesRecv)
			txBytes := float64(counter.BytesSent - lastCounter.BytesSent)
			rxRate := rxBytes / duration
			txRate := txBytes / duration
			fmt.Printf("Interface: %s, RX: %.2f KB/s, TX: %.2f KB/s\n", counter.Name, rxRate/1024, txRate/1024)
		}

		lastCounters = currentCounters
		lastTime = currentTime
	}
}

func findCounter(counters []net.IOCountersStat, name string) *net.IOCountersStat {
	for i := range counters {
		if counters[i].Name == name {
			return &counters[i]
		}
	}
	return nil
}
