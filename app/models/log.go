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

var LogDBColumns = []string{
	"logs.id",
	"logs.client_id",
	"logs.level",
	"logs.logged_at",
	"logs.message",
	"logs.attributes",
}

func (l *Log) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"client_id":  l.ClientID,
		"level":      l.Level,
		"logged_at":  l.LoggedAt,
		"message":    l.Message,
		"attributes": l.Attributes,
	}
}

func (l Log) ToSlice() []interface{} {
	return []interface{}{
		l.ClientID,
		l.Level,
		l.LoggedAt,
		l.Message,
		l.Attributes,
	}
}
