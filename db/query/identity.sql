-- name: CreatePerson :one
INSERT INTO persons (
    nic,
    name,
    address
) VALUES (
    $1,
    $2,
    $3)
RETURNING *;

-- name: GetPerson :one
SELECT * FROM persons
WHERE nic = $1;


