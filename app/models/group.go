package models

type Groups []Group

type Group struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
