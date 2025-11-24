package article

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, a *Article) error
	FindByID(ctx context.Context, id uint) (*Article, error)
	FindAll(ctx context.Context, offset, limit int) ([]Article, int64, error)
	Update(ctx context.Context, a *Article) error
	Delete(ctx context.Context, id uint) error
}

type gormRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &gormRepository{db: db}
}

func (r *gormRepository) Create(ctx context.Context, a *Article) error {
	return r.db.WithContext(ctx).Create(a).Error
}

func (r *gormRepository) FindByID(ctx context.Context, id uint) (*Article, error) {
	var a Article
	if err := r.db.WithContext(ctx).
		Preload("Photos", func(tx *gorm.DB) *gorm.DB {
			return tx.Order("photos.order ASC")
		}).
		First(&a, id).Error; err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *gormRepository) FindAll(ctx context.Context, offset, limit int) ([]Article, int64, error) {
	var (
		articles []Article
		total    int64
	)

	q := r.db.WithContext(ctx).Model(&Article{})
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := q.
		Preload("Photos", func(tx *gorm.DB) *gorm.DB {
			return tx.Order("photos.order ASC")
		}).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&articles).Error; err != nil {
		return nil, 0, err
	}

	return articles, total, nil
}

func (r *gormRepository) Update(ctx context.Context, a *Article) error {
	// Strategy: delete existing photos, then recreate (simple & clear)
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("article_id = ?", a.ID).Delete(&Photo{}).Error; err != nil {
			return err
		}
		if err := tx.Model(&Article{}).Where("id = ?", a.ID).Updates(map[string]interface{}{
			"title":        a.Title,
			"slug":         a.Slug,
			"content":      a.Content,
			"published_at": a.PublishedAt,
		}).Error; err != nil {
			return err
		}
		if len(a.Photos) > 0 {
			if err := tx.Model(a).Association("Photos").Replace(a.Photos); err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *gormRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&Article{}, id).Error
}
