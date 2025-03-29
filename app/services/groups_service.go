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

type GroupsService struct {
	raptor.Service
}

func (gs *GroupsService) All() (models.Groups, error) {
	rows, err := gs.DB.Conn().(*pgxpool.Pool).
		Query(context.Background(), sql.AllGroups)

	if err != nil {
		gs.Log.Error("Error getting groups", "error", err)
		return models.Groups{}, errs.NewErrorInternal(err.Error())
	}
	defer rows.Close()

	groups, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Group])
	if err != nil {
		gs.Log.Error("Error collecting groups", "error", err)
		return models.Groups{}, errs.NewErrorInternal(err.Error())
	}

	return groups, nil
}
