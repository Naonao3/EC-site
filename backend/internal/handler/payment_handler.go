package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Naonao3/EC-site/backend/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/webhook"
)

type PaymentHandler struct {
	paymentService service.PaymentService
	webhookSecret  string
}

func NewPaymentHandler(paymentService service.PaymentService, webhookSecret string) *PaymentHandler {
	return &PaymentHandler{
		paymentService: paymentService,
		webhookSecret:  webhookSecret,
	}
}

// CreatePaymentIntentRequest Payment Intent作成リクエスト
type CreatePaymentIntentRequest struct {
	OrderID uint `json:"order_id" binding:"required"`
}

// CreatePaymentIntent Payment Intent作成
func (h *PaymentHandler) CreatePaymentIntent(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var req CreatePaymentIntentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	clientSecret, err := h.paymentService.CreatePaymentIntent(req.OrderID, userID.(uint))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"client_secret": clientSecret,
	})
}

// HandleWebhook Stripe Webhook処理
func (h *PaymentHandler) HandleWebhook(c *gin.Context) {
	const MaxBodyBytes = int64(65536)
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MaxBodyBytes)

	payload, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Error reading request body"})
		return
	}

	// Webhook署名検証
	event, err := webhook.ConstructEvent(payload, c.GetHeader("Stripe-Signature"), h.webhookSecret)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Webhook signature verification failed"})
		return
	}

	// イベント処理
	switch event.Type {
	case "payment_intent.succeeded":
		// PaymentIntentオブジェクトを取得
		var paymentIntent stripe.PaymentIntent
		if err := json.Unmarshal(event.Data.Raw, &paymentIntent); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error parsing webhook JSON"})
			return
		}

		// 決済成功処理
		if err := h.paymentService.HandlePaymentSuccess(paymentIntent.ID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

	case "payment_intent.payment_failed":
		// 決済失敗処理（オプション）
		// 必要に応じて実装

	default:
		// その他のイベントは無視
	}

	c.JSON(http.StatusOK, gin.H{"received": true})
}

// GetPaymentByOrderID 注文の決済情報取得
func (h *PaymentHandler) GetPaymentByOrderID(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var req struct {
		OrderID uint `form:"order_id" binding:"required"`
	}

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payment, err := h.paymentService.GetPaymentByOrderID(req.OrderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if payment == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "payment not found"})
		return
	}

	// ユーザー確認（セキュリティ）
	if payment.Order.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"payment": payment})
}