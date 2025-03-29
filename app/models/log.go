package models

import "time"

type Log struct {
	ID         int64                  `json:"id"`
	ClientID   int                    `json:"clientId"`
	LoggedAt   time.Time              `json:"loggedAt"`
	Level      int                    `json:"level"`
	Message    string                 `json:"message"`
	Attributes map[string]interface{} `json:"attributes"`
}
