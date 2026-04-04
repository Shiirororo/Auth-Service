package product_entity

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/datatypes"
)

type SKUStatus int8

type SKU struct {
	ID         int             `gorm:"column:id;primaryKey;autoIncrement"`
	SkuNo      string          `gorm:"column:skuNo;size:32;uniqueIndex:uk_sku_no"`
	SkuName    string          `gorm:"column:skuName;size:50"`
	SkuDesc    string          `gorm:"column:skuDesc;size:256"`
	SkuType    SKUStatus       `gorm:"column:skuType"`
	Status     SKUStatus       `gorm:"column:status;not null"`
	Sort       int             `gorm:"column:sort;default:0"`
	SkuStock   int             `gorm:"column:skuStock;not null;default:0"`
	SkuPrice   decimal.Decimal `gorm:"column:skuPrice;type:decimal(8,2);not null"`
	CreateTime time.Time       `gorm:"column:createTime;autoCreateTime"`
	LastUpdate time.Time       `gorm:"column:lastUpdate;autoUpdateTime"`
}

func (SKU) TableName() string { return "sku" }

type SKUAttr struct {
	ID           int            `gorm:"column:id;primaryKey;autoIncrement"`
	SkuNo        string         `gorm:"column:skuNo;size:32;uniqueIndex:uk_sku_no"`
	SkuAttribute datatypes.JSON `gorm:"column:skuAttribute"`
	CreateTime   time.Time      `gorm:"column:createTime;autoCreateTime"`
	LastUpdate   time.Time      `gorm:"column:lastUpdate;autoUpdateTime"`
}

func (SKUAttr) TableName() string { return "sku_attr" }

type SPUToSKU struct {
	ID         uint64    `gorm:"column:id;primaryKey;autoIncrement"`
	SkuNo      string    `gorm:"column:skuNo;size:32;not null;index:idx_sku"`
	SpuNo      string    `gorm:"column:spuNo;size:32;not null;uniqueIndex:uk_spu_sku"`
	IsDeleted  uint8     `gorm:"column:isDeleted;default:0"`
	CreateTime time.Time `gorm:"column:createTime;autoCreateTime"`
	LastUpdate time.Time `gorm:"column:lastUpdate;autoUpdateTime"`
}

func (SPUToSKU) TableName() string { return "spu_to_sku" }
