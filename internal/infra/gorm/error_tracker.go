package gorm

import (
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
	"strings"
)

func IsDuplicateError(err error) (is bool, columnName string) {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		columnName = extractColumnName(pgErr.ConstraintName)
		is = true
		return
	}
	return
}

func extractColumnName(constraintName string) string {
	parts := strings.Split(constraintName, "_")
	if len(parts) > 1 {
		return parts[len(parts)-2]
	}
	return "unknown"
}
