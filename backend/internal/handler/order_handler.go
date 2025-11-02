package handler

import (
	"net/http"
	"strconv"

	"github.com/Naonao3/EC-site/backend/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderHandler struct {
	orderService service.OrderService
	db           *gorm.DB
}

func NewOrderHandler(orderService service.OrderService, db *gorm.DB) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
		db:           db,
	}
}

// UpdateOrderStatusRequest 注文ステータス更新リクエスト
type UpdateOrderStatusRequest struct {
	Status string `json:"status" binding:"required"`
}

// CreateOrder 注文作成
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	order, err := h.orderService.CreateOrder(userID.(uint), h.db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Order created successfully",
		"order":   order,
	})
}

// GetOrderByID 注文詳細取得
func (h *OrderHandler) GetOrderByID(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	orderID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	order, err := h.orderService.GetOrderByID(userID.(uint), uint(orderID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}

// GetUserOrders ユーザーの注文一覧取得
func (h *OrderHandler) GetUserOrders(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	orders, total, err := h.orderService.GetUserOrders(userID.(uint), page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"orders":    orders,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}

// GetAllOrders 全注文取得（管理者用）
func (h *OrderHandler) GetAllOrders(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	orders, total, err := h.orderService.GetAllOrders(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"orders":    orders,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}

// UpdateOrderStatus 注文ステータス更新（管理者用）
func (h *OrderHandler) UpdateOrderStatus(c *gin.Context) {
	orderID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	var req UpdateOrderStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.orderService.UpdateOrderStatus(uint(orderID), req.Status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order status updated successfully"})
}