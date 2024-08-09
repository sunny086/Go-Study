package yk

import "time"

// Warehouse  仓库信息
type Warehouse struct {
	ID             int64     `gorm:"column:id;primaryKey" json:"id"`
	FactoryCode    string    `gorm:"factory_code" json:"factory_code"` // 工厂编码
	FactoryName    string    `gorm:"factory_name" json:"factory_name"` // 工厂名称
	Code           string    `gorm:"column:code" json:"code"`          // 仓库编码
	WarehouseCode  string    `gorm:"warehouse_code" json:"warehouse_code"`
	Name           string    `gorm:"column:name" json:"name"`                                          // 仓库名称
	Area           string    `gorm:"column:area" json:"area"`                                          // 库房面积
	Address        string    `gorm:"column:address" json:"address"`                                    // 仓库地址
	Telephone      string    `gorm:"column:telephone" json:"telephone"`                                // 仓库电话
	Principal      string    `gorm:"column:principal" json:"principal"`                                // 负责人
	Comment        string    `gorm:"column:comment" json:"comment"`                                    // 备注信息
	NeedStock      int64     `gorm:"need_stock" json:"need_stock"`                                     // 是否需要计算库存 0 需要 1不需要
	ManufactCode   string    `gorm:"manufact_code" json:"manufact_code"`                               // 工厂编码
	AllocateType   int       `gorm:"allocate_type" json:"allocate_type"`                               // 调拨类型 1:一步调拨
	ExpenseType    int       `gorm:"column:expense_type;default:0;NOT NULL" json:"expense_type"`       // 费用状态：1已费用化 2.未费用化
	WarehouseState int       `gorm:"column:warehouse_state;default:0;NOT NULL" json:"warehouse_state"` // 仓库状态 1开启 2关闭
	WarehouseType  int       `gorm:"column:warehouse_type;default:0;NOT NULL" json:"warehouse_type"`   // 仓库类型 1 普通仓库 2三方仓库
	CreatedTime    time.Time `gorm:"column:created_time" json:"created_time"`                          // 创建时间
	UpdatedTime    time.Time `gorm:"column:updated_time" json:"updated_time"`                          // 更新时间
}

func (Warehouse) TableName() string {
	return `warehouse`
}
