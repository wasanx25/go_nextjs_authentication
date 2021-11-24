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

func TestUserRepository_FindByUsername(t *testing.T) {
	// setup
	username := "test1"

	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	connection, err := db.DB()
	defer connection.Close()

	db.Create([]*model2.User{
		{Model: gorm.Model{ID: 1, CreatedAt: time.Unix(0, 0), UpdatedAt: time.Unix(0, 0)}, Username: username, Auth0UserID: "test_auth0_user_id_1"},
	})

	sut := repository.NewUserRepository(db)

	// when
	actual, err := sut.FindByUsername(username)
	if err != nil {
		t.Fatal(err)
	}

	// then
	expected := &model2.User{
		Model: gorm.Model{ID: 1, CreatedAt: time.Unix(0, 0), UpdatedAt: time.Unix(0, 0)}, Username: "test1", Auth0UserID: "test_auth0_user_id_1",
	}
	assert.Equal(t, expected, actual)

	// cleanup
	cleanupTables(db)
}

func TestUserRepository_FindByVendorUserID(t *testing.T) {
	// setup
	vendorUserID := "test_auth0_user_id_1"

	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	connection, err := db.DB()
	defer connection.Close()

	db.Create([]*model2.User{
		{Model: gorm.Model{ID: 1, CreatedAt: time.Unix(0, 0), UpdatedAt: time.Unix(0, 0)}, Username: "test1", Auth0UserID: vendorUserID},
	})

	sut := repository.NewUserRepository(db)

	// when
	actual, err := sut.FindByAuth0UserID(vendorUserID)
	if err != nil {
		t.Fatal(err)
	}

	// then
	expected := &model2.User{
		Model: gorm.Model{ID: 1, CreatedAt: time.Unix(0, 0), UpdatedAt: time.Unix(0, 0)}, Username: "test1", Auth0UserID: "test_auth0_user_id_1",
	}
	assert.Equal(t, expected, actual)

	// cleanup
	cleanupTables(db)
}

func TestUserRepository_FindNoFollowUsersByUserID(t *testing.T) {
	// setup
	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	connection, err := db.DB()
	defer connection.Close()

	userID := uint(1)

	db.Create([]*model2.User{
		{Model: gorm.Model{ID: userID, CreatedAt: time.Unix(0, 0), UpdatedAt: time.Unix(0, 0)}, Username: "test1", Auth0UserID: "test_auth0_user_id_1"},
		{Model: gorm.Model{ID: 2, CreatedAt: time.Unix(0, 0), UpdatedAt: time.Unix(0, 0)}, Username: "test2", Auth0UserID: "test_auth0_user_id_2"},
		{Model: gorm.Model{ID: 3, CreatedAt: time.Unix(0, 0), UpdatedAt: time.Unix(0, 0)}, Username: "test3", Auth0UserID: "test_auth0_user_id_3"},
		{Model: gorm.Model{ID: 4, CreatedAt: time.Unix(0, 0), UpdatedAt: time.Unix(0, 0)}, Username: "test4", Auth0UserID: "test_auth0_user_id_4"},
		{Model: gorm.Model{ID: 5, CreatedAt: time.Unix(0, 0), UpdatedAt: time.Unix(0, 0)}, Username: "test5", Auth0UserID: "test_auth0_user_id_5"},
	})

	db.Create([]*model2.Follow{
		{From: userID, To: 2, FollowedAt: time.Unix(0, 0)},
		{From: userID, To: 3, FollowedAt: time.Unix(0, 0)},
	})

	sut := repository.NewUserRepository(db)

	// when
	actual, err := sut.FindNoFollowUsersByUserID(userID)
	if err != nil {
		t.Fatal(err)
	}

	// then
	expected := []*model2.User{
		{Model: gorm.Model{ID: 4, CreatedAt: time.Unix(0, 0), UpdatedAt: time.Unix(0, 0)}, Username: "test4", Auth0UserID: "test_auth0_user_id_4"},
		{Model: gorm.Model{ID: 5, CreatedAt: time.Unix(0, 0), UpdatedAt: time.Unix(0, 0)}, Username: "test5", Auth0UserID: "test_auth0_user_id_5"},
	}
	assert.Equal(t, expected, actual)

	// cleanup
	cleanupTables(db)
}

func TestUserRepository_FindFollowUsersByUserID(t *testing.T) {
	// setup
	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	connection, err := db.DB()
	defer connection.Close()

	userID := uint(1)

	db.Create([]*model2.User{
		{Model: gorm.Model{ID: userID, CreatedAt: time.Unix(0, 0), UpdatedAt: time.Unix(0, 0)}, Username: "test1", Auth0UserID: "test_auth0_user_id_1"},
		{Model: gorm.Model{ID: 2, CreatedAt: time.Unix(0, 0), UpdatedAt: time.Unix(0, 0)}, Username: "test2", Auth0UserID: "test_auth0_user_id_2"},
		{Model: gorm.Model{ID: 3, CreatedAt: time.Unix(0, 0), UpdatedAt: time.Unix(0, 0)}, Username: "test3", Auth0UserID: "test_auth0_user_id_3"},
		{Model: gorm.Model{ID: 4, CreatedAt: time.Unix(0, 0), UpdatedAt: time.Unix(0, 0)}, Username: "test4", Auth0UserID: "test_auth0_user_id_4"},
		{Model: gorm.Model{ID: 5, CreatedAt: time.Unix(0, 0), UpdatedAt: time.Unix(0, 0)}, Username: "test5", Auth0UserID: "test_auth0_user_id_5"},
	})

	db.Create([]*model2.Follow{
		{From: userID, To: 2, FollowedAt: time.Unix(0, 0)},
		{From: userID, To: 3, FollowedAt: time.Unix(1000, 0)},
	})

	sut := repository.NewUserRepository(db)

	// when
	actual, err := sut.FindFollowUsersByUserID(userID)
	if err != nil {
		t.Fatal(err)
	}

	// then
	expected := []*model2.User{
		{Model: gorm.Model{ID: 3, CreatedAt: time.Unix(0, 0), UpdatedAt: time.Unix(0, 0)}, Username: "test3", Auth0UserID: "test_auth0_user_id_3"},
		{Model: gorm.Model{ID: 2, CreatedAt: time.Unix(0, 0), UpdatedAt: time.Unix(0, 0)}, Username: "test2", Auth0UserID: "test_auth0_user_id_2"},
	}
	assert.Equal(t, expected, actual)

	// cleanup
	cleanupTables(db)
}

func TestUserRepository_CreateIfNotExists(t *testing.T) {
	// setup
	username := "test1"

	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	connection, err := db.DB()
	defer connection.Close()

	user := model2.User{Username: username, Auth0UserID: "test_auth0_user_id_1"}

	sut := repository.NewUserRepository(db)

	// when
	actual := sut.CreateIfNotExists(user)

	// then
	assert.Equal(t, nil, actual)

	// cleanup
	cleanupTables(db)
}
