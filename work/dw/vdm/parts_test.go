package vdm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
)

type VehicleParts struct {
	Id              uint64    `gorm:"id" json:"id"`
	VehicleApplyNo  string    `json:"vehicle_apply_no" gorm:"vehicle_apply_no"`                               // 车辆申请单号
	PartName        string    `json:"part_name" gorm:"part_name"`                                             // 信息类型名称
	PartCode        int       `json:"part_code" gorm:"part_code"`                                             // 信息类型编码
	Content         string    `json:"content" gorm:"content"`                                                 // 信息内容
	Description     string    `json:"description" gorm:"description"`                                         // 其他描述
	CreatedUserId   int64     `json:"created_user_id" gorm:"created_user_id"`                                 // 创建人ID
	CreatedUserName string    `json:"created_user_name" gorm:"created_user_name"`                             // 创建人名称
	CreatedAt       time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"created_at"` // 创建时间
	UpdatedAt       time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"updated_at"` // 更新时间
}

func (VehicleParts) TableName() string {
	return `vehicle_parts`
}

func TestPartsSql(t *testing.T) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:@tcp(127.0.0.1:3306)/dw_test?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         256,                                                                       // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                      // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                      // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                      // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                     // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})

	var parts []VehicleParts
	err = db.Model(&VehicleParts{}).Find(&parts).Error
	if err != nil {
		t.Error(err)
	}
	fmt.Println(parts)

}
