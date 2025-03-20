package repository

import (
	"context"
	"database/sql"
	"fmt"
	"sandbox/internal/storage/database"
)

type Repo struct {
	db database.DBConnector
}

var AllTables = []string{
	TableChatMessages,
	TableObserveChats,
}

func InitRepo(db database.DBConnector) *Repo {
	return &Repo{db: db}
}

func (r *Repo) CleanupTables(ctx context.Context) error {
	for _, table := range AllTables {
		q := fmt.Sprintf("DELETE FROM %s", table)
		if _, err := r.db.Client().ExecContext(ctx, q); err != nil {
			return err
		}
	}
	return nil
}

func checkUpdated(result sql.Result, err error) error {
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}
	return nil
}
