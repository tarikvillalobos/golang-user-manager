package domain

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	Email     string    `gorm:"size:120;uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"not null" json:"-"`
	Role      string    `gorm:"size:30;default:user" json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Profile   Profile   `json:"profile"`
}

type Profile struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	UserID uint   `gorm:"uniqueIndex" json:"user_id"`
	Phone  string `gorm:"size:20" json:"phone"`
	Avatar string `json:"avatar"`
	Bio    string `gorm:"size:255" json:"bio"`
}
