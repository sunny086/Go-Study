package yk

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

var (
	supplierMap            = map[string]string{}
	warehouseMap           = map[string]string{}
	factoryList            []*Factory
	supplierList           []*Supplier
	warehouseList          []*Warehouse
	materielRelocationList []*MaterielRelocation
	flowList               []*MaterielFlow
	inboundOrderList       []*InboundOrder
	inboundMaterielList    []*InboundMateriels
	outboundOrderList      []*OutboundOrder
	outboundMaterielList   []*OutboundMateriels

	db *gorm.DB
)

func init() {
	db, _ = gorm.Open(mysql.New(mysql.Config{
		DSN:                       "deepway:nBI85b1ogucJfZlv5y0R@tcp(10.51.101.12:3306)/xujs?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         256,                                                                                              // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                                             // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                                             // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                                             // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                                            // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	db.Model(&MaterielRelocation{}).Find(&materielRelocationList)
	db.Model(&Factory{}).Find(&factoryList)
	db.Model(&Supplier{}).Find(&supplierList)
	db.Model(&Warehouse{}).Find(&warehouseList)
	for _, item := range supplierList {
		supplierMap[item.SupplierCode] = item.SupplierName
	}
	for _, item := range warehouseList {
		warehouseMap[item.WarehouseCode] = item.Name
	}
}
func TestDataClean(t *testing.T) {

	var noPrefix = "YK240527" //YK2405270017
	//noStart++ 转四位字符串 左侧补0
	noStart := 17
	for _, item := range materielRelocationList {
		//入库
		noStart = noStart + 1
		no := fmt.Sprintf("%s%04d", noPrefix, noStart)
		var inboundOrder = &InboundOrder{
			InboundNo:       no, //todo
			RelationNo:      item.RelocationNo,
			FactoryCode:     item.FactoryCode,
			FactoryName:     item.FactoryName,
			WarehouseCode:   item.RelocationWarehouseIn,
			SupplierCode:    item.SupplierCode,
			SupplierName:    item.SupplierName,
			InboundType:     27,
			CreatedUserId:   item.CreatedUserId,
			CreatedUserName: item.CreatedUserName,
			CreatedTime:     item.CreatedTime,
			UpdatedTime:     item.UpdatedTime,
			InboundTime:     item.CreatedTime,
		}
		var inboundMateriel = &InboundMateriels{
			InboundNo:    inboundOrder.InboundNo, //todo
			MaterielCode: item.MaterielCode,
			SupplierCode: item.SupplierCode,
			MaterielName: item.MaterielName,
			SupplierName: item.SupplierName,
			MeasureUnit:  item.Unit,
			ActualNum:    item.RelocationNum,
			ReceivedNum:  item.RelocationNum,
			InboundInfo:  "{}",
			CreatedTime:  item.CreatedTime,
			UpdatedTime:  item.UpdatedTime,
		}
		inboundOrderList = append(inboundOrderList, inboundOrder)
		inboundMaterielList = append(inboundMaterielList, inboundMateriel)

		//出库
		noStart = noStart + 1
		no = fmt.Sprintf("%s%04d", noPrefix, noStart)
		var outboundOrder = &OutboundOrder{
			OutboundNo:      no, //todo
			FactoryCode:     item.FactoryCode,
			FactoryName:     item.FactoryName,
			DeliveryNo:      item.RelocationNo,
			DeliveryType:    9,
			RequiredType:    25,
			OrderType:       12,
			OutboundState:   3,
			GeneratedSource: 1,
			WarehouseCode:   item.WarehouseCode,
			Warehouse:       warehouseMap[item.WarehouseCode],
			OutboundTime:    item.CreatedTime,
			CreatedUserID:   item.CreatedUserId,
			CreatedUserName: item.CreatedUserName,
			RequiredTime:    item.CreatedTime.Format("2006-01-02 15:04:05"),
			CreatedTime:     item.CreatedTime,
			UpdatedTime:     item.UpdatedTime,
		}
		var outboundMateriel = &OutboundMateriels{
			OutboundNo:   outboundOrder.OutboundNo, //todo
			MaterielCode: item.MaterielCode,
			MaterielName: item.MaterielName,
			UsedCode:     item.UsedCode,
			SupplierCode: item.SupplierCode,
			SupplierName: item.SupplierName,
			Num:          item.RelocationNum,
			Remark:       item.Remark,
			OutboundInfo: "{}",
			OutLineNo:    1, //todo
			ReceivedNum:  item.RelocationNum,
			RequiredNum:  item.RelocationNum,
			ReceivedTime: item.CreatedTime,
			RequiredTime: item.CreatedTime,
			CreatedTime:  item.CreatedTime,
			UpdatedTime:  item.CreatedTime,
		}
		outboundOrderList = append(outboundOrderList, outboundOrder)
		outboundMaterielList = append(outboundMaterielList, outboundMateriel)

		// 移出
		flowList = append(flowList, &MaterielFlow{
			RelationNo:    item.RelocationNo,
			FlowOrderNo:   outboundOrder.OutboundNo,
			MaterielCode:  item.MaterielCode,
			MaterielName:  item.MaterielName,
			UsedCode:      item.UsedCode,
			SupplierCode:  item.SupplierCode,
			SupplierName:  item.SupplierName,
			WarehouseCode: item.RelocationOut,
			WarehouseName: warehouseMap[item.RelocationOut],
			FactoryCode:   item.FactoryCode,
			FactoryName:   item.FactoryName,
			Num:           item.RelocationNum,
			FlowType:      2,
			RequiredType:  25,
			OrderType:     12,
			FlowDate:      item.CreatedTime,
			CreatedAt:     item.CreatedTime,
			UpdatedAt:     item.CreatedTime,
		})
		// 移入
		flowList = append(flowList, &MaterielFlow{
			RelationNo:    item.RelocationNo,
			FlowOrderNo:   inboundOrder.InboundNo,
			MaterielCode:  item.MaterielCode,
			MaterielName:  item.MaterielName,
			UsedCode:      item.UsedCode,
			SupplierCode:  item.SupplierCode,
			SupplierName:  item.SupplierName,
			WarehouseCode: item.RelocationWarehouseIn,
			WarehouseName: warehouseMap[item.RelocationWarehouseIn],
			FactoryCode:   item.FactoryCode,
			FactoryName:   item.FactoryName,
			Num:           item.RelocationNum,
			FlowType:      1,
			RequiredType:  27,
			OrderType:     27,
			FlowDate:      item.CreatedTime,
			CreatedAt:     item.CreatedTime,
			UpdatedAt:     item.CreatedTime,
		})
	}
	db.Model(&InboundOrder{}).Create(inboundOrderList)
	db.Model(&InboundMateriels{}).Create(inboundMaterielList)
	db.Model(&OutboundOrder{}).Create(outboundOrderList)
	db.Model(&OutboundMateriels{}).Create(outboundMaterielList)
	db.Model(&MaterielFlow{}).Create(flowList)

}
