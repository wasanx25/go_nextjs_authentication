package model

import (
	"time"
)

type Follow struct {
	From       uint `gorm:"primaryKey"`
	To         uint `gorm:"primaryKey"`
	FollowedAt time.Time
}
