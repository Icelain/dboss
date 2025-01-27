package transport

import (
	"net/http"
	"sync"
)

type BufferedTransport struct {
	defaultTransport http.RoundTripper

	currentConnections uint64
	maxConnections     uint64
	mu                 *sync.Mutex
}

func (b *BufferedTransport) RoundTrip(req *http.Request) (*http.Response, error) {

	for b.maxConnections > b.currentConnections {

	}

	return b.defaultTransport.RoundTrip(req)

}
