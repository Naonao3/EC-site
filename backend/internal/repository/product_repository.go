package repository

import (
	"errors"

	"github.com/Naonao3/EC-site/backend/internal/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *model.Product) error
	GetByID(id uint) (*model.Product, error)
	Update(product *model.Product) error
	Delete(id uint) error
	List(page, pageSize int) ([]model.Product, int64, error)
	GetByCategory(category string, page, pageSize int) ([]model.Product, int64, error)
	Search(keyword string, page, pageSize int) ([]model.Product, int64, error)
	UpdateStock(id uint, quantity int) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(product *model.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) GetByID(id uint) (*model.Product, error) {
	var product model.Product
	err := r.db.First(&product, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) Update(product *model.Product) error {
	return r.db.Save(product).Error
}

func (r *productRepository) Delete(id uint) error {
	return r.db.Delete(&model.Product{}, id).Error
}

func (r *productRepository) List(page, pageSize int) ([]model.Product, int64, error) {
	var products []model.Product
	var total int64

	offset := (page - 1) * pageSize

	if err := r.db.Model(&model.Product{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.Offset(offset).Limit(pageSize).Find(&products).Error
	return products, total, err
}

func (r *productRepository) GetByCategory(category string, page, pageSize int) ([]model.Product, int64, error) {
	var products []model.Product
	var total int64

	offset := (page - 1) * pageSize

	query := r.db.Model(&model.Product{}).Where("category = ?", category)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Offset(offset).Limit(pageSize).Find(&products).Error
	return products, total, err
}

func (r *productRepository) Search(keyword string, page, pageSize int) ([]model.Product, int64, error) {
	var products []model.Product
	var total int64

	offset := (page - 1) * pageSize

	query := r.db.Model(&model.Product{}).Where("name ILIKE ? OR description ILIKE ?", "%"+keyword+"%", "%"+keyword+"%")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Offset(offset).Limit(pageSize).Find(&products).Error
	return products, total, err
}

func (r *productRepository) UpdateStock(id uint, quantity int) error {
	return r.db.Model(&model.Product{}).Where("id = ?", id).Update("stock", gorm.Expr("stock + ?", quantity)).Error
}