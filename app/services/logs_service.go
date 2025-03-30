package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/go-logistiq/api/app/models"
	"github.com/go-raptor/raptor/v3"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nats-io/nats.go"
)

type LogsService struct {
	raptor.Service

	Clients *ClientsService
}

func (ls *LogsService) ParseNATSMessage(msg *nats.Msg) (*models.Logs, error) {
	subjectParts := strings.Split(msg.Subject, ".")
	if len(subjectParts) != 3 {
		return nil, errors.New("invalid subject format")
	}

	groupSlug := subjectParts[1]
	clientSlug := subjectParts[2]

	client, err := ls.Clients.GetBySlug(groupSlug, clientSlug)
	clientID := client.ID
	if err != nil {
		return nil, errors.New("client not found")
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

	return &logs, nil
}

func (ls *LogsService) Save(logs *models.Logs) error {
	if logs == nil || len(*logs) == 0 {
		return nil
	}

	rows := make([][]interface{}, len(*logs))
	for i, log := range *logs {
		attrs, err := json.Marshal(log.Attributes)
		if err != nil {
			ls.Log.Error("Failed to marshal attributes", "error", err, "log_id", log.ID)
			return fmt.Errorf("marshal attributes: %w", err)
		}

		rows[i] = []interface{}{
			log.ClientID,
			log.Level,
			log.LoggedAt,
			log.Message,
			attrs,
		}
	}

	_, err := ls.DB.Conn().(*pgxpool.Pool).
		CopyFrom(
			context.Background(),
			pgx.Identifier{"logs"},
			[]string{
				"client_id",
				"level",
				"logged_at",
				"message",
				"attributes",
			},
			pgx.CopyFromRows(rows),
		)
	if err != nil {
		ls.Log.Error("Failed to copy logs to database", "error", err)
		return fmt.Errorf("copy logs: %w", err)
	}

	ls.Log.Info("Successfully saved logs", "count", len(*logs))
	return nil
}
