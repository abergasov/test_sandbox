package repository

import (
	"context"
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
