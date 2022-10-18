// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: user.query.sql

package repositories

import (
	"context"
	"time"
)

const countUserEmail = `-- name: CountUserEmail :one
SELECT COUNT(*) FROM Users WHERE email = $1
`

func (q *Queries) CountUserEmail(ctx context.Context, email string) (int64, error) {
	row := q.db.QueryRowContext(ctx, countUserEmail, email)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countUserUsername = `-- name: CountUserUsername :one
SELECT COUNT(*) FROM Users WHERE username = $1
`

func (q *Queries) CountUserUsername(ctx context.Context, username string) (int64, error) {
	row := q.db.QueryRowContext(ctx, countUserUsername, username)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM Users WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUserEmail = `-- name: GetUserEmail :one
SELECT id, username, email, password, age, created_at, updated_at FROM Users WHERE email = $1
`

func (q *Queries) GetUserEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.Age,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const insertUser = `-- name: InsertUser :one
INSERT INTO Users (username, email, password, age) VALUES ($1, $2, $3, $4) 
RETURNING id, username, email, age
`

type InsertUserParams struct {
	Username string
	Email    string
	Password string
	Age      int32
}

type InsertUserRow struct {
	ID       int32
	Username string
	Email    string
	Age      int32
}

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) (InsertUserRow, error) {
	row := q.db.QueryRowContext(ctx, insertUser,
		arg.Username,
		arg.Email,
		arg.Password,
		arg.Age,
	)
	var i InsertUserRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Age,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE Users SET email = $1, username = $2, updated_at = NOW() WHERE id = $3 
RETURNING id, username, email, age, updated_at
`

type UpdateUserParams struct {
	Email    string
	Username string
	ID       int32
}

type UpdateUserRow struct {
	ID        int32
	Username  string
	Email     string
	Age       int32
	UpdatedAt time.Time
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (UpdateUserRow, error) {
	row := q.db.QueryRowContext(ctx, updateUser, arg.Email, arg.Username, arg.ID)
	var i UpdateUserRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Age,
		&i.UpdatedAt,
	)
	return i, err
}