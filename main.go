package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ncaak/pifiabot/server"
)

func webhook(w http.ResponseWriter, r *http.Request) {
	log.Println("INFO :: Request received")
	fmt.Fprintf(w, "OK")
}

func main() {

	fmt.Println("INFO :: Starting the server...")

	var service = server.Build()

	service.AddRoute("/v1/bot-api", webhook)

	log.Fatal(
		service.Listen("8443"),
	)
}
