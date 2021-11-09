package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
)

func getCertificates() []tls.Certificate {
	var cert, err = tls.LoadX509KeyPair("localhost.crt", "localhost.key")
	if err != nil {
		log.Fatalf("ERROR :: Certicate could not be retrieved : %v", err.Error())
	}
	return []tls.Certificate{cert}
}

func getMuxHandler() *http.ServeMux {
	var mux = http.NewServeMux()
	// API Methods
	mux.HandleFunc("/v1/bot-api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("que entro")
	})

	return mux
}

func getServer() *http.Server {
	return &http.Server{
		Addr:    ":8443",
		Handler: getMuxHandler(),
		TLSConfig: &tls.Config{
			Certificates: getCertificates(),
		},
	}
}

func main() {

	fmt.Println("INFO :: Starting the server...")

	log.Fatal(
		getServer().ListenAndServeTLS("", ""),
	)

}
