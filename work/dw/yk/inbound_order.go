package yk

import "time"

// InboundOrder  入库单
type InboundOrder struct {
	ID              int64     `gorm:"column:id;primaryKey" json:"id"`
	InboundNo       string    `gorm:"column:inbound_no" json:"inbound_no"`             //  入库单号
	RelationNo      string    `gorm:"column:relation_no" json:"relation_no"`           //  关联
	FactoryCode     string    `gorm:"column:factory_code" json:"factory_code"`         // 工厂编码
	FactoryName     string    `gorm:"column:factory_name" json:"factory_name"`         // 工厂名称
	SupplierCode    string    `gorm:"column:supplier_code" json:"supplier_code"`       //  供应商编码
	SupplierName    string    `gorm:"column:supplier_name" json:"supplier_name"`       //  供应商名称
	RequirementDate time.Time `gorm:"column:requirement_date" json:"requirement_date"` //  需求日期
	InboundTime     time.Time `gorm:"column:inbound_time" json:"inbound_time"`         //  入库时间
	DeliveryWay     int64     `gorm:"column:delivery_way" json:"delivery_way"`         //  配送方式 1:安吉配送 2:直投
	DeliveryOrigin  int64     `gorm:"column:delivery_origin" json:"delivery_origin"`   //  配送单来源
	Comment         string    `gorm:"column:comment" json:"comment"`                   //  备注信息
	CreatedTime     time.Time `gorm:"column:created_time" json:"created_time"`         //  创建时间
	UpdatedTime     time.Time `gorm:"column:updated_time" json:"updated_time"`         //  更新时间
	InboundType     int64     `gorm:"inbound_type" json:"inbound_type"`                // 入库类型
	WarehouseCode   string    `gorm:"warehouse_code" json:"warehouse_code"`            // 仓库
	WarehouseName   string    `gorm:"warehouse_name" json:"warehouse_name"`            // 仓库名称
	CreatedUserId   int64     `gorm:"created_user_id" json:"created_user_id"`          // 创建人ID
	CreatedUserName string    `gorm:"created_user_name" json:"created_user_name"`      // 创建人名称
}

func (InboundOrder) TableName() string {
	return `inbound_order`
}
