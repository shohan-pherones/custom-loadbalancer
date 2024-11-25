package main

import (
	"io"
	"log"
	"net/http"
	"sync/atomic"
)

type Backend struct {
	URL    string
	Active bool
}

type Loadbalancer struct {
	backends []Backend
	counter  uint64
}

func NewLoadBalancer(backends []Backend) *Loadbalancer {
	return &Loadbalancer{backends: backends}
}

func (lb *Loadbalancer) GetNextBackend() *Backend {
	for i := 0; i < len(lb.backends); i++ {
		index := atomic.AddUint64(&lb.counter, 1) % uint64(len(lb.backends))
		backend := &lb.backends[index]
		if backend.Active {
			return backend
		}
	}
	return nil
}

func (lb *Loadbalancer) ProxyResquest(rw http.ResponseWriter, r *http.Request) {
	backend := lb.GetNextBackend()
	if backend == nil {
		http.Error(rw, "No active backend available", http.StatusUnavailableForLegalReasons)
		return
	}

	res, err := http.Get(backend.URL + r.URL.Path)
	if err != nil {
		http.Error(rw, "Failed to reach backend", http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	for key, values := range res.Header {
		for _, value := range values {
			rw.Header().Add(key, value)
		}
	}
	rw.WriteHeader(res.StatusCode)
	io.Copy(rw, res.Body)
}

func main() {
	backends := []Backend{
		{URL: "http://server1:8081", Active: true},
		{URL: "http://server2:8082", Active: true},
		{URL: "http://server3:8083", Active: true},
	}

	lb := NewLoadBalancer(backends)

	http.HandleFunc("/", lb.ProxyResquest)
	log.Println("Load Balancer running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
