package cache

import (
	"net/http"
)

type CachedPage struct {
	StatusCode int
	Header http.Header
	Body []byte
}

var CacheMemory map[string]CachedPage = map[string]CachedPage{}

var (
	OriginServer string
	Port string
)