package repository

import (
	sd "golang-user-manager/internal/subscription/domain"
	"gorm.io/gorm"
)

type SubscriptionRepository interface {
	Create(s *sd.Subscription) error
}

type GormSubscriptionRepository struct{ db *gorm.DB }

func NewGormSubscriptionRepository(db *gorm.DB) *GormSubscriptionRepository {
	return &GormSubscriptionRepository{db: db}
}

func (r *GormSubscriptionRepository) Create(s *sd.Subscription) error { return r.db.Create(s).Error }
