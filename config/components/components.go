package components

import (
	"github.com/go-logistiq/api/db"
	"github.com/go-raptor/connectors/pgx"
	"github.com/go-raptor/raptor/v3"
)

func New(utils *raptor.Utils) *raptor.Components {
	return &raptor.Components{
		DatabaseConnector: pgx.NewPgxConnector(utils.Config.DatabaseConfig, db.Migrations()),
		Controllers:       Controllers(),
		Services:          Services(utils),
		Middlewares:       Middlewares(utils),
	}
}
