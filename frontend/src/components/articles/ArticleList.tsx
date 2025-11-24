import type { Article } from '../../types/article'
import { ArticleCard } from './ArticleCard'

interface ArticleListProps {
  articles: Article[]
  onDelete?: (id: number) => void
}

export function ArticleList({ articles, onDelete }: ArticleListProps) {
  if (articles.length === 0) {
    return <p style={{ color: '#6b7280' }}>Belum ada artikel.</p>
  }

  return (
    <div
      style={{
        display: 'grid',
        gridTemplateColumns: 'repeat(auto-fit, minmax(260px, 1fr))',
        gap: 16
      }}
    >
      {articles.map(article => (
        <ArticleCard key={article.id} article={article} onDelete={onDelete} />
      ))}
    </div>
  )
}
