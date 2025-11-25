package main

import (
	"log"

	"github.com/bagusshndr/linknau-article-test/internal/config"
	"github.com/bagusshndr/linknau-article-test/internal/database"
	httpserver "github.com/bagusshndr/linknau-article-test/internal/http"
)

func main() {
	cfg := config.Load()

	db := database.NewPostgres(cfg)

	srv := httpserver.NewServer(cfg, db)

	if err := srv.Run(cfg.AppPort); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
