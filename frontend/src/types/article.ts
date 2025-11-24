export interface Photo {
  id?: number
  articleId?: number
  url: string
  caption?: string
  order?: number
  createdAt?: string
  updatedAt?: string
}

export interface Article {
  id: number
  title: string
  slug: string
  content: string
  publishedAt?: string | null
  createdAt?: string
  updatedAt?: string
  photos?: Photo[]
}

export interface ArticleListResponse {
  data: Article[]
  page: number
  size: number
  total: number
}
