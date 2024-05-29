package yk

import "time"

// Supplier  供应商信息
type Supplier struct {
	ID               int64     `gorm:"column:id;primaryKey" json:"id"`
	SupplierName     string    `gorm:"column:supplier_name" json:"supplier_name"`           // 供应商名称
	SupplierCode     string    `gorm:"column:supplier_code" json:"supplier_code"`           // 供应商编码
	IsOfficial       int64     `gorm:"column:is_official" json:"is_official"`               // 正式供应商 1:是 2:否
	Type             string    `gorm:"column:type" json:"type"`                             // 供应商分类 1:零件
	JacName          string    `gorm:"column:jac_name" json:"jac_name"`                     // 供应商 jac 名称
	JacCode          string    `gorm:"column:jac_code" json:"jac_code"`                     // 供应商 jac 编码
	Contact          string    `gorm:"column:contact" json:"contact"`                       // 联系人
	Telephone        string    `gorm:"column:telephone" json:"telephone"`                   // 联系电话
	NeedQualityCheck int64     `gorm:"column:need_quality_check" json:"need_quality_check"` // 是否需要质检 1:需要 2:不需要
	TaxNo            string    `gorm:"column:tax_no" json:"tax_no"`                         // 税号
	CreatedTime      time.Time `gorm:"column:created_time" json:"created_time"`             // 创建时间
	UpdatedTime      time.Time `gorm:"column:updated_time" json:"updated_time"`             // 更新时间
	DeliveryAddress  string    `gorm:"delivery_address" json:"delivery_address"`            // 发货地址
}

func (Supplier) TableName() string {
	return `supplier`
}
