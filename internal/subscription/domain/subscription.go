package domain

import "time"

type Subscription struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	PlanID    uint      `gorm:"not null" json:"plan_id"`
	Status    string    `gorm:"size:20;default:active" json:"status"`
	StartedAt time.Time `json:"started_at"`
	EndsAt    time.Time `json:"ends_at"`
}
