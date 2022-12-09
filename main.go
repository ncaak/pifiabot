package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ncaak/pifiabot/client"
	"github.com/ncaak/pifiabot/config"
	"github.com/ncaak/pifiabot/models"
	"github.com/ncaak/pifiabot/server"
)

func handleUpdate(w http.ResponseWriter, r *http.Request) {

	log.Println("INFO :: Request received")

	defer r.Body.Close()

	var input = server.GetInput(r.Body)

	if input.IsCommand {

		// TODO : Handle command

		// stub
		var output = models.Output{
			ChatId:    input.ChatId,
			MessageId: input.MessageId,
			Text:      "pong",
		}

		var telegram = client.Build(os.Getenv("BOT_TOKEN")) // TODO : Move ENV to configuration
		telegram.SendMessage(output)
	}
}

func main() {

	log.Println("INFO :: Setting up Configuration")
	var cfg, errCfg = config.Setup()
	if errCfg != nil {
		log.Println("ERROR :: Setting up configuration")
		log.Fatal(errCfg)
	}

	log.Println("INFO :: Starting the server...")
	var service, err = server.Build(cfg)
	if err != nil {
		log.Println("ERROR :: There was an error when building the service")
		log.Fatal(err)
	}

	service.AddRoute("/v1/bot-api", handleUpdate) // TODO : Path should be configurable

	log.Fatal(
		service.Listen("8443"),
	)
}
