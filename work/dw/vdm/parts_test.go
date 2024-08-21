package vdm

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sort"
	"strconv"
	"strings"
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

func TestReadPartsExcel(t *testing.T) {
	// 读取excel
	// 1. 读取excel
	// 2. 解析excel
	// 3. 写入数据库

	// 读取Excel文件
	filePath := "零件信息.xlsx"
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Fatalf("无法打开文件: %v", err)
	}

	// 读取Sheet1中的所有行
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		log.Fatalf("无法读取行: %v", err)
	}
	// 初始化一个结构体的数组
	var vehiclePartsList []VehicleParts

	// 遍历每一行，从第二行开始读取，因为第一行是标题
	for _, row := range rows[1:] {
		partCode, _ := strconv.Atoi(row[2]) // 将信息类型编码转换为整数

		// row[3] 是 / 替换空
		row[3] = strings.Replace(row[3], "/", "", -1)

		vehiclePart := VehicleParts{
			VehicleApplyNo: row[0],
			PartName:       row[1],
			PartCode:       partCode,
			Content:        row[3],
		}
		vehiclePartsList = append(vehiclePartsList, vehiclePart)
	}

	var applyNo = []string{"VE2024060001", "VE2024060002", "VE2024060003", "VE2024060004", "VE2024060005", "VE2024060006", "VE2024060007", "VE2024060008", "VE2024060009", "VE2024060010", "VE2024060011", "VE2024060012", "VE2024060013", "VE2024060014", "VE2024060015", "VE2024060016", "VE2024060017", "VE2024060018", "VE2024060019", "VE2024060020", "VE2024060021", "VE2024060022", "VE2024060023", "VE2024060024", "VE2024060025", "VE2024060026", "VE2024060027", "VE2024060028", "VE2024060029", "VE2024060030", "VE2024060031", "VE2024060032", "VE2024060033", "VE2024060034", "VE2024060035", "VE2024060036", "VE2024060037", "VE2024060038", "VE2024060039", "VE2024060040", "VE2024060041", "VE2024060042", "VE2024060043", "VE2024060044", "VE2024060045", "VE2024060046", "VE2024060047", "VE2024060048", "VE2024060049", "VE2024060050", "VE2024060051", "VE2024060052", "VE2024060053", "VE2024060054", "VE2024060055", "VE2024060056", "VE2024060057", "VE2024060058", "VE2024060059", "VE2024060060", "VE2024060061", "VE2024060062", "VE2024060063", "VE2024060064", "VE2024060065", "VE2024060066", "VE2024060067", "VE2024060068", "VE2024060069", "VE2024060070", "VE2024060071", "VE2024060072", "VE2024060073", "VE2024060074", "VE2024060075", "VE2024060076", "VE2024060077", "VE2024060078", "VE2024060079", "VE2024060080", "VE2024060081", "VE2024060082", "VE2024060083", "VE2024060084", "VE2024060085", "VE2024070001", "VE2024070002", "VE2024070003", "VE2024070004", "VE2024070005", "VE2024070006"}

	var vehiclePartsMap = make(map[string][]VehicleParts)
	for _, item := range vehiclePartsList {
		vehiclePartsMap[item.VehicleApplyNo] = append(vehiclePartsMap[item.VehicleApplyNo], item)
	}

	// 不存在的车辆申请单号 组装数据
	for _, no := range applyNo {
		if _, ok := vehiclePartsMap[no]; !ok {
			vehiclePartsList = append(vehiclePartsList,
				VehicleParts{
					VehicleApplyNo: no,
					PartCode:       1,
					PartName:       "电池型号/状态",
				}, VehicleParts{
					VehicleApplyNo: no,
					PartCode:       2,
					PartName:       "电池组编号",
				}, VehicleParts{
					VehicleApplyNo: no,
					PartCode:       3,
					PartName:       "电池编码1",
				}, VehicleParts{
					VehicleApplyNo: no,
					PartCode:       4,
					PartName:       "电池编码2",
				}, VehicleParts{
					VehicleApplyNo: no,
					PartCode:       5,
					PartName:       "电池编码3",
				}, VehicleParts{
					VehicleApplyNo: no,
					PartCode:       6,
					PartName:       "电池编码4",
				}, VehicleParts{
					VehicleApplyNo: no,
					PartCode:       7,
					PartName:       "电池编码5",
				}, VehicleParts{
					VehicleApplyNo: no,
					PartCode:       8,
					PartName:       "电池编码6",
				}, VehicleParts{
					VehicleApplyNo: no,
					PartCode:       9,
					PartName:       "电池编码7",
				}, VehicleParts{
					VehicleApplyNo: no,
					PartCode:       10,
					PartName:       "驱动电机型号",
				}, VehicleParts{
					VehicleApplyNo: no,
					PartCode:       11,
					PartName:       "驱动电机1（6x4中桥或4x2后桥1）",
				}, VehicleParts{
					VehicleApplyNo: no,
					PartCode:       12,
					PartName:       "驱动电机2（6x4后桥或4x2后桥2）",
				}, VehicleParts{
					VehicleApplyNo: no,
					PartCode:       13,
					PartName:       "TBOX编号",
				}, VehicleParts{
					VehicleApplyNo: no,
					PartCode:       14,
					PartName:       "SIM卡号",
				}, VehicleParts{
					VehicleApplyNo: no,
					PartCode:       15,
					PartName:       "前桥编号",
				}, VehicleParts{
					VehicleApplyNo: no,
					PartCode:       16,
					PartName:       "中桥编号",
				}, VehicleParts{
					VehicleApplyNo: no,
					PartCode:       17,
					PartName:       "后桥编号",
				}, VehicleParts{
					VehicleApplyNo: no,
					PartCode:       18,
					PartName:       "驾驶室编号",
				})
		}
	}

	// 按照车辆申请单号降序排序
	// 可以使用 sort.Slice 来进行排序
	sort.Slice(vehiclePartsList, func(i, j int) bool {
		if vehiclePartsList[i].VehicleApplyNo == vehiclePartsList[j].VehicleApplyNo {
			return vehiclePartsList[i].PartCode < vehiclePartsList[j].PartCode
		}
		return vehiclePartsList[i].VehicleApplyNo > vehiclePartsList[j].VehicleApplyNo
	})

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:@tcp(127.0.0.1:3306)/dw_test?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         256,                                                                       // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                      // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                      // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                      // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                     // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	err = db.Model(&VehicleParts{}).Create(&vehiclePartsList).Error
	if err != nil {
		t.Error(err)
	}
	var parts []VehicleParts
	err = db.Model(&VehicleParts{}).Find(&parts).Error
	if err != nil {
		t.Error(err)
	}
	fmt.Println(parts)
}
