package controller

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/wasanx25/go_nextjs_authentication/authentication"
	"github.com/wasanx25/go_nextjs_authentication/repository"
)

type UserController struct {
	userRepository repository.UserRepositoryInterface
}

func NewUserController(userRepository repository.UserRepositoryInterface) *UserController {
	return &UserController{
		userRepository: userRepository,
	}
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

type UsersResponse struct {
	Users []UserResponse `json:"users"`
}

func (u *UserController) ListNoFollows(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*authentication.JWTCustomClaims)
	subject := claims.Subject

	user, err := u.userRepository.FindByAuth0UserID(subject)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusUnauthorized, "不正なユーザーのリクエストを受け付けました")
	}

	users, err := u.userRepository.FindNoFollowUsersByUserID(user.ID)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "システムエラーが発生しました")
	}

	var userResponse []UserResponse
	userResponse = []UserResponse{}
	for _, user := range users {
		userResponse = append(userResponse, UserResponse{
			ID:       user.ID,
			Username: user.Username,
		})
	}

	response := UsersResponse{Users: userResponse}

	return c.JSON(http.StatusOK, response)
}

func (u *UserController) ListFollows(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*authentication.JWTCustomClaims)
	subject := claims.Subject

	user, err := u.userRepository.FindByAuth0UserID(subject)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusUnauthorized, "不正なユーザーのリクエストを受け付けました")
	}

	users, err := u.userRepository.FindFollowUsersByUserID(user.ID)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "システムエラーが発生しました")
	}

	var userResponse []UserResponse
	userResponse = []UserResponse{}
	for _, user := range users {
		userResponse = append(userResponse, UserResponse{
			ID:       user.ID,
			Username: user.Username,
		})
	}

	response := UsersResponse{Users: userResponse}

	return c.JSON(http.StatusOK, response)
}
