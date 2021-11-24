package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wasanx25/go_nextjs_authentication/backend/controller"
	"github.com/wasanx25/go_nextjs_authentication/backend/model"
	"gorm.io/gorm"
)

func TestFollowController_Create(t *testing.T) {
	// given
	loginUserID := uint(1)
	loginAuth0UserID := "test_auth0_user_id"

	userRepositoryMock := new(userRepositoryMock)
	userRepositoryMock.On("FindByAuth0UserID", loginAuth0UserID).Return(&model.User{Model: gorm.Model{ID: loginUserID}}, nil)

	followRepositoryMock := new(followRepositoryMock)
	followRepositoryMock.On("Create", loginUserID, uint(2), mock.Anything).Return(nil)

	request := httptest.NewRequest(http.MethodPost, "/", nil)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(request, recorder)

	c.Set("user", createJWTToken("test1", loginAuth0UserID))
	c.SetPath("/follow/:user_id")
	c.SetParamNames("user_id")
	c.SetParamValues("2")

	sut := controller.NewFollowController(followRepositoryMock, userRepositoryMock)

	// when
	actual := sut.Create(c)

	// then
	if assert.NoError(t, actual) {
		assert.Equal(t, http.StatusCreated, recorder.Code)
	}
}

func TestFollowController_Create_InternalServerError(t *testing.T) {
	// given
	loginUserID := uint(1)
	loginAuth0UserID := "test_auth0_user_id"

	userRepositoryMock := new(userRepositoryMock)
	userRepositoryMock.On("FindByAuth0UserID", loginAuth0UserID).Return(&model.User{Model: gorm.Model{ID: loginUserID}}, nil)

	followRepositoryMock := new(followRepositoryMock)
	followRepositoryMock.On("Create", loginUserID, uint(2), mock.Anything).Return(gorm.ErrInvalidDB)

	request := httptest.NewRequest(http.MethodPost, "/", nil)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(request, recorder)

	c.Set("user", createJWTToken("test1", loginAuth0UserID))
	c.SetPath("/follow/:user_id")
	c.SetParamNames("user_id")
	c.SetParamValues("2")

	sut := controller.NewFollowController(followRepositoryMock, userRepositoryMock)

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

func TestFollowController_Delete(t *testing.T) {
	// given
	loginUserID := uint(1)
	loginAuth0UserID := "test_auth0_user_id"

	userRepositoryMock := new(userRepositoryMock)
	userRepositoryMock.On("FindByAuth0UserID", loginAuth0UserID).Return(&model.User{Model: gorm.Model{ID: loginUserID}}, nil)

	followRepositoryMock := new(followRepositoryMock)
	followRepositoryMock.On("Delete", loginUserID, uint(2)).Return(nil)

	request := httptest.NewRequest(http.MethodDelete, "/", nil)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(request, recorder)

	c.Set("user", createJWTToken("test1", loginAuth0UserID))
	c.SetPath("/follow/:user_id")
	c.SetParamNames("user_id")
	c.SetParamValues("2")

	sut := controller.NewFollowController(followRepositoryMock, userRepositoryMock)

	// when
	actual := sut.Delete(c)

	// then
	if assert.NoError(t, actual) {
		assert.Equal(t, http.StatusNoContent, recorder.Code)
	}
}

func TestFollowController_Delete_InternalServerError(t *testing.T) {
	// given
	loginUserID := uint(1)
	loginAuth0UserID := "test_auth0_user_id"

	userRepositoryMock := new(userRepositoryMock)
	userRepositoryMock.On("FindByAuth0UserID", loginAuth0UserID).Return(&model.User{Model: gorm.Model{ID: loginUserID}}, nil)

	followRepositoryMock := new(followRepositoryMock)
	followRepositoryMock.On("Delete", loginUserID, uint(2)).Return(gorm.ErrInvalidDB)

	request := httptest.NewRequest(http.MethodDelete, "/", nil)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(request, recorder)

	c.Set("user", createJWTToken("test1", loginAuth0UserID))
	c.SetPath("/follow/:user_id")
	c.SetParamNames("user_id")
	c.SetParamValues("2")

	sut := controller.NewFollowController(followRepositoryMock, userRepositoryMock)

	// when
	actual := sut.Delete(c)

	// then
	if assert.Error(t, actual) {
		httpError, ok := actual.(*echo.HTTPError)
		if ok {
			assert.Equal(t, http.StatusInternalServerError, httpError.Code)
		}
	}
}

func TestFollowController_Create_InvalidBody(t *testing.T) {
	// given
	loginAuth0UserID := "test_auth0_user_id"

	userRepositoryMock := new(userRepositoryMock)
	userRepositoryMock.On("FindByAuth0UserID", loginAuth0UserID).Return(&model.User{}, nil)

	followRepositoryMock := new(followRepositoryMock)
	sut := controller.NewFollowController(followRepositoryMock, userRepositoryMock)

	e := echo.New()

	tests := []struct {
		testName string
		input    string
		call     func(c echo.Context) error
	}{
		{"Nothing values Create(..)", "", sut.Create},
		{"Invalid to_user_id value Create(..)", "two", sut.Create},
		{"Nothing values Delete(..)", "", sut.Delete},
		{"Invalid to_user_id value Delete(..)", "two", sut.Delete},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/", nil)
			request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			recorder := httptest.NewRecorder()

			c := e.NewContext(request, recorder)

			c.Set("user", createJWTToken("test1", loginAuth0UserID))
			c.SetPath("/follow/:user_id")
			c.SetParamNames("user_id")
			c.SetParamValues(tt.input)

			// when
			actual := tt.call(c)

			// then
			if assert.Error(t, actual) {
				httpError, ok := actual.(*echo.HTTPError)
				if ok {
					assert.Equal(t, http.StatusBadRequest, httpError.Code)
				}
			}
		})
	}
}
