package gorm

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"testing"
	"time"
)

// FlowAudit 结构体
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

// TableName FlowAudit 表名
func (FlowAudit) TableName() string {
	return "audit_flow_11"
}

var client *gorm.DB

func init() {

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:Netvine123#@!@tcp(10.25.10.126:3306)/firewall?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         256,                                                                                        // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                                       // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                                       // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                                       // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                                      // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		fmt.Println("gorm.Open err:", err)
	}
	client = db
}

var CrashedError = errors.New("is marked as crashed and should be repaired")
var CrashedErrorStr = "is marked as crashed and should be repaired"

func TestMyisamCrashed1(t *testing.T) {
	defer func() {
		//捕获test抛出的panic
		if err := recover(); err != nil {
			fmt.Println("panic recover:", err)
		}
	}()
	var flowAudit []FlowAudit
	//err := client.Model(&FlowAudit{}).Limit(10).Find(&flowAudit).Error
	err := client.Table("audit_flow_11").Limit(10).Find(&flowAudit).Error
	if err != nil {
		fmt.Println("client.Find err:", err.Error())
	}
	if strings.Contains(err.Error(), CrashedErrorStr) {
		fmt.Println("err is CrashedError")
		panic(err)
	}
}

//firewall.audit_flow_11	check	warning	Table is marked as crashed
//firewall.audit_flow_11	check	warning	5 clients are using or haven't closed the table properly
//firewall.audit_flow_11	check	error	Size of indexfile is: 253952        Should be: 254976
//firewall.audit_flow_11	check	error	Corrupt

func TestMyisamCrashed2(t *testing.T) {
	var tableNames []string
	client.Table("information_schema.TABLES").Where("table_schema = 'firewall' AND ENGINE = 'MyISAM'").Pluck("table_name", &tableNames)
	fmt.Println("tableNames:", tableNames)
	var id int64
	for _, tableName := range tableNames {
		var err error
		if tableName == "session_statistic" {
			err = client.Table(tableName).Limit(1).Offset(1).Pluck("conc_count", &id).Error
		} else {
			err = client.Table(tableName).Limit(1).Offset(1).Pluck("id", &id).Error
		}
		if err != nil && strings.Contains(err.Error(), CrashedErrorStr) {
			fmt.Println("err is CrashedError", tableName)
			//修复表
			err = client.Exec("repair table " + tableName).Error
			if err != nil {
				fmt.Println(tableName, "修复异常：", err.Error())
			} else {
				fmt.Println(tableName, "修复成功")
			}

			var result string
			client.Raw("CREATE TABLE " + tableName + "_temp" + " LIKE " + tableName).Scan(&result)
			fmt.Println("result:", result)
			client.Exec("DROP TABLE " + tableName).Scan(&result)
			fmt.Println("result:", result)
			client.Exec("RENAME TABLE " + tableName + "_temp" + " TO " + tableName).Scan(&result)

		}
	}
}

// FlowAudit 结构体
type Flow struct {
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

// TableName FlowAudit 表名
func (Flow) TableName() string {
	return "flow"
}

func TestMigrate(t *testing.T) {
	// 自动迁移数据表
	err := client.AutoMigrate(&Flow{})
	if err != nil {
		panic(err)
	}
}

type Result struct {
	Name    string
	Comment string
}

func TestCheckCrashed(t *testing.T) {
	var res Result
	err := client.Raw("check table audit_flow_1").Scan(&res).Error
	if err != nil {
		t.Error(err)
	}
}

func TestFindTableName(t *testing.T) {
	var dropSql []string
	err := client.Table("information_schema.TABLES").Select("table_name").
		Where("table_schema = 'firewall' AND ENGINE = 'MyISAM' AND table_name like ?", "audit_flow_%").Scan(&dropSql).Error
	if err != nil {
		t.Error(err)
	}
	fmt.Println("dropSql:", dropSql)
}
