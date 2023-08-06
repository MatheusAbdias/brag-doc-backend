package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/MatheusAbdias/brag-doc-backend/config"

	eventControllers "github.com/MatheusAbdias/brag-doc-backend/internal/controllers/events"
	tagControllers "github.com/MatheusAbdias/brag-doc-backend/internal/controllers/tags"
	"github.com/MatheusAbdias/brag-doc-backend/internal/db/events"
	"github.com/MatheusAbdias/brag-doc-backend/internal/db/tags"

	"github.com/MatheusAbdias/brag-doc-backend/internal/routers"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var (
	server  *gin.Engine
	db      *sql.DB
	tagsDB  *tags.Queries
	eventDB *events.Queries
)

func init() {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	db, err := sql.Open(config.Driver, config.Source)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	tagsDB = tags.New(db)
	eventDB = events.New(db)

	fmt.Println("Database connected successfully...")

	server = gin.Default()
}

func main() {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	tagsController := tagControllers.Controller{Repo: tagsDB}
	eventController := eventControllers.Controller{Repo: eventDB}

	viRouter := server.Group("/v1")

	routers.RegisterTagRouters(viRouter, tagsController)
	routers.RegisterEventRouters(viRouter, eventController)
	log.Fatalf("could not run server %v", server.Run(":"+config.Port))
}
