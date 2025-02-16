CREATE TABLE flashcard_sets (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    slug TEXT NOT NULL,
    title TEXT NOT NULL,
    description TEXT NOT NULL DEFAULT "",
    edit_token TEXT NOT NULL
);

CREATE TABLE cards (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    term TEXT NOT NULL,
    definition TEXT NOT NULL,
    set_id INTEGER NOT NULL,
    FOREIGN KEY (set_id) REFERENCES flashcard_sets (id)
);
