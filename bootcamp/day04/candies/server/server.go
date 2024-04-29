package server

import (
	"candies/server/packages/security"
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"os"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	caFile, err := os.ReadFile("./certs/minica.pem")
	if err != nil {
		return err
	}
	systemCAs := x509.NewCertPool()
	ok := systemCAs.AppendCertsFromPEM(caFile)
	if !ok {
		panic("error [server]")
	}

	tlsConfig := tls.Config{
		ClientCAs:             systemCAs,
		ClientAuth:            tls.RequireAndVerifyClientCert,
		GetCertificate:        security.LoadServerCert("./certs/candies/cert.pem", "./certs/candies/key.pem"),
		VerifyPeerCertificate: security.VerifyCertChains,
	}

	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		TLSConfig:      &tlsConfig,
	}

	return s.httpServer.ListenAndServeTLS("", "")
}
