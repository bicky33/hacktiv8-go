// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: photo.query.sql

package repositories

import (
	"context"
	"time"
)

const deletePhoto = `-- name: DeletePhoto :exec
DELETE FROM Photos WHERE id = $1
`

func (q *Queries) DeletePhoto(ctx context.Context, id uint32) error {
	_, err := q.db.ExecContext(ctx, deletePhoto, id)
	return err
}

const getPhotoById = `-- name: GetPhotoById :one
SELECT id, title, caption, photo_url, user_id, created_at, updated_at FROM Photos WHERE id = $1
`

func (q *Queries) GetPhotoById(ctx context.Context, id uint32) (Photo, error) {
	row := q.db.QueryRowContext(ctx, getPhotoById, id)
	var i Photo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Caption,
		&i.PhotoUrl,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserPhoto = `-- name: GetUserPhoto :many
SELECT 
    photos.id, photos.title, photos.caption, photos.photo_url, photos.user_id, photos.created_at, photos.updated_at, 
    users.email, 
    users.username 
FROM Photos as photos 
JOIN Users as users ON photos.user_id = users.id
`

type GetUserPhotoRow struct {
	ID        uint32
	Title     string
	Caption   string
	PhotoUrl  string
	UserID    uint32
	CreatedAt time.Time
	UpdatedAt time.Time
	Email     string
	Username  string
}

func (q *Queries) GetUserPhoto(ctx context.Context) ([]GetUserPhotoRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserPhoto)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserPhotoRow
	for rows.Next() {
		var i GetUserPhotoRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Caption,
			&i.PhotoUrl,
			&i.UserID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Email,
			&i.Username,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertPhoto = `-- name: InsertPhoto :one
INSERT INTO Photos (title, caption, photo_url, user_id) VALUES ($1, $2, $3, $4)
RETURNING id, title, caption, photo_url, user_id, created_at
`

type InsertPhotoParams struct {
	Title    string
	Caption  string
	PhotoUrl string
	UserID   uint32
}

type InsertPhotoRow struct {
	ID        uint32
	Title     string
	Caption   string
	PhotoUrl  string
	UserID    uint32
	CreatedAt time.Time
}

func (q *Queries) InsertPhoto(ctx context.Context, arg InsertPhotoParams) (InsertPhotoRow, error) {
	row := q.db.QueryRowContext(ctx, insertPhoto,
		arg.Title,
		arg.Caption,
		arg.PhotoUrl,
		arg.UserID,
	)
	var i InsertPhotoRow
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Caption,
		&i.PhotoUrl,
		&i.UserID,
		&i.CreatedAt,
	)
	return i, err
}

const updatePhoto = `-- name: UpdatePhoto :one
UPDATE Photos SET title = $1, caption = $2, photo_url = $3, updated_at = NOW() WHERE id = $4
RETURNING id, title, caption, photo_url, user_id, updated_at
`

type UpdatePhotoParams struct {
	Title    string
	Caption  string
	PhotoUrl string
	ID       uint32
}

type UpdatePhotoRow struct {
	ID        uint32
	Title     string
	Caption   string
	PhotoUrl  string
	UserID    uint32
	UpdatedAt time.Time
}

func (q *Queries) UpdatePhoto(ctx context.Context, arg UpdatePhotoParams) (UpdatePhotoRow, error) {
	row := q.db.QueryRowContext(ctx, updatePhoto,
		arg.Title,
		arg.Caption,
		arg.PhotoUrl,
		arg.ID,
	)
	var i UpdatePhotoRow
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Caption,
		&i.PhotoUrl,
		&i.UserID,
		&i.UpdatedAt,
	)
	return i, err
}
