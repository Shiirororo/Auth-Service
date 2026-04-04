package repository

import (
	"context"

	product_entity "github.com/user_service/internal/product/domain/model/entity"
)

type ProductRepository interface {
	Create(ctx context.Context, p *product_entity.Product) error
	GetByID(ctx context.Context, id uint64) (*product_entity.Product, error)
	Update(ctx context.Context, p *product_entity.Product) error
	SoftDelete(ctx context.Context, id uint64) error
	ListByShop(ctx context.Context, shopID int64, offset, limit int) ([]*product_entity.Product, error)
}

type SKURepository interface {
	Create(ctx context.Context, sku *product_entity.SKU) error
	GetBySkuNo(ctx context.Context, skuNo string) (*product_entity.SKU, error)
	UpdateStock(ctx context.Context, skuNo string, delta int) error
}

type SPUToSKURepository interface {
	Bind(ctx context.Context, rel *product_entity.SPUToSKU) error
	GetSKUsBySpuNo(ctx context.Context, spuNo string) ([]*product_entity.SPUToSKU, error)
}
