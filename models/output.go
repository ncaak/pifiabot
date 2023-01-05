package models

type Output struct {
	ChatId    int
	MessageId int
	Text      string
}

type SetWebhook struct {
	Url         string
	Certificate []byte
}
