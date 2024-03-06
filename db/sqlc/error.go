package db

import (
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

const (
	ForeignKeyViolation = "23503"
	UniqueKeyViolation  = "23505"
)

var ErrRecordNotFound = pgx.ErrNoRows

var ErrorUniqueViolation = &pgconn.PgError{Code: UniqueKeyViolation}
var ErrorForeignViolation = &pgconn.PgError{Code: ForeignKeyViolation}

func ErrorCode(err error) string {
	var pgErr *pgconn.PgError

	if errors.As(err, &pgErr) {
		return pgErr.Message
	}
	return ""
}
