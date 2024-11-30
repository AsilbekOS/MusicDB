CREATE TABLE IF NOT EXISTS songs (
    id SERIAL PRIMARY KEY,
    group_name VARCHAR(255),
    song_name VARCHAR(255),
    release_date DATE,
    link VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS lyrics (
    id SERIAL PRIMARY KEY,
    song_id INTEGER REFERENCES songs(id),
    verse_number INTEGER,
    verse_text TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
