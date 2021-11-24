package repository

import (
	"github.com/wasanx25/go_nextjs_authentication/backend/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type UserRepositoryInterface interface {
	FindByUsername(username string) (user *model.User, err error)
	FindByAuth0UserID(id string) (user *model.User, err error)
	FindNoFollowUsersByUserID(id uint) (users []*model.User, err error)
	FindFollowUsersByUserID(id uint) (users []*model.User, err error)
	CreateIfNotExists(user model.User) error
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) FindByUsername(username string) (user *model.User, err error) {
	err = u.db.Where("username = ?", username).Find(&user).Error
	return
}

func (u *UserRepository) FindByAuth0UserID(id string) (user *model.User, err error) {
	err = u.db.Where("auth0_user_id = ?", id).Find(&user).Error
	return
}

func (u *UserRepository) FindNoFollowUsersByUserID(id uint) (users []*model.User, err error) {
	subQuery := u.db.Select("id").
		Joins("LEFT JOIN follows ON users.id = follows.to").
		Where("follows.from = ?", id).
		Table("users")
	err = u.db.Where("id NOT IN (?) AND id != ?", subQuery, id).Find(&users).Error
	return
}

func (u *UserRepository) FindFollowUsersByUserID(id uint) (users []*model.User, err error) {
	err = u.db.Joins("LEFT JOIN follows ON users.id = follows.to").
		Where("follows.from = ?", id).
		Order("follows.followed_at DESC").
		Find(&users).
		Error
	return
}

func (u *UserRepository) CreateIfNotExists(user model.User) error {
	return u.db.Where(model.User{Auth0UserID: user.Auth0UserID}).Attrs(user).FirstOrCreate(&user).Error
}
