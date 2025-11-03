package repository

import (
	ud "golang-user-manager/internal/user/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	List() ([]ud.User, error)
	GetByID(id uint) (*ud.User, error)
	GetByEmail(email string) (*ud.User, error)
	Create(u *ud.User) error
	Update(u *ud.User) error
	Delete(id uint) error
	Exists(id uint) (bool, error)
	EmailExists(email string) (bool, error)
}

type GormUserRepository struct{ db *gorm.DB }

func NewGormUserRepository(db *gorm.DB) *GormUserRepository { return &GormUserRepository{db: db} }

func (r *GormUserRepository) List() ([]ud.User, error) {
	var users []ud.User
	return users, r.db.Preload("Profile").Find(&users).Error
}

func (r *GormUserRepository) GetByID(id uint) (*ud.User, error) {
	var u ud.User
	if err := r.db.Preload("Profile").First(&u, id).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *GormUserRepository) GetByEmail(email string) (*ud.User, error) {
	var u ud.User
	if err := r.db.Where("email = ?", email).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *GormUserRepository) Create(u *ud.User) error { return r.db.Create(u).Error }
func (r *GormUserRepository) Update(u *ud.User) error { return r.db.Save(u).Error }
func (r *GormUserRepository) Delete(id uint) error    { return r.db.Delete(&ud.User{}, id).Error }

func (r *GormUserRepository) Exists(id uint) (bool, error) {
	var count int64
	if err := r.db.Model(&ud.User{}).Where("id = ?", id).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *GormUserRepository) EmailExists(email string) (bool, error) {
	var count int64
	if err := r.db.Model(&ud.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
