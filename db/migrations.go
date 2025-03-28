package db

import (
	"github.com/go-raptor/connector/pgx"
)

func Migrations() pgx.Migrations {
	return pgx.Migrations{}
}
