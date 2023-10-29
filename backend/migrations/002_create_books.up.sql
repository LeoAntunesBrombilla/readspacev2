CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    subtitle VARCHAR(255),
    authors TEXT[],
    publisher VARCHAR(255),
    description TEXT,
    page_count INTEGER,
    categories TEXT,
    language VARCHAR(10),
    small_thumbnail VARCHAR(255),
    thumbnail VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
