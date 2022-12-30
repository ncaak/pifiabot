package main

import (
	"log"
	"net/http"

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

		client.Get().SendMessage(output)
	}
}

func main() {

	log.Println("INFO :: Setting up Configuration")
	if err := config.Setup(); err != nil {
		log.Fatal(err)
	}

	log.Println("INFO :: Setting up Webhook")
	client.Setup(config.Get().BotToken)
	var webhook = models.SetWebhook{
		Url:         config.Get().GetEndpoint(),
		Certificate: config.Get().CertBytes,
	}
	if err := client.Get().SetWebhook(webhook); err != nil {
		log.Fatal(err)
	}

	log.Println("INFO :: Starting the server...")
	http.HandleFunc("/"+config.Get().Url.Path, handleUpdate)

	log.Fatal(
		http.ListenAndServeTLS(
			":"+config.Get().Url.Port,
			config.Get().CertFilePath,
			config.Get().KeyFilePath,
			nil),
	)
}
