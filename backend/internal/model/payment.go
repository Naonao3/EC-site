package model

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	ID                    uint           `gorm:"primarykey" json:"id"`
	OrderID               uint           `gorm:"not null;uniqueIndex" json:"order_id"`
	StripePaymentIntentID string         `gorm:"size:255;index" json:"stripe_payment_intent_id"`
	StripePaymentMethodID string         `gorm:"size:255" json:"stripe_payment_method_id,omitempty"`
	Amount                int            `gorm:"not null" json:"amount"` // 最小通貨単位（日本円の場合は円単位）
	Currency              string         `gorm:"default:'jpy'" json:"currency"`
	Status                string         `gorm:"default:'pending'" json:"status"` // pending, succeeded, failed, canceled
	CreatedAt             time.Time      `json:"created_at"`
	UpdatedAt             time.Time      `json:"updated_at"`
	DeletedAt             gorm.DeletedAt `gorm:"index" json:"-"`

	// リレーション
	Order Order `gorm:"foreignKey:OrderID" json:"order,omitempty"`
}