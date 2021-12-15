package models

type Update struct {
	Message struct {
		Id   int    `json:"message_id"`
		Text string `json:"text"`
		Chat struct {
			Id int `json:"id"`
		} `json:"chat"`
		Entities []struct {
			Type string `json:"type"`
		} `json:"entities"`
	} `json:"message,omitempty"`
}
