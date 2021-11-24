package repository_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	model2 "github.com/wasanx25/go_nextjs_authentication/backend/model"
	"github.com/wasanx25/go_nextjs_authentication/backend/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestFollowRepository_Create(t *testing.T) {
	// setup
	fromUserID := uint(1)
	toUserID := uint(2)

	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	connection, err := db.DB()
	defer connection.Close()

	db.Create([]*model2.User{
		{Model: gorm.Model{ID: fromUserID}, Username: "test1", Auth0UserID: "test_auth0_user_id_1"},
		{Model: gorm.Model{ID: toUserID}, Username: "test2", Auth0UserID: "test_auth0_user_id_2"},
	})

	sut := repository.NewFollowRepository(db)

	// when
	actual := sut.Create(fromUserID, toUserID, time.Unix(0, 0))

	// then
	assert.Equal(t, nil, actual)

	// cleanup
	cleanupTables(db)
}

func TestFollowRepository_Delete(t *testing.T) {
	// setup
	fromUserID := uint(1)
	toUserID := uint(2)

	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	connection, err := db.DB()
	defer connection.Close()

	db.Create([]*model2.User{
		{Model: gorm.Model{ID: fromUserID}, Username: "test1", Auth0UserID: "test_auth0_user_id_1"},
		{Model: gorm.Model{ID: toUserID}, Username: "test2", Auth0UserID: "test_auth0_user_id_2"},
	})
	db.Create([]*model2.Follow{
		{From: fromUserID, To: toUserID, FollowedAt: time.Unix(0, 0)},
	})

	sut := repository.NewFollowRepository(db)

	// when
	actual := sut.Delete(fromUserID, toUserID)

	// then
	assert.Equal(t, nil, actual)

	// cleanup
	cleanupTables(db)
}
