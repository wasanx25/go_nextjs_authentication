package model

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	UserID   uint
	Text     string
	PostedAt time.Time
}
