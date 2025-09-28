package models

type Groups []Group

type Group struct {
	ID   int    `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}

var GroupDBColumns = []string{
	"groups.id",
	"groups.slug",
	"groups.name",
}

func (g *Group) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"slug": g.Slug,
		"name": g.Name,
	}
}
