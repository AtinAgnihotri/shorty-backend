-- +goose Up

CREATE TABLE links (
    id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    short_link_blob VARCHAR(7) NOT NULL UNIQUE,
    long_link TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down

DROP TABLE links;