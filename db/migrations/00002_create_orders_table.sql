-- SQL-Dialect: SQLite

-- +goose Up
CREATE TABLE orders (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    
    state INTEGER NOT NULL DEFAULT 0,
    
    author_id INTEGER NOT NULL,

    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME,

    FOREIGN KEY (author_id) REFERENCES users(ID)
);

-- +goose Down
DROP TABLE orders;