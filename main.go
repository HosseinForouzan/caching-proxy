package main

import (
	"net/http"

	"github.com/HosseinForouzan/caching-proxy/proxy"
)

func main() {

	http.HandleFunc("/", proxy.ForwardRequest)
	http.ListenAndServe(":8080", nil)
}




