package services

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/go-raptor/raptor/v4"
	"github.com/go-raptor/raptor/v4/errs"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DatabaseService struct {
	raptor.Service
}

var ctx = context.Background()

func (d *DatabaseService) Conn() *pgxpool.Pool {
	return d.Database.Conn().(*pgxpool.Pool)
}

func Select[T any](db *DatabaseService, query squirrel.SelectBuilder) ([]T, error) {
	sql, args, sqlErr := query.ToSql()
	if sqlErr != nil {
		db.Log.Error("Error building SQL query", "error", sqlErr)
		return nil, errs.NewErrorInternal("Error building SQL query", "error", sqlErr.Error())
	}

	rows, err := db.Conn().Query(ctx, sql, args...)
	if err != nil {
		db.Log.Error("Error executing SQL query", "error", err)
		return nil, PgError(err)
	}
	defer rows.Close()

	results, err := pgx.CollectRows(rows, pgx.RowToStructByName[T])
	if err != nil {
		db.Log.Error("Error collecting rows", "error", err)
		return nil, PgError(err)
	}

	return results, nil
}

func SelectOne[T any](db *DatabaseService, query squirrel.SelectBuilder) (T, error) {
	sql, args, sqlErr := query.ToSql()
	if sqlErr != nil {
		db.Log.Error("Error building SQL query", "error", sqlErr)
		return *new(T), errs.NewErrorInternal("Error building SQL query", "error", sqlErr.Error())
	}

	rows, err := db.Conn().Query(ctx, sql, args...)
	if err != nil {
		db.Log.Error("Error executing SQL query", "error", err)
		return *new(T), PgError(err)
	}
	defer rows.Close()

	result, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[T])
	if err != nil {
		db.Log.Error("Error collecting row", "error", err)
		return *new(T), PgError(err)
	}

	return result, nil
}

func Insert[T any](db *DatabaseService, query squirrel.InsertBuilder) (T, error) {
	sql, args, sqlErr := query.ToSql()
	if sqlErr != nil {
		db.Log.Error("Error building SQL query", "error", sqlErr)
		return *new(T), errs.NewErrorInternal("Error building SQL query", "error", sqlErr.Error())
	}

	rows, err := db.Conn().Query(ctx, sql, args...)
	if err != nil {
		db.Log.Error("Error executing SQL query", "error", err)
		return *new(T), PgError(err)
	}
	defer rows.Close()

	result, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[T])
	if err != nil {
		db.Log.Error("Error collecting row on insert", "error", err)
		return *new(T), PgError(err)
	}

	return result, nil
}

func Update[T any](db *DatabaseService, query squirrel.UpdateBuilder) (T, error) {
	sql, args, sqlErr := query.ToSql()
	if sqlErr != nil {
		db.Log.Error("Error building SQL query", "error", sqlErr)
		return *new(T), errs.NewErrorInternal("Error building SQL query", "error", sqlErr.Error())
	}

	rows, err := db.Conn().Query(ctx, sql, args...)
	if err != nil {
		db.Log.Error("Error executing SQL query", "error", err)
		return *new(T), PgError(err)
	}
	defer rows.Close()

	result, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[T])
	if err != nil {
		db.Log.Error("Error collecting row on update", "error", err)
		return *new(T), PgError(err)
	}

	return result, nil
}

func Delete[T any](db *DatabaseService, query squirrel.DeleteBuilder) (T, error) {
	sql, args, sqlErr := query.ToSql()
	if sqlErr != nil {
		db.Log.Error("Error building SQL query", "error", sqlErr)
		return *new(T), errs.NewErrorInternal("Error building SQL query", "error", sqlErr.Error())
	}

	rows, err := db.Conn().Query(ctx, sql, args...)
	if err != nil {
		db.Log.Error("Error executing SQL query", "error", err)
		return *new(T), PgError(err)
	}
	defer rows.Close()

	result, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[T])
	if err != nil {
		db.Log.Error("Error collecting row on delete", "error", err)
		return *new(T), PgError(err)
	}

	return result, nil
}

func Exec(db *DatabaseService, query squirrel.Sqlizer) (int64, error) {
	sql, args, sqlErr := query.ToSql()
	if sqlErr != nil {
		db.Log.Error("Error building SQL query", "error", sqlErr)
		return 0, errs.NewErrorInternal("Error building SQL query", "error", sqlErr.Error())
	}

	result, err := db.Conn().Exec(ctx, sql, args...)
	if err != nil {
		db.Log.Error("Error executing SQL query", "error", err)
		return 0, PgError(err)
	}

	return result.RowsAffected(), nil
}

func BatchInsert(db *DatabaseService, queries []squirrel.InsertBuilder) error {
	batch := &pgx.Batch{}
	for _, query := range queries {
		sql, args, err := query.ToSql()
		if err != nil {
			db.Log.Error("Error building SQL query for batch insert", "error", err)
			return errs.NewErrorInternal("Error building SQL query for batch insert", "error", err.Error())
		}
		batch.Queue(sql, args...)
	}

	results := db.Conn().SendBatch(ctx, batch)
	defer results.Close()

	_, err := results.Exec()
	if err != nil {
		db.Log.Error("Error executing batch insert", "error", err)
		return PgError(err)
	}

	return nil
}

func PgError(err error) error {
	if err == pgx.ErrNoRows {
		return errs.NewErrorNotFound("Resource not found")
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code[:2] {
		case "23": // Integrity constraint violations (e.g., unique, foreign key)
			return errs.NewErrorConflict("Conflict while executing query", "state", pgErr.Code, "message", pgErr.Message, "detail", pgErr.Detail, "hint", pgErr.Hint)
		case "42": // Syntax or access rule violations
			return errs.NewErrorBadRequest("Invalid query syntax or access violation", "state", pgErr.Code, "message", pgErr.Message, "detail", pgErr.Detail)
		case "08": // Connection exceptions
			return errs.NewErrorInternal("Database connection error", "state", pgErr.Code, "message", pgErr.Message)
		case "53": // Insufficient resources (e.g., too many connections)
			return errs.NewErrorInternal("Insufficient database resources", "state", pgErr.Code, "message", pgErr.Message)
		default:
			return errs.NewErrorInternal("Database error", "state", pgErr.Code, "message", pgErr.Message, "detail", pgErr.Detail, "hint", pgErr.Hint)
		}
	}

	return errs.NewErrorInternal("Unknown database error", "error", err.Error())
}
