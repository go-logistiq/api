package models

type Clients []Client

type Client struct {
	ID      int    `json:"id"`
	Slug    string `json:"slug"`
	Name    string `json:"name"`
	GroupID int    `json:"groupId"`
}

var ClientDBColumns = []string{
	"clients.id",
	"clients.slug",
	"clients.name",
	"clients.group_id",
}

func (c *Client) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"slug":     c.Slug,
		"name":     c.Name,
		"group_id": c.GroupID,
	}
}
