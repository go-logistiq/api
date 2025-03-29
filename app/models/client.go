package models

type Clients []Client

type Client struct {
	ID      int    `json:"id"`
	Slug    string `json:"slug"`
	Name    string `json:"name"`
	GroupID int    `json:"groupId"`
}
