package domain

type Plan struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	Name        string  `gorm:"size:100;uniqueIndex;not null" json:"name"`
	Description string  `gorm:"size:255" json:"description"`
	Price       float64 `json:"price"`
	LimitUsers  int     `json:"limit_users"`
}
