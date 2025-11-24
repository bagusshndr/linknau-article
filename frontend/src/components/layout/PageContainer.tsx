import type { PropsWithChildren } from 'react'

export function PageContainer({ children }: PropsWithChildren) {
  return (
    <div
      style={{
        maxWidth: '1120px',
        margin: '0 auto',
        padding: '24px 16px 40px'
      }}
    >
      {children}
    </div>
  )
}
