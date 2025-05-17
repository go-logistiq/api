package services

import (
	"context"

	"github.com/go-logistiq/api/app/models"
	"github.com/go-logistiq/api/db/sql"
	"github.com/go-raptor/raptor/v4"
	"github.com/go-raptor/raptor/v4/errs"
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

func (gs *GroupsService) GetBySlug(slug string) (models.Group, error) {
	rows, err := gs.DB.Conn().(*pgxpool.Pool).
		Query(context.Background(), sql.GetGroupBySlug, slug)

	if err != nil {
		gs.Log.Error("Error getting group by name", "error", err)
		return models.Group{}, errs.NewErrorInternal(err.Error())
	}
	defer rows.Close()

	group, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Group])
	if err != nil {
		if err == pgx.ErrNoRows {
			return models.Group{}, errs.NewErrorNotFound("Group not found")
		}
		gs.Log.Error("Error collecting group", "error", err)
		return models.Group{}, errs.NewErrorInternal(err.Error())
	}

	return group, nil
}
