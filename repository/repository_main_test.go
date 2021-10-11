package repository_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/wasanx25/go_nextjs_authentication/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbURL string

func TestMain(m *testing.M) {
	password := "root"
	database := "testdb"

	ctx := context.Background()
	request := testcontainers.ContainerRequest{
		Image:        "mysql:8",
		ExposedPorts: []string{"3306/tcp"},
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": password,
			"MYSQL_DATABASE":      database,
		},
		WaitingFor: wait.ForSQL("3306/tcp", "mysql", func(port nat.Port) string {
			return fmt.Sprintf("root:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", password, "127.0.0.1", port.Port(), database)
		}),
	}
	mysqlContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: request,
		Started:          true,
	})
	if err != nil {
		log.Fatalln(err)
	}
	defer mysqlContainer.Terminate(ctx)

	_, err = mysqlContainer.ContainerIP(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	port, err := mysqlContainer.MappedPort(ctx, "3306")
	if err != nil {
		log.Fatalln(err)
	}
	dbURL = fmt.Sprintf("root:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", password, "127.0.0.1", port.Port(), database)
	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	connection, err := db.DB()
	defer connection.Close()

	err = db.AutoMigrate(&model.User{}, &model.Post{}, &model.Follow{})
	if err != nil {
		log.Fatalln(err)
	}

	code := m.Run()
	os.Exit(code)
}

func cleanupTables(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS=0;")
	db.Exec("TRUNCATE TABLE posts;")
	db.Exec("TRUNCATE TABLE follows;")
	db.Exec("TRUNCATE TABLE users;")
	db.Exec("SET FOREIGN_KEY_CHECKS=1;")
}
