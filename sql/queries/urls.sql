-- name: CreateURL :one

INSERT INTO urls(id, created_at, updated_at, original_url, tiny_url)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetUrlByTyniUrl :one

SELECT *
FROM urls
WHERE tiny_url = $1;