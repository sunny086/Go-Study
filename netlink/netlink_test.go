package netlink

import (
	"fmt"
	"github.com/vishvananda/netlink"
	"testing"
	"time"
)

// 就这个吧 貌似可以 但是只有linux可以编译启动
func TestNetSpeedByNetlink(t *testing.T) {
	link, err := netlink.LinkByName("enp1s0")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 获取收包速率、发包速率和实时速率
	lastRx := link.Attrs().Statistics.RxBytes
	lastTx := link.Attrs().Statistics.TxBytes

	time.Sleep(time.Second)
	link2, err := netlink.LinkByName("enp1s0")
	if err != nil {
		fmt.Println(err)
	}
	rx := link2.Attrs().Statistics.RxBytes
	tx := link2.Attrs().Statistics.TxBytes
	rxRate := float64(rx-lastRx) / 1024
	txRate := float64(tx-lastTx) / 1024
	lastRx = rx
	lastTx = tx
	fmt.Printf("收包速率：%.2f KB/s, 发包速率：%.2f KB/s\n", rxRate, txRate)
}
