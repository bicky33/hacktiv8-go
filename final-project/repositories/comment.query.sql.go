// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: comment.query.sql

package repositories

import (
	"context"
	"time"
)

const deleteComment = `-- name: DeleteComment :exec
DELETE FROM Comments WHERE id = $1
`

func (q *Queries) DeleteComment(ctx context.Context, id uint32) error {
	_, err := q.db.ExecContext(ctx, deleteComment, id)
	return err
}

const getComment = `-- name: GetComment :many
SELECT 
    comments.id, comments.user_id, comments.photo_id, comments.message, comments.created_at, comments.updated_at,  
    users.id, 
    users.email, 
    users.username, 
    photos.id, 
    photos.title, 
    photos.caption, 
    photos.photo_url, 
    photos.user_id 
FROM Comments as comments 
JOIN Users as users ON users.id = comments.user_id 
JOIN Photos as photos ON photos.id = comments.photo_id
`

type GetCommentRow struct {
	ID        uint32
	UserID    uint32
	PhotoID   uint32
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
	ID_2      uint32
	Email     string
	Username  string
	ID_3      uint32
	Title     string
	Caption   string
	PhotoUrl  string
	UserID_2  uint32
}

func (q *Queries) GetComment(ctx context.Context) ([]GetCommentRow, error) {
	rows, err := q.db.QueryContext(ctx, getComment)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCommentRow
	for rows.Next() {
		var i GetCommentRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.PhotoID,
			&i.Message,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ID_2,
			&i.Email,
			&i.Username,
			&i.ID_3,
			&i.Title,
			&i.Caption,
			&i.PhotoUrl,
			&i.UserID_2,
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

const getCommentById = `-- name: GetCommentById :one
SELECT id, user_id, photo_id, message, created_at, updated_at FROM Comments WHERE id = $1
`

func (q *Queries) GetCommentById(ctx context.Context, id uint32) (Comment, error) {
	row := q.db.QueryRowContext(ctx, getCommentById, id)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.PhotoID,
		&i.Message,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const insertComment = `-- name: InsertComment :one
INSERT INTO Comments (message, user_id, photo_id) VALUES($1, $2, $3) 
RETURNING id, message, photo_id, user_id, created_at
`

type InsertCommentParams struct {
	Message string
	UserID  uint32
	PhotoID uint32
}

type InsertCommentRow struct {
	ID        uint32
	Message   string
	PhotoID   uint32
	UserID    uint32
	CreatedAt time.Time
}

func (q *Queries) InsertComment(ctx context.Context, arg InsertCommentParams) (InsertCommentRow, error) {
	row := q.db.QueryRowContext(ctx, insertComment, arg.Message, arg.UserID, arg.PhotoID)
	var i InsertCommentRow
	err := row.Scan(
		&i.ID,
		&i.Message,
		&i.PhotoID,
		&i.UserID,
		&i.CreatedAt,
	)
	return i, err
}

const updateComment = `-- name: UpdateComment :one
UPDATE Comments SET message = $1, updated_at = NOW() 
WHERE id = $2 
RETURNING id, message, photo_id, user_id, updated_at
`

type UpdateCommentParams struct {
	Message string
	ID      uint32
}

type UpdateCommentRow struct {
	ID        uint32
	Message   string
	PhotoID   uint32
	UserID    uint32
	UpdatedAt time.Time
}

func (q *Queries) UpdateComment(ctx context.Context, arg UpdateCommentParams) (UpdateCommentRow, error) {
	row := q.db.QueryRowContext(ctx, updateComment, arg.Message, arg.ID)
	var i UpdateCommentRow
	err := row.Scan(
		&i.ID,
		&i.Message,
		&i.PhotoID,
		&i.UserID,
		&i.UpdatedAt,
	)
	return i, err
}
