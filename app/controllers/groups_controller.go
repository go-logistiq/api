package controllers

import (
	"github.com/go-logistiq/api/app/services"
	"github.com/go-raptor/raptor/v3"
)

type GroupsController struct {
	raptor.Controller

	Groups *services.GroupsService
}

func (gc *GroupsController) All(s raptor.State) error {
	groups, err := gc.Groups.All()
	if err != nil {
		return s.JSONError(err)
	}
	return s.JSONResponse(groups)
}

func (gc *GroupsController) GetBySlug(s raptor.State) error {
	slug := s.Param("slug")
	group, err := gc.Groups.GetBySlug(slug)
	if err != nil {
		return s.JSONError(err)
	}
	return s.JSONResponse(group)
}
