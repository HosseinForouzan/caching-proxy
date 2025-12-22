package cache

import (
	"fmt"
	"io"
	"net/http"
)

func Get(url string) (*CachedPage, error) {
	resp, err := http.Get(url)
	if err != nil {
		return &CachedPage{}, fmt.Errorf("can't get the url: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &CachedPage{}, fmt.Errorf("error in reading body %w", err)
	}

	return &CachedPage{
		StatusCode: resp.StatusCode,
		Header:     resp.Header.Clone(),
		Body:       body,
	}, nil

}

func Set(url string, cachedPage *CachedPage){

		CacheMemory[url] = *cachedPage
}