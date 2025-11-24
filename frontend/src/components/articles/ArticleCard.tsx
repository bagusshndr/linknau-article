import type { Article } from '../../types/article'

interface ArticleCardProps {
  article: Article
  onDelete?: (id: number) => void
}

export function ArticleCard({ article, onDelete }: ArticleCardProps) {
  const mainPhoto =
    article.photos && article.photos.length > 0 ? article.photos[0].url : undefined
  const created = article.createdAt ? new Date(article.createdAt).toLocaleDateString() : ''
  const preview =
    article.content.length > 160 ? article.content.slice(0, 160) + '...' : article.content

  return (
    <div
      style={{
        backgroundColor: '#ffffff',
        borderRadius: 12,
        overflow: 'hidden',
        boxShadow: '0 1px 4px rgba(0,0,0,0.08)',
        display: 'flex',
        flexDirection: 'column'
      }}
    >
      {mainPhoto && (
        <div style={{ height: 160, overflow: 'hidden' }}>
          <img
            src={mainPhoto}
            alt={article.title}
            style={{ width: '100%', height: '100%', objectFit: 'cover', display: 'block' }}
          />
        </div>
      )}

      <div style={{ padding: '12px 14px', flex: 1 }}>
        <div style={{ fontSize: 12, color: '#6b7280', marginBottom: 4 }}>{created}</div>
        <h3
          style={{
            fontSize: 16,
            margin: '0 0 6px',
            lineHeight: 1.3
          }}
        >
          {article.title}
        </h3>
        <p
          style={{
            fontSize: 14,
            color: '#4b5563',
            margin: 0
          }}
        >
          {preview}
        </p>
      </div>

      {onDelete && (
        <div
          style={{
            padding: '8px 10px',
            borderTop: '1px solid #e5e7eb',
            display: 'flex',
            justifyContent: 'flex-end'
          }}
        >
          <button
            type="button"
            onClick={() => onDelete(article.id)}
            style={{
              padding: '4px 10px',
              fontSize: 13,
              borderRadius: 6,
              border: 'none',
              backgroundColor: '#ef4444',
              color: '#ffffff',
              cursor: 'pointer'
            }}
          >
            Delete
          </button>
        </div>
      )}
    </div>
  )
}
