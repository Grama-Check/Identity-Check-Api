-- name: CreatePerson :one
INSERT INTO persons (
    id,
    name,
    address
) VALUES (
    $1,
    $2,
    $3)
RETURNING *;

-- name: GetPerson :one
SELECT * FROM persons
WHERE id = $1;


