package repository

import (
	pd "golang-user-manager/internal/plan/domain"
	"gorm.io/gorm"
)

type PlanRepository interface{ List() ([]pd.Plan, error) }

type GormPlanRepository struct{ db *gorm.DB }

func NewGormPlanRepository(db *gorm.DB) *GormPlanRepository { return &GormPlanRepository{db: db} }

func (r *GormPlanRepository) List() ([]pd.Plan, error) {
	var p []pd.Plan
	return p, r.db.Find(&p).Error
}
