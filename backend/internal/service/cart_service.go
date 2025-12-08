package service

import (
	"errors"

	"github.com/Naonao3/EC-site/backend/internal/model"
	"github.com/Naonao3/EC-site/backend/internal/repository"
)

type CartService interface {
	GetCart(userID uint) (*model.Cart, error)
	AddToCart(userID, productID uint, quantity int) (*model.CartItem, error)
	UpdateCartItem(userID, cartItemID uint, quantity int) (*model.CartItem, error)
	RemoveFromCart(userID, cartItemID uint) error
	ClearCart(userID uint) error
}

type cartService struct {
	cartRepo    repository.CartRepository
	productRepo repository.ProductRepository
}

func NewCartService(cartRepo repository.CartRepository, productRepo repository.ProductRepository) CartService {
	return &cartService{
		cartRepo:    cartRepo,
		productRepo: productRepo,
	}
}

// カート取得
func (s *cartService) GetCart(userID uint) (*model.Cart, error) {
	items, err := s.cartRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	// カート全体の情報を計算
	cart := &model.Cart{
		Items:      items,
		TotalItems: 0,
		TotalPrice: 0,
	}

	for _, item := range items {
		cart.TotalItems += item.Quantity
		cart.TotalPrice += item.Product.Price * float64(item.Quantity)
	}

	return cart, nil
}

// カートに追加
func (s *cartService) AddToCart(userID, productID uint, quantity int) (*model.CartItem, error) {
	// バリデーション
	if quantity <= 0 {
		return nil, errors.New("quantity must be greater than 0")
	}

	// 商品の存在確認
	product, err := s.productRepo.GetByID(productID)
	if err != nil {
		return nil, err
	}

	// 在庫確認
	if product.Stock < quantity {
		return nil, errors.New("insufficient stock")
	}

	// 既に同じ商品がカートにあるか確認
	existingItem, err := s.cartRepo.GetByUserAndProduct(userID, productID)
	if err != nil {
		return nil, err
	}

	var itemID uint
	if existingItem != nil {
		// 既にある場合は数量を更新
		newQuantity := existingItem.Quantity + quantity

		// 在庫確認
		if product.Stock < newQuantity {
			return nil, errors.New("insufficient stock")
		}

		existingItem.Quantity = newQuantity
		if err := s.cartRepo.Update(existingItem); err != nil {
			return nil, err
		}
		itemID = existingItem.ID
	} else {
		// 新規追加
		cartItem := &model.CartItem{
			UserID:    userID,
			ProductID: productID,
			Quantity:  quantity,
		}

		if err := s.cartRepo.Create(cartItem); err != nil {
			return nil, err
		}
		itemID = cartItem.ID
	}

	// 追加/更新されたアイテムを商品情報とともに取得
	items, err := s.cartRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	// 該当のアイテムを探して返す
	for _, item := range items {
		if item.ID == itemID {
			return &item, nil
		}
	}

	return nil, errors.New("failed to retrieve cart item")
}

// カートアイテム更新
func (s *cartService) UpdateCartItem(userID, cartItemID uint, quantity int) (*model.CartItem, error) {
	// バリデーション
	if quantity <= 0 {
		return nil, errors.New("quantity must be greater than 0")
	}

	// カートアイテム取得
	cartItem, err := s.cartRepo.GetByID(cartItemID)
	if err != nil {
		return nil, err
	}

	// ユーザーの所有確認
	if cartItem.UserID != userID {
		return nil, errors.New("unauthorized")
	}

	// 商品の在庫確認
	product, err := s.productRepo.GetByID(cartItem.ProductID)
	if err != nil {
		return nil, err
	}

	if product.Stock < quantity {
		return nil, errors.New("insufficient stock")
	}

	// 数量更新
	cartItem.Quantity = quantity
	if err := s.cartRepo.Update(cartItem); err != nil {
		return nil, err
	}

	// 更新されたアイテムを商品情報とともに取得
	items, err := s.cartRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	// 該当のアイテムを探して返す
	for _, item := range items {
		if item.ID == cartItemID {
			return &item, nil
		}
	}

	return nil, errors.New("failed to retrieve cart item")
}

// カートから削除
func (s *cartService) RemoveFromCart(userID, cartItemID uint) error {
	// カートアイテム取得
	cartItem, err := s.cartRepo.GetByID(cartItemID)
	if err != nil {
		return err
	}

	// ユーザーの所有確認
	if cartItem.UserID != userID {
		return errors.New("unauthorized")
	}

	return s.cartRepo.Delete(cartItemID)
}

// カートをクリア
func (s *cartService) ClearCart(userID uint) error {
	return s.cartRepo.DeleteByUserID(userID)
}