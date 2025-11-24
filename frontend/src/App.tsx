import { useEffect, useState } from 'react'
import { PageContainer } from './components/layout/PageContainer'
import { ArticleForm } from './components/articles/ArticleForm'
import { ArticleList } from './components/articles/ArticleList'
import type { Article } from './types/article'
import { createArticle, deleteArticle, fetchArticles } from './api/articlesApi.ts'
import type { CreateArticleRequest } from './api/articlesApi.ts'

function App() {
  const [articles, setArticles] = useState<Article[]>([])
  const [loading, setLoading] = useState(false)
  const [saving, setSaving] = useState(false)
  const [error, setError] = useState<string | null>(null)

  const loadArticles = async () => {
    try {
      setLoading(true)
      setError(null)
      const res = await fetchArticles(1, 50)
      setArticles(res.data ?? [])
    } catch (err: unknown) {
      console.error(err)
      const msg = err instanceof Error ? err.message : 'Failed to load articles'
      setError(msg)
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    void loadArticles()
  }, [])

  const handleCreate = async (payload: CreateArticleRequest) => {
    try {
      setSaving(true)
      setError(null)
      await createArticle(payload)
      await loadArticles()
    } catch (err: unknown) {
      console.error(err)
      const msg = err instanceof Error ? err.message : 'Failed to create article'
      setError(msg)
    } finally {
      setSaving(false)
    }
  }

  const handleDelete = async (id: number) => {
    const confirmed = window.confirm('Yakin ingin menghapus artikel ini?')
    if (!confirmed) return

    try {
      setError(null)
      await deleteArticle(id)
      await loadArticles()
    } catch (err: unknown) {
      console.error(err)
      const msg = err instanceof Error ? err.message : 'Failed to delete article'
      setError(msg)
    }
  }

  return (
    <PageContainer>
      <header style={{ marginBottom: 24 }}>
        <h1 style={{ margin: 0, fontSize: 26 }}>Linknau Articles</h1>
        <p style={{ margin: '4px 0 0', color: '#6b7280', fontSize: 14 }}>
          Simple article feature – admin create/delete & user article listing.
        </p>
      </header>

      {error && (
        <div
          style={{
            backgroundColor: '#fee2e2',
            border: '1px solid #fecaca',
            padding: '8px 12px',
            borderRadius: 8,
            marginBottom: 16,
            color: '#b91c1c',
            fontSize: 14
          }}
        >
          Error: {error}
        </div>
      )}

      {/* Admin section */}
      <section
        style={{
          backgroundColor: '#ffffff',
          borderRadius: 12,
          padding: '16px 18px',
          boxShadow: '0 1px 4px rgba(0,0,0,0.08)',
          marginBottom: 24
        }}
      >
        <h2 style={{ fontSize: 18, margin: '0 0 12px' }}>Admin – Create Article</h2>
        <ArticleForm onCreate={handleCreate} loading={saving} />
      </section>

      {/* User view section */}
      <section>
        <div
          style={{
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'space-between',
            marginBottom: 12
          }}
        >
          <h2 style={{ fontSize: 18, margin: 0 }}>User View – Articles</h2>
          <button
            type="button"
            onClick={() => void loadArticles()}
            disabled={loading}
            style={{
              padding: '6px 14px',
              borderRadius: 6,
              border: 'none',
              backgroundColor: '#10b981',
              color: '#ffffff',
              fontSize: 14,
              cursor: loading ? 'default' : 'pointer'
            }}
          >
            {loading ? 'Refreshing...' : 'Refresh'}
          </button>
        </div>

        {loading && <p style={{ fontSize: 14 }}>Loading articles...</p>}

        {!loading && <ArticleList articles={articles} onDelete={handleDelete} />}
      </section>
    </PageContainer>
  )
}

export default App
