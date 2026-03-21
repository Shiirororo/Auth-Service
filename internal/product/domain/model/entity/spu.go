package product_entity

import (
	"time"

	"gorm.io/datatypes"
)

type SPU struct {
	ID               uint64         `gorm:"column:id;primaryKey;autoIncrement;comment:id"`
	ProductName      *string        `gorm:"column:productName;type:varchar(64);comment:spu name"`
	ProductDesc      *string        `gorm:"column:productDesc;type:varchar(256);comment:spu desc"`
	ProductStatus    *int8          `gorm:"column:productStatus;type:tinyint;comment:0: out of stock, 1: in stock"`
	ProductAttribute datatypes.JSON `gorm:"column:productAttribute;type:json;comment:json type attribute"`
	ProductShopID    *int64         `gorm:"column:productShopID;comment:shop id"`
	IsDeleted        uint8          `gorm:"column:isDeleted;default:0;comment:0: active, 1: deleted"`
	Sort             int            `gorm:"column:sort;default:0;comment:priority sort"`
	CreateTime       time.Time      `gorm:"column:createTime;autoCreateTime"`
	LastUpdate       time.Time      `gorm:"column:lastUpdate;autoUpdateTime"`
}

func (s *SPU) TableName() string {
	return "sd_product"
}
