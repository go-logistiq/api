package models

import "time"

type LogRecords []LogRecord
type LogRecord struct {
	Level      int            `json:"level"`
	LoggedAt   time.Time      `json:"loggedAt"`
	Message    string         `json:"message"`
	Attributes map[string]any `json:"attributes"`
}

type Logs []Log
type Log struct {
	ID       int64 `json:"id"`
	ClientID int   `json:"clientId"`

	LogRecord
}
