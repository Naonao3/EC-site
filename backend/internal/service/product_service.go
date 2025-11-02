package service

import (
	"errors"

	"github.com/Naonao3/EC-site/backend/internal/model"
	"github.com/Naonao3/EC-site/backend/internal/repository"
)

type ProductService interface {
	CreateProduct(product *model.Product) error
	GetProductByID(id uint) (*model.Product, error)
	UpdateProduct(product *model.Product) error
	DeleteProduct(id uint) error
	ListProducts(page, pageSize int) ([]model.Product, int64, error)
	GetProductsByCategory(category string, page, pageSize int) ([]model.Product, int64, error)
	SearchProducts(keyword string, page, pageSize int) ([]model.Product, int64, error)
	UpdateStock(id uint, quantity int) error
}

type productService struct {
	productRepo repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
	return &productService{
		productRepo: productRepo,
	}
}

func (s *productService) CreateProduct(product *model.Product) error {
	// バリデーション
	if product.Name == "" {
		return errors.New("product name is required")
	}
	if product.Price <= 0 {
		return errors.New("product price must be greater than 0")
	}
	if product.Stock < 0 {
		return errors.New("product stock cannot be negative")
	}

	return s.productRepo.Create(product)
}

func (s *productService) GetProductByID(id uint) (*model.Product, error) {
	return s.productRepo.GetByID(id)
}

func (s *productService) UpdateProduct(product *model.Product) error {
	// 商品の存在確認
	_, err := s.productRepo.GetByID(product.ID)
	if err != nil {
		return err
	}

	// バリデーション
	if product.Name == "" {
		return errors.New("product name is required")
	}
	if product.Price <= 0 {
		return errors.New("product price must be greater than 0")
	}
	if product.Stock < 0 {
		return errors.New("product stock cannot be negative")
	}

	return s.productRepo.Update(product)
}

func (s *productService) DeleteProduct(id uint) error {
	// 商品の存在確認
	_, err := s.productRepo.GetByID(id)
	if err != nil {
		return err
	}

	return s.productRepo.Delete(id)
}

func (s *productService) ListProducts(page, pageSize int) ([]model.Product, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	return s.productRepo.List(page, pageSize)
}

func (s *productService) GetProductsByCategory(category string, page, pageSize int) ([]model.Product, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	return s.productRepo.GetByCategory(category, page, pageSize)
}

func (s *productService) SearchProducts(keyword string, page, pageSize int) ([]model.Product, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	return s.productRepo.Search(keyword, page, pageSize)
}

func (s *productService) UpdateStock(id uint, quantity int) error {
	// 商品の存在確認
	product, err := s.productRepo.GetByID(id)
	if err != nil {
		return err
	}

	// 在庫不足チェック
	if product.Stock+quantity < 0 {
		return errors.New("insufficient stock")
	}

	return s.productRepo.UpdateStock(id, quantity)
}