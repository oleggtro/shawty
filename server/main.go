package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/cloudybyte/shawty/server/db"
	"github.com/cloudybyte/shawty/server/routes"
	"github.com/cloudybyte/shawty/server/util"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4"
)

const (
	API_VERSION = "v1"
	DB_URL      = "postgres://postgres:1234@localhost:5432/shawty_dev?sslmode=disable"
)

func main() {
	//TODO: make db url configurable
	conn, err := pgx.Connect(context.Background(), DB_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	m, err := migrate.New(
		"file://migrations",
		DB_URL,
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && (err.Error() == "no change") {
		log.Println("Migrations ran; no change")
	} else {
		log.Fatal(err)
	}

	state := util.State{
		Db: conn,
	}

	r := gin.Default()

	//r.Use(util.AuthMiddleware())
	r.Use(gin.Recovery())
	r.Use(util.Site(state))

	r.POST("/signup", routes.CreateUser)

	r.POST("/session", routes.CreateSession)

	secured := r.Group("/sec", db.AuthMiddleware())

	secured.GET("/hello", hello)

	secured.POST("/red", routes.CreateShortlink)

	r.GET("/redirect/:id", routes.UseShortlink)

	r.Run("localhost:8080")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func hello(c *gin.Context) {
	c.String(200, "hello")
}