package yk

type Factory struct {
	Id          int64  `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	FactoryCode string `gorm:"column:factory_code;NOT NULL" json:"factory_code"` // 工厂编码
	FactoryName string `gorm:"column:factory_name;NOT NULL" json:"factory_name"` // 工厂名称
}

func (Factory) TableName() string {
	return `factory`
}
