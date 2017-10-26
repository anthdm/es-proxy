package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	endpoint := "http://localhost:9200"
	url, err := url.Parse(endpoint)
	if err != nil {
		panic(err)
	}
	id := os.Getenv("NEW_RELIC_ID")

	log.Fatal(http.ListenAndServe(":3000", NewEsProxy(url, id)))
}

type EsProxy struct {
	p          *httputil.ReverseProxy
	newRelicID string
}

func NewEsProxy(target *url.URL, id string) *EsProxy {
	return &EsProxy{
		p:          httputil.NewSingleHostReverseProxy(target),
		newRelicID: id,
	}
}

func (proxy *EsProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if len(r.Header["X-Newrelic-Id"]) > 0 && r.Header["X-Newrelic-Id"][0] == proxy.newRelicID {
		proxy.p.ServeHTTP(w, r)
		return
	}
	http.Error(w, "Not authorized", 403)
}
