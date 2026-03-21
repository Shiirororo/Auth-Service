package product_entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type SKU struct {
	ID         int64           `gorm:"column:id;primaryKey;autoIncrement"`
	SKUNo      string          `gorm:"column:skuNo;type:varchar(32);uniqueIndex:uk_sku_no;default:'';comment:sku no"`
	SKUName    *string         `gorm:"column:skuName;type:varchar(50);comment:sku name"`
	SKUDesc    *string         `gorm:"column:skuDesc;type:varchar(256);comment:sku desc"`
	SKUType    *int8           `gorm:"column:skuType;type:tinyint;comment:sku type"`
	Status     int8            `gorm:"column:status;type:tinyint;not null;comment:status"`
	Sort       int             `gorm:"column:sort;default:0;comment:priority sort"`
	SKUStock   int             `gorm:"column:skuStock;not null;default:0;comment:sku stock"`
	SKUPrice   decimal.Decimal `gorm:"column:skuPrice;type:decimal(10,2);not null;comment:sku price"`
	CreateTime time.Time       `gorm:"column:createTime;autoCreateTime"`
	LastUpdate time.Time       `gorm:"column:lastUpdate;autoUpdateTime"`
}

func (s *SKU) TableName() string {
	return "sku"
}
