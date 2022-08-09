package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
)

func main() {
	caCert, _ := ioutil.ReadFile("ca.crt")
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	cert, _ := tls.LoadX509KeyPair("client.crt", "client.key")

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
				Certificates: []tls.Certificate{cert},
			},
		},
	}

	// Make a request
	r, err := client.Post("https://localhost:8080", "text/plain", nil)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}
