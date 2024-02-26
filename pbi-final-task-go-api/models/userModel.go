package models

import "time"

type User struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
}
