package queries

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/go-logistiq/api/app/models"
)

func GetClients() sq.SelectBuilder {
	return psql.
		Select(models.ClientDBColumns...).
		From("clients").
		OrderBy("name")
}

func GetClientBySlug(groupID int, slug string) sq.SelectBuilder {
	return psql.
		Select(models.ClientDBColumns...).
		From("clients").
		Where(sq.Eq{
			"group_id": groupID,
			"slug":     slug},
		)
}

func GetClientByID(id int) sq.SelectBuilder {
	return psql.
		Select(models.ClientDBColumns...).
		From("clients").
		Where(sq.Eq{"id": id})
}
