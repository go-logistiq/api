package components

import (
	"github.com/go-logistiq/api/db"
	"github.com/go-raptor/connectors/pgx"
	"github.com/go-raptor/raptor/v4"
)

func New(c *raptor.Config) *raptor.Components {
	return &raptor.Components{
		DatabaseConnector: pgx.NewPgxConnector(c.DatabaseConfig, db.Migrations()),
		Controllers:       Controllers(),
		Services:          Services(c),
		Middlewares:       Middlewares(c),
	}
}
