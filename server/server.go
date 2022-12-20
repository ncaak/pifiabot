package server

import (
	"crypto/tls"
	"errors"
	"net/http"
	"time"
)

type server struct {
	certs []tls.Certificate
	mux   *http.ServeMux
}

func Build(certificate []byte, privateKey []byte) (*server, error) {
	var cert, err = tls.X509KeyPair(certificate, privateKey)
	if err != nil {
		return &server{}, errors.New("ERROR :: Security keys could not be retrieved : " + err.Error())
	}

	return &server{
		certs: []tls.Certificate{cert},
		mux:   http.NewServeMux(),
	}, nil
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
