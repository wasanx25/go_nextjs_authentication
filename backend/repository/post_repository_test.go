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

func TestPostRepository_FindAllByFolloweeId(t *testing.T) {
	// setup
	fromUserID := uint(1)
	toUserID := uint(2)

	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	connection, err := db.DB()
	defer connection.Close()

	if err != nil {
		t.Fatal(err)
	}

	db.Create([]*model2.User{
		{Model: gorm.Model{ID: fromUserID}, Username: "test1", Auth0UserID: "test_auth0_user_id_1"},
		{Model: gorm.Model{ID: toUserID}, Username: "test2", Auth0UserID: "test_auth0_user_id_2"},
	})
	db.Create([]*model2.Follow{
		{From: fromUserID, To: toUserID, FollowedAt: time.Unix(0, 0)},
	})
	db.Create([]*model2.Post{
		{Model: gorm.Model{ID: 1, CreatedAt: time.Unix(0, 0), UpdatedAt: time.Unix(0, 0)}, UserID: toUserID, Text: "最初の投稿です！", PostedAt: time.Unix(1000*1000, 0)},
		{Model: gorm.Model{ID: 2, CreatedAt: time.Unix(0, 0), UpdatedAt: time.Unix(0, 0)}, UserID: toUserID, Text: "2回目の投稿です！", PostedAt: time.Unix(2000*1000, 0)},
		{Model: gorm.Model{ID: 3, CreatedAt: time.Unix(0, 0), UpdatedAt: time.Unix(0, 0)}, UserID: toUserID, Text: "3回目の投稿です！", PostedAt: time.Unix(3000*1000, 0)},
	})

	sut := repository.NewPostRepository(db)

	// when
	actual, err := sut.FindAllByFolloweeId(uint(fromUserID))
	if err != nil {
		t.Fatal(err)
	}

	// then
	expected := []*model2.Post{
		{Model: gorm.Model{ID: 3, CreatedAt: time.Unix(0, 0), UpdatedAt: time.Unix(0, 0)}, UserID: 2, Text: "3回目の投稿です！", PostedAt: time.Unix(3000*1000, 0)},
		{Model: gorm.Model{ID: 2, CreatedAt: time.Unix(0, 0), UpdatedAt: time.Unix(0, 0)}, UserID: 2, Text: "2回目の投稿です！", PostedAt: time.Unix(2000*1000, 0)},
		{Model: gorm.Model{ID: 1, CreatedAt: time.Unix(0, 0), UpdatedAt: time.Unix(0, 0)}, UserID: 2, Text: "最初の投稿です！", PostedAt: time.Unix(1000*1000, 0)},
	}

	assert.Equal(t, expected, actual)

	// cleanup
	cleanupTables(db)
}

func TestPostRepository_Create(t *testing.T) {
	// setup
	postUserID := uint(1)

	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	connection, err := db.DB()
	defer connection.Close()

	if err != nil {
		t.Fatal(err)
	}

	db.Create([]*model2.User{
		{Model: gorm.Model{ID: postUserID}, Username: "test1", Auth0UserID: "test_auth0_user_id_1"},
	})

	sut := repository.NewPostRepository(db)

	// when
	actual := sut.Create(postUserID, "テスト投稿です", time.Unix(0, 0))

	// then
	assert.Equal(t, nil, actual)

	// cleanup
	cleanupTables(db)
}
