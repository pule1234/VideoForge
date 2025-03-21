// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Copywriting struct {
	ID        int64            `json:"id"`
	Source    string           `json:"source"`
	Title     string           `json:"title"`
	Content   string           `json:"content"`
	Date      pgtype.Timestamp `json:"date"`
	CreatedAt time.Time        `json:"created_at"`
	DeleteAt  pgtype.Timestamp `json:"delete_at"`
}
