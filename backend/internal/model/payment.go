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
	Amount                float64        `gorm:"not null" json:"amount"`
	Currency              string         `gorm:"default:'jpy'" json:"currency"`
	Status                string         `gorm:"default:'pending'" json:"status"` // pending, succeeded, failed, canceled
	CreatedAt             time.Time      `json:"created_at"`
	UpdatedAt             time.Time      `json:"updated_at"`
	DeletedAt             gorm.DeletedAt `gorm:"index" json:"-"`

	// リレーション
	Order Order `gorm:"foreignKey:OrderID" json:"order,omitempty"`
}