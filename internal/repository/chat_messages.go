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

func (r *Repo) GetAllChatMessages(ctx context.Context) ([]*entities.ChatMessage, error) {
	q := fmt.Sprintf("SELECT %s FROM %s", fieldsChatMessagesColumns, TableChatMessages)
	return utils.QueryRowsToStruct[entities.ChatMessage](ctx, r.db.Client(), q)
}
