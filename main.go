package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/HosseinForouzan/caching-proxy/cache"
	"github.com/HosseinForouzan/caching-proxy/proxy"
)



func main() {

	flag.StringVar(&cache.OriginServer, "origin", "", "URL of the origin server")
	flag.StringVar(&cache.Port, "port", "8080", "the port of the proxy server")
	flag.Parse()

	if cache.OriginServer == "" {
		log.Fatal("origin server must be filled")
	}
	

	http.HandleFunc("/", proxy.ForwardRequest)
	fmt.Println("Listening on port ",cache.Port, "..." )
	http.ListenAndServe(":" + cache.Port, nil)
}




