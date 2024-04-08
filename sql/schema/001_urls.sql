-- +goose Up


CREATE TABLE urls(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    original_url TEXT NOT NULL,
    tiny_url TEXT NOT NULL
);

-- +goose Down

DROP TABLE urls;