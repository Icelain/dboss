package transport

import (
	"log"
	"net/http"
	"sync"
	"sync/atomic"
)

type BufferedTransport struct {
	defaultTransport http.RoundTripper

	currentConnections atomic.Uint64
	maxConnections     uint64
	cond               *sync.Cond
}

func (b *BufferedTransport) RoundTrip(req *http.Request) (*http.Response, error) {

	b.currentConnections.Add(1)
	log.Println("connect")

	defer func() {

		cur := b.currentConnections.Load()
		b.currentConnections.Store(cur - 1)

		log.Printf("disconnect at %d\n", cur)

	}()

	b.cond.Signal()

	b.cond.L.Lock()
	defer b.cond.L.Unlock()

	for b.maxConnections < b.currentConnections.Load() {

		b.cond.Wait()

	}

	return b.defaultTransport.RoundTrip(req)

}

func NewBufferedTransport(maxConnections uint64) *BufferedTransport {

	mutex := &sync.Mutex{}
	cond := sync.NewCond(mutex)

	return &BufferedTransport{

		defaultTransport:   http.DefaultTransport,
		currentConnections: atomic.Uint64{}, // default value of an atomic uint64 is 0
		maxConnections:     maxConnections,
		cond:               cond,
	}

}
