CREATE TABLE reading_time
(
    session_id INT PRIMARY KEY,  -- Ensures one-to-one relationship
    date       DATE NOT NULL,
    reading_time   INT NOT NULL, -- Duration in minutes
    FOREIGN KEY (session_id) REFERENCES read_sessions (id) ON DELETE CASCADE
);