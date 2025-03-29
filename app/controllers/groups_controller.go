package controllers

import (
	"github.com/go-logistiq/api/app/services"
	"github.com/go-raptor/raptor/v3"
)

type GroupsController struct {
	raptor.Controller

	Groups *services.GroupsService
}

func (gc *GroupsController) All(c *raptor.Context) error {
	groups, err := gc.Groups.All()
	if err != nil {
		return c.JSONError(err)
	}
	return c.JSONResponse(groups)
}

func (gc *GroupsController) GetBySlug(c *raptor.Context) error {
	slug := c.Param("slug")
	group, err := gc.Groups.GetBySlug(slug)
	if err != nil {
		return c.JSONError(err)
	}
	return c.JSONResponse(group)
}
