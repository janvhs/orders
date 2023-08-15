-- SQL-Dialect: SQLite

-- +goose Up
CREATE TABLE users (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    guid TEXT NOT NULL UNIQUE,

    username TEXT,
    user_dn TEXT,

    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME DEFAULT NULL
);

-- +goose Down
DROP TABLE users;
