CREATE TABLE read_sessions
(
    id         SERIAL PRIMARY KEY,
    user_id    INT NOT NULL,
    book_id    INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (book_id) REFERENCES books (id) ON DELETE CASCADE
);

CREATE TYPE reading_duration AS
(
    date DATE,
    time INT
);

CREATE TABLE session_details
(
    session_id INT                NOT NULL,
    duration   reading_duration[] NOT NULL,
    FOREIGN KEY (session_id) REFERENCES read_sessions (id) ON DELETE CASCADE
);

