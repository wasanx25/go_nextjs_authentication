package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/wasanx25/go_nextjs_authentication/backend/authentication"
	repository2 "github.com/wasanx25/go_nextjs_authentication/backend/repository"
)

type FollowController struct {
	followRepository repository2.FollowRepositoryInterface
	userRepository   repository2.UserRepositoryInterface
}

func NewFollowController(followRepository repository2.FollowRepositoryInterface, userRepository repository2.UserRepositoryInterface) *FollowController {
	return &FollowController{
		followRepository: followRepository,
		userRepository:   userRepository,
	}
}

type FollowCreateBody struct {
	ToUserID json.Number `json:"to_user_id"`
}

func (f *FollowController) Create(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*authentication.JWTCustomClaims)
	subject := claims.Subject

	user, err := f.userRepository.FindByAuth0UserID(subject)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusUnauthorized, "不正なユーザーのリクエストを受け付けました")
	}

	to, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "不正なパラメータを受け付けました")
	}

	err = f.followRepository.Create(user.ID, uint(to), time.Now())
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "システムエラーが発生しました")
	}

	return c.JSON(http.StatusCreated, "")
}

func (f *FollowController) Delete(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*authentication.JWTCustomClaims)
	subject := claims.Subject

	user, err := f.userRepository.FindByAuth0UserID(subject)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusUnauthorized, "不正なユーザーのリクエストを受け付けました")
	}

	to, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "不正なパラメータを受け付けました")
	}

	err = f.followRepository.Delete(user.ID, uint(to))
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "システムエラーが発生しました")
	}

	return c.JSON(http.StatusNoContent, "")
}
