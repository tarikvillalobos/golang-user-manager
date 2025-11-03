package domain

type Profile struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	UserID uint   `gorm:"uniqueIndex" json:"user_id"`
	Phone  string `gorm:"size:20" json:"phone"`
	Avatar string `json:"avatar"`
	Bio    string `gorm:"size:255" json:"bio"`
}
