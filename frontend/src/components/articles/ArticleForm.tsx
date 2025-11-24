import type { FormEvent } from 'react'
import type { CSSProperties } from 'react'
import type { CreateArticleRequest } from '../../api/articlesApi.ts'

interface ArticleFormProps {
  onCreate: (payload: CreateArticleRequest) => Promise<void> | void
  loading?: boolean
}

export function ArticleForm({ onCreate, loading }: ArticleFormProps) {
  const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()

    const formData = new FormData(e.currentTarget)
    const title = String(formData.get('title') || '')
    const slug = String(formData.get('slug') || '')
    const content = String(formData.get('content') || '')
    const mainPhotoUrl = String(formData.get('mainPhotoUrl') || '')

    const payload: CreateArticleRequest = {
      title,
      slug: slug || undefined,
      content,
      photos: mainPhotoUrl
        ? [
            {
              url: mainPhotoUrl,
              caption: 'Main photo',
              order: 1
            }
          ]
        : []
    }

    await onCreate(payload)
    e.currentTarget.reset()
  }

  return (
    <form onSubmit={handleSubmit}>
      <div style={{ marginBottom: 10 }}>
        <label style={labelStyle}>
          Title <span style={{ color: 'red' }}>*</span>
        </label>
        <input
          name="title"
          type="text"
          required
          placeholder="Judul artikel"
          style={inputStyle}
        />
      </div>

      <div style={{ marginBottom: 10 }}>
        <label style={labelStyle}>
          Slug{' '}
          <span style={{ fontSize: 12, color: '#6b7280' }}>
            (optional – kalau kosong akan di-generate dari title)
          </span>
        </label>
        <input
          name="slug"
          type="text"
          placeholder="cara-meningkatkan-brand-awareness"
          style={inputStyle}
        />
      </div>

      <div style={{ marginBottom: 10 }}>
        <label style={labelStyle}>
          Content <span style={{ color: 'red' }}>*</span>
        </label>
        <textarea
          name="content"
          required
          rows={4}
          placeholder="Isi artikel..."
          style={{
            ...inputStyle,
            resize: 'vertical'
          }}
        />
      </div>

      <div style={{ marginBottom: 16 }}>
        <label style={labelStyle}>
          Main Photo URL{' '}
          <span style={{ fontSize: 12, color: '#6b7280' }}>(optional)</span>
        </label>
        <input
          name="mainPhotoUrl"
          type="url"
          placeholder="https://example.com/image.jpg"
          style={inputStyle}
        />
      </div>

      <button
        type="submit"
        disabled={loading}
        style={{
          padding: '8px 18px',
          borderRadius: 6,
          border: 'none',
          backgroundColor: '#2563eb',
          color: '#fff',
          cursor: loading ? 'default' : 'pointer',
          fontSize: 14
        }}
      >
        {loading ? 'Saving...' : 'Create Article'}
      </button>
    </form>
  )
}

const inputStyle: CSSProperties = {
  width: '100%',
  padding: '8px 10px',
  borderRadius: 6,
  border: '1px solid #d1d5db',
  fontSize: 14
}

const labelStyle: CSSProperties = {
  display: 'block',
  marginBottom: 4,
  fontSize: 14
}
