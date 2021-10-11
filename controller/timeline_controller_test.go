package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/wasanx25/go_nextjs_authentication/controller"
	"github.com/wasanx25/go_nextjs_authentication/model"
	"gorm.io/gorm"
)

func TestTimelineController_Index(t *testing.T) {
	// given
	loginUserID := uint(1)
	loginAuth0UserID := "test_auth0_user_id"

	userRepositoryMock := new(userRepositoryMock)
	userRepositoryMock.On("FindByAuth0UserID", loginAuth0UserID).Return(&model.User{Model: gorm.Model{ID: loginUserID}}, nil)

	posts := []*model.Post{
		{
			Model: gorm.Model{
				ID: 1,
			},
			UserID:   loginUserID,
			Text:     "一人目の投稿です",
			PostedAt: time.Unix(1000, 0),
		},
		{
			Model: gorm.Model{
				ID: 2,
			},
			UserID:   2,
			Text:     "二人目の投稿です",
			PostedAt: time.Unix(2000, 0),
		},
	}

	postRepositoryMock := new(postRepositoryMock)
	postRepositoryMock.On("FindAllByFolloweeId", loginUserID).Return(posts, nil)

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(request, recorder)

	c.Set("user", createJWTToken("test1", loginAuth0UserID))

	sut := controller.NewTimelineController(postRepositoryMock, userRepositoryMock)

	// when
	actual := sut.Index(c)

	// then
	firstPostedAt := time.Unix(1000, 0).Format(time.RFC3339)
	secondPostedAt := time.Unix(2000, 0).Format(time.RFC3339)
	expectedJson := `{"posts":[{"post_id":1,"user_id":1,"text":"一人目の投稿です","posted_at":"` + firstPostedAt + `"},{"post_id":2,"user_id":2,"text":"二人目の投稿です","posted_at":"` + secondPostedAt + `"}]}
`

	if assert.NoError(t, actual) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, expectedJson, recorder.Body.String())
	}
}

func TestTimelineController_Index_EmptySlice(t *testing.T) {
	// given
	loginUserID := uint(1)
	loginAuth0UserID := "test_auth0_user_id"

	userRepositoryMock := new(userRepositoryMock)
	userRepositoryMock.On("FindByAuth0UserID", loginAuth0UserID).Return(&model.User{Model: gorm.Model{ID: loginUserID}}, nil)

	var posts []*model.Post
	postRepositoryMock := new(postRepositoryMock)
	postRepositoryMock.On("FindAllByFolloweeId", loginUserID).Return(posts, nil)

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(request, recorder)

	c.Set("user", createJWTToken("test1", loginAuth0UserID))

	sut := controller.NewTimelineController(postRepositoryMock, userRepositoryMock)

	// when
	actual := sut.Index(c)

	// then
	expectedJson := `{"posts":[]}
`

	if assert.NoError(t, actual) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, expectedJson, recorder.Body.String())
	}
}

func TestTimelineController_Index_InternalServerError(t *testing.T) {
	// given
	loginUserID := uint(1)
	loginAuth0UserID := "test_auth0_user_id"

	userRepositoryMock := new(userRepositoryMock)
	userRepositoryMock.On("FindByAuth0UserID", loginAuth0UserID).Return(&model.User{Model: gorm.Model{ID: loginUserID}}, nil)

	postRepositoryMock := new(postRepositoryMock)
	postRepositoryMock.On("FindAllByFolloweeId", loginUserID).Return([]*model.Post{}, gorm.ErrInvalidDB)

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(request, recorder)

	c.Set("user", createJWTToken("test1", loginAuth0UserID))

	sut := controller.NewTimelineController(postRepositoryMock, userRepositoryMock)

	// when
	actual := sut.Index(c)

	// then
	if assert.Error(t, actual) {
		httpError, ok := actual.(*echo.HTTPError)
		if ok {
			assert.Equal(t, http.StatusInternalServerError, httpError.Code)
		}
	}
}
