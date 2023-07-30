package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/MatheusAbdias/brag-doc-backend/config"
	controllers "github.com/MatheusAbdias/brag-doc-backend/controllers/tags"
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
		log.Fatalf("could not connect to database: %v", err)
	}

	db = dbConn.New(conn)

	fmt.Println("Database connected successfully...")

	server = gin.Default()
}

func main() {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	tagController := controllers.TagController{Repo: db}

	routerV1 := server.Group("/v1")

	tagsRouter := routerV1.Group("tags")

	tagsRouter.POST("", tagController.CreateTag)
	tagsRouter.GET(":id", tagController.GetTag)
	tagsRouter.GET("", tagController.ListTags)
	tagsRouter.PATCH(":id", tagController.UpdateTag)
	tagsRouter.DELETE(":id", tagController.DeleteTag)

	log.Fatalf("could not run server %v", server.Run(":"+config.Port))
}
