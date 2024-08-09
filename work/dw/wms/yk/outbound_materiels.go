package yk

import "time"

// OutboundMateriels  出库单物料信息
type OutboundMateriels struct {
	ID                       int64     `gorm:"column:id;primaryKey" json:"id"`
	OutboundNo               string    `gorm:"column:outbound_no" json:"outbound_no"`                               // 出库单号
	DeliveryMaterielID       int64     `gorm:"column:delivery_materiel_id" json:"delivery_materiel_id"`             // 配送单物料 ID
	OutLineNo                int       `gorm:"column:out_line_no" json:"out_line_no"`                               // 物料行号
	ItemNo                   int       `gorm:"column:item_no" json:"item_no"`                                       // 来源行项目
	PickMaterielID           int64     `gorm:"column:pick_materiel_id" json:"pick_materiel_id"`                     // 拣货单物料 ID
	MaterielCode             string    `gorm:"column:materiel_code" json:"materiel_code"`                           // 零件编码
	UsedCode                 string    `gorm:"column:used_code" json:"used_code"`                                   // 曾用零件编码
	MaterielName             string    `gorm:"column:materiel_name" json:"materiel_name"`                           // 零件名称
	DWMaterielCode           string    `gorm:"column:dw_materiel_code" json:"dw_materiel_code"`                     // SAP 零件编码
	DWMaterielName           string    `gorm:"column:dw_materiel_name" json:"dw_materiel_name"`                     // SAP 零件名称
	ManufacturerMaterielCode string    `gorm:"column:manufacturer_materiel_code" json:"manufacturer_materiel_code"` // 代加工工厂零件编码
	ManufacturerMaterielName string    `gorm:"column:manufacturer_materiel_name" json:"manufacturer_materiel_name"` // 代加工工厂零件名称
	SupplierCode             string    `gorm:"column:supplier_code" json:"supplier_code"`                           // 供应商编码
	SupplierName             string    `gorm:"column:supplier_name" json:"supplier_name"`                           // 供应商名称
	Vin                      string    `gorm:"column:vin" json:"vin"`                                               // vin码
	ManufacturerSupplierCode string    `gorm:"column:manufacturer_supplier_code" json:"manufacturer_supplier_code"` // 代加工工厂供应商编码
	ManufacturerSupplierName string    `gorm:"column:manufacturer_supplier_name" json:"manufacturer_supplier_name"` // 代加工工厂供应商名称
	MapSupplierCode          string    `gorm:"column:map_supplier_code" json:"map_supplier_code"`                   // 临时映射供应商编码
	MapSupplierName          string    `gorm:"column:map_supplier_name" json:"map_supplier_name"`                   // 临时映射供应商名称
	VDMID                    string    `gorm:"column:vdm_id" json:"vdm_id"`                                         // vdm 宜搭 id
	VehicleNo                string    `gorm:"column:vehicle_no" json:"vehicle_no"`                                 // vdm 车辆流水号
	Station                  string    `gorm:"column:station" json:"station"`                                       // 工位
	ProductionLine           string    `gorm:"column:production_line" json:"production_line"`                       // 产线
	Num                      int64     `gorm:"column:num" json:"num"`                                               // 数量
	RequiredNum              int64     `gorm:"column:required_num" json:"required_num"`                             // 需求数量
	ReceivedNum              int64     `gorm:"column:received_num" json:"received_num"`                             // 收货数量
	ReceivedTime             time.Time `gorm:"received_time" json:"received_time"`                                  // 收获时间
	RequiredTime             time.Time `gorm:"column:required_time" json:"required_time"`                           // 需求时间
	OutboundInfo             string    `gorm:"column:outbound_info" json:"outbound_info"`                           // 库存物料扣减信息
	Remark                   string    `gorm:"column:remark" json:"remark"`                                         // 料单备注信息
	CreatedTime              time.Time `gorm:"column:created_time" json:"created_time"`                             // 创建时间
	UpdatedTime              time.Time `gorm:"column:updated_time" json:"updated_time"`                             // 更新时间
}

func (OutboundMateriels) TableName() string {
	return `outbound_materiels`
}
