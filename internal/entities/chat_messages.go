package entities

type EditChatMessage struct {
	NewText string `json:"new_text"`
	IsBot   bool   `json:"is_bot"`
}
type ChatMessage struct {
	ChatID       uint64 `db:"chat_id" json:"chat_id"`
	MessageID    uint64 `db:"message_id" json:"message_id"`
	Timestamp    string `db:"timestamp" json:"timestamp"`
	TimestampNum uint64 `db:"timestamp_num" json:"timestamp_num"`
	UserID       uint64 `db:"user_id" json:"user_id"`
	Message      string `db:"message" json:"message"`
	ReplyTo      uint64 `db:"reply_to" json:"reply_to"`
	IsBot        bool   `db:"is_bot" json:"is_bot"`
	IsDeleted    bool   `db:"is_deleted" json:"is_deleted"`
	IsEdited     bool   `db:"is_edited" json:"is_edited"`
}
