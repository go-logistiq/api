package services

import (
	"context"

	"github.com/go-logistiq/api/app/models"
	"github.com/go-logistiq/api/db/sql"
	"github.com/go-raptor/errs"
	"github.com/go-raptor/raptor/v3"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ClientsService struct {
	raptor.Service
}

func (gs *ClientsService) All() (models.Clients, error) {
	rows, err := gs.DB.Conn().(*pgxpool.Pool).
		Query(context.Background(), sql.AllClients)

	if err != nil {
		gs.Log.Error("Error getting clients", "error", err)
		return models.Clients{}, errs.NewErrorInternal(err.Error())
	}
	defer rows.Close()

	clients, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Client])
	if err != nil {
		gs.Log.Error("Error collecting clients", "error", err)
		return models.Clients{}, errs.NewErrorInternal(err.Error())
	}

	return clients, nil
}
