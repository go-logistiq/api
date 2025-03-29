package models

type Clients []Client

type Client struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	GroupID int    `json:"groupId"`
}
