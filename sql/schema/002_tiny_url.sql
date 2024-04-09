-- +goose Up

ALTER TABLE urls ADD CONSTRAINT tiny_url_unique UNIQUE (tiny_url);

-- +goose Down

ALTER TABLE urls
DROP CONSTRAINT IF EXISTS tiny_url_unique;

