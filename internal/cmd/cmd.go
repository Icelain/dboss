package cmd

import (
	"dboss/internal/proxy"
	"flag"
	"log"
	"net/url"
)

func Execute() {

	maxconn := flag.Uint("maxconn", 2000, "Maximum number of connections for the rate limiter to handle")
	port := flag.Uint("port", 8080, "The port to run the rate limiting reverse proxy on")
	address := flag.String("address", "", "Address to the server the traffic is being passed to")

	flag.Parse()

	if *address == "" {

		log.Fatal("Enter a valid address")

	}

	socketAddressUrl, err := url.Parse(*address)
	if err != nil {

		log.Fatal(err)

	}

	proxy.Serve(*port, socketAddressUrl)

}
