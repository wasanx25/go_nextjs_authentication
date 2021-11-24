package main

import (
	"time"

	"github.com/pelletier/go-toml"
	"github.com/wasanx25/go_nextjs_authentication/backend/config"
	"github.com/wasanx25/go_nextjs_authentication/backend/model"
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

	users := []*model.User{
		{
			Username:    "test1",
			Auth0UserID: "test_auth0_user_id_1",
		},
		{
			Username:    "test2",
			Auth0UserID: "test_auth0_user_id_2",
		},
		{
			Username:    "test3",
			Auth0UserID: "test_auth0_user_id_3",
		},
		{
			Username:    "test4",
			Auth0UserID: "test_auth0_user_id_4",
		},
	}

	posts := []*model.Post{
		{
			UserID:   1,
			Text:     "このサービスを初めて見ました！",
			PostedAt: time.Now(),
		},
		{
			UserID:   1,
			Text:     "テストで投稿しました！",
			PostedAt: time.Now().Add(10000000000),
		},
		{
			UserID:   1,
			Text:     "すごい！",
			PostedAt: time.Now().Add(20000000000),
		},
		{
			UserID:   2,
			Text:     "test2です！やりました！",
			PostedAt: time.Now().Add(30000000000),
		},
		{
			UserID:   2,
			Text:     "test1さんはすごいですね！",
			PostedAt: time.Now().Add(40000000000),
		},
		{
			UserID:   3,
			Text:     "これは一体...！",
			PostedAt: time.Now().Add(50000000000),
		},
		{
			UserID:   4,
			Text:     "うおおおお",
			PostedAt: time.Now().Add(60000000000),
		},
		{
			UserID:   4,
			Text:     "わおおおお",
			PostedAt: time.Now().Add(70000000000),
		},
	}

	follows := []*model.Follow{
		{
			From:       1,
			To:         2,
			FollowedAt: time.Now(),
		},
		{
			From:       1,
			To:         3,
			FollowedAt: time.Now(),
		},
		{
			From:       1,
			To:         4,
			FollowedAt: time.Now(),
		},
		{
			From:       2,
			To:         3,
			FollowedAt: time.Now(),
		},
		{
			From:       2,
			To:         4,
			FollowedAt: time.Now(),
		},
		{
			From:       3,
			To:         4,
			FollowedAt: time.Now(),
		},
	}

	db.Create(users)
	db.Create(posts)
	db.Create(follows)
}
