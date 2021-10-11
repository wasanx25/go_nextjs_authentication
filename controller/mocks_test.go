package controller_test

import (
	"html/template"
	"io"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"github.com/wasanx25/go_nextjs_authentication/authentication"
	"github.com/wasanx25/go_nextjs_authentication/model"
)

func createJWTToken(username string, auth0UserID string) *jwt.Token {
	return &jwt.Token{
		Claims: &authentication.JWTCustomClaims{
			Name:           username,
			StandardClaims: jwt.StandardClaims{Subject: auth0UserID},
		},
	}
}

type templateMock struct {
	templates *template.Template
}

func (t *templateMock) Render(w io.Writer, name string, data interface{}, _c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func createTemplateMock(htmlName string, text string) *templateMock {
	return &templateMock{
		templates: template.Must(template.New(htmlName).Parse(text)),
	}
}

type followRepositoryMock struct {
	mock.Mock
}

func (f *followRepositoryMock) Create(fromUserID uint, toUserID uint, followedAt time.Time) error {
	args := f.Called(fromUserID, toUserID, followedAt)
	return args.Error(0)
}

func (f *followRepositoryMock) Delete(fromUserID uint, toUserID uint) error {
	args := f.Called(fromUserID, toUserID)
	return args.Error(0)
}

type userRepositoryMock struct {
	mock.Mock
}

func (u *userRepositoryMock) FindByUsername(username string) (user *model.User, err error) {
	args := u.Called(username)
	return args.Get(0).(*model.User), args.Error(1)
}

func (u *userRepositoryMock) FindByAuth0UserID(id string) (user *model.User, err error) {
	args := u.Called(id)
	return args.Get(0).(*model.User), args.Error(1)
}

func (u *userRepositoryMock) FindNoFollowUsersByUserID(id uint) (users []*model.User, err error) {
	args := u.Called(id)
	return args.Get(0).([]*model.User), args.Error(1)
}

func (u *userRepositoryMock) FindFollowUsersByUserID(id uint) (users []*model.User, err error) {
	args := u.Called(id)
	return args.Get(0).([]*model.User), args.Error(1)
}

func (u *userRepositoryMock) CreateIfNotExists(user model.User) error {
	args := u.Called(user)
	return args.Error(0)
}

type postRepositoryMock struct {
	mock.Mock
}

func (t *postRepositoryMock) FindAllByFolloweeId(id uint) (posts []*model.Post, err error) {
	args := t.Called(id)
	return args.Get(0).([]*model.Post), args.Error(1)
}

func (t *postRepositoryMock) Create(userID uint, text string, postedAt time.Time) error {
	args := t.Called(userID, text, postedAt)
	return args.Error(0)
}
