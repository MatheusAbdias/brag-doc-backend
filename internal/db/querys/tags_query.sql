-- name: CreateTag :exec

INSERT INTO tags (name) VALUES ($1) RETURNING *;

-- name: GetTag :one

SELECT * FROM tags WHERE id = $1 LIMIT 1;

-- name: GetTags :many

SELECT * FROM tags ORDER BY id LIMIT $1 OFFSET $2;

-- name: UpdateTag :exec

UPDATE tags SET name = $1 WHERE id = $2;

-- name: DeleteTag :exec

DELETE FROM tags WHERE id = $1;