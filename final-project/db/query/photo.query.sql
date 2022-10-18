-- name: InsertPhoto :one 
INSERT INTO Photos (title, caption, photo_url, user_id) VALUES ($1, $2, $3, $4)
RETURNING id, title, caption, photo_url, user_id, created_at;

-- name: GetUserPhoto :many 
SELECT 
    photos.*, 
    users.email, 
    users.username 
FROM Photos as photos 
JOIN Users as users ON photos.user_id = users.id;

-- name: UpdatePhoto :one
UPDATE Photos SET title = $1, caption = $2, photo_url = $3, updated_at = NOW() 
RETURNING id, title, caption, photo_url, user_id, updated_at;

-- name: DeletePhoto :exec 
DELETE FROM Photos WHERE id = $1;