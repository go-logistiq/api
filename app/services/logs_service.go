package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-logistiq/api/app/models"
	"github.com/go-raptor/raptor/v4"
	"github.com/jackc/pgx/v5"
	"github.com/nats-io/nats.go"
)

type LogsService struct {
	raptor.Service

	Clients *ClientsService
	DB      *DatabaseService
}

func (ls *LogsService) ParseNATSMessage(msg *nats.Msg) (models.Logs, error) {
	clientID, err := ls.Clients.GetIDBySubject(msg.Subject)
	if err != nil {
		return nil, err
	}

	var logRecords models.LogRecords
	if err := json.Unmarshal(msg.Data, &logRecords); err != nil {
		return nil, errors.New("failed to unmarshal log records")
	}

	logs := make(models.Logs, len(logRecords))
	for i, record := range logRecords {
		logs[i] = models.Log{
			ID:        0,
			ClientID:  clientID,
			LogRecord: record,
		}
	}

	return logs, nil
}

func (ls *LogsService) Save(logs models.Logs) error {
	if len(logs) == 0 {
		return nil
	}

	rows := make([][]interface{}, len(logs))
	for i, log := range logs {
		rows[i] = log.ToSlice()
	}

	_, err := ls.DB.Conn().
		CopyFrom(
			context.Background(),
			pgx.Identifier{"logs"},
			models.LogDBColumns[1:],
			pgx.CopyFromRows(rows),
		)
	if err != nil {
		ls.Log.Error("Failed to copy logs to database", "error", err)
		return fmt.Errorf("copy logs: %w", err)
	}

	ls.Log.Info("Successfully saved logs", "client", logs[0].ClientID, "count", len(logs))
	return nil
}
