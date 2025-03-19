package repository

import (
	"context"
	"fmt"
	"sandbox/internal/entities"
	"sandbox/internal/utils"
)

const (
	TableObserveChats = "observe_chats"
)

func (r *Repo) SaveObserveChats(ctx context.Context, chats []entities.ObserveChat) error {
	chunks := utils.ChunkSlice(chats, 10)
	for _, chunk := range chunks {
		q, p := utils.GenerateBulkInsertSQL(TableObserveChats, utils.PQParamPlaceholder, chunk, func(e entities.ObserveChat) map[string]any {
			return map[string]any{
				"chat_id":   e.ChatID,
				"chat_name": e.ChatName,
				"chat_nick": e.ChatNick,
				"active":    e.Active,
			}
		})
		if _, err := r.db.Client().ExecContext(ctx, q, p...); err != nil {
			return fmt.Errorf("failed to save observe chats: %w", err)
		}
	}
	return nil
}
