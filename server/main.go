package main

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudybyte/shawty/server/routes"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

const API_VERSION = "v1"

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "shawty_dev"
)

func main() {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:1234@localhost:5432/shawty_dev")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	router := gin.Default()
	router.GET("/redirect", routes.CreateShortlink)

	router.Run("localhost:8080")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}