package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/wasanx25/go_nextjs_authentication/controller"
	"github.com/wasanx25/go_nextjs_authentication/model"
	"gorm.io/gorm"
)

func TestUserController_ListNoFollows(t *testing.T) {
	// given
	loginUserID := uint(1)
	loginAuth0UserID := "test_auth0_user_id"

	users := []*model.User{
		{
			Model: gorm.Model{
				ID: 1,
			},
			Username: "test1",
		},
		{
			Model: gorm.Model{
				ID: 2,
			},
			Username: "test2",
		},
	}
	userRepositoryMock := new(userRepositoryMock)
	userRepositoryMock.On("FindByAuth0UserID", loginAuth0UserID).Return(&model.User{Model: gorm.Model{ID: loginUserID}}, nil)
	userRepositoryMock.On("FindNoFollowUsersByUserID", loginUserID).Return(users, nil)

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(request, recorder)

	c.Set("user", createJWTToken("test1", loginAuth0UserID))

	sut := controller.NewUserController(userRepositoryMock)

	// when
	actual := sut.ListNoFollows(c)

	// then
	expectedJson := `{"users":[{"id":1,"username":"test1"},{"id":2,"username":"test2"}]}
`

	if assert.NoError(t, actual) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, expectedJson, recorder.Body.String())
	}
}

func TestUserController_ListNoFollows_EmptySlice(t *testing.T) {
	// given
	loginUserID := uint(1)
	loginAuth0UserID := "test_auth0_user_id"

	var users []*model.User
	userRepositoryMock := new(userRepositoryMock)
	userRepositoryMock.On("FindByAuth0UserID", loginAuth0UserID).Return(&model.User{Model: gorm.Model{ID: loginUserID}}, nil)
	userRepositoryMock.On("FindNoFollowUsersByUserID", loginUserID).Return(users, nil)

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(request, recorder)

	c.Set("user", createJWTToken("test1", loginAuth0UserID))

	sut := controller.NewUserController(userRepositoryMock)

	// when
	actual := sut.ListNoFollows(c)

	// then
	expectedJson := `{"users":[]}
`

	if assert.NoError(t, actual) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, expectedJson, recorder.Body.String())
	}
}

func TestUserController_ListNoFollows_InternalServerError(t *testing.T) {
	// given
	loginUserID := uint(1)
	loginAuth0UserID := "test_auth0_user_id"

	userRepositoryMock := new(userRepositoryMock)
	userRepositoryMock.On("FindByAuth0UserID", loginAuth0UserID).Return(&model.User{Model: gorm.Model{ID: loginUserID}}, nil)
	userRepositoryMock.On("FindNoFollowUsersByUserID", loginUserID).Return([]*model.User{}, gorm.ErrInvalidDB)

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(request, recorder)

	c.Set("user", createJWTToken("test1", loginAuth0UserID))

	sut := controller.NewUserController(userRepositoryMock)

	// when
	actual := sut.ListNoFollows(c)

	// then
	if assert.Error(t, actual) {
		httpError, ok := actual.(*echo.HTTPError)
		if ok {
			assert.Equal(t, http.StatusInternalServerError, httpError.Code)
		}
	}
}

func TestUserController_ListFollows(t *testing.T) {
	// given
	loginUserID := uint(1)
	loginAuth0UserID := "test_auth0_user_id"

	users := []*model.User{
		{
			Model: gorm.Model{
				ID: 1,
			},
			Username: "test1",
		},
		{
			Model: gorm.Model{
				ID: 2,
			},
			Username: "test2",
		},
	}
	userRepositoryMock := new(userRepositoryMock)
	userRepositoryMock.On("FindByAuth0UserID", loginAuth0UserID).Return(&model.User{Model: gorm.Model{ID: loginUserID}}, nil)
	userRepositoryMock.On("FindFollowUsersByUserID", loginUserID).Return(users, nil)

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(request, recorder)

	c.Set("user", createJWTToken("test1", loginAuth0UserID))

	sut := controller.NewUserController(userRepositoryMock)

	// when
	actual := sut.ListFollows(c)

	// then
	expectedJson := `{"users":[{"id":1,"username":"test1"},{"id":2,"username":"test2"}]}
`

	if assert.NoError(t, actual) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, expectedJson, recorder.Body.String())
	}
}

func TestUserController_ListFollows_EmptySlice(t *testing.T) {
	// given
	loginUserID := uint(1)
	loginAuth0UserID := "test_auth0_user_id"

	var users []*model.User
	userRepositoryMock := new(userRepositoryMock)
	userRepositoryMock.On("FindByAuth0UserID", loginAuth0UserID).Return(&model.User{Model: gorm.Model{ID: loginUserID}}, nil)
	userRepositoryMock.On("FindFollowUsersByUserID", loginUserID).Return(users, nil)

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(request, recorder)

	c.Set("user", createJWTToken("test1", loginAuth0UserID))

	sut := controller.NewUserController(userRepositoryMock)

	// when
	actual := sut.ListFollows(c)

	// then
	expectedJson := `{"users":[]}
`

	if assert.NoError(t, actual) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, expectedJson, recorder.Body.String())
	}
}

func TestUserController_ListFollows_InternalServerError(t *testing.T) {
	// given
	loginUserID := uint(1)
	loginAuth0UserID := "test_auth0_user_id"

	userRepositoryMock := new(userRepositoryMock)
	userRepositoryMock.On("FindByAuth0UserID", loginAuth0UserID).Return(&model.User{Model: gorm.Model{ID: loginUserID}}, nil)
	userRepositoryMock.On("FindFollowUsersByUserID", loginUserID).Return([]*model.User{}, gorm.ErrInvalidDB)

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(request, recorder)

	c.Set("user", createJWTToken("test1", loginAuth0UserID))

	sut := controller.NewUserController(userRepositoryMock)

	// when
	actual := sut.ListFollows(c)

	// then
	if assert.Error(t, actual) {
		httpError, ok := actual.(*echo.HTTPError)
		if ok {
			assert.Equal(t, http.StatusInternalServerError, httpError.Code)
		}
	}
}
