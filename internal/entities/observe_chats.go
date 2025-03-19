package entities

type ObserveChat struct {
	ChatID   uint64 `db:"chat_id" json:"ChatID"`
	ChatName string `db:"chat_name" json:"ChatName"`
	ChatNick string `db:"chat_nick" json:"ChatNick"`
	Active   bool   `db:"active" json:"Active"`
}
