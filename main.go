package main

import (
	"fmt"
	"io"
	"net/http"
)

type cachedPage struct {
	StatusCode int
	Header http.Header
	Body []byte
}


func main() {

	cache := make(map[string]cachedPage)
	url := "https://motamem.org/"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Errorf("Can't open the site")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	cache[url] = cachedPage{
		StatusCode: resp.StatusCode,
		Header: resp.Header.Clone(),
		Body: body,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cached := cache[url]
		w.WriteHeader(cached.StatusCode)

		for k, v := range cached.Header {
			for _, val := range v {
				w.Header().Add(k, val)
			}
		}

		w.Write(cached.Body)
	})

	http.ListenAndServe(":8080", nil)




}