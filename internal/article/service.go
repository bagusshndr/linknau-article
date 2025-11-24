package article

import (
	"context"
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Service interface {
	Create(ctx context.Context, req *ArticleRequest) (*Article, error)
	GetByID(ctx context.Context, id uint) (*Article, error)
	List(ctx context.Context, page, pageSize int) ([]Article, int64, error)
	Update(ctx context.Context, id uint, req *ArticleRequest) (*Article, error)
	Delete(ctx context.Context, id uint) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) Create(ctx context.Context, req *ArticleRequest) (*Article, error) {
	slug := req.Slug
	if slug == "" {
		slug = generateSlug(req.Title)
	}

	a := &Article{
		Title:       req.Title,
		Slug:        slug,
		Content:     req.Content,
		PublishedAt: req.PublishedAt,
	}

	for _, p := range req.Photos {
		a.Photos = append(a.Photos, Photo{
			URL:     p.URL,
			Caption: p.Caption,
			Order:   p.Order,
		})
	}

	if err := s.repo.Create(ctx, a); err != nil {
		return nil, err
	}

	return a, nil
}

func (s *service) GetByID(ctx context.Context, id uint) (*Article, error) {
	a, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}
	return a, nil
}

func (s *service) List(ctx context.Context, page, pageSize int) ([]Article, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize
	return s.repo.FindAll(ctx, offset, pageSize)
}

func (s *service) Update(ctx context.Context, id uint, req *ArticleRequest) (*Article, error) {
	existing, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	existing.Title = req.Title
	if req.Slug != "" {
		existing.Slug = req.Slug
	} else {
		existing.Slug = generateSlug(req.Title)
	}
	existing.Content = req.Content
	existing.PublishedAt = req.PublishedAt
	existing.UpdatedAt = time.Now()

	// rebuild photos
	existing.Photos = nil
	for _, p := range req.Photos {
		existing.Photos = append(existing.Photos, Photo{
			URL:     p.URL,
			Caption: p.Caption,
			Order:   p.Order,
		})
	}

	if err := s.repo.Update(ctx, existing); err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *service) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func generateSlug(title string) string {
	slug := strings.ToLower(title)
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.ReplaceAll(slug, "/", "-")
	slug = strings.ReplaceAll(slug, "\\", "-")
	return slug
}
