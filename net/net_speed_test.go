package net

import (
	"fmt"
	"github.com/shirou/gopsutil/net"
	"testing"
	"time"
)

// 这个不对 有问题
func TestNetSpeed(t *testing.T) {
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
