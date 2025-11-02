package repository

import (
	"errors"

	"github.com/Naonao3/EC-site/backend/internal/model"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	Create(payment *model.Payment) error
	GetByID(id uint) (*model.Payment, error)
	GetByOrderID(orderID uint) (*model.Payment, error)
	GetByPaymentIntentID(paymentIntentID string) (*model.Payment, error)
	Update(payment *model.Payment) error
	List(page, pageSize int) ([]model.Payment, int64, error)
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db: db}
}

// 決済作成
func (r *paymentRepository) Create(payment *model.Payment) error {
	return r.db.Create(payment).Error
}

// IDで決済取得
func (r *paymentRepository) GetByID(id uint) (*model.Payment, error) {
	var payment model.Payment
	err := r.db.Preload("Order").First(&payment, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("payment not found")
		}
		return nil, err
	}
	return &payment, nil
}

// 注文IDで決済取得
func (r *paymentRepository) GetByOrderID(orderID uint) (*model.Payment, error) {
	var payment model.Payment
	err := r.db.Where("order_id = ?", orderID).Preload("Order").First(&payment).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 見つからない場合はnilを返す
		}
		return nil, err
	}
	return &payment, nil
}

// Stripe Payment Intent IDで決済取得
func (r *paymentRepository) GetByPaymentIntentID(paymentIntentID string) (*model.Payment, error) {
	var payment model.Payment
	err := r.db.Where("stripe_payment_intent_id = ?", paymentIntentID).Preload("Order").First(&payment).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &payment, nil
}

// 決済更新
func (r *paymentRepository) Update(payment *model.Payment) error {
	return r.db.Save(payment).Error
}

// 決済一覧取得（管理者用）
func (r *paymentRepository) List(page, pageSize int) ([]model.Payment, int64, error) {
	var payments []model.Payment
	var total int64

	offset := (page - 1) * pageSize

	if err := r.db.Model(&model.Payment{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.Preload("Order").
		Order("created_at DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&payments).Error

	return payments, total, err
}