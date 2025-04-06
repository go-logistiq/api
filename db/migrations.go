package db

import (
	"github.com/go-logistiq/api/db/migrations"
	"github.com/go-raptor/connectors/pgx"
)

func Migrations() pgx.Migrations {
	return pgx.Migrations{
		"20250328225258_create_groups":  migrations.CreateGroups{},
		"20250328231103_create_clients": migrations.CreateClients{},
		"20250328233211_create_logs":    migrations.CreateLogs{},
	}
}
