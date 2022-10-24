-- name: InsertSocialMedia :one 
INSERT INTO SocialMedias (name, social_media_url, user_id) VALUES ($1, $2, $3) 
RETURNING id, name, social_media_url, user_id, created_at;

-- name: GetSocialMedia :many 
SELECT 
    socialMedia.*, 
    users.id, 
    users.username
FROM SocialMedias as socialMedia 
JOIN Users as users ON users.id = socialMedia.user_id;

-- name: GetSocialMediaById :one
SELECT * FROM SocialMedias WHERE id = $1;

-- name: UpdateSocialMedia :one 
UPDATE SocialMedias SET name = $1, social_media_url = $2, updated_at = NOW() 
WHERE id = $3 
RETURNING id, name, social_media_url, user_id, updated_at; 

-- name: DeleteSocialMedia :exec 
DELETE FROM SocialMedias WHERE id = $1;

