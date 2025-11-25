package article

import (
	"context"
	"errors"
	"regexp"
	"strings"
	"time"
)

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) Create(ctx context.Context, req *ArticleRequest) (*Article, error) {
	if strings.TrimSpace(req.Title) == "" {
		return nil, errors.New("title is required")
	}
	if strings.TrimSpace(req.Content) == "" {
		return nil, errors.New("content is required")
	}

	slug := req.Slug
	if slug == "" {
		slug = generateSlug(req.Title)
	}

	now := time.Now()

	art := &Article{
		Title:       req.Title,
		Slug:        slug,
		Content:     req.Content,
		PublishedAt: req.PublishedAt,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	for _, p := range req.Photos {
		if strings.TrimSpace(p.URL) == "" {
			continue
		}
		art.Photos = append(art.Photos, Photo{
			URL:       p.URL,
			Caption:   p.Caption,
			Order:     p.Order,
			CreatedAt: now,
			UpdatedAt: now,
		})
	}

	if err := s.repo.Create(ctx, art); err != nil {
		return nil, err
	}
	return art, nil
}

func (s *service) GetByID(ctx context.Context, id uint) (*Article, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *service) List(ctx context.Context, page, pageSize int) ([]Article, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
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

	if strings.TrimSpace(req.Title) != "" {
		existing.Title = req.Title
	}
	if req.Slug != "" {
		existing.Slug = req.Slug
	} else if req.Title != "" {
		existing.Slug = generateSlug(req.Title)
	}
	if strings.TrimSpace(req.Content) != "" {
		existing.Content = req.Content
	}
	existing.PublishedAt = req.PublishedAt
	existing.UpdatedAt = time.Now()

	// replace photos
	existing.Photos = make([]Photo, 0, len(req.Photos))
	for _, p := range req.Photos {
		if strings.TrimSpace(p.URL) == "" {
			continue
		}
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

var nonSlugChar = regexp.MustCompile(`[^a-z0-9\-]+`)

func generateSlug(title string) string {
	s := strings.ToLower(strings.TrimSpace(title))
	s = strings.ReplaceAll(s, " ", "-")
	s = nonSlugChar.ReplaceAllString(s, "-")
	s = strings.Trim(s, "-")
	if s == "" {
		return "article"
	}
	return s
}
