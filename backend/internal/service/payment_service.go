package service

import (
	"errors"
	"fmt"

	"github.com/Naonao3/EC-site/backend/internal/model"
	"github.com/Naonao3/EC-site/backend/internal/repository"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/paymentintent"
)

type PaymentService interface {
	CreatePaymentIntent(orderID uint, userID uint) (string, error)
	HandlePaymentSuccess(paymentIntentID string) error
	GetPaymentByOrderID(orderID uint) (*model.Payment, error)
}

type paymentService struct {
	paymentRepo repository.PaymentRepository
	orderRepo   repository.OrderRepository
	stripeKey   string
}

func NewPaymentService(
	paymentRepo repository.PaymentRepository,
	orderRepo repository.OrderRepository,
	stripeKey string,
) PaymentService {
	// Stripeの秘密鍵を設定
	stripe.Key = stripeKey

	return &paymentService{
		paymentRepo: paymentRepo,
		orderRepo:   orderRepo,
		stripeKey:   stripeKey,
	}
}

// Payment Intent作成
func (s *paymentService) CreatePaymentIntent(orderID uint, userID uint) (string, error) {
	// 注文取得
	order, err := s.orderRepo.GetByID(orderID)
	if err != nil {
		return "", err
	}

	// ユーザーの所有確認
	if order.UserID != userID {
		return "", errors.New("unauthorized")
	}

	// 既に決済が存在するか確認
	existingPayment, err := s.paymentRepo.GetByOrderID(orderID)
	if err != nil {
		return "", err
	}

	// 既に決済が存在する場合は、既存のPayment IntentからClient Secretを取得して返す
	if existingPayment != nil {
		// Stripeから既存のPayment Intentを取得
		pi, err := paymentintent.Get(existingPayment.StripePaymentIntentID, nil)
		if err != nil {
			return "", fmt.Errorf("failed to retrieve existing payment intent: %w", err)
		}

		// Client Secretを返す
		return pi.ClientSecret, nil
	}

	// 金額を整数に変換（Stripeは最小通貨単位で扱う。日本円の場合は円単位）
	amount := int64(order.TotalAmount)

	// Stripe Payment Intent作成
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(amount),
		Currency: stripe.String(string(stripe.CurrencyJPY)),
		Metadata: map[string]string{
			"order_id": fmt.Sprintf("%d", orderID),
		},
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		return "", fmt.Errorf("failed to create payment intent: %w", err)
	}

	// Paymentレコード作成
	payment := &model.Payment{
		OrderID:               orderID,
		StripePaymentIntentID: pi.ID,
		Amount:                int(order.TotalAmount),
		Currency:              "jpy",
		Status:                "pending",
	}

	if err := s.paymentRepo.Create(payment); err != nil {
		return "", err
	}

	// Client Secretを返す（フロントエンドで使用）
	return pi.ClientSecret, nil
}

// 決済成功時の処理
func (s *paymentService) HandlePaymentSuccess(paymentIntentID string) error {
	// Payment取得
	payment, err := s.paymentRepo.GetByPaymentIntentID(paymentIntentID)
	if err != nil {
		return err
	}

	if payment == nil {
		return errors.New("payment not found")
	}

	// 既に処理済みの場合はスキップ
	if payment.Status == "succeeded" {
		return nil
	}

	// Payment更新
	payment.Status = "succeeded"
	if err := s.paymentRepo.Update(payment); err != nil {
		return err
	}

	// 注文ステータス更新
	order, err := s.orderRepo.GetByID(payment.OrderID)
	if err != nil {
		return err
	}

	order.Status = "confirmed"
	if err := s.orderRepo.Update(order); err != nil {
		return err
	}

	return nil
}

// 注文IDで決済取得
func (s *paymentService) GetPaymentByOrderID(orderID uint) (*model.Payment, error) {
	return s.paymentRepo.GetByOrderID(orderID)
}