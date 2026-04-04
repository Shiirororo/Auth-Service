package product_entity

import (
	"time"

	"gorm.io/datatypes"
)

type ProductStatus int8

const (
	ProductOutOfStock ProductStatus = 0
	ProductInStock    ProductStatus = 1
)

type Product struct {
	ProductID        uint64         `gorm:"column:productId;primaryKey;autoIncrement"`
	ProductName      string         `gorm:"column:productName;size:64"`
	ProductDesc      string         `gorm:"column:productDesc;size:256"`
	ProductStatus    ProductStatus  `gorm:"column:productStatus"`
	ProductAttribute datatypes.JSON `gorm:"column:productAttribute"`
	ProductShopID    int64          `gorm:"column:productShopID"`
	IsDeleted        uint8          `gorm:"column:isDeleted;default:0"`
	Sort             int            `gorm:"column:sort;default:0"`
	CreateTime       time.Time      `gorm:"column:createTime;autoCreateTime"`
	LastUpdate       time.Time      `gorm:"column:lastUpdate;autoUpdateTime"`
}

func (Product) TableName() string { return "sd_product" }
