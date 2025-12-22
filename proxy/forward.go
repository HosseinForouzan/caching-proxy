package proxy

import (
	"fmt"
	"net/http"
	"time"

	"github.com/HosseinForouzan/caching-proxy/cache"
)

func GetUrl() string {
	url := "https://motamem.org"

	return url
}

func checkExistenceOfUrl(url string) bool {
	if _, ok := cache.CacheMemory[url]; ok {
		return true
	}

	return false
}

func ForwardRequest(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	url := GetUrl()
	if checkExistenceOfUrl(url) {
		w.Header().Set("X-cache", "HIT")
		fmt.Println("HIT")
	}else {
		w.Header().Set("X-cache", "MISS")

		fmt.Println("MISS")
		getPage, _ := cache.Get(url)
		cache.Set(url, getPage)
	}

	w.WriteHeader(cache.CacheMemory[url].StatusCode)
		for k, v := range cache.CacheMemory[url].Header {
			for _, val := range v {
				w.Header().Add(k, val)
			}
		}

		w.Write(cache.CacheMemory[url].Body)

		elapsed := time.Since(start).Seconds()

		fmt.Println(elapsed)
}