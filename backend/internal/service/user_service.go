package service

import (
	"errors"
	"time"

	"github.com/Naonao3/EC-site/backend/internal/model"
	"github.com/Naonao3/EC-site/backend/internal/repository"
	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	Register(email, password, name string) (*model.User, error)
	Login(email, password string) (string, error)
	GetUserByID(id uint) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id uint) error
	ListUsers(page, pageSize int) ([]model.User, int64, error)
}

type userService struct {
	userRepo  repository.UserRepository
	jwtSecret string
}

func NewUserService(userRepo repository.UserRepository, jwtSecret string) UserService {
	return &userService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

func (s *userService) Register(email, password, name string) (*model.User, error) {
	// メールアドレスの重複チェック
	existingUser, _ := s.userRepo.GetByEmail(email)
	if existingUser != nil {
		return nil, errors.New("email already exists")
	}

	// ユーザー作成
	user := &model.User{
		Email: email,
		Name:  name,
		Role:  "customer",
	}

	// パスワードのハッシュ化
	if err := user.HashPassword(password); err != nil {
		return nil, err
	}

	// データベースに保存
	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) Login(email, password string) (string, error) {
	// ユーザー取得
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// パスワード検証
	if err := user.CheckPassword(password); err != nil {
		return "", errors.New("invalid email or password")
	}

	// JWTトークン生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *userService) GetUserByID(id uint) (*model.User, error) {
	return s.userRepo.GetByID(id)
}

func (s *userService) UpdateUser(user *model.User) error {
	// ユーザーの存在確認
	_, err := s.userRepo.GetByID(user.ID)
	if err != nil {
		return err
	}

	return s.userRepo.Update(user)
}

func (s *userService) DeleteUser(id uint) error {
	// ユーザーの存在確認
	_, err := s.userRepo.GetByID(id)
	if err != nil {
		return err
	}

	return s.userRepo.Delete(id)
}

func (s *userService) ListUsers(page, pageSize int) ([]model.User, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	return s.userRepo.List(page, pageSize)
}