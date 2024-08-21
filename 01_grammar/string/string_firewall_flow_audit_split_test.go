package string

import (
	"math"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestFlowAuditSplit(t *testing.T) {
	message := "25;2023-03-13 14:26:12;59986;22;10.25.16.6;a4:1a:3a:91:a4:1a;273;0----\n"
	//message := "0;2022-12-28 10:15:20;14336;12288;172.21.1.254;c8:5b:76:ef:ad:ee;67;0-------"
	messageContent := strings.Split(message, "----")
	split := strings.Split(messageContent[0], ";")
	deviceIp := split[4]
	deviceMac := split[5]
	//异常流量丢弃
	if len(deviceMac) > 0 && deviceIp == "NULL" {
		return
	}
	srcPort, _ := strconv.Atoi(split[2])
	destPort, _ := strconv.Atoi(split[3])
	sendSpeed, _ := strconv.Atoi(split[6])
	recvSpeed, _ := strconv.Atoi(split[7])
	sendSpeedNum := int(math.Ceil(float64(sendSpeed)))
	recvSpeedNum := int(math.Ceil(float64(recvSpeed)))
	var Entity FlowAudit
	Entity.DeviceIp = deviceIp
	Entity.DeviceMac = deviceMac
	TrafficType, _ := strconv.Atoi(split[0])
	Entity.TrafficType = TrafficType
	if TrafficType == 40 {
		//最后一个
		Entity.TrafficName = "unknown"
	} else {
		Entity.TrafficName = ProtocolTuple[TrafficType]
	}
	Entity.SourcePort = srcPort
	Entity.DestinationPort = destPort
	Entity.SendSpeed = sendSpeedNum
	Entity.RecvSpeed = recvSpeedNum
	Entity.RecvBytes, _ = strconv.Atoi(split[7])
	Entity.SendBytes, _ = strconv.Atoi(split[6])
	location, _ := time.ParseInLocation("2006-01-02 15:04:05", split[1], time.Local)
	Entity.Timestamp = location
	splitMac := strings.Split(split[5], ":")
	macPrefix := splitMac[0] + splitMac[1] + splitMac[2]
	upper := strings.ToUpper(macPrefix)
	get := DeviceMacMap[upper]
	if get == "" {
		Entity.DeviceName = "未知"
	} else {
		Entity.DeviceName = get
	}

}

var ProtocolTuple = []string{"equipment", "modbus", "opcda", "iec104", "dnp3", "fins", "mms", "s7", "profnetio", "goose", "sv", "enip", "opcua", "pnrt_dcp", "http", "ftp", "bacnet", "tls", "ssh", "imap", "smb", "smb2", "network", "pn_rt", "baosight", "dns", "netbios", "smtp", "telnet", "pop3", "snmp", "unknown"}

var DeviceMacMap = map[string]string{
	"000000": "未知",
}

type FlowAudit struct {
	ID              int64     `gorm:"primarykey" json:"id"` // 主键ID
	DeviceIp        string    `json:"deviceIp" form:"deviceIp" gorm:"column:device_ip;comment:设备IP;size:25;"`
	DeviceMac       string    `json:"deviceMac" form:"deviceMac" gorm:"column:device_mac;comment:设备mac;size:25;"`
	DeviceId        string    `json:"deviceId" form:"deviceId" gorm:"column:device_id;comment:设备id;size:25;"`
	DeviceName      string    `json:"deviceName" form:"deviceName" gorm:"column:device_name;comment:设备名称;size:255;"`
	TrafficType     int       `json:"trafficType" form:"trafficType" gorm:"column:traffic_type;comment:协议类型;size:10;"`
	TrafficName     string    `json:"trafficName" form:"trafficName" gorm:"column:traffic_name;comment:协议名称;size:255;"`
	SourcePort      int       `json:"sourcePort" form:"sourcePort" gorm:"column:source_port;comment:源端口;size:10;"`
	DestinationPort int       `json:"destinationPort" form:"destinationPort" gorm:"column:destination_port;comment:目的端口;size:10;"`
	SendBytes       int       `json:"sendBytes" form:"sendBytes" gorm:"column:send_bytes;comment:发送字节数;size:10;"`
	RecvBytes       int       `json:"recvBytes" form:"recvBytes" gorm:"column:recv_bytes;comment:接收字节数;size:10;"`
	SendSpeed       int       `json:"sendSpeed" form:"sendSpeed" gorm:"column:send_speed;comment:发送速率;size:10;"`
	RecvSpeed       int       `json:"recvSpeed" form:"recvSpeed" gorm:"column:recv_speed;comment:接收速率;size:10;"`
	Timestamp       time.Time `json:"timestamp" form:"timestamp" gorm:"column:timestamp;comment:时间戳;"`
}
