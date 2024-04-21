CREATE TABLE session_durations
(
    session_id INT  NOT NULL,
    date       DATE NOT NULL,
    duration   INT  NOT NULL, -- Duration in minutes
    FOREIGN KEY (session_id) REFERENCES read_sessions (id) ON DELETE CASCADE
);