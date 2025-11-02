package repository

import (
	"errors"

	"github.com/Naonao3/EC-site/backend/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *model.User) error
	GetByID(id uint) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Update(user *model.User) error
	Delete(id uint) error
	List(page, pageSize int) ([]model.User, int64, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetByID(id uint) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(user *model.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&model.User{}, id).Error
}

func (r *userRepository) List(page, pageSize int) ([]model.User, int64, error) {
	var users []model.User
	var total int64

	offset := (page - 1) * pageSize

	if err := r.db.Model(&model.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.Offset(offset).Limit(pageSize).Find(&users).Error
	return users, total, err
}