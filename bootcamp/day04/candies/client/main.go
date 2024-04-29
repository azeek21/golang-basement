package main

import (
	"candies/server/packages/security"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	client := getClient()
	resp, err := client.Get("https://candies:8082/buy-candy")
	must(err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	must(err)

	fmt.Printf("Status: %s  Body: %s\n", resp.Status, string(body))
}

func getClient() *http.Client {
	cp := x509.NewCertPool()
	data, err := os.ReadFile("./certs/minica.pem")
	if err != nil {
		panic(err.Error())
	}
	ok := cp.AppendCertsFromPEM(data)
	if !ok {
		panic("Failed to append ceritifcate to pool")
	}

	config := &tls.Config{
		RootCAs:               cp,
		GetClientCertificate:  security.LoadClientCert("./certs/candies/cert.pem", "./certs/candies/key.pem"),
		VerifyPeerCertificate: security.VerifyCertChains,
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: config,
		},
	}
	return client
}

func must(err error) {
	if err != nil {
		fmt.Printf("CLIENT: [fatal]: %v\n", err.Error())
		os.Exit(1)
	}
}
