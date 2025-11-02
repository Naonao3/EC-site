package repository

import (
	"errors"

	"github.com/Naonao3/EC-site/backend/internal/model"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order *model.Order) error
	GetByID(id uint) (*model.Order, error)
	GetByUserID(userID uint, page, pageSize int) ([]model.Order, int64, error)
	Update(order *model.Order) error
	List(page, pageSize int) ([]model.Order, int64, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

// 注文作成
func (r *orderRepository) Create(order *model.Order) error {
	return r.db.Create(order).Error
}

// IDで注文取得
func (r *orderRepository) GetByID(id uint) (*model.Order, error) {
	var order model.Order
	err := r.db.Preload("OrderItems.Product").Preload("User").First(&order, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("order not found")
		}
		return nil, err
	}
	return &order, nil
}

// ユーザーIDで注文一覧取得
func (r *orderRepository) GetByUserID(userID uint, page, pageSize int) ([]model.Order, int64, error) {
	var orders []model.Order
	var total int64

	offset := (page - 1) * pageSize

	query := r.db.Model(&model.Order{}).Where("user_id = ?", userID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Preload("OrderItems.Product").
		Order("created_at DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&orders).Error

	return orders, total, err
}

// 注文更新
func (r *orderRepository) Update(order *model.Order) error {
	return r.db.Save(order).Error
}

// 注文一覧取得（管理者用）
func (r *orderRepository) List(page, pageSize int) ([]model.Order, int64, error) {
	var orders []model.Order
	var total int64

	offset := (page - 1) * pageSize

	if err := r.db.Model(&model.Order{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.Preload("OrderItems.Product").
		Preload("User").
		Order("created_at DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&orders).Error

	return orders, total, err
}