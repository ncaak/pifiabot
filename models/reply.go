package models

type Reply struct {
	ChatId  int    `json:"chat_id"`
	ReplyId int    `json:"reply_to_message_id"`
	Text    string `json:"text"`
}
