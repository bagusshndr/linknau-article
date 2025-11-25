package article

import "context"

type Repository interface {
	Create(ctx context.Context, a *Article) error
	FindByID(ctx context.Context, id uint) (*Article, error)
	FindAll(ctx context.Context, offset, limit int) ([]Article, int64, error)
	Update(ctx context.Context, a *Article) error
	Delete(ctx context.Context, id uint) error
}

type Service interface {
	Create(ctx context.Context, req *ArticleRequest) (*Article, error)
	GetByID(ctx context.Context, id uint) (*Article, error)
	List(ctx context.Context, page, pageSize int) ([]Article, int64, error)
	Update(ctx context.Context, id uint, req *ArticleRequest) (*Article, error)
	Delete(ctx context.Context, id uint) error
}
