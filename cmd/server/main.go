package main

import (
	"fmt"
	"log"

	"github.com/bagusshndr/linknau-article-test/internal/article"
	"github.com/bagusshndr/linknau-article-test/internal/config"
	"github.com/bagusshndr/linknau-article-test/internal/db"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	dbConn := db.NewGormDB(cfg)

	r := gin.Default()

	// health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	articleRepo := article.NewRepository(dbConn)
	articleSvc := article.NewService(articleRepo)
	articleHandler := article.NewHandler(articleSvc)

	api := r.Group("/api/v1")
	articleHandler.RegisterRoutes(api)

	addr := fmt.Sprintf(":%s", cfg.AppPort)
	log.Printf("Server running on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
