package repository

import (
	pd "golang-user-manager/internal/profile/domain"
	"gorm.io/gorm"
)

type ProfileRepository interface {
	Update(p *pd.Profile) error
}

type GormProfileRepository struct{ db *gorm.DB }

func NewGormProfileRepository(db *gorm.DB) *GormProfileRepository {
	return &GormProfileRepository{db: db}
}

func (r *GormProfileRepository) Update(p *pd.Profile) error { return r.db.Save(p).Error }
