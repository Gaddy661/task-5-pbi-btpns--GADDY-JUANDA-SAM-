package models

type Photo struct {
	ID       uint `gorm:"primaryKey"`
	Title    string
	Caption  string
	PhotoURL string `gorm:"not null"`
	UserID   uint   `gorm:"not null"`
	User     User   `gorm:"constraint:OnDelete:CASCADE;"`
}
