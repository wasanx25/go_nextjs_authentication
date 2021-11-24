package repository

import (
	"time"

	"github.com/wasanx25/go_nextjs_authentication/backend/model"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

type PostRepositoryInterface interface {
	FindAllByFolloweeId(id uint) (posts []*model.Post, err error)
	Create(userID uint, text string, postedAt time.Time) error
}

func NewPostRepository(db *gorm.DB) PostRepositoryInterface {
	return &PostRepository{
		db: db,
	}
}

func (t *PostRepository) Create(userID uint, text string, postedAt time.Time) error {
	post := &model.Post{
		UserID:   userID,
		Text:     text,
		PostedAt: postedAt,
	}

	return t.db.Select("UserID", "Text", "PostedAt").Create(post).Error
}

// FindAllByFolloweeId TODO: FolloweeIdのみのnamingだが、実際には自分の投稿の取得のパラメータとしても使用しているので、namingに検討の余地がある
func (t *PostRepository) FindAllByFolloweeId(id uint) (posts []*model.Post, err error) {
	err = t.db.
		Joins("LEFT JOIN users ON users.id = posts.user_id").
		Joins("LEFT JOIN follows ON follows.to = users.id").
		Where("follows.from = ? OR posts.user_id = ?", id, id).
		Order("posts.posted_at DESC").
		Find(&posts).
		Error
	return
}
