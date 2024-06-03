// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: author.sql

package db

import (
	"context"
	"database/sql"
)

const countAuthors = `-- name: CountAuthors :one
SELECT count(*) FROM authors
`

func (q *Queries) CountAuthors(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countAuthors)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createAuthor = `-- name: CreateAuthor :one
INSERT INTO authors (
  first_name,
  last_name,
  middle_name
) VALUES (
  ?1, ?2, ?3
) RETURNING author_id, first_name, last_name, middle_name
`

type CreateAuthorParams struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name"`
}

func (q *Queries) CreateAuthor(ctx context.Context, arg CreateAuthorParams) (Author, error) {
	row := q.db.QueryRowContext(ctx, createAuthor, arg.FirstName, arg.LastName, arg.MiddleName)
	var i Author
	err := row.Scan(
		&i.AuthorID,
		&i.FirstName,
		&i.LastName,
		&i.MiddleName,
	)
	return i, err
}

const deleteAuthor = `-- name: DeleteAuthor :exec
DELETE FROM authors WHERE author_id = ?1
`

func (q *Queries) DeleteAuthor(ctx context.Context, authorID int64) error {
	_, err := q.db.ExecContext(ctx, deleteAuthor, authorID)
	return err
}

const getAuthor = `-- name: GetAuthor :one
SELECT author_id, first_name, last_name, middle_name FROM authors
WHERE author_id = ?1 LIMIT 1
`

func (q *Queries) GetAuthor(ctx context.Context, authorID int64) (Author, error) {
	row := q.db.QueryRowContext(ctx, getAuthor, authorID)
	var i Author
	err := row.Scan(
		&i.AuthorID,
		&i.FirstName,
		&i.LastName,
		&i.MiddleName,
	)
	return i, err
}

const getAuthorByName = `-- name: GetAuthorByName :one
SELECT author_id, first_name, last_name, middle_name FROM authors
WHERE
  first_name = ?1 AND
  last_name = ?2 AND
  middle_name = COALESCE(?3, middle_name)
LIMIT 1
`

type GetAuthorByNameParams struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name"`
}

func (q *Queries) GetAuthorByName(ctx context.Context, arg GetAuthorByNameParams) (Author, error) {
	row := q.db.QueryRowContext(ctx, getAuthorByName, arg.FirstName, arg.LastName, arg.MiddleName)
	var i Author
	err := row.Scan(
		&i.AuthorID,
		&i.FirstName,
		&i.LastName,
		&i.MiddleName,
	)
	return i, err
}

const listAuthors = `-- name: ListAuthors :many
SELECT author_id, first_name, last_name, middle_name FROM authors
ORDER BY author_id
LIMIT ?1
OFFSET ?2
`

type ListAuthorsParams struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

func (q *Queries) ListAuthors(ctx context.Context, arg ListAuthorsParams) ([]Author, error) {
	rows, err := q.db.QueryContext(ctx, listAuthors, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Author{}
	for rows.Next() {
		var i Author
		if err := rows.Scan(
			&i.AuthorID,
			&i.FirstName,
			&i.LastName,
			&i.MiddleName,
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

const updateAuthor = `-- name: UpdateAuthor :one
UPDATE authors
SET
  first_name = COALESCE(?1, first_name),
  last_name = COALESCE(?2, last_name),
  middle_name = COALESCE(?3, middle_name)
WHERE
  author_id = ?4
RETURNING author_id, first_name, last_name, middle_name
`

type UpdateAuthorParams struct {
	FirstName  sql.NullString `json:"first_name"`
	LastName   sql.NullString `json:"last_name"`
	MiddleName sql.NullString `json:"middle_name"`
	AuthorID   int64          `json:"author_id"`
}

func (q *Queries) UpdateAuthor(ctx context.Context, arg UpdateAuthorParams) (Author, error) {
	row := q.db.QueryRowContext(ctx, updateAuthor,
		arg.FirstName,
		arg.LastName,
		arg.MiddleName,
		arg.AuthorID,
	)
	var i Author
	err := row.Scan(
		&i.AuthorID,
		&i.FirstName,
		&i.LastName,
		&i.MiddleName,
	)
	return i, err
}
