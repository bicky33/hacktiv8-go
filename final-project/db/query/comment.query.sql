-- name: InsertComment :one
INSERT INTO Comments (message, user_id, photo_id) VALUES($1, $2, $3) 
RETURNING id, message, photo_id, user_id, created_at;

-- name: GetComment :many
SELECT 
    comments.*,  
    users.id, 
    users.email, 
    users.username, 
    photos.id, 
    photos.title, 
    photos.caption, 
    photos.photo_url, 
    photos.user_id 
FROM Comments as comments 
JOIN Users as users ON users.user_id = comments.user_id 
JOIN Photos as photos ON photos.id = comments.photo_id;

-- name: UpdateComment :one
UPDATE Comments SET message = $1, updated_at = NOW() 
WHERE id = $2 
RETURNING id, message, photo_id, user_id, updated_at; 

-- name: DeleteComment :exec
DELETE FROM Comments WHERE id = $1;



