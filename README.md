# Linknau Article Feature — Installation Guide (Ringkas)

## 🚀 Backend (Golang)

### 1. Install Go
Download & install Go:
https://go.dev/dl/

Cek versi:
```bash
go version
```

### 2. Setup PostgreSQL (Opsional)
Buat database:
```sql
CREATE DATABASE linknau_articles;
```

Buat tabel:
```sql
CREATE TABLE IF NOT EXISTS articles (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    slug VARCHAR(255) UNIQUE NOT NULL,
    content TEXT NOT NULL,
    published_at TIMESTAMP NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS photos (
    id SERIAL PRIMARY KEY,
    article_id INTEGER NOT NULL REFERENCES articles(id) ON DELETE CASCADE ON UPDATE CASCADE,
    url VARCHAR(512) NOT NULL,
    caption VARCHAR(255),
    "order" INTEGER DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_photos_article_id ON photos(article_id);
```

### 3. Setup Environment
Buat file `.env`:
```
APP_PORT=8080

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=linknau_articles
DB_SSLMODE=disable
```

### 4. Run Backend
```bash
go mod tidy
go run ./cmd/server
```

Backend berjalan di:
```
http://localhost:8080
```

---

## 💻 Frontend (React + Vite)

### 1. Install Node.js
Download:
https://nodejs.org/en/download

Cek versi:
```bash
node -v
npm -v
```

### 2. Install Dependency
Masuk folder frontend:
```bash
cd frontend
npm install
```

### 3. Jalankan Frontend
```bash
npm run dev
```

Frontend berjalan di:
```
http://localhost:5173
```

Proxy otomatis menghubungkan ke backend melalui `/api`.

---

## ✔ Testing
1. Create Article dari frontend.
2. Lihat daftar artikel pada section "User View".
3. Delete artikel dari tombol **Delete**.
4. Semua data disimpan via API Golang ke PostgreSQL.

---

## 🎉 Selesai!
