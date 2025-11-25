package article

import (
	"context"

	"gorm.io/gorm"
)

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
	var art Article
	if err := r.db.WithContext(ctx).
		Preload("Photos").
		First(&art, id).Error; err != nil {
		return nil, err
	}
	return &art, nil
}

func (r *gormRepository) FindAll(ctx context.Context, offset, limit int) ([]Article, int64, error) {
	var (
		articles []Article
		total    int64
	)

	tx := r.db.WithContext(ctx).Model(&Article{})
	if err := tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := tx.Preload("Photos").
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&articles).Error; err != nil {
		return nil, 0, err
	}

	return articles, total, nil
}

func (r *gormRepository) Update(ctx context.Context, a *Article) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("article_id = ?", a.ID).Delete(&Photo{}).Error; err != nil {
			return err
		}
		for i := range a.Photos {
			a.Photos[i].ArticleID = a.ID
		}
		if err := tx.Save(a).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *gormRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&Article{}, id).Error
}
