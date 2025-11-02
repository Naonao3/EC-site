package service

import (
	"errors"

	"github.com/Naonao3/EC-site/backend/internal/model"
	"github.com/Naonao3/EC-site/backend/internal/repository"
)

type CartService interface {
	GetCart(userID uint) (*model.Cart, error)
	AddToCart(userID, productID uint, quantity int) error
	UpdateCartItem(userID, cartItemID uint, quantity int) error
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
func (s *cartService) AddToCart(userID, productID uint, quantity int) error {
	// バリデーション
	if quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}

	// 商品の存在確認
	product, err := s.productRepo.GetByID(productID)
	if err != nil {
		return err
	}

	// 在庫確認
	if product.Stock < quantity {
		return errors.New("insufficient stock")
	}

	// 既に同じ商品がカートにあるか確認
	existingItem, err := s.cartRepo.GetByUserAndProduct(userID, productID)
	if err != nil {
		return err
	}

	if existingItem != nil {
		// 既にある場合は数量を更新
		newQuantity := existingItem.Quantity + quantity
		
		// 在庫確認
		if product.Stock < newQuantity {
			return errors.New("insufficient stock")
		}

		existingItem.Quantity = newQuantity
		return s.cartRepo.Update(existingItem)
	}

	// 新規追加
	cartItem := &model.CartItem{
		UserID:    userID,
		ProductID: productID,
		Quantity:  quantity,
	}

	return s.cartRepo.Create(cartItem)
}

// カートアイテム更新
func (s *cartService) UpdateCartItem(userID, cartItemID uint, quantity int) error {
	// バリデーション
	if quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}

	// カートアイテム取得
	cartItem, err := s.cartRepo.GetByID(cartItemID)
	if err != nil {
		return err
	}

	// ユーザーの所有確認
	if cartItem.UserID != userID {
		return errors.New("unauthorized")
	}

	// 商品の在庫確認
	product, err := s.productRepo.GetByID(cartItem.ProductID)
	if err != nil {
		return err
	}

	if product.Stock < quantity {
		return errors.New("insufficient stock")
	}

	// 数量更新
	cartItem.Quantity = quantity
	return s.cartRepo.Update(cartItem)
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