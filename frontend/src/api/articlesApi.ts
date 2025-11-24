import { apiDelete, apiGet, apiPost } from './client'
import type { Article, ArticleListResponse, Photo } from '../types/article'

export interface CreateArticleRequest {
  title: string
  slug?: string
  content: string
  photos?: Pick<Photo, 'url' | 'caption' | 'order'>[]
  publishedAt?: string | null
}

export async function fetchArticles(
  page = 1,
  pageSize = 20
): Promise<ArticleListResponse> {
  return apiGet<ArticleListResponse>(`/api/v1/articles?page=${page}&pageSize=${pageSize}`)
}

export async function createArticle(payload: CreateArticleRequest): Promise<Article> {
  return apiPost<CreateArticleRequest, Article>('/api/v1/articles', payload)
}

export async function deleteArticle(id: number): Promise<void> {
  return apiDelete(`/api/v1/articles/${id}`)
}
