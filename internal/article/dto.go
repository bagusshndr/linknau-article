package article

import "time"

type PhotoRequest struct {
	URL     string `json:"url"`
	Caption string `json:"caption"`
	Order   int    `json:"order"`
}

type ArticleRequest struct {
	Title       string         `json:"title"`
	Slug        string         `json:"slug"`
	Content     string         `json:"content"`
	PublishedAt *time.Time     `json:"publishedAt"`
	Photos      []PhotoRequest `json:"photos"`
}

type PhotoResponse struct {
	ID        uint      `json:"id"`
	URL       string    `json:"url"`
	Caption   string    `json:"caption"`
	Order     int       `json:"order"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ArticleResponse struct {
	ID          uint            `json:"id"`
	Title       string          `json:"title"`
	Slug        string          `json:"slug"`
	Content     string          `json:"content"`
	PublishedAt *time.Time      `json:"publishedAt"`
	CreatedAt   time.Time       `json:"createdAt"`
	UpdatedAt   time.Time       `json:"updatedAt"`
	Photos      []PhotoResponse `json:"photos"`
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
		Photos:      make([]PhotoResponse, 0, len(a.Photos)),
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
