package proxy

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func Serve(port uint, to *url.URL) {

	proxy := httputil.NewSingleHostReverseProxy(to)
	proxy.Transport =

		http.ListenAndServe(fmt.Sprintf(":%d", port), proxy)
}
