package db

import (
	"fmt"
	"log"

	"github.com/bagusshndr/linknau-article-test/internal/article"
	"github.com/bagusshndr/linknau-article-test/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormDB(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBSSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err := db.AutoMigrate(&article.Article{}, &article.Photo{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	return db
}
