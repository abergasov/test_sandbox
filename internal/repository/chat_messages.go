package repository

import (
	"context"
	"fmt"
	"sandbox/internal/entities"
	"sandbox/internal/utils"
	"strings"
)

const (
	TableChatMessages = "chat_messages"
)

var (
	fieldsChatMessages = []string{
		"chat_id",
		"message_id",
		"timestamp",
		"timestamp_num",
		"user_id",
		"message",
		"reply_to",
		"is_bot",
		"is_deleted",
		"is_edited",
	}
	fieldsChatMessagesColumns = strings.Join(fieldsChatMessages, ",")
)

func (r *Repo) SaveChatMessages(ctx context.Context, chats []entities.ChatMessage) error {
	chunks := utils.ChunkSlice(chats, 20)
	for _, chunk := range chunks {
		q, p := utils.GenerateBulkInsertSQL(TableChatMessages, utils.PQParamPlaceholder, chunk, func(e entities.ChatMessage) map[string]interface{} {
			return map[string]interface{}{
				"chat_id":       e.ChatID,
				"message_id":    e.MessageID,
				"timestamp":     e.Timestamp,
				"timestamp_num": e.TimestampNum,
				"user_id":       e.UserID,
				"message":       e.Message,
				"reply_to":      e.ReplyTo,
				"is_bot":        e.IsBot,
				"is_deleted":    e.IsDeleted,
				"is_edited":     e.IsEdited,
			}
		})
		if _, err := r.db.Client().ExecContext(ctx, q, p...); err != nil {
			return fmt.Errorf("failed to save chat messages: %w", err)
		}
	}
	return nil
}

func (r *Repo) GetChatMessages(ctx context.Context, pagination *entities.Pagination) ([]*entities.ChatMessage, error) {
	q := fmt.Sprintf("SELECT %s FROM %s ORDER BY message_id DESC LIMIT $1 OFFSET $2", fieldsChatMessagesColumns, TableChatMessages)
	return utils.QueryRowsToStruct[entities.ChatMessage](ctx, r.db.Client(), q, pagination.PerPage, pagination.Page)
}

func (r *Repo) GetAllChatMessages(ctx context.Context) ([]*entities.ChatMessage, error) {
	q := fmt.Sprintf("SELECT %s FROM %s", fieldsChatMessagesColumns, TableChatMessages)
	return utils.QueryRowsToStruct[entities.ChatMessage](ctx, r.db.Client(), q)
}

func (r *Repo) GetChatMessageByID(ctx context.Context, messageID uint64) (*entities.ChatMessage, error) {
	q := fmt.Sprintf("SELECT %s FROM %s WHERE message_id = $1", fieldsChatMessagesColumns, TableChatMessages)
	return utils.QueryRowToStruct[entities.ChatMessage](ctx, r.db.Client(), q, messageID)
}

func (r *Repo) UpdateChatMessageByID(ctx context.Context, messageID uint64, message *entities.EditChatMessage) error {
	q := fmt.Sprintf("UPDATE %s SET message = $1, is_bot = $2 WHERE message_id = $3", TableChatMessages)
	return checkUpdated(r.db.Client().ExecContext(ctx, q, message.NewText, message.IsBot, messageID))
}

func (r *Repo) DeleteChatMessageByID(ctx context.Context, messageID uint64) error {
	q := fmt.Sprintf("DELETE FROM %s WHERE message_id = $1", TableChatMessages)
	_, err := r.db.Client().ExecContext(ctx, q, messageID)
	return err
}
