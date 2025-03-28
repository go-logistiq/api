package migrations

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type CreateLogs struct{}

func (m CreateLogs) Name() string {
	return "create_logs_table"
}

func (m CreateLogs) Up(tx pgx.Tx) error {
	_, err := tx.Exec(context.Background(), `
		CREATE TABLE logs (
			id BIGSERIAL PRIMARY KEY,
			client_id INT NOT NULL,
			logged_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
			level SMALLINT NOT NULL CHECK (level BETWEEN -10 AND 20),
			message VARCHAR(512) NOT NULL,
			attributes JSONB NOT NULL DEFAULT '{}',
			CONSTRAINT fk_client
				FOREIGN KEY (client_id)
				REFERENCES clients(id)
				ON DELETE CASCADE
		)
	`)
	if err != nil {
		return err
	}

	_, err = tx.Exec(context.Background(), `
		CREATE INDEX logs_logged_at_idx ON logs (logged_at);
	`)
	if err != nil {
		return err
	}

	_, err = tx.Exec(context.Background(), `
		CREATE INDEX logs_client_id_idx ON logs (client_id);
	`)
	if err != nil {
		return err
	}

	_, err = tx.Exec(context.Background(), `
		CREATE INDEX logs_level_idx ON logs (level);
	`)
	if err != nil {
		return err
	}

	return nil
}

func (m CreateLogs) Down(tx pgx.Tx) error {
	_, err := tx.Exec(context.Background(), `
		DROP TABLE IF EXISTS logs
	`)
	return err
}
