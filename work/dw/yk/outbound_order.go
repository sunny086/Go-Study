package yk

import "time"

// OutboundOrder  出库单详情
type OutboundOrder struct {
	ID                int64     `gorm:"column:id;primaryKey" json:"id"`
	OutboundNo        string    `gorm:"column:outbound_no" json:"outbound_no"`           // 出库单号
	DeliveryNo        string    `gorm:"column:delivery_no" json:"delivery_no"`           // 关联配送单号
	PickNo            string    `gorm:"column:pick_no" json:"pick_no"`                   // 关联拣货单号
	DeliveryType      int       `gorm:"column:delivery_type" json:"delivery_type"`       // 配送方式 1: 安吉配送
	RequiredType      int       `gorm:"column:required_type" json:"required_type"`       // 类型 1:日料单 2:补货单 3:看板单
	OrderType         int       `gorm:"column:order_type" json:"order_type"`             // 订单类型 1:料单 2:调拨单 3:售后配件 4:采购退货 5:整车订单
	WorkCenter        string    `gorm:"column:work_center" json:"work_center"`           // 工作中心
	ManufactWarehouse string    `gorm:"column:manufact_house" json:"manufact_house"`     // 生产商仓库
	Warehouse         string    `gorm:"column:warehouse" json:"warehouse"`               // 发货仓库
	WarehouseCode     string    `gorm:"column:warehouse_code" json:"warehouse_code"`     // 仓库编码
	FactoryCode       string    `gorm:"column:factory_code" json:"factory_code"`         // 工厂编码
	FactoryName       string    `gorm:"column:factory_name" json:"factory_name"`         // 工厂名称
	RequiredTime      string    `gorm:"column:required_time" json:"required_time"`       // 需求时间
	OutboundTime      time.Time `gorm:"column:outbound_time" json:"outbound_time"`       // 出库时间
	OutboundState     int64     `gorm:"column:outbound_state" json:"outbound_state"`     // 关联单类型 1:待收货 2:已收货
	GeneratedSource   int       `gorm:"column:generated_source" json:"generated_source"` // 生成来源 1: 系统料单 2: 手动提报
	CustomName        string    `gorm:"column:custom_name" json:"custom_name"`           // 客户名称/收货方
	Remark            string    `gorm:"column:remark" json:"remark"`                     // 备注信息
	CreatedUserID     int64     `gorm:"created_user_id" json:"created_user_id"`          // 创建人ID
	CreatedUserName   string    `gorm:"created_user_name" json:"created_user_name"`      // 创建人名称
	CreatedTime       time.Time `gorm:"column:created_time" json:"created_time"`         // 创建时间
	UpdatedTime       time.Time `gorm:"column:updated_time" json:"updated_time"`         // 更新时间
}

func (OutboundOrder) TableName() string {
	return `outbound_order`
}
