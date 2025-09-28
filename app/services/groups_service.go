package services

import (
	"github.com/go-logistiq/api/app/models"
	"github.com/go-logistiq/api/app/queries"
	"github.com/go-raptor/raptor/v4"
)

type GroupsService struct {
	raptor.Service

	DB *DatabaseService
}

func (s *GroupsService) All() (models.Groups, error) {
	return Select[models.Group](s.DB, queries.GetGroups())
}

func (s *GroupsService) GetBySlug(slug string) (models.Group, error) {
	return SelectOne[models.Group](s.DB, queries.GetGroupBySlug(slug))
}
