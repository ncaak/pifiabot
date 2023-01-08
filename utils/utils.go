package utils

import (
	"encoding/json"
	"io"
	"log"

	"github.com/ncaak/pifiabot/models"
)

func GetRequestInput(body io.ReadCloser) (models.Input, error) {
	var update models.Update
	defer body.Close()

	var decoder = json.NewDecoder(body)
	if err := decoder.Decode(&update); err != nil {
		log.Println("ERROR :: Decoding Update structure")
		return models.Input{}, err
	}

	input := getInputFromUpdate(update)
	return input, nil
}

func getInputFromUpdate(update models.Update) models.Input {
	var entities = update.Message.Entities

	return models.Input{
		ChatId:    update.Message.Chat.Id,
		IsCommand: len(entities) > 0 && entities[0].Type == "bot_command",
		MessageId: update.Message.Id,
		Text:      update.Message.Text,
	}
}
