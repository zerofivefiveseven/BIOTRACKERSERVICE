// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package repo

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Plot struct {
	ID      int64
	UserID  int32
	Name    string
	Content []byte
}

type UsersCredential struct {
	ID       int64
	Email    string
	Password string
}

type UsersInfo struct {
	ID             pgtype.Int4
	FirstName      string
	LastName       string
	Dateofbirthday pgtype.Timestamp
}
