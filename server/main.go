package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
)

func main() {
	// ...

	caCert, _ := ioutil.ReadFile("ca.crt")
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		ClientCAs:  caCertPool,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}
	tlsConfig.BuildNameToCertificate()

	server := &http.Server{
		Addr:      ":9443",
		TLSConfig: tlsConfig,
	}

	server.ListenAndServe()

	server.ListenAndServeTLS("server.crt", "server.key")

	// ...
}
