package repository

import (
	"time"

	"github.com/wasanx25/go_nextjs_authentication/backend/model"
	"gorm.io/gorm"
)

type FollowRepository struct {
	db *gorm.DB
}

type FollowRepositoryInterface interface {
	Create(fromUserID uint, toUserID uint, followedAt time.Time) error
	Delete(fromUserID uint, toUserID uint) error
}

func NewFollowRepository(db *gorm.DB) FollowRepositoryInterface {
	return &FollowRepository{
		db: db,
	}
}

func (f *FollowRepository) Create(fromUserID uint, toUserID uint, followedAt time.Time) error {
	follow := &model.Follow{
		From:       fromUserID,
		To:         toUserID,
		FollowedAt: followedAt,
	}
	return f.db.Select("From", "To", "FollowedAt").Create(follow).Error
}

func (f *FollowRepository) Delete(fromUserID uint, toUserID uint) error {
	return f.db.Delete(model.Follow{}, "`from` = ? AND `to` = ?", fromUserID, toUserID).Error
}
