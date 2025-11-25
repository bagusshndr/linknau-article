package article

import "time"

type Photo struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ArticleID uint      `gorm:"index;not null" json:"articleId"`
	URL       string    `gorm:"size:512;not null" json:"url"`
	Caption   string    `gorm:"size:255" json:"caption"`
	Order     int       `gorm:"default:0" json:"order"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Article struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Title       string     `gorm:"size:255;not null" json:"title"`
	Slug        string     `gorm:"size:255;uniqueIndex;not null" json:"slug"`
	Content     string     `gorm:"type:text;not null" json:"content"`
	PublishedAt *time.Time `json:"publishedAt"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	Photos      []Photo    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"photos"`
}
