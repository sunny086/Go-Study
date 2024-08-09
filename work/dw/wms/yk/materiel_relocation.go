package yk

import "time"

type MaterielRelocation struct {
	Id                    int64     `gorm:"column:id" json:"id"`
	RelocationNo          string    `gorm:"column:relocation_no" json:"relocation_no"`                     // 移库单号
	RelocationTime        time.Time `gorm:"column:relocation_time" json:"relocation_time"`                 // 移库时间
	BatchNo               string    `gorm:"column:batch_no" json:"batch_no"`                               // 批次单号
	SupplierCode          string    `gorm:"column:supplier_code" json:"supplier_code"`                     // 供应商编码
	SupplierName          string    `gorm:"column:supplier_name" json:"supplier_name"`                     // 供应商名称
	MaterielCode          string    `gorm:"column:materiel_code" json:"materiel_code"`                     // 零件编码
	UsedCode              string    `gorm:"column:used_code" json:"used_code"`                             // 曾用零件编码
	MaterielName          string    `gorm:"column:materiel_name" json:"materiel_name"`                     // 零件名称
	RelocationWarehouseIn string    `gorm:"column:relocation_warehouse_in" json:"relocation_warehouse_in"` // 移入仓库
	RelocationAreaIn      string    `gorm:"column:relocation_area_in" json:"relocation_area_in"`           // 移入库区
	RelocationLocationIn  string    `gorm:"column:relocation_location_in" json:"relocation_location_in"`   // 移入库位
	RelocationOut         string    `gorm:"column:relocation_out" json:"relocation_out"`                   // 移出仓库
	RelocationNum         int64     `gorm:"column:relocation_num" json:"relocation_num"`                   // 移动数量
	Unit                  string    `gorm:"column:unit" json:"unit"`                                       // 单位
	Remark                string    `gorm:"column:remark" json:"remark"`                                   // 备注
	WarehouseCode         string    `gorm:"column:warehouse_code" json:"warehouse_code"`                   // 仓库编码
	FactoryCode           string    `gorm:"column:factory_code" json:"factory_code"`                       // 工厂编码
	FactoryName           string    `gorm:"column:factory_name" json:"factory_name"`                       // 工厂名称
	CreatedUserId         int64     `gorm:"column:created_user_id" json:"created_user_id"`                 // 创建人ID
	CreatedUserName       string    `gorm:"column:created_user_name" json:"created_user_name"`             // 创建人名称
	CreatedTime           time.Time `gorm:"column:created_time" json:"created_time"`                       // 创建时间
	UpdatedTime           time.Time `gorm:"column:updated_time" json:"updated_time"`                       // 更新时间
}

func (MaterielRelocation) TableName() string {
	return `materiel_relocation`
}
