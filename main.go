package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/ncaak/pifiabot/models"
	"github.com/ncaak/pifiabot/server"
)

func getInput(body io.ReadCloser) (input models.Input) {
	var update models.Update
	var decoder = json.NewDecoder(body)
	if err := decoder.Decode(&update); err != nil {
		log.Println("INFO :: Decoding request body : " + err.Error())
		return
	}

	if entities := update.Message.Entities; len(entities) == 0 || entities[0].Type != "bot_command" {
		log.Println("INFO :: Request body was not a Bot Command")
		return
	}

	return models.Input{
		ChatId:    strconv.Itoa(update.Message.Chat.Id),
		MessageId: strconv.Itoa(update.Message.Id),
		Text:      update.Message.Text,
	}
}

func webhook(w http.ResponseWriter, r *http.Request) {
	log.Println("INFO :: Request received")

	var input = getInput(r.Body)
	log.Println(input)

	// TODO : Handle command

}

func main() {

	fmt.Println("INFO :: Starting the server...")

	var service, err = server.Build()
	if err != nil {
		log.Println("ERROR :: There was an error when building the service")
		log.Fatal(err)
	}

	service.AddRoute("/v1/bot-api", webhook)

	log.Fatal(
		service.Listen("8443"),
	)
}
