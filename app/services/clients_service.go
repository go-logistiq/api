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

	Groups *GroupsService
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

func (gs *ClientsService) GetBySlug(groupSlug, clientSlug string) (models.Client, error) {
	group, err := gs.Groups.GetBySlug(groupSlug)
	if err != nil {
		return models.Client{}, errs.NewErrorNotFound("Group not found")
	}

	rows, err := gs.DB.Conn().(*pgxpool.Pool).
		Query(context.Background(), sql.GetClientBySlug, group.ID, clientSlug)

	if err != nil {
		gs.Log.Error("Error getting client by name", "error", err)
		return models.Client{}, errs.NewErrorInternal(err.Error())
	}
	defer rows.Close()

	client, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Client])
	if err != nil {
		if err == pgx.ErrNoRows {
			return models.Client{}, errs.NewErrorNotFound("Client not found")
		}
		gs.Log.Error("Error collecting client", "error", err)
		return models.Client{}, errs.NewErrorInternal(err.Error())
	}

	return client, nil
}
