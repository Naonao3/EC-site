package service

import (
	"errors"

	"github.com/Naonao3/EC-site/backend/internal/model"
	"github.com/Naonao3/EC-site/backend/internal/repository"
	"gorm.io/gorm"
)

type OrderService interface {
	CreateOrder(userID uint, db *gorm.DB) (*model.Order, error)
	GetOrderByID(userID, orderID uint) (*model.Order, error)
	GetUserOrders(userID uint, page, pageSize int) ([]model.Order, int64, error)
	GetAllOrders(page, pageSize int) ([]model.Order, int64, error)
	UpdateOrderStatus(orderID uint, status string) error
}

type orderService struct {
	orderRepo   repository.OrderRepository
	cartRepo    repository.CartRepository
	productRepo repository.ProductRepository
}

func NewOrderService(
	orderRepo repository.OrderRepository,
	cartRepo repository.CartRepository,
	productRepo repository.ProductRepository,
) OrderService {
	return &orderService{
		orderRepo:   orderRepo,
		cartRepo:    cartRepo,
		productRepo: productRepo,
	}
}

// 注文作成（トランザクション処理）
func (s *orderService) CreateOrder(userID uint, db *gorm.DB) (*model.Order, error) {
	// トランザクション開始
	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	// エラー時のロールバック
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// カートアイテム取得
	cartItems, err := s.cartRepo.GetByUserID(userID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if len(cartItems) == 0 {
		tx.Rollback()
		return nil, errors.New("cart is empty")
	}

	// 注文作成
	order := &model.Order{
		UserID:      userID,
		TotalAmount: 0,
		Status:      "pending",
	}

	// 注文明細を作成し、合計金額を計算
	var orderItems []model.OrderItem
	totalAmount := 0.0

	for _, cartItem := range cartItems {
		// 商品取得
		product, err := s.productRepo.GetByID(cartItem.ProductID)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		// 在庫確認
		if product.Stock < cartItem.Quantity {
			tx.Rollback()
			return nil, errors.New("insufficient stock for product: " + product.Name)
		}

		// 在庫減少
		if err := s.productRepo.UpdateStock(product.ID, -cartItem.Quantity); err != nil {
			tx.Rollback()
			return nil, err
		}

		// 注文明細作成
		orderItem := model.OrderItem{
			ProductID: cartItem.ProductID,
			Quantity:  cartItem.Quantity,
			Price:     product.Price,
		}
		orderItems = append(orderItems, orderItem)

		// 合計金額計算
		totalAmount += product.Price * float64(cartItem.Quantity)
	}

	order.TotalAmount = totalAmount
	order.OrderItems = orderItems

	// 注文を保存
	if err := tx.Create(order).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// カートをクリア
	if err := s.cartRepo.DeleteByUserID(userID); err != nil {
		tx.Rollback()
		return nil, err
	}

	// トランザクションコミット
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	// 注文の完全な情報を取得して返す
	return s.orderRepo.GetByID(order.ID)
}

// 注文詳細取得
func (s *orderService) GetOrderByID(userID, orderID uint) (*model.Order, error) {
	order, err := s.orderRepo.GetByID(orderID)
	if err != nil {
		return nil, err
	}

	// ユーザーの所有確認
	if order.UserID != userID {
		return nil, errors.New("unauthorized")
	}

	return order, nil
}

// ユーザーの注文一覧取得
func (s *orderService) GetUserOrders(userID uint, page, pageSize int) ([]model.Order, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	return s.orderRepo.GetByUserID(userID, page, pageSize)
}

// 全注文取得（管理者用）
func (s *orderService) GetAllOrders(page, pageSize int) ([]model.Order, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	return s.orderRepo.List(page, pageSize)
}

// 注文ステータス更新
func (s *orderService) UpdateOrderStatus(orderID uint, status string) error {
	order, err := s.orderRepo.GetByID(orderID)
	if err != nil {
		return err
	}

	// ステータスのバリデーション
	validStatuses := map[string]bool{
		"pending":   true,
		"confirmed": true,
		"shipped":   true,
		"delivered": true,
		"cancelled": true,
	}

	if !validStatuses[status] {
		return errors.New("invalid status")
	}

	order.Status = status
	return s.orderRepo.Update(order)
}