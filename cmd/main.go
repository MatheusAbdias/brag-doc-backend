package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/MatheusAbdias/brag-doc-backend/config"
	dbConn "github.com/MatheusAbdias/brag-doc-backend/internal/db/tags"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var (
	server *gin.Engine
	db     *dbConn.Queries
)

func init() {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	conn, err := sql.Open(config.PostgresDriver, config.PostgresSource)
	if err != nil {
		log.Fatalf("could not connect to postgres database: %v", err)
	}

	db = dbConn.New(conn)

	fmt.Println("PostgreSQL connected successfully...")

	server = gin.Default()
}

func main() {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	router := server.Group("/api")

	router.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Welcome to Golang with PostgreSQL"})
	})

	server.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": fmt.Sprintf("Route %s not found", ctx.Request.URL)})
	})

	log.Fatal(server.Run(":" + config.Port))
}
