package controller_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wasanx25/go_nextjs_authentication/controller"
	"github.com/wasanx25/go_nextjs_authentication/model"
	"gorm.io/gorm"
)

func TestPostController_Create(t *testing.T) {
	// given
	loginUserID := uint(1)
	loginAuth0UserID := "test_auth0_user_id"
	text := "テスト投稿です"

	userRepositoryMock := new(userRepositoryMock)
	userRepositoryMock.On("FindByAuth0UserID", loginAuth0UserID).Return(&model.User{Model: gorm.Model{ID: loginUserID}}, nil)

	postRepositoryMock := new(postRepositoryMock)
	postRepositoryMock.On("Create", loginUserID, text, mock.Anything).Return(nil)

	json := `{"text":"` + text + `"}`

	request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(json))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(request, recorder)

	c.Set("user", createJWTToken("test1", loginAuth0UserID))

	sut := controller.NewPostController(postRepositoryMock, userRepositoryMock)

	// when
	actual := sut.Create(c)

	// then
	if assert.NoError(t, actual) {
		assert.Equal(t, http.StatusCreated, recorder.Code)
	}
}

func TestPostController_Create_InvalidText(t *testing.T) {
	// given
	loginUserID := uint(1)
	loginAuth0UserID := "test_auth0_user_id"
	invalidText := "テストテストテストテストテストテストテストテストテストテストテストテストテストテストテスト" +
		"テストテストテストテストテストテストテストテストテストテストテストテストテストテストテストテストテストテスト" +
		"テストテストテストテストテストテストテストテストテストテストテストテストテストテスト" // 141文字

	userRepositoryMock := new(userRepositoryMock)
	userRepositoryMock.On("FindByAuth0UserID", loginAuth0UserID).Return(&model.User{Model: gorm.Model{ID: loginUserID}}, nil)
	postRepositoryMock := new(postRepositoryMock)

	json := `{"text":"` + invalidText + `"}`

	request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(json))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(request, recorder)

	c.Set("user", createJWTToken("test1", loginAuth0UserID))

	sut := controller.NewPostController(postRepositoryMock, userRepositoryMock)

	// when
	actual := sut.Create(c)

	// then
	if assert.Error(t, actual) {
		httpError, ok := actual.(*echo.HTTPError)
		if ok {
			assert.Equal(t, http.StatusBadRequest, httpError.Code)
		}
	}
}

func TestPostController_Create_InternalServerError(t *testing.T) {
	// given
	loginUserID := uint(1)
	loginAuth0UserID := "test_auth0_user_id"
	text := "テスト投稿です"

	userRepositoryMock := new(userRepositoryMock)
	userRepositoryMock.On("FindByAuth0UserID", loginAuth0UserID).Return(&model.User{Model: gorm.Model{ID: loginUserID}}, nil)

	postRepositoryMock := new(postRepositoryMock)
	postRepositoryMock.On("Create", loginUserID, text, mock.Anything).Return(gorm.ErrInvalidDB)

	json := `{"text":"` + text + `"}`

	request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(json))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(request, recorder)

	c.Set("user", createJWTToken("test1", loginAuth0UserID))

	sut := controller.NewPostController(postRepositoryMock, userRepositoryMock)

	// when
	actual := sut.Create(c)

	// then
	if assert.Error(t, actual) {
		httpError, ok := actual.(*echo.HTTPError)
		if ok {
			assert.Equal(t, http.StatusInternalServerError, httpError.Code)
		}
	}
}
