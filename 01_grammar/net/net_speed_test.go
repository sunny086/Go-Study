package net

import (
	"bufio"
	"fmt"
	"github.com/shirou/gopsutil/net"
	"os"
	"strconv"
	"strings"
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

// 这个也行
func TestNetSpeed4(t *testing.T) {
	interfaces := []string{"enp1s0", "enp2s0", "enp3s0", "enp4s0", "enp5s0"}

	for _, iface := range interfaces {
		rxBytes, txBytes, err := getInterfaceBytes(iface)
		if err != nil {
			fmt.Printf("Error getting bytes for %s: %s\n", iface, err)
			continue
		}

		fmt.Printf("%s: RX %s/s, TX %s/s\n", iface, formatBytes(rxBytes), formatBytes(txBytes))
	}
}

func getInterfaceBytes(iface string) (uint64, uint64, error) {
	path := fmt.Sprintf("/sys/class/net/%s/statistics/", iface)
	rxFile, err := os.Open(path + "rx_bytes")
	if err != nil {
		return 0, 0, err
	}
	defer rxFile.Close()

	txFile, err := os.Open(path + "tx_bytes")
	if err != nil {
		return 0, 0, err
	}
	defer txFile.Close()

	rxScanner := bufio.NewScanner(rxFile)
	if !rxScanner.Scan() {
		return 0, 0, fmt.Errorf("No data for rx_bytes")
	}
	rxBytes, err := strconv.ParseUint(strings.TrimSpace(rxScanner.Text()), 10, 64)
	if err != nil {
		return 0, 0, err
	}

	txScanner := bufio.NewScanner(txFile)
	if !txScanner.Scan() {
		return 0, 0, fmt.Errorf("No data for tx_bytes")
	}
	txBytes, err := strconv.ParseUint(strings.TrimSpace(txScanner.Text()), 10, 64)
	if err != nil {
		return 0, 0, err
	}

	return rxBytes, txBytes, nil
}

func formatBytes(bytes uint64) string {
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}
	size := float64(bytes)
	for _, unit := range units {
		if size < 1024 {
			return fmt.Sprintf("%.2f %s", size, unit)
		}
		size /= 1024
	}
	return fmt.Sprintf("%.2f %s", size, "YB")
}
