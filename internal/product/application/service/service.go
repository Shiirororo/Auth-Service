package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/shopspring/decimal"
	product_entity "github.com/user_service/internal/product/domain/model/entity"
	"github.com/user_service/internal/product/domain/repository"
)

type ProductServiceInterface interface {
	CreateProduct(ctx context.Context, p *product_entity.Product) error
	GetProduct(ctx context.Context, id uint64) (*product_entity.Product, error)
	UpdateProduct(ctx context.Context, p *product_entity.Product) error
	DeleteProduct(ctx context.Context, id uint64) error
	ListByShop(ctx context.Context, shopID int64, offset, limit int) ([]*product_entity.Product, error)

	CreateSKU(ctx context.Context, sku *product_entity.SKU, attr *product_entity.SKUAttr) error
	GetSKU(ctx context.Context, skuNo string) (*product_entity.SKU, error)
	UpdateStock(ctx context.Context, skuNo string, delta int) error

	BindSKUToProduct(ctx context.Context, spuNo, skuNo string) error
	GetProductSKUs(ctx context.Context, spuNo string) ([]*product_entity.SPUToSKU, error)
}

type ProductService struct {
	productRepo repository.ProductRepository
	skuRepo     repository.SKURepository
	spuSkuRepo  repository.SPUToSKURepository
}

func NewProductService(
	productRepo repository.ProductRepository,
	skuRepo repository.SKURepository,
	spuSkuRepo repository.SPUToSKURepository,
) ProductServiceInterface {
	return &ProductService{
		productRepo: productRepo,
		skuRepo:     skuRepo,
		spuSkuRepo:  spuSkuRepo,
	}
}

func (s *ProductService) CreateProduct(ctx context.Context, p *product_entity.Product) error {
	if p.ProductName == "" {
		return errors.New("product name is required")
	}
	return s.productRepo.Create(ctx, p)
}

func (s *ProductService) GetProduct(ctx context.Context, id uint64) (*product_entity.Product, error) {
	return s.productRepo.GetByID(ctx, id)
}

func (s *ProductService) UpdateProduct(ctx context.Context, p *product_entity.Product) error {
	if p.ProductID == 0 {
		return errors.New("product id is required")
	}
	return s.productRepo.Update(ctx, p)
}

func (s *ProductService) DeleteProduct(ctx context.Context, id uint64) error {
	return s.productRepo.SoftDelete(ctx, id)
}

func (s *ProductService) ListByShop(ctx context.Context, shopID int64, offset, limit int) ([]*product_entity.Product, error) {
	return s.productRepo.ListByShop(ctx, shopID, offset, limit)
}

func (s *ProductService) CreateSKU(ctx context.Context, sku *product_entity.SKU, attr *product_entity.SKUAttr) error {
	if sku.SkuNo == "" {
		return errors.New("skuNo is required")
	}
	if sku.SkuPrice.LessThanOrEqual(decimal.Zero) {
		return errors.New("sku price must be greater than zero")
	}
	if err := s.skuRepo.Create(ctx, sku); err != nil {
		return fmt.Errorf("create sku: %w", err)
	}
	return nil
}

func (s *ProductService) GetSKU(ctx context.Context, skuNo string) (*product_entity.SKU, error) {
	return s.skuRepo.GetBySkuNo(ctx, skuNo)
}

func (s *ProductService) UpdateStock(ctx context.Context, skuNo string, delta int) error {
	if delta == 0 {
		return nil
	}
	return s.skuRepo.UpdateStock(ctx, skuNo, delta)
}

func (s *ProductService) BindSKUToProduct(ctx context.Context, spuNo, skuNo string) error {
	return s.spuSkuRepo.Bind(ctx, &product_entity.SPUToSKU{SpuNo: spuNo, SkuNo: skuNo})
}

func (s *ProductService) GetProductSKUs(ctx context.Context, spuNo string) ([]*product_entity.SPUToSKU, error) {
	return s.spuSkuRepo.GetSKUsBySpuNo(ctx, spuNo)
}
