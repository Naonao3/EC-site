package model

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	UserID      uint           `gorm:"not null" json:"user_id"`
	TotalAmount float64        `gorm:"not null" json:"total_amount"`
	Status      string         `gorm:"default:'pending'" json:"status"` // pending, confirmed, shipped, delivered, cancelled
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// リレーション
	User       User        `gorm:"foreignKey:UserID" json:"user,omitempty"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID" json:"order_items,omitempty"`
}

type OrderItem struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	OrderID   uint      `gorm:"not null" json:"order_id"`
	ProductID uint      `gorm:"not null" json:"product_id"`
	Quantity  int       `gorm:"not null" json:"quantity"`
	Price     float64   `gorm:"not null" json:"price"` // 注文時の価格
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// リレーション
	Product Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
}