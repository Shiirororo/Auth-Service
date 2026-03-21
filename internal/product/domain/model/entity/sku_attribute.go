package product_entity

import (
	"time"

	"gorm.io/datatypes"
)

type SKUAttr struct {
	ID           int64          `gorm:"column:id;primaryKey"`
	SKUNo        string         `gorm:"column:skuNo;index"`
	SKUAttribute datatypes.JSON `gorm:"column:skuAttribute;type:json"`

	CreateTime time.Time `gorm:"autoCreateTime"`
	LastUpdate time.Time `gorm:"autoUpdateTime"`
}

func (skuA *SKUAttr) TableName() string {
	return "sku_attribute"
}
