package proxy

import (
	"dboss/internal/transport"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func Serve(port, maxconn uint64, to *url.URL) {

	proxy := httputil.NewSingleHostReverseProxy(to)
	proxy.Transport = transport.NewBufferedTransport(maxconn)

	log.Printf("listening on port %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), proxy)
}
