package migrations

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type CreateGroups struct{}

func (m CreateGroups) Name() string {
	return "create_groups_table"
}

func (m CreateGroups) Up(tx pgx.Tx) error {
	_, err := tx.Exec(context.Background(), `
		CREATE TABLE groups (
			id SERIAL PRIMARY KEY,
			name VARCHAR(50) NOT NULL UNIQUE
		)
	`)
	if err != nil {
		return err
	}

	return nil
}

func (m CreateGroups) Down(tx pgx.Tx) error {
	_, err := tx.Exec(context.Background(), `
		DROP TABLE IF EXISTS groups
	`)
	return err
}
