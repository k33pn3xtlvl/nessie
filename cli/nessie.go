package main

import (
	"flag"
	"fmt"
	"github.com/attwad/nessie"
)

var apiURL, username, password, caCertPath, loginFile string
var ignoreSSLCertsErrors bool

func init() {
	flag.StringVar(&apiURL, "api_url", "", "")
	flag.StringVar(&username, "username", "", "Username to login with, in production read that from a file, do not set from the command line or it will end up in your history.")
	flag.StringVar(&password, "password", "", "Password that matches the provided username, in production read that from a file, do not set from the command line or it will end up in your history.")
	flag.BoolVar(&ignoreSSLCertsErrors, "ignore_ssl_certs_errors", false, "If true, transport will not verify the certificate of the Nessus server exposing the API, seriously, do not do this outside of a dev environment...")
	flag.StringVar(&caCertPath, "ca_cert_path", "", "Path to the public certificate of the Nessus server exposing the API, provide this instead of setting ignore_ssl_certs_errors to true in production environments")
	flag.Parse()
}

func main() {
	nessus, err := nessie.NewNessus(apiURL, caCertPath, ignoreSSLCertsErrors)
	if err != nil {
		panic(err)
	}

	if err := nessus.Login(username, password); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("Logged-in!\n")
	defer fmt.Println(nessus.Logout())
}