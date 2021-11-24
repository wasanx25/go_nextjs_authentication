package controller

import (
	"fmt"
	"net/http"
	"time"
	"unicode/utf8"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/wasanx25/go_nextjs_authentication/backend/authentication"
	repository2 "github.com/wasanx25/go_nextjs_authentication/backend/repository"
)

type PostController struct {
	postRepository repository2.PostRepositoryInterface
	userRepository repository2.UserRepositoryInterface
}

func NewPostController(postRepository repository2.PostRepositoryInterface, userRepository repository2.UserRepositoryInterface) *PostController {
	return &PostController{
		postRepository: postRepository,
		userRepository: userRepository,
	}
}

const TEXT_LIMIT = 140

type PostCreateBody struct {
	Text string `json:"text"`
}

func (t *PostController) Create(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*authentication.JWTCustomClaims)
	subject := claims.Subject

	user, err := t.userRepository.FindByAuth0UserID(subject)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusUnauthorized, "不正なユーザーのリクエストを受け付けました")
	}

	body := &PostCreateBody{}
	err = c.Bind(body)

	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "不正なパラメータを受け付けました")
	}

	if utf8.RuneCountInString(body.Text) > TEXT_LIMIT {
		return echo.NewHTTPError(http.StatusBadRequest, "投稿は140文字以内で入力してください")
	}

	postedAt := time.Now()

	err = t.postRepository.Create(user.ID, body.Text, postedAt)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "システムエラーが発生しました")
	}

	return c.JSON(http.StatusCreated, "")
}
