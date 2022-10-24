-- name: InsertUser :one
INSERT INTO Users (username, email, password, age) VALUES ($1, $2, $3, $4) 
RETURNING id, username, email, age ;

-- name: GetUserByEmail :one 
SELECT * FROM Users WHERE email = $1;

-- name: GetUserById :one
SELECT * FROM Users WHERE id = $1;

-- name: UpdateUser :one
UPDATE Users SET email = $1, username = $2, updated_at = NOW()::date WHERE id = $3 
RETURNING id, username, email, age, updated_at;

-- name: DeleteUser :exec 
DELETE FROM Users WHERE id = $1;

-- name: CountUserEmail :one
SELECT COUNT(*) FROM Users WHERE email = $1;

-- name: CountUserUsername :one
SELECT COUNT(*) FROM Users WHERE username = $1;

