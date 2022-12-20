-- name: GetUser :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = ? LIMIT 1;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY name ASC;

-- name: CreateUser :one
INSERT INTO users (
  name, email, password
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateUser :one
UPDATE users SET name = ?, email = ?, password = ?, avatar = ?, location = ?, title = ?, status = ?, tags = ?, meta = ?, last_login = ?, created_at = ?, updated_at = ? 
WHERE id = ?
RETURNING *;


-- name: DeleteUser :execresult
DELETE FROM users
WHERE id = ?;




