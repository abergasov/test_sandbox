package repository

import "sandbox/internal/storage/database"

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
