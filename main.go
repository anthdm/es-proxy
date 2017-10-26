package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	endpoint := "http://localhost:9200"
	url, err := url.Parse(endpoint)
	if err != nil {
		panic(err)
	}

	log.Fatal(http.ListenAndServe(":3000", NewEsProxy(url)))
}

type EsProxy struct {
	p *httputil.ReverseProxy
}

func NewEsProxy(target *url.URL) *EsProxy {
	return &EsProxy{httputil.NewSingleHostReverseProxy(target)}
}

func (proxy *EsProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Host)
	proxy.p.ServeHTTP(w, r)
}
