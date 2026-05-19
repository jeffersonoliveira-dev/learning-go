CREATE TABLE IF NOT EXISTS short_urls (
    id BIGSERIAL PRIMARY KEY,
    code TEXT UNIQUE,
    original_url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
