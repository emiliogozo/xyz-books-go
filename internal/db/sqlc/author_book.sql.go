// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: author_book.sql

package db

import (
	"context"
)

const createAuthorBookRel = `-- name: CreateAuthorBookRel :exec
INSERT INTO author_book (
  author_id,
  book_id
) VALUES (
  ?1, ?2
) RETURNING author_id, book_id
`

type CreateAuthorBookRelParams struct {
	AuthorID int64 `json:"author_id"`
	BookID   int64 `json:"book_id"`
}

func (q *Queries) CreateAuthorBookRel(ctx context.Context, arg CreateAuthorBookRelParams) error {
	_, err := q.db.ExecContext(ctx, createAuthorBookRel, arg.AuthorID, arg.BookID)
	return err
}

const listAuthorsWithBookID = `-- name: ListAuthorsWithBookID :many
SELECT a.author_id, a.first_name, a.last_name, a.middle_name
FROM
  author_book ab
  JOIN authors a ON ab.author_id = a.author_id
WHERE ab.book_id = ?1
`

type ListAuthorsWithBookIDRow struct {
	Author Author `json:"author"`
}

func (q *Queries) ListAuthorsWithBookID(ctx context.Context, bookID int64) ([]ListAuthorsWithBookIDRow, error) {
	rows, err := q.db.QueryContext(ctx, listAuthorsWithBookID, bookID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListAuthorsWithBookIDRow{}
	for rows.Next() {
		var i ListAuthorsWithBookIDRow
		if err := rows.Scan(
			&i.Author.AuthorID,
			&i.Author.FirstName,
			&i.Author.LastName,
			&i.Author.MiddleName,
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