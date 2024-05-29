package yk

import "time"

type InboundMateriels struct {
	Id             int64     `gorm:"column:id" json:"id"`
	InboundNo      string    `gorm:"column:inbound_no" json:"inbound_no"`             // 入库单号
	ItemNo         string    `gorm:"item_no" json:"item_no"`                          // 业务行号
	MaterielCode   string    `gorm:"column:materiel_code" json:"materiel_code"`       // 物料主数据编码
	MaterielName   string    `gorm:"column:materiel_name" json:"materiel_name"`       // 零件名称
	DWMaterielCode string    `gorm:"column:dw_materiel_code" json:"dw_materiel_code"` // SAP 零件编码
	DWMaterielName string    `gorm:"column:dw_materiel_name" json:"dw_materiel_name"` // SAP 零件名称
	MeasureUnit    string    `gorm:"column:measure_unit" json:"measure_unit"`         // 度量单位
	SupplierCode   string    `gorm:"column:supplier_code" json:"supplier_code"`       // 供应商编码
	SupplierName   string    `gorm:"column:supplier_name" json:"supplier_name"`       // 供应商名称
	Vin            string    `gorm:"column:vin"  json:"vin"`                          // vin码
	ActualNum      int64     `gorm:"column:actual_num" json:"actual_num"`             // 实发数
	ReceivedNum    int64     `gorm:"column:received_num" json:"received_num"`         // 实收数
	InboundInfo    string    `gorm:"column:inbound_info" json:"inbound_info"`         // 入库JSON
	Comment        string    `gorm:"column:comment" json:"comment"`                   // 备注信息
	CreatedTime    time.Time `gorm:"column:created_time" json:"created_time"`         // 创建时间
	UpdatedTime    time.Time `gorm:"column:updated_time" json:"updated_time"`         // 更新时间
}

func (InboundMateriels) TableName() string {
	return `inbound_materiels`
}
