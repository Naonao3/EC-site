package model

import (
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID              uint           `gorm:"primarykey" json:"id"`
	UserID          uint           `gorm:"not null" json:"user_id"`
	OrderNumber     string         `gorm:"type:varchar(50);uniqueIndex;not null" json:"order_number"`
	TotalAmount     float64        `gorm:"not null" json:"total_amount"`
	Status          string         `gorm:"default:'pending'" json:"status"` // pending, confirmed, shipped, delivered, cancelled
	ShippingAddress string         `gorm:"type:text" json:"shipping_address"`                       // nullable に変更
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`

	// リレーション
	User       User        `gorm:"foreignKey:UserID" json:"user,omitempty"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID" json:"order_items,omitempty"`
}

// BeforeCreate 注文作成前のフック（注文番号の自動生成）
func (o *Order) BeforeCreate(tx *gorm.DB) error {
	if o.OrderNumber == "" {
		// 注文番号の生成: ORD-{TIMESTAMP}-{RANDOM}
		timestamp := time.Now().Unix()
		random := rand.Intn(10000)
		o.OrderNumber = fmt.Sprintf("ORD-%d-%04d", timestamp, random)
	}
	return nil
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