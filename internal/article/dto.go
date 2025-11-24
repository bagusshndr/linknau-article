package article

import "time"

type PhotoRequest struct {
	URL     string `json:"url" binding:"required"`
	Caption string `json:"caption"`
	Order   int    `json:"order"`
}

type ArticleRequest struct {
	Title       string         `json:"title" binding:"required"`
	Slug        string         `json:"slug"`
	Content     string         `json:"content" binding:"required"`
	PublishedAt *time.Time     `json:"publishedAt"`
	Photos      []PhotoRequest `json:"photos"`
}

type ArticleResponse struct {
	ID          uint            `json:"id"`
	Title       string          `json:"title"`
	Slug        string          `json:"slug"`
	Content     string          `json:"content"`
	PublishedAt *time.Time      `json:"publishedAt,omitempty"`
	CreatedAt   time.Time       `json:"createdAt"`
	UpdatedAt   time.Time       `json:"updatedAt"`
	Photos      []PhotoResponse `json:"photos,omitempty"`
}

type PhotoResponse struct {
	ID        uint      `json:"id"`
	URL       string    `json:"url"`
	Caption   string    `json:"caption"`
	Order     int       `json:"order"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func ToArticleResponse(a *Article) *ArticleResponse {
	res := &ArticleResponse{
		ID:          a.ID,
		Title:       a.Title,
		Slug:        a.Slug,
		Content:     a.Content,
		PublishedAt: a.PublishedAt,
		CreatedAt:   a.CreatedAt,
		UpdatedAt:   a.UpdatedAt,
	}

	for _, p := range a.Photos {
		res.Photos = append(res.Photos, PhotoResponse{
			ID:        p.ID,
			URL:       p.URL,
			Caption:   p.Caption,
			Order:     p.Order,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		})
	}

	return res
}
