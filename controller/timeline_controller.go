package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/wasanx25/go_nextjs_authentication/authentication"
	"github.com/wasanx25/go_nextjs_authentication/repository"
)

type TimelineController struct {
	postRepository repository.PostRepositoryInterface
	userRepository repository.UserRepositoryInterface
}

type PostViewModel struct {
	PostID   uint      `json:"post_id"`
	UserID   uint      `json:"user_id"`
	Text     string    `json:"text"`
	PostedAt time.Time `json:"posted_at"`
}

type TimelineViewModel struct {
	Posts []PostViewModel `json:"posts"`
}

func NewTimelineController(postRepository repository.PostRepositoryInterface, userRepository repository.UserRepositoryInterface) *TimelineController {
	return &TimelineController{
		postRepository: postRepository,
		userRepository: userRepository,
	}
}

func (t *TimelineController) Index(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*authentication.JWTCustomClaims)
	subject := claims.Subject

	user, err := t.userRepository.FindByAuth0UserID(subject)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusUnauthorized, "不正なユーザーのリクエストを受け付けました")
	}

	posts, err := t.postRepository.FindAllByFolloweeId(user.ID)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "システムエラーが発生しました")
	}

	var postViewModels []PostViewModel
	postViewModels = []PostViewModel{}
	for _, post := range posts {
		postViewModels = append(postViewModels, PostViewModel{
			PostID:   post.ID,
			UserID:   post.UserID,
			Text:     post.Text,
			PostedAt: post.PostedAt,
		})
	}

	viewModel := TimelineViewModel{
		Posts: postViewModels,
	}

	return c.JSON(http.StatusOK, viewModel)
}
