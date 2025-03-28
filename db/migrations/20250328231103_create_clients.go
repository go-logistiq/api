package migrations

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type CreateClients struct{}

func (m CreateClients) Name() string {
	return "create_clients_table"
}

func (m CreateClients) Up(tx pgx.Tx) error {
	_, err := tx.Exec(context.Background(), `
		CREATE TABLE clients (
			id SERIAL PRIMARY KEY,
			name VARCHAR(50) NOT NULL,
			group_id INT NOT NULL,
			CONSTRAINT fk_group
				FOREIGN KEY (group_id)
				REFERENCES groups(id)
				ON DELETE CASCADE,
			CONSTRAINT unique_name_in_group
				UNIQUE (name, group_id)
		)
	`)
	if err != nil {
		return err
	}

	return nil
}

func (m CreateClients) Down(tx pgx.Tx) error {
	_, err := tx.Exec(context.Background(), `
		DROP TABLE IF EXISTS clients
	`)
	return err
}
