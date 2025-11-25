package http

import (
	"github.com/bagusshndr/linknau-article-test/internal/article"
	"github.com/bagusshndr/linknau-article-test/internal/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	Engine *gin.Engine
}

func NewServer(cfg *config.Config, db *gorm.DB) *Server {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	repo := article.NewRepository(db)
	svc := article.NewService(repo)
	handler := article.NewHTTPHandler(svc)

	api := r.Group("/api/v1")
	handler.RegisterRoutes(api)

	return &Server{Engine: r}
}

func (s *Server) Run(port string) error {
	return s.Engine.Run(":" + port)
}
