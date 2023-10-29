CREATE TABLE book_list_books (
    list_id INT NOT NULL,
    book_id INT NOT NULL,
    PRIMARY KEY (list_id, book_id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (list_id) REFERENCES book_lists(id) ON DELETE CASCADE,
    FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE
);
