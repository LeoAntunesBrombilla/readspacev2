CREATE TABLE read_sessions
(
    id         SERIAL PRIMARY KEY,
    user_id    INT NOT NULL,
    book_id    INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (book_id) REFERENCES books (id) ON DELETE CASCADE
);
