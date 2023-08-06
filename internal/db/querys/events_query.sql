-- name: CreateEvent :exec

INSERT INTO
    events (name, description, date)
VALUES ($1, $2, $3) RETURNING id;

-- name: GetEvent :one

SELECT * FROM events WHERE id = $1 LIMIT 1;

-- name: GetEvents :many

SELECT * FROM events ORDER BY id LIMIT $1 OFFSET $2;

-- name: UpdateEvent :exec

UPDATE events
SET
    name = $1,
    description = $2,
    date = $3
WHERE id = $4;

-- name: DeleteEvent :exec

DELETE FROM events WHERE id = $1;