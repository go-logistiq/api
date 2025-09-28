package queries

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/go-logistiq/api/app/models"
)

func GetGroups() sq.SelectBuilder {
	return psql.
		Select(models.GroupDBColumns...).
		From("groups").
		OrderBy("name")
}

func GetGroupBySlug(slug string) sq.SelectBuilder {
	return psql.
		Select(models.GroupDBColumns...).
		From("groups").
		Where(sq.Eq{"slug": slug})
}
