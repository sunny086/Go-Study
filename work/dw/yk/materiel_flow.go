package yk

import "time"

type MaterielFlow struct {
	ID            int64     `json:"id" gorm:"column:id"`
	FlowOrderNo   string    `json:"flow_order_no" gorm:"column:flow_order_no"`   // 流水单号-出入库单号
	OrderID       int64     `json:"order_id" gorm:"column:order_id"`             // 流水单号-主键ID
	RelationNo    string    `json:"relation_no" gorm:"column:relation_no"`       // 关联单号
	MaterielCode  string    `json:"materiel_code" gorm:"column:materiel_code"`   // 物料编码
	MaterielName  string    `json:"materiel_name" gorm:"column:materiel_name"`   // 物料名称
	UsedCode      string    `json:"used_code" gorm:"column:used_code"`           // 曾用物料编码
	Vin           string    `json:"vin" gorm:"column:vin"`                       // vin码
	SupplierCode  string    `json:"supplier_code" gorm:"column:supplier_code"`   // 供应商编码
	SupplierName  string    `json:"supplier_name" gorm:"column:supplier_name"`   // 供应商名称
	WarehouseCode string    `json:"warehouse_code" gorm:"column:warehouse_code"` // 库存仓库编码
	WarehouseName string    `json:"warehouse_name" gorm:"column:warehouse_name"` // 库存仓库名称
	FactoryCode   string    `json:"factory_code" gorm:"column:factory_code"`     // 工厂编码
	FactoryName   string    `json:"factory_name" gorm:"column:factory_name"`     // 工厂名称
	Num           int64     `json:"num" gorm:"column:num"`                       // 数量
	FlowType      int       `json:"flow_type" gorm:"column:flow_type"`           // 出入库流水类型。1:入库 2:出库
	RequiredType  int       `json:"required_type" gorm:"column:required_type"`   // 需求类型
	OrderType     int       `json:"order_type" gorm:"column:order_type"`         // 订单类型
	FlowDate      time.Time `json:"flow_date" gorm:"column:flow_date"`           // 物料出入库时间
	CreatedAt     time.Time `json:"created_at" gorm:"column:created_at"`         // 创建时间
	UpdatedAt     time.Time `json:"updated_at" gorm:"column:updated_at"`         // 更新时间
}

// TableName 表名称
func (MaterielFlow) TableName() string {
	return "materiel_flow"
}
