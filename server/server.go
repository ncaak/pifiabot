package server

import (
	"crypto/tls"
	"log"
	"net/http"
	"time"
)

type server struct {
	certs []tls.Certificate
	mux   *http.ServeMux
}

func Build() *server {
	var cert, err = tls.LoadX509KeyPair("cert.pem", "private.key")
	if err != nil {
		log.Fatalf("ERROR :: Security keys could not be retrieved : %v", err.Error())
	}

	return &server{
		certs: []tls.Certificate{cert},
		mux:   http.NewServeMux(),
	}
}

func (s server) AddRoute(method string, logic func(http.ResponseWriter, *http.Request)) {
	s.mux.HandleFunc(method, logic)
}

func (s server) Listen(port string) error {
	var httpServer = &http.Server{
		Addr:    ":" + port,
		Handler: s.mux,
		TLSConfig: &tls.Config{
			Certificates: s.certs,
		},
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return httpServer.ListenAndServeTLS("", "")
}
