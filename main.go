package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pelletier/go-toml"
	"github.com/wasanx25/go_nextjs_authentication/authentication"
	"github.com/wasanx25/go_nextjs_authentication/config"
	"github.com/wasanx25/go_nextjs_authentication/controller"
	"github.com/wasanx25/go_nextjs_authentication/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	conf := config.Config{}
	tree, err := toml.LoadFile("config.toml")
	if err != nil {
		panic(err)
	}

	err = tree.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(mysql.Open(conf.DatabaseDSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	connection, err := db.DB()
	defer connection.Close()

	if err != nil {
		panic(err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: conf.AllowOrigins,
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	postRepository := repository.NewPostRepository(db)
	userRepository := repository.NewUserRepository(db)
	followRepository := repository.NewFollowRepository(db)

	timelineController := controller.NewTimelineController(postRepository, userRepository)
	postController := controller.NewPostController(postRepository, userRepository)
	userController := controller.NewUserController(userRepository)
	followController := controller.NewFollowController(followRepository, userRepository)

	r := e.Group("")

	r.Use(middleware.JWTWithConfig(authentication.JWTConfig(conf, userRepository)))
	r.GET("/timeline", timelineController.Index)
	r.POST("/post", postController.Create)
	r.GET("/no_follow_users", userController.ListNoFollows)
	r.GET("/follow_users", userController.ListFollows)
	r.POST("/follow/:user_id", followController.Create)
	r.DELETE("/follow/:user_id", followController.Delete)

	e.Logger.Fatal(e.Start(":1323"))
}
