package repository

import (
	ud "golang-user-manager/internal/user/domain"
	"gorm.io/gorm"
)

type AuthRepository interface {
	GetByEmail(email string) (*ud.User, error)
	CreateUser(u *ud.User) error
}

type GormAuthRepository struct{ db *gorm.DB }

func NewGormAuthRepository(db *gorm.DB) *GormAuthRepository { return &GormAuthRepository{db: db} }

func (r *GormAuthRepository) GetByEmail(email string) (*ud.User, error) {
	var u ud.User
	if err := r.db.Where("email=?", email).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}
func (r *GormAuthRepository) CreateUser(u *ud.User) error { return r.db.Create(u).Error }
