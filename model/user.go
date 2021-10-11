package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string   `gorm:"size:255;uniqueIndex"`
	Auth0UserID string   `gorm:"size:255;uniqueIndex"`
	Posts       []Post   `gorm:"foreignKey:UserID"`
	Follower    []Follow `gorm:"foreignKey:To"`
	Followee    []Follow `gorm:"foreignKey:From"`
}
