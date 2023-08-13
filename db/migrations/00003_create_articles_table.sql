-- SQL-Dialect: SQLite

-- +goose Up
CREATE TABLE articles (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    amount INTEGER NOT NULL,
    url TEXT NOT NULL,
    price REAL NOT NULL,
    cost_centre TEXT NOT NULL,
    state INTEGER NOT NULL DEFAULT 0,

    order_id INTEGER NOT NULL,
    author_id INTEGER NOT NULL,
    for_user_id INTEGER NOT NULL,

    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME DEFAULT NULL,

    FOREIGN KEY (order_id) REFERENCES orders(ID),
    FOREIGN KEY (author_id) REFERENCES users(ID),
    FOREIGN KEY (for_user_id) REFERENCES users(ID)
);

-- +goose Down
DROP TABLE articles;