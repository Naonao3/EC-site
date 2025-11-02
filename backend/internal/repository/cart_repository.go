package repository

import (
	"errors"

	"github.com/Naonao3/EC-site/backend/internal/model"
	"gorm.io/gorm"
)

type CartRepository interface {
	GetByUserID(userID uint) ([]model.CartItem, error)
	GetByID(id uint) (*model.CartItem, error)
	GetByUserAndProduct(userID, productID uint) (*model.CartItem, error)
	Create(cartItem *model.CartItem) error
	Update(cartItem *model.CartItem) error
	Delete(id uint) error
	DeleteByUserID(userID uint) error
}

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db: db}
}

// ユーザーのカートアイテム全取得
func (r *cartRepository) GetByUserID(userID uint) ([]model.CartItem, error) {
	var items []model.CartItem
	err := r.db.Where("user_id = ?", userID).Preload("Product").Find(&items).Error
	return items, err
}

// IDでカートアイテム取得
func (r *cartRepository) GetByID(id uint) (*model.CartItem, error) {
	var item model.CartItem
	err := r.db.First(&item, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("cart item not found")
		}
		return nil, err
	}
	return &item, nil
}

// ユーザーと商品でカートアイテム取得
func (r *cartRepository) GetByUserAndProduct(userID, productID uint) (*model.CartItem, error) {
	var item model.CartItem
	err := r.db.Where("user_id = ? AND product_id = ?", userID, productID).First(&item).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 見つからない場合はnilを返す（エラーではない）
		}
		return nil, err
	}
	return &item, nil
}

// カートアイテム作成
func (r *cartRepository) Create(cartItem *model.CartItem) error {
	return r.db.Create(cartItem).Error
}

// カートアイテム更新
func (r *cartRepository) Update(cartItem *model.CartItem) error {
	return r.db.Save(cartItem).Error
}

// カートアイテム削除
func (r *cartRepository) Delete(id uint) error {
	return r.db.Delete(&model.CartItem{}, id).Error
}

// ユーザーのカート全削除（注文完了時など）
func (r *cartRepository) DeleteByUserID(userID uint) error {
	return r.db.Where("user_id = ?", userID).Delete(&model.CartItem{}).Error
}