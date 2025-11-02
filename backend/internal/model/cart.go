package model

import (
	"time"
)

type CartItem struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	UserID    uint      `gorm:"not null;index" json:"user_id"`
	ProductID uint      `gorm:"not null" json:"product_id"`
	Quantity  int       `gorm:"not null" json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// リレーション
	User    User    `gorm:"foreignKey:UserID" json:"-"`
	Product Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
}

// カート全体の情報を返す構造体
type Cart struct {
	Items      []CartItem `json:"items"`
	TotalItems int        `json:"total_items"`
	TotalPrice float64    `json:"total_price"`
}