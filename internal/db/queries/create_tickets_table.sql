CREATE TABLE IF NOT EXISTS tickets (
    id TEXT PRIMARY KEY,
    site_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    content TEXT NOT NULL
);
