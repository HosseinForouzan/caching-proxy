package main

import (
	"net/http"

	"github.com/HosseinForouzan/caching-proxy/cache"
)

func main() {

	url := "https://sanjesh.org/"
	getPage, _ := cache.Get(url)

	cache.Set(url, getPage)


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(cache.CacheMemory[url].StatusCode)

		for k, v := range cache.CacheMemory[url].Header {
			for _, val := range v {
				w.Header().Add(k, val)
			}
		}

		w.Write(cache.CacheMemory[url].Body)
	})

	http.ListenAndServe(":8080", nil)
}




